FROM golang:1.11.4-alpine3.8

RUN mkdir -p /go/src/tutor_platform
RUN mkdir -p /var/log/tutor_platform/
WORKDIR /go/src/tutor_platform
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tutor_platform .

FROM alpine:3.8
RUN mkdir -p /go/src/tutor_platform
WORKDIR /go/src/tutor_platform
COPY static /go/src/tutor_platform/static
COPY views /go/src/tutor_platform/views
COPY --from=0 /go/src/tutor_platform/tutor_platform .
COPY --from=0 /go/src/tutor_platform/localtime /etc/localtime

EXPOSE 8080

CMD exec ./tutor_platform >>/var/log/tutor_platform/web.log 2>>/var/log/tutor_platform/web.log