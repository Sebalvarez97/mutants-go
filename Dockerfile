FROM golang:1.15
RUN mkdir /app
ADD /api /app
WORKDIR /app
RUN go build -o server .
CMD ["/app/server"]