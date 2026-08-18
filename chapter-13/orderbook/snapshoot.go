package orderbook

type Snapshot struct {
	LastEventID int
	Bids        map[float64]*PriceLevel
	Asks        map[float64]*PriceLevel
}

func (ob *OrderBook) clonePriceLevels(
	levels map[float64]*PriceLevel,
) map[float64]*PriceLevel {

	cloned := make(map[float64]*PriceLevel)

	for price, level := range levels {
		newLevel := &PriceLevel{
			Price:  level.Price,
			Orders: make([]*Order, len(level.Orders)),
		}

		for i, order := range level.Orders {
			copied := *order
			newLevel.Orders[i] = &copied
		}

		cloned[price] = newLevel
	}

	return cloned
}

func (ob *OrderBook) SaveSnapshot() error {
	snapshot := Snapshot{
		LastEventID: ob.eventStore.LastEventID(),
		Bids:        ob.clonePriceLevels(ob.Bids),
		Asks:        ob.clonePriceLevels(ob.Asks),
	}
	return ob.snapshotStore.Save(snapshot)
}

func (ob *OrderBook) LoadSnapshot() error {
	snapshot, err := ob.snapshotStore.Load()
	if err != nil {
		return err
	}

	if snapshot == nil {
		return nil
	}

	ob.Bids = snapshot.Bids
	ob.Asks = snapshot.Asks

	ob.orders = make(map[int]*Order)

	for _, level := range ob.Bids {
		for _, order := range level.Orders {
			ob.orders[order.ID] = order
		}
	}

	for _, level := range ob.Asks {
		for _, order := range level.Orders {
			ob.orders[order.ID] = order
		}
	}

	return nil
}

func (ob *OrderBook) Recover() error {
	if err := ob.LoadSnapshot(); err != nil {
		return err
	}
	snapshot, err := ob.snapshotStore.Load()
	if err != nil {
		return err
	}
	if snapshot == nil {
		ob.Replay(ob.eventStore.Load())
		return nil
	}
	ob.Replay(ob.eventStore.LoadAfter(snapshot.LastEventID))
	return nil
}
