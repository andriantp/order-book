package marketdata

type ReadModel struct {
	Bids map[float64]int
	Asks map[float64]int

	BestBid float64
	BestAsk float64

	LastPrice float64
	LastQty   int
	Trades    []Trade
}

func NewReadModel() *ReadModel {
	return &ReadModel{
		Bids: make(map[float64]int),
		Asks: make(map[float64]int),
	}
}
