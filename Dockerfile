FROM golang:1.11

RUN apt-get update

ENV PORT=8000

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
