FROM golang:alpine

RUN mkdir /src

ADD . /src/

WORKDIR /src/app

RUN go build -o main .

RUN adduser -S -D -H -h /app appuser

USER appuser

CMD ["./main"]