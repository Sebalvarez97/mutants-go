FROM golang:1.15
RUN mkdir /app
ADD . /app
WORKDIR /app/src
RUN go build -o server .
CMD ["/app/src/server"]