FROM golang:1.14

WORKDIR /go/src/github.com/voluntariado-ucc-ing/volunteer_api
COPY . .

RUN go get -d  -v ./...
RUN go install -v ./...

CMD ["volunteer_api"]

# Database Credentials
ENV DB_HOST=172.17.0.2:5432

ENV DB_USER=postgres

ENV DB_PASS=ysl*gzzjic4Taok

ENV DB_NAME=voluntariado_ing

EXPOSE 8080

# Commands for running in docker
# docker build -t volunteer_api .
# docker run -it --rm -p 8080:8080 --name api  volunteer_api