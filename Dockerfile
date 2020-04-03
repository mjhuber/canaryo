FROM golang:1.14 AS build

WORKDIR /go/src/github.com/mjhuber/canaryo
ADD . /go/src/github.com/mjhuber/canaryo
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o app

FROM gcr.io/distroless/base
EXPOSE 8080
COPY --from=build /go/src/github.com/mjhuber/canaryo/app /app
CMD ["./app"]
