run:
	ENV=DEV go run cmd/main/main.go

clean:
	rm -rf build

build:
	go build -o build/app cmd/main/main.go

run-prod:
	ENV=PROD ./build/app

docker-build:
	sudo docker build --no-cache -t hermes-chat .
