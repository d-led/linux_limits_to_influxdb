# FROM golang:1.9-alpine3.7

ARG golang_tag=1.9-alpine3.7
FROM golang:$golang_tag

WORKDIR /go/src/llti

COPY . .

RUN go get -d -v ./...
RUN go build
RUN go test ./...

CMD ["./llti"]
