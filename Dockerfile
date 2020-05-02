FROM golang:1.14

WORKDIR /go/src/github.com/voluntariado-ucc-ing/volunteer_api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["volunteer_api"]