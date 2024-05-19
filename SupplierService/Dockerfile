FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY .. .

RUN go build -ldflags="-s -w" -o /app/grpc_server grpcServer/server.go

RUN go build -ldflags="-s -w" -o /app/http_server httpServer/server.go


FROM scratch

WORKDIR /app

COPY --from=builder /app/grpc_server /app/grpc_server
COPY --from=builder /app/http_server /app/http_server

EXPOSE 50053
EXPOSE 8082

CMD ["./grpc_server"]

CMD ["./http_server"]
