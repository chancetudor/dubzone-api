build:
	go build -o bin/main cmd/server/main.go

test:
	go test

run:
	go build -o bin/main cmd/server/main.go && ./bin/main
	
swagger:
	swagger generate spec -o ./swagger.yaml --scan-models