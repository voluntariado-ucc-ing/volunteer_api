FROM golang:1.14

WORKDIR /go/src/github.com/voluntariado-ucc-ing/voluntarios_api/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["voluntarios_api"]