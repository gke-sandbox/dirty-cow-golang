FROM golang:1.10.3
RUN apt-get update && apt-get upgrade -y
WORKDIR /go/src/app
COPY . .
RUN go build -o dirtyc0w .
CMD ["./dirtyc0w"]
