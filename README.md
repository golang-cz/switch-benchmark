# Benchmark comparison between Golang switch-case statements vs. string hashmap for an input string 

## Rationale
We're code generating REST API server via [gospeak](https://github.com/golang-cz/gospeak) or
[webrpc](https://github.com/webrpc/webrpc). We're interested in making the generated Go HTTP router code
to match incoming requests to the corresponding handlers as fast as possible. Given that all
webrpc routes are code-generated and known ahead of time, we can choose between generating
`map[string]handler` or a code with a bunch of switch-case statements. What's faster?

## Benchmark
See the full [benchmark file](./benchmark_test.go).

### Switch case on string
```go
switch r.URL.Path {
case "/rpc/ExampleService/Ping":
    pingHandlerJSON()
case "/rpc/ExampleService/Status":
    statusHandlerPING()
default:
    notFoundHandler()
}
```

### vs. map[string]handler
```go
routes := map[string]handler{
    "/rpc/ExampleService/Ping": pingHandlerJSON,
    "/rpc/ExampleService/Status": statusHandlerPING,
}

handler, ok := routes[r.URL.Path]
if !ok {
    notFoundHandler()
}
handler()
```

## RESULT: The winner is the code-generated SWITCH
Turns out, the switch case statements are a little faster compared to a string hashmap, thanks to built-in
`switch string` binary search algorithm.

```
$ go version
go version go1.21.0 darwin/arm64

$ go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/golang-cz/switch-benchmark
BenchmarkStringSliceRange-8   	  532791	      2254 ns/op
BenchmarkStringMap-8          	 2152723	       518.1 ns/op
BenchmarkStringSwitch-8       	 2780476	       433.3 ns/op
PASS
ok  	github.com/golang-cz/switch-benchmark	5.386s
```

# License
[MIT License](./LICENSE)
