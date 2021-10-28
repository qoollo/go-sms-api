# Go SMS API
Qoollo SMS service with GRPC and HTTP API on Golang

## Usage
### Preparations
Edit `config.yaml`

### Installation
```bash
$ go build
```

### Running
```bash
$ sudo ./go-sms-api # sudo is required
```

### Running using Docker
Build an image
```bash
$ docker build -t smsapi .
```

..and run it mounting config file and /dev folder as volumes to access serial ports from the docker container
```bash
$ docker run -v ${PWD}/config.yaml:/config.yaml -v /dev:/dev --privileged -p 3201:3201 -p 3200:3200 smsapi
```

## Development
### Dependencies
```bash
$ export GO111MODULE=off

$ go get github.com/grpc-ecosystem/grpc-gateway

$ go get github.com/envoyproxy/protoc-gen-validate

$ cd ~/go/src/github.com/grpc-ecosystem/grpc-gateway
$ mkdir third_party && cd third_party
$ git clone https://github.com/googleapis/googleapis.git
```