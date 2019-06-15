FROM golang:1.12
WORKDIR /go/src/app/
COPY . .
ENV GO111MODULE=on
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build cmd/go-watch-s3/go-watch-s3.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/app/go-watch-s3 ./go-watch-s3
CMD ["./go-watch-s3"]
