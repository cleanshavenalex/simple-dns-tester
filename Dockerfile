FROM golang:1.7.3
WORKDIR /go/src/github.com/cleanshavenalex/simple-dns-tester
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o dns-tester ./main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root

COPY --from=0 /go/src/github.com/cleanshavenalex/simple-dns-tester .
CMD ["./dns-tester"]  
