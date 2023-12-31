FROM golang:1.20

EXPOSE 3000
EXPOSE 6000

WORKDIR /go/src/app
COPY . .

ARG GO111MODULE=off
RUN go get -d -v ./...
RUN go build -v -o app ./...
RUN mv ./app /go/bin/

CMD ["app"]
