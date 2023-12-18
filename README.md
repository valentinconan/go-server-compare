# compare server frameworks

## Build
Install all dependencies `go mod tidy`

Build project `go build main.go`

Run `go run main.go`

## Test
Using bombardier in order to test the different frameworks

Install bombardier ` go install github.com/codesenberg/bombardier@latest`

Then run on different ports, with different connections: 
> bombardier -c 50 http://localhost:8080/health
> 
> bombardier -c 200 http://localhost:8081/health
> 
> bombardier -c 1000 http://localhost:8082/health