FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /docker-user-service

EXPOSE 8888

CMD [ "/docker-user-service" ]