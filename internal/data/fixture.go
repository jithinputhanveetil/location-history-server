package data

import (
	"location-history-server/internal/e"
	"sort"
	"sync"
)

func s() {
	Fx = &Fixture{}
	Fx.LocationHistory = sync.Map{}
}

var Fx *Fixture

type Fixture struct {
	LocationHistory sync.Map
}

// LocationHistory implements Repo.
var _ Repo = (*Fixture)(nil)

// GetLocationHistoryByOrderID lists the location history.
func (f *Fixture) GetLocationHistoryByOrderID(orderID string, max int) ([]*History, error) {
	val, ok := f.LocationHistory.Load(orderID)
	if !ok {
		return nil, e.ErrResourceNotFound
	}
	locations := val.([]*History)
	sort.Slice(locations, func(i, j int) bool {
		return locations[i].InsertionTime.Before(*locations[j].InsertionTime)
	})
	if max == 0 {
		return locations, nil
	}
	if max > len(locations) {
		max = len(locations)
	}
	locations = locations[len(locations)-max:]
	return locations, nil
}

// AddHistoryByOrderID adds the location history.
func (f *Fixture) AddHistoryByOrderID(orderID string, history *History) error {
	histories := []*History{}
	val, ok := f.LocationHistory.Load(orderID)
	if !ok {
		histories = append(histories, history)
	} else {
		histories = val.([]*History)
		histories = append(histories, history)
	}
	f.LocationHistory.Store(orderID, histories)
	return nil
}

// DeleteLocationHistoryByOrderID deletes the location history.
func (f *Fixture) DeleteLocationHistoryByOrderID(orderID string) error {
	_, ok := f.LocationHistory.Load(orderID)
	if !ok {
		return e.ErrResourceNotFound
	}
	f.LocationHistory.Delete(orderID)
	return nil
}
