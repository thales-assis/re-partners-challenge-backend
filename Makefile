tests:
	go test -p 4 -parallel 4 -count=1 ./...

wire:
	wire ./app

swagger:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go -o docs

watch:
	air server --port 8080

run:
	docker-compose build && docker-compose up