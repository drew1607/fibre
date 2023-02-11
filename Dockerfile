FROM golang:1.20.0-alpine3.17

RUN mkdir /app
ADD . /app

WORKDIR /app
RUN make run .
CMD ["/app/main"]