FROM golang:1.20.2-alpine as builder

WORKDIR /src

COPY app/. .

RUN GOOS=linux GOARCH=386 go build -v -o ./build/run main.go

FROM alpine
RUN apk add --no-cache tzdata
ENV TZ="Asia/Almaty"

WORKDIR /app

COPY --from=builder /src/build/run .
#COPY --from=builder /src/config/. config/
#
#EXPOSE 8001
#
#CMD ["./run"]
