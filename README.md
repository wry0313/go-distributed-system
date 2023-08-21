# Go Distributed Systems
This is a miniature, Go-based distributed system that supports the microservices architecture. It provides features such as service registration, discovery, dependency updates, and a basic load balancing mechanism. You can integrate your own microservice by adding it to the cmd directory. This distributed system combines elements of both the hub-and-spoke model and the peer-to-peer model.
## Prerequisites
Go v1.21 or higher

Docker
## Getting Started
Start the micro services with Docker
``` bash
docker-compose up
```
## Using Package Service API Endpoints
Use the API endpoints provided by the `package service` to create your own microservice. 

For example, to register a service with the `registry`:
``` golang
host, port := "yourservicename", "4000"
serviceAddr := fmt.Sprintf("http://%s:%s", host, port)
r := registry.Registration{
	ServiceName:      registry.LogService,
	ServiceURL:       serviceAddr,
	RequiredServices: make([]registry.ServiceName, 0),
	ServiceUpdateURL: serviceAddr + "/services",
	HeartbeatURL:     serviceAddr + "/heartbeat",
}
ctx, err := service.Start(
	context.Background(),
	host,
	port,
	r,
	log.RegisterHandlers,
)
```
Afterward, create the corresponding Dockerfile and docker-compose.yml files for the microservice, using formats similar to the ones presented here:
### Dockerfile.yourservicename
``` Dockerfile
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o yourservicename ./cmd/yourservicename

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/yourservicename .

EXPOSE portnumber

CMD ["./yourservicename"]
```
### docker-compose.yml
``` yml

  yourservicename:
    build:
      context: .
      dockerfile: Dockerfile.yourservicename
    ports:
      - "portnumber:portnumber"
    stdin_open: true
    tty: true
    depends_on:
      - registryservice
      - logservice
      - ...
```
