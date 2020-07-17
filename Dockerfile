FROM golang:latest
ENV GO111MODULE=on
WORKDIR /go/src/github.com/Cherry0202/RamenWikiStoreFunction
COPY . .
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
ENTRYPOINT ["/app"]


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