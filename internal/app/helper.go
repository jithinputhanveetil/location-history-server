package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	portEnv         = "HISTORY_SERVER_LISTEN_ADDR"
	defaultPort     = ":8080"
	useInMemory     = "USE_IN_MEMORY"
	defaultInMemory = true
)

func mustReadPort() string {
	port := os.Getenv(portEnv)
	if len(port) == 0 {
		log.Printf("HISTORY_SERVER_LISTEN_ADDR config is empty, using the default value %s", defaultPort)
		port = defaultPort
	}

	return fixPrefix(port, ":")
}

func fixPrefix(val, prefix string) string {
	if !strings.HasPrefix(val, ":") {
		val = ":" + val
	}

	return val
}

func mustReadInMemoryMode() bool {
	var inMemory bool
	m := os.Getenv(useInMemory)
	if len(m) == 0 {
		log.Printf("USE_IN_MEMORY config is empty, using the default value %s", defaultInMemory)
		inMemory = defaultInMemory
	} else {
		modeVal, err := strconv.ParseBool(m)
		if err != nil {
			inMemory = defaultInMemory
			log.Printf("unsupported value for USE_IN_MEMORY config: %s, using the default value: %t", m, inMemory)
		} else {
			inMemory = modeVal
			log.Printf("USE_IN_MEMORY: ", inMemory)
		}
	}
	return inMemory
}

const (
	statusOK   = "ok"
	statusFail = "nok"
)

// Response implements standard JSON response payload structure.
type Response struct {
	Status     string          `json:"status"`
	Result     json.RawMessage `json:"result,omitempty"`
	Error      *ResponseError  `json:"error,omitempty"`
	Pagination json.RawMessage `json:"pagination,omitempty"`
}

// ResponseError implements the standard FR Error response structure.
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// send sends a successful JSON response using
// the standard success format
func send(w http.ResponseWriter, status int, result interface{}) {
	rj, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	r := &Response{
		Status: statusOK,
		Result: rj,
	}

	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

// fail ends an unsuccessful JSON response with the standard failure format.
// fail overrides gep Fail and allows to send custom error with the response.
func fail(w http.ResponseWriter, status int, msg string) {
	r := &Response{
		Status: statusFail,
		Error: &ResponseError{
			Code:    status,
			Message: msg,
		},
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}
