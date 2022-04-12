FROM golang:latest AS builder

RUN apt-get update -y &&\
    apt-get upgrade -y &&\
    apt-get install protobuf-compiler -y &&\
    apt-get install golang-goprotobuf-dev -y
    
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /go/src/workdir/src

COPY ./src /go/src/workdir/src

RUN go mod download

RUN go build -o main ./

WORKDIR /dist

RUN cp /go/src/workdir/src/main .

FROM scratch

COPY --from=builder /dist/main .

ENTRYPOINT ["/main"]