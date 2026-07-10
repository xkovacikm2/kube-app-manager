FROM golang:1.23-bookworm AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags='-s -w' -o /out/kube-app-manager ./cmd/kube-app-manager

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /
COPY --from=builder /out/kube-app-manager /kube-app-manager

EXPOSE 8080

ENTRYPOINT ["/kube-app-manager"]
