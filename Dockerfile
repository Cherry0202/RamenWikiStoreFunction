FROM golang:latest
ENV GO111MODULE=on
WORKDIR /go/src/github.com/Cherry0202/RamenWikiStoreFunction
COPY . .
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ADD . /go/src/github.com/Cherry0202/RamenWikiStoreFunction
CMD ["go", "run", "server.go"]