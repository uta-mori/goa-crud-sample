FROM golang:1.15.3-buster AS build
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -mod=readonly -v -o book ./cmd/book

# k8s、cloud-run用ビルド --target deploy
FROM alpine:3 AS deploy
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN apk add --no-cache ca-certificates
COPY --from=build /app/book /book
CMD ["/book"]
