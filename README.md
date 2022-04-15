# Rss reader service

- Http service that parse multiple RSS feed urls and return feed items
- The parser can be accessed on {localhost}:{port}/parse. JSON request in {["url1, "url2"]} format

### Runing:
- Building and running:
    `go build cmd/reader/main.go` or
    `go run cmd/reader/main.go`
- Docker:
    `docker build .`

You can configure the PORT on which the server is using by setting PORT environment variable. Default value is 9000