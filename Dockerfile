FROM golang:1.14

WORKDIR /go/src/github.com/voluntariado-ucc-ing/volunteer_api
COPY . .

RUN go install -v ./...

CMD ["volunteer_api"]

# Database Credentials


EXPOSE 8080
EXPOSE 587

# Commands for running in docker
# docker build -t volunteer_api .
# docker run -it --rm -p 8080:8080 --name api  volunteer_api
