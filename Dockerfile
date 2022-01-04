FROM alpine:latest as builder
RUN apk add ca-certificates

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./findrss /findrss
COPY ./index.html /index.html
ENTRYPOINT ["/findrss"]
