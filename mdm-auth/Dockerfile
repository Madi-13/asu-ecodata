FROM reg.appleseed.kz/docker-proxy/library/golang:1.18.1-alpine as builder

WORKDIR /src

COPY ./ .

RUN GOOS=linux GOARCH=386 go mod vendor && go build -v -o ./build/run cmd/web/main.go

FROM alpine
WORKDIR /app

COPY --from=builder /src/build/run .

EXPOSE 9090

ENTRYPOINT ["/build/run"]
