package data

type Repo interface {
	GetLocationHistoryByOrderID(orderID string, max int) ([]*History, error)
	DeleteLocationHistoryByOrderID(orderID string) error
	AddHistoryByOrderID(orderID string, history *History) error
}

type DB struct {
	// TODO: db struct for operations
}

// DB implements Repo.
var _ Repo = (*DB)(nil)

func NewRepo(inMemStorage bool) Repo {
	if inMemStorage {
		s()
		return Fx
	}
	return &DB{}
}

func (db *DB) GetLocationHistoryByOrderID(string, int) ([]*History, error) { return nil, nil }
func (db *DB) DeleteLocationHistoryByOrderID(string) error                 { return nil }
func (db *DB) AddHistoryByOrderID(string, *History) error                  { return nil }
