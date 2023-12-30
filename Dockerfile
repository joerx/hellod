FROM golang:1.21-alpine3.18 AS builder

ADD ./ /go/src/github.com/joerx/hellod
WORKDIR /go/src/github.com/joerx/hellod
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hellod .

FROM alpine:3.18
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/joerx/hellod/hellod /usr/local/bin
ENTRYPOINT ["hellod"]
