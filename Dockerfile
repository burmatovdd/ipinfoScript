FROM golang:latest

RUN mkdir -p /go/src/ipinfoScript


WORKDIR /ipinfoScript

RUN go version

#COPY cmd/server ./
COPY go.mod go.sum ./
COPY *.go ./

RUN go mod download
RUN go build -o ipinfoScript .
EXPOSE 8080
CMD ["./ipinfoScript"]