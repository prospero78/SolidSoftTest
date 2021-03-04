run.debug:
	clear
	go build -race -o ./bin/solidtest ./cmd/solidtest/main.go
	cd ./bin && \
	./solidtest debug run
run.help:
	clear
	go build -o ./bin/solidtest ./cmd/solidtest/main.go
	cd ./bin && \
	./solidtest
run:
	clear
	go build -o ./bin/solidtest ./cmd/solidtest/main.go
	cd ./bin && \
	./solidtest run
fmt:
	clear
	go fmt ./...
lint:
	clear
	golangci-lint run ./...
	gocritic check
	gocyclo -top 7 ./cmd ./internal
