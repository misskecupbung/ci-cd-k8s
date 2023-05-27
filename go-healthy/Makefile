NAME=healthy
TAG=dieple/go-$(NAME)
VER=v1.0.0

all: clean build run

build:
	go get github.com/gorilla/mux
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o healthy
	docker build -t $(TAG) -t $(TAG):$(VER) .

run:
	docker run -d -p 8000:8000 -e PORT=8000 --name=$(NAME) $(TAG)

clean:
	-docker stop $(NAME)
	-docker rm $(NAME)

push:
	docker push $(TAG)
	docker push $(TAG):$(VER)
