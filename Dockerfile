FROM golang:1.15
# create a working directory
WORKDIR /go/src/app
# add source code
ADD src/app app
# run main.go
CMD ["go", "run", "app/server.go"]