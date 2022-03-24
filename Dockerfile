# This Dockerfile is optimized to
# occupy as low disk space as possible.
#
# For the optimization by building speed
# see my article on Medium:
# https://medium.com/@maximgradan/microservices-with-go-modules-9fa1e82c35b8
FROM golang:1.17.8-alpine3.15 AS builder
COPY . /go/src/github.com/zergon321/task
RUN cd /go/src/github.com/zergon321/task && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/task /go/src/github.com/zergon321/task

FROM alpine:3.14.4
COPY --from=builder /go/bin/task /bin/task
ENTRYPOINT [ "/bin/task" ]