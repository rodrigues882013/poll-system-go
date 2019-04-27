FROM golang:alpine

ENV PROJECT_NAME vote-consumer-service
ENV GOPATH /app/

RUN mkdir /app
ADD . /app/./src/github.com/felipe_rodrigues/${PROJECT_NAME}

WORKDIR /app

RUN go build -o main ./src/github.com/felipe_rodrigues/${PROJECT_NAME}/cmd/${PROJECT_NAME}
RUN rm -rf /app/src

CMD ["/app/main"]