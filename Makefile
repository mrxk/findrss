
all: build

build:
	dev-env run -e pi -- CGO_ENABLED=0 go build -a -v

docker:
	docker build . -t findrss

install:
	docker run -d --restart unless-stopped --name findrss -p 3006:80 findrss

