package marketdata

import "order-book/orderbook"

func (rm *ReadModel) PlaceOrder(side orderbook.Side, price float64, qty int) {
	switch side {
	case orderbook.Buy:
		rm.Bids[price] += qty
		rm.RefreshBestBid()

	case orderbook.Sell:
		rm.Asks[price] += qty
		rm.RefreshBestAsk()
	}
}

func (rm *ReadModel) ExecuteTrade(price float64, qty int) {
	rm.LastPrice = price
	rm.LastQty = qty

	rm.Trades = append(rm.Trades, Trade{
		Price: price,
		Qty:   qty,
	})
}

func (rm *ReadModel) RefreshBestBid() {
	var best float64

	for price, qty := range rm.Bids {
		if qty <= 0 {
			continue
		}
		if price > best {
			best = price
		}
	}

	rm.BestBid = best
}

func (rm *ReadModel) RefreshBestAsk() {
	var best float64

	for price, qty := range rm.Asks {
		if qty <= 0 {
			continue
		}
		if best == 0 || price < best {
			best = price
		}
	}

	rm.BestAsk = best
}
