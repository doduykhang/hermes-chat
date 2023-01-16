run:
	go run cmd/main/main.go

build:
	go build -o build/app cmd/main/main.go

docker-build:
	sudo docker build --no-cache -t hermes-char .
