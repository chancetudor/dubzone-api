build:
	go build -o bin/main cmd/server/main.go

test:
	go test

run:
	go build -o bin/main cmd/server/main.go && ./bin/main
	
swagger:
	swagger generate spec -o ./openapi/swagger.yaml --scan-models
	
clean:
	rm -r bin