dev:
	air -c .air.toml

build:
	go build -o ./tmp/main .

run:
	go run .