GOPATH:=$(shell go env GOPATH)
APP_NAME:=service-ecom

.PHONY: build test docker

build-order: 
	CGO_ENABLED=0 go build -a -o order services/orderservice/cmd/server/main.go

build-payment: 
	CGO_ENABLED=0 go build -a -o payment services/paymentservice/cmd/server/main.go

docker:
	docker build -t order-service -f Dockerfile.order .
	docker build -t payment-service -f Dockerfile.payment .

# running docker image
run-order-service:
	docker run -p 4000:4000 order-service

run-payment-service:
	docker run -p 4001:4001 payment-service