FROM golang:1.8.3

WORKDIR /go/src/github.com/denkhaus/microservices/greeter
COPY . .

# Run a gofmt and exclude all vendored code.
RUN test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*"))" || { echo "Run \"gofmt -s -w\" on your Golang code"; exit 1; }

RUN VERSION=$(git describe --all --exact-match `git rev-parse HEAD` | grep tags | sed 's/tags\///') \
 && GIT_COMMIT=$(git rev-list -1 HEAD) \
 && go get -v . \
 && CGO_ENABLED=0 GOOS=linux go build --ldflags "-s -w -X github.com/denkhaus/microservices/greeter/main.GitCommit=${GIT_COMMIT} -X github.com/denkhaus/microservices/greeter/main.Version=${VERSION}" -a -installsuffix cgo -o service . \
 && go test $(go list ./... | grep -v /vendor/ | grep -v /template/) -cover


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /go/src/github.com/denkhaus/microservices/greeter/service .

ENTRYPOINT ["./service"]
