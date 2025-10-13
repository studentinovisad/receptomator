FROM docker.io/library/golang:1.25.2-trixie AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o receptomator .

FROM scratch
COPY --from=builder /app/receptomator /receptomator
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER 65532
ENTRYPOINT ["/receptomator"]
EXPOSE 8080

LABEL org.opencontainers.image.title="receptomator"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="https://github.com/studentinovisad/receptomator"
LABEL org.opencontainers.image.source="https://github.com/studentinovisad/receptomator"
LABEL org.opencontainers.image.documentation="https://github.com/studentinovisad/receptomator#readme"
LABEL org.opencontainers.image.vendor="studentinovisad"
