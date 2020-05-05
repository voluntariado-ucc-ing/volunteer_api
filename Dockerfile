FROM golang:1.14

WORKDIR /go/src/github.com/voluntariado-ucc-ing/volunteer_api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["volunteer_api"]

# Database Credentials
ENV DB_HOST=ec2-52-87-135-240.compute-1.amazonaws.com

ENV DB_USER=rumxsiovwviqfx

ENV DB_PASS=096f56199709e39cf83d39a209d46657d35047b97b8f8c173028830b5a9fe207

ENV DB_NAME=d9fuu92algrdm0

EXPOSE 8080

# Commands for running in docker
# docker build -t volunteer_api .
# docker run -it --rm -p 8080:8080 --name api  volunteer_api