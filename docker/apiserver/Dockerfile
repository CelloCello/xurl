FROM golang:1.18.6-alpine3.16 AS builder

COPY ./apiserver/ /app

WORKDIR /app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apiserver .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o apiserver .


FROM alpine:3.16

# RUN apk add --no-cache libc6-compat

COPY --from=builder /app/apiserver /app/apiserver
COPY ./apiserver/views/templates/. /app/views/templates/
COPY ./docker/apiserver/*.sh /
RUN chmod +x /*.sh
RUN chmod +x /app/apiserver

WORKDIR /app

ENTRYPOINT ["/entrypoint.sh"]
