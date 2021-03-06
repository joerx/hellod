FROM golang:alpine AS builder

ADD ./ /go/src/github.com/joerx/hellod
WORKDIR /go/src/github.com/joerx/hellod
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hellod .

FROM alpine:latest  
ENV ADDRESS=:9000
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/joerx/hellod/hellod .
CMD ["./hellod"]
