FROM golang:1.14.2 as builder

WORKDIR /go/src/
COPY . .
RUN go test -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main basic-rest

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ENV PORT 9002
COPY --from=builder /go/src/main ./basic-rest
CMD ["./basic-rest"]