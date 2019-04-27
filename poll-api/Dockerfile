FROM golang:alpine

ENV PROJECT_NAME poll-api
ENV GOPATH /app/

RUN mkdir /app
ADD . /app/src/github.com/felipe_rodrigues/${PROJECT_NAME}

WORKDIR /app

RUN go build -o main ./src/github.com/felipe_rodrigues/${PROJECT_NAME}/cmd/${PROJECT_NAME}
RUN rm -rf /app/src

EXPOSE 3000

CMD ["/app/main"]
