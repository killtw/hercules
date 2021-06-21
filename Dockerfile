FROM golang:alpine AS build

ADD . /src

RUN apk -U add git gcc musl-dev && \
    cd /src && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app

FROM alpine

RUN apk -U add tzdata && \
    cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime
COPY --from=build /src/app /
COPY --from=build /src/public /public
COPY --from=build /src/views /views

EXPOSE 8080

ENTRYPOINT ["/app"]
