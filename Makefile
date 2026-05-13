port = 6380

run:
	go run ./cmd/main.go --port $(port)

build:
	go build -o ./out/kv-store ./cmd/main.go
