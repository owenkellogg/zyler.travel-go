FROM golang:1.18

RUN mkdir -p /go/src/github.com/zylerdj/zyler.travel-go/public/uploads

RUN mkdir -p /var/zyler.travel/public/uploads

WORKDIR /go/src/github.com/zylerdj/zyler.travel-go

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o build/app github.com/zylerdj/zyler.travel-go

EXPOSE 5200 5200

ENTRYPOINT ["build/app"]
