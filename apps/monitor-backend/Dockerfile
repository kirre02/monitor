# Build stage
FROM golang:1.20-alpine3.19 as Build-stage

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPATH=/monitor-backend

WORKDIR /build

COPY . .

RUN go mod download

RUN go build ./cmd/monitor-backend/  

# Final stage
FROM alpine:3.19.0

WORKDIR /api/

COPY --from=Build-stage /build/monitor-backend ./bin/monitor-backend
COPY --from=Build-stage /build/.env.dev /api/

EXPOSE 9090

CMD ["./bin/monitor-backend"]
