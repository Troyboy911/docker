# ---- builder ----
FROM golang:1.21 as builder
WORKDIR /src
# Use modules
COPY go.mod ./
RUN go mod download
COPY . .
# Build a static binary, with minimal symbols
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /out/hardened-health ./main.go

# ---- runtime ----
FROM gcr.io/distroless/static:nonroot
# Nonroot image has no shell; exposes port and runs user 'nonroot'
COPY --from=builder /out/hardened-health /usr/local/bin/hardened-health
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/hardened-health"]
