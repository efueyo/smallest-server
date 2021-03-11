FROM golang:alpine AS builder
RUN apk update && apk add --no-cache alpine-sdk

WORKDIR /src
ADD . .
RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s"  -o ./bin/server

FROM scratch as production
ENV PATH=/bin
COPY --from=builder /src/bin/server /src/bin/server
ENTRYPOINT ["/src/bin/server"]
