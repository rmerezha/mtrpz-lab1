FROM golang:1.23.5 AS build

WORKDIR /usr/app

ENV CGO_ENABLED=0

COPY . .

RUN go build -o solver

FROM alpine:3.21.1

COPY --from=build /usr/app/solver .

ENTRYPOINT ["./solver"]
