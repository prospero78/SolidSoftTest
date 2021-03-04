run.app:
	go build -o ./bin/solidtest ./cmd/solidtest/main.go
	cd ./bin && \
	./solidtest