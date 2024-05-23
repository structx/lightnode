
FROM golang:1.22-bookworm as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /usr/bin/server ./cmd/server

FROM gcr.io/distroless/static-debian12

COPY --from=builder /usr/bin/ /app/bin/

VOLUME [ "/var/log/lightnode", "/local/lightnode", "/opt/lightnode/raft", "/opt/lightnode/data" ]

EXPOSE 8080 50051

ENTRYPOINT [ "/app/bin/server" ]

