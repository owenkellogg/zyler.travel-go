FROM golang:1.18

RUN mkdir -p /go/src/github.com/zylerdj/zyler.travel-go

WORKDIR /go/src/github.com/zylerdj/zyler.travel-go

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o build/zyler.travel github.com/zylerdj/zyler.travel-go

RUN cp build/zyler.travel /usr/bin/zyler.travel

EXPOSE 5200 5200

ENTRYPOINT ["/usr/bin/zyler.travel"]
