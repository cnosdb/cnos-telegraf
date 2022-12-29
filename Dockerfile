FROM golang:latest as builder
WORKDIR /go/src/github.com/cnosdb/telegraf
COPY . /go/src/github.com/cnosdb/telegraf
RUN go env -w GOPROXY="https://goproxy.cn"
RUN go build -o telegraf ./cmd/telegraf

FROM ubuntu:latest
COPY --from=builder /go/src/github.com/cnosdb/telegraf/telegraf /usr/bin/
RUN mkdir -p /etc/telegraf/
RUN /usr/bin/telegraf config > /etc/telegraf/telegraf.conf
CMD ["telegraf"]
