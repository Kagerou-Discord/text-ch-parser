FROM golang:1.17.6 AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .
RUN CGO_ENABLED=0 go build -o ./out/parser .

FROM gcr.io/distroless/static
USER nonroot
WORKDIR /app
COPY --from=builder --chown=nonroot:nonroot /app/out/parser .

CMD ["/app/parser"]
