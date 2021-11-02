# base image that we are using 
FROM golang:1.17-alpine

# create a directory inside a image
# this instructs docker to use this 
# dir as default directory
WORKDIR /app

# copy dependencies and download them 
COPY go.mod ./
COPY go.sum ./

RUN go mod download 

COPY . .