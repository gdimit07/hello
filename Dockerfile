FROM golang:1.23
WORKDIR /hello
RUN mkdir customer
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY customer/id.go ./customer
RUN go build -o customerctl
CMD ["./customerctl"]
