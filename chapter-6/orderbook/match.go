package orderbook

import "fmt"

func (ob *OrderBook) Match(order *Order) {
	if order.Type == Market {
		fmt.Printf(
			"Incoming %s %s #%d Qty:%d\n",
			order.Type,
			order.Side,
			order.ID,
			order.Qty,
		)
	} else {
		fmt.Printf(
			"Incoming %s %s #%d @%.0f Qty:%d\n",
			order.Type,
			order.Side,
			order.ID,
			order.Price,
			order.Qty,
		)
	}

	// part-1: tentukan sisi lawan
	var levels map[float64]*PriceLevel
	if order.Side == Buy {
		levels = ob.Asks
	} else {
		levels = ob.Bids
	}

	// part-2: cari harga terbaik
	var bestLevel *PriceLevel
	if order.Side == Buy {
		bestLevel = bestAsk(levels)
	} else {
		bestLevel = bestBid(levels)
	}

	// part-3: ada kandidat?
	if bestLevel == nil {
		fmt.Println("No price level available")
		return
	}

	// part-3.5: harga benar-benar match?
	if order.Type == Limit {
		if order.Side == Buy && order.Price < bestLevel.Price {
			fmt.Println("Price does not match")
			return
		}
		if order.Side == Sell && order.Price > bestLevel.Price {
			fmt.Println("Price does not match")
			return
		}
	}

	// part-4: FIFO
	restingOrder := bestLevel.Orders[0]

	// part-5: handling partial fills
	tradedQty := order.Qty
	if restingOrder.Qty < tradedQty {
		tradedQty = restingOrder.Qty
	}

	fmt.Printf(
		"Matched #%d Qty:%d Price: %.0f\n",
		restingOrder.ID,
		tradedQty,
		bestLevel.Price,
	)

	order.Qty -= tradedQty
	restingOrder.Qty -= tradedQty

	// part-6: cleanup dan remaining order
	if restingOrder.Qty == 0 {
		bestLevel.Orders = bestLevel.Orders[1:]

		if len(bestLevel.Orders) == 0 {
			delete(levels, bestLevel.Price)
		}
	}

	if order.Type == Limit && order.Qty > 0 {
		fmt.Printf(
			"Remaining Qty:%d returned to book\n",
			order.Qty,
		)
		ob.AddOrder(order)
	}
	
	if order.Type == Market && order.Qty > 0 {
		fmt.Printf(
			"Unfilled Qty:%d discarded\n",
			order.Qty,
		)
	}

}

func bestAsk(levels map[float64]*PriceLevel) *PriceLevel {
	var best *PriceLevel
	for _, level := range levels {
		if best == nil || level.Price < best.Price {
			best = level
		}
	}
	return best
}

func bestBid(levels map[float64]*PriceLevel) *PriceLevel {
	var best *PriceLevel
	for _, level := range levels {
		if best == nil || level.Price > best.Price {
			best = level
		}
	}
	return best
}
