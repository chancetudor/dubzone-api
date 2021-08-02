build:
	go build -o bin/main cmd/api/main.go

#test:

run:
	go build -o bin/main cmd/api/main.go && ./bin/main