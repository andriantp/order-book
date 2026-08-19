package orderbook

type SnapshotStore interface {
	Save(snapshot Snapshot) error
	Load() (*Snapshot, error)
}

type MemorySnapshotStore struct {
	snapshot *Snapshot
}

func NewMemorySnapshotStore() *MemorySnapshotStore {
	return &MemorySnapshotStore{}
}

func (s *MemorySnapshotStore) Save(snapshot Snapshot) error {
	s.snapshot = &snapshot
	return nil
}

func (s *MemorySnapshotStore) Load() (*Snapshot, error) {
	return s.snapshot, nil
}
