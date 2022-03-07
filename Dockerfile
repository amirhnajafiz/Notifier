# build from ubuntu image
FROM golang:alpine

# switch to root
USER root

# main working directory
WORKDIR /app

# copy all file from current directory
COPY go.mod go.sum ./

# copy all of the files
COPY . ./

# building our file
RUN go build -o /notifier

# running execute file
CMD ["/notifier"]
