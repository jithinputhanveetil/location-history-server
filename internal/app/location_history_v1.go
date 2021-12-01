package app

import (
	"encoding/json"
	"location-history-server/internal/data"
	"location-history-server/internal/e"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

// initLocationHistoryRouter sets up location-history router.
func (s *Server) initLocationHistoryRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/{order_id}", s.handleGetLocationHistoryV1)
	r.Put("/{order_id}", s.handlePutLocationHistoryV1)
	r.Delete("/{order_id}", s.handleDeleteLocationHistoryV1)
	return r
}

func (s *Server) handleGetLocationHistoryV1(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "order_id")
	maxStr := r.URL.Query().Get("max")
	var (
		max int
		err error
	)
	if len(maxStr) > 0 {
		max, err = strconv.Atoi(maxStr)
		if err != nil {
			log.Println("handleGetLocationHistoryV1: invalid max value: ", err)
			fail(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	history, err := s.repo.GetLocationHistoryByOrderID(orderID, max)
	if err == e.ErrResourceNotFound {
		log.Println("requested resource not found for oderID: ", orderID)
		fail(
			w,
			http.StatusNotFound,
			"requested resource not found",
		)
		return
	}
	if err != nil {
		log.Println(err)
		fail(w, http.StatusInternalServerError, "")
		return
	}
	send(w, http.StatusOK, map[string]interface{}{
		"order_id": orderID,
		"history":  history,
	})
}

func (s *Server) handlePutLocationHistoryV1(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "order_id")
	history := new(data.History)
	if err := json.NewDecoder(r.Body).Decode(history); err != nil {
		log.Println("handlePutLocationHistoryV1: json decoding failed:", err)
		fail(w, http.StatusBadRequest, err.Error())
		return
	}
	t := time.Now()
	history.InsertionTime = &t

	err := s.repo.AddHistoryByOrderID(orderID, history)
	if err != nil {
		log.Println(err)
		fail(w, http.StatusInternalServerError, "")
		return
	}
	log.Println("handlePutLocationHistoryV1: hostory added to orderID: ", orderID)
	send(w, http.StatusOK, "success")
}

func (s *Server) handleDeleteLocationHistoryV1(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "order_id")
	err := s.repo.DeleteLocationHistoryByOrderID(orderID)
	if err == e.ErrResourceNotFound {
		log.Println("requested resource not found for oderID: ", orderID)
		fail(
			w,
			http.StatusNotFound,
			"requested resource not found",
		)
		return
	}
	if err != nil {
		log.Println(err)
		fail(w, http.StatusInternalServerError, "")
		return
	}
	send(w, http.StatusOK, "success")
}
