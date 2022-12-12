FROM golang:alpine
LABEL maintainer="sivops"

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]