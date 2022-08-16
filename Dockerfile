FROM golang:1.14.6
RUN mkdir -p /go/src/github.com/zylerdj/zyler.travel-go
COPY go.mod go.sum /go/src/github.com/zylerdj/zyler.travel-go
WORKDIR /go/src/github.com/zylerdj/zyler.travel-go
RUN go mod download
COPY . /go/src/github.com/zylerdj/zyler.travel-go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/zyler.travel github.com/zylerdj/zyler.travel-go

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/zyler.travel"]
