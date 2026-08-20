package orderbook

import "fmt"

func (ob *OrderBook) Replay(events []Event) error {
	for _, event := range events {
		switch e := event.(type) {
		case OrderPlaced:
			ob.applyOrderPlaced(e)

		case TradeExecuted:
			if err := ob.applyTradeExecuted(e); err != nil {
				return err
			}

		case OrderCancelled:
			if err := ob.applyOrderCancelled(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ob *OrderBook) applyOrderPlaced(event OrderPlaced) error {
	ob.insertOrder(event.Order)
	return nil
}

func (ob *OrderBook) applyTradeExecuted(event TradeExecuted) error {
	if err := ob.applyTrade(event.BuyOrderID, event.Quantity); err != nil {
		return err
	}
	if err := ob.applyTrade(event.SellOrderID, event.Quantity); err != nil {
		return err
	}
	return nil
}

func (ob *OrderBook) applyOrderCancelled(event OrderCancelled) error {
	order, ok := ob.orders[event.OrderID]
	if !ok {
		return fmt.Errorf("order %d not found", event.OrderID)
	}

	ob.removeOrder(order)
	return nil
}

func (ob *OrderBook) applyTrade(orderID, qty int) error {
	order, ok := ob.orders[orderID]
	if !ok {
		return nil
	}

	order.Qty -= qty

	if order.Qty < 0 {
		return fmt.Errorf("negative quantity")
	}

	if order.Qty == 0 {
		ob.removeOrder(order)
	}

	return nil
}

func (ob *OrderBook) insertOrder(order *Order) {
	var levels map[float64]*PriceLevel

	switch order.Side {
	case Sell:
		levels = ob.Asks
	default:
		levels = ob.Bids
	}

	level, ok := levels[order.Price]
	if !ok {
		level = &PriceLevel{
			Price: order.Price,
		}

		levels[order.Price] = level
	}

	level.Orders = append(level.Orders, order)

	ob.orders[order.ID] = order
}

func (ob *OrderBook) removeOrder(order *Order) {

	var levels map[float64]*PriceLevel

	switch order.Side {
	case Sell:
		levels = ob.Asks
	default:
		levels = ob.Bids
	}

	level, ok := levels[order.Price]
	if !ok {
		return
	}

	for i, o := range level.Orders {
		if o.ID != order.ID {
			continue
		}

		level.Orders = append(level.Orders[:i], level.Orders[i+1:]...)
		break
	}

	delete(ob.orders, order.ID)

	if len(level.Orders) == 0 {
		delete(levels, order.Price)
	}
}
