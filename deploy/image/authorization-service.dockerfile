FROM alpine:latest

RUN mkdir /app

COPY ../../bin/authorizationApp /app

CMD ["/app/authorizationApp"]