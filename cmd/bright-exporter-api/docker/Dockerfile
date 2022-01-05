FROM golang:1.17-bullseye as build

WORKDIR /go/src/github.com/rk295/bright-golang/cmd/bright-exporter-api
ADD . /go/src/github.com/rk295/bright-golang/cmd/bright-exporter-api

RUN go get -d -v ./...

RUN go build -o /go/bin/bright-exporter-api

# FROM gcr.io/distroless/base-debian10
FROM golang:1.17-bullseye
COPY --from=build /go/bin/bright-exporter-api /
CMD ["/bright-exporter-api"]
