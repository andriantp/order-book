package orderbook

import "fmt"

func (ob *OrderBook) Cancel(side Side, id int) error {

	// part-1: choose side
	var levels map[float64]*PriceLevel

	switch side {
	case Sell:
		levels = ob.Asks
	default:
		levels = ob.Bids
	}

	// part-2: search price level
	for _, level := range levels {

		// part-3: search order
		for i, order := range level.Orders {

			if order.ID != id {
				continue
			}

			// part-4: remove order
			level.Orders = append(
				level.Orders[:i],
				level.Orders[i+1:]...,
			)

			// part-5: cleanup
			if len(level.Orders) == 0 {
				delete(levels, level.Price)
			}

			// part-6: Recording OrderCancelled Events
			ob.eventStore.Append(OrderCancelled{
				OrderID: order.ID,
				Price:   order.Price,
				Side:    order.Side,
			})

			return nil
		}
	}

	return fmt.Errorf("order #%d not found", id)
}
