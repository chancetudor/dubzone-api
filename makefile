build:
	go build -o bin/main cmd/server/main.go

#test:

run:
	go build -o bin/main cmd/server/main.go && ./bin/main