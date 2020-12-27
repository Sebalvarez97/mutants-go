FROM golang:1.15
# create a working directory
WORKDIR /go/src/app
# add source code
ADD src src
# run main.go
CMD ["go", "run", "src/server.go"]