FROM golang:1.20-alpine3.17 AS build-env
WORKDIR /app
COPY . .
RUN go build -o skyauc cmd/auctions.go

FROM alpine:3.17 AS prod
WORKDIR /app
COPY --from=build-env /app/skyauc .
RUN ls
CMD ["./skyauc"]