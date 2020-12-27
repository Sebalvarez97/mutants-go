FROM golang:1.15
RUN mkdir /app
ADD src /app
WORKDIR /app
RUN go build -o server .
CMD ["/app/server"]