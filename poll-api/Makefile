tests:
	go test -v ./...
build:
	cd cmd/pool-api/ && go build
run:
	PORT=${PORT} \
    	MAX_CONN=${MAX_CONN} \
    	REQUEST_TIMEOUT=${REQUEST_TIMEOUT} \
    	MAX_IDLE_CONN_PER_HOST=${MAX_IDLE_CONN_PER_HOST}
	./cmd/poll-api/poll-api
