# Location History Server

- Simple backend server for location history server

## Technologies used.

- Go - 1.12

## How to run ?

- go build
- HISTORY_SERVER_LISTEN_ADDR=8080 ./location-history-server

## Sample CURL APIs

### PUT

```
curl --location --request PUT 'http://localhost:8080/location/def456' \
--header 'Content-Type: application/json' \
--data-raw '{
	"lat": 19.49,
	"lng": 56.79
}'
```

### GET

```
curl --location --request GET 'http://localhost:8080/location/def456?max=2'
```

### DELETE

```
curl --location --request DELETE 'http://localhost:8080/location/def456'
```