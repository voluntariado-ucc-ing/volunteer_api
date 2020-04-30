FROM golang:1.14

WORKDIR /go/src/voluntarios_api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["voluntarios_api"]