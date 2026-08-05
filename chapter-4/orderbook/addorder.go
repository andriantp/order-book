package orderbook

func (ob *OrderBook) AddOrder(order *Order) {
	var levels map[float64]*PriceLevel

	if order.Side == Buy {
		levels = ob.Bids
	} else {
		levels = ob.Asks
	}

	level := levels[order.Price]

	if level == nil {
		level = &PriceLevel{
			Price: order.Price,
		}
		levels[order.Price] = level
	}

	level.Orders = append(level.Orders, order)
}
