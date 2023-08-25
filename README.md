# Benchmark comparison between Golang switch-case statements vs. string hashmap for an input string 

## Rationale
We're code generating REST API server via [gospeak](https://github.com/golang-cz/gospeak) or
[webrpc](https://github.com/webrpc/webrpc). We're interested in making the generated Go HTTP router code
to match incoming requests to the corresponding handlers as fast as possible. Given that all
webrpc routes are code-generated and known ahead of time, we can choose between generating
`map[string]handler` or a code with a bunch of switch-case statements. What's faster?

## Benchmark
See the full [benchmark file](./benchmark_test.go).

See some [shorter benchmarks in stdlib](https://github.com/golang/go/blob/ddad9b618cce0ed91d66f0470ddb3e12cfd7eeac/src/cmd/compile/internal/test/switch_test.go#L86).

### Switch case on string
```go
switch r.URL.Path {
case "/rpc/ExampleService/Ping":
    pingHandler()
case "/rpc/ExampleService/Status":
    statusHandler()
default:
    notFoundHandler()
}
```

### vs. map[string]handler
```go
routes := map[string]handler{
	"/rpc/ExampleService/Ping": pingHandler,
	"/rpc/ExampleService/Status": statusHandler,
}

handler, ok := routes[r.URL.Path]
if !ok {
	notFoundHandler()
	return
}
handler()
```

## RESULT: The winner is the code-generated SWITCH
Turns out, the switch-case statements are a bit faster compared to a string hashmap in Go 1.20+.

```
$ go version
go version go1.21.0 linux/amd64

$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/golang-cz/switch-case-benchmark
cpu: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
BenchmarkStringSliceRange-2   	  263060	      4608 ns/op
BenchmarkStringMap-2          	  904586	      1469 ns/op
BenchmarkStringSwitch-2       	 1382715	       856.8 ns/op
PASS
ok  	github.com/golang-cz/switch-case-benchmark	4.676s
```

```
$ go version
go version go1.21.0 darwin/arm64

$ go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/golang-cz/switch-case-benchmark
BenchmarkStringSliceRange-8   	  532791	      2254 ns/op
BenchmarkStringMap-8          	 2152723	       518.1 ns/op
BenchmarkStringSwitch-8       	 2780476	       433.3 ns/op
PASS
ok  	github.com/golang-cz/switch-case-benchmark	5.386s
```

## More details..
Go 1.19 and older used to use two-level binary search for switch-case statements matching 2+
string values. The algorithm sorted strings by length and then by value, as it is much cheaper
to compare lengths than values.

In Go 1.20, Keith Randall implemented a jump table for comparing strings that have at least
two different lengths to improve the performance even further. This effectively mean that the
compiler generates two switches, the outer for string length and the inner for string values.
See [cmd/compile: modify switches of strings to use jump table for lengths](https://go-review.googlesource.com/c/go/+/395714)
https://github.com/golang/go/issues/34381).

In Go 1.21, Keith Randall went even further and submitted [cmd/compile: implement jump tables](https://go-review.googlesource.com/c/go/+/357330),
with a great description of why the new jump tables implementation are likely to behave better
than binary search algorithm given the less consumption of branch predictor resources. Keith
claims that predictable switch microbenchmarks won't see the benefit, but a real program would.
On top of that, he mentions that [cmd/compile: feedback-guided optimization](https://github.com/golang/go/issues/28262)
might help increase the performance even further.

There is some more great discussion at Go issue [cmd/compile: use hashing for switch statements](https://github.com/golang/go/issues/34381)
submitted by Matthew Dempsky. The proposal is to use a minimal perfect hash to improve the switch
case performance. There are some proof-of-concepts, ie. [cmd/compile: hash strings in walk/switch.go](https://go-review.googlesource.com/c/go/+/307191),
that didn't make it to the stdlib as of Aug 25 2023. Let's keep looking at changes made in https://github.com/golang/go/blob/master/src/cmd/compile/internal/walk/switch.go.

Anyway, the above progress and recent discussions make me believe that switch-case on strings
will always perform better than string hashmap in the future releases of Go. 

# License
[MIT License](./LICENSE)
