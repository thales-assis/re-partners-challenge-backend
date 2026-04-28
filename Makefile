tests:
	go test -p 4 -parallel 4 -count=1 ./...

wire:
	wire ./app

watch:
	air server --port 8080