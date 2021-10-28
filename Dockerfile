FROM golang:alpine3.14 AS build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o api .

FROM alpine:3.14 AS bin
COPY --from=build /go/src/app/api /api
# creates empty config.yaml to prevent docker's
# "mounting file on directory" error
RUN touch config.yaml
# running with sudo access as serial 
# port requires super user access
RUN su -
CMD ["/api"]
