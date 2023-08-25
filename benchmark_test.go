package switch_benchmark

import "testing"

func BenchmarkStringSliceRange(b *testing.B) {
	var match int
	for n := 0; n < b.N; n++ {
		for _, v := range requestPaths {
			for _, route := range knownRoutes {
				if v == route {
					match++
					break
				}
			}
		}
	}
}

func BenchmarkStringMap(b *testing.B) {
	routeMap := map[string]struct{}{}
	for _, route := range knownRoutes {
		routeMap[route] = struct{}{}
	}

	var match int
	for n := 0; n < b.N; n++ {
		for _, v := range requestPaths {
			_, ok := routeMap[v]
			if ok {
				match++
			}
		}
	}
}

func BenchmarkStringSwitch(b *testing.B) {
	var match int
	for n := 0; n < b.N; n++ {
		for _, v := range requestPaths {
			switch v {
			case "/rpc/ExampleService/Ping":
				match++
			case "/rpc/ExampleService/Status":
				match++
			case "/rpc/ExampleService/Version":
				match++
			case "/rpc/ExampleService/GetUser":
				match++
			case "/rpc/ExampleService/FindUser":
				match++
			case "/rpc/ExampleService/ListUsers":
				match++
			case "/rpc/ExampleService/CreateUser":
				match++
			case "/rpc/ExampleService/UpdateUser":
				match++
			case "/rpc/ExampleService/DeleteUser":
				match++
			case "/rpc/AnotherService/Ping":
				match++
			case "/rpc/AnotherService/Status":
				match++
			case "/rpc/AnotherService/Version":
				match++
			case "/rpc/AnotherService/GetUser":
				match++
			case "/rpc/AnotherService/FindUser":
				match++
			case "/rpc/AnotherService/ListUsers":
				match++
			case "/rpc/AnotherService/CreateUser":
				match++
			case "/rpc/AnotherService/UpdateUser":
				match++
			case "/rpc/AnotherService/DeleteUser":
				match++
			case "/rpc/OtherService/Ping":
				match++
			case "/rpc/OtherService/Status":
				match++
			case "/rpc/OtherService/Version":
				match++
			case "/rpc/OtherService/GetUser":
				match++
			case "/rpc/OtherService/FindUser":
				match++
			case "/rpc/OtherService/ListUsers":
				match++
			case "/rpc/OtherService/CreateUser":
				match++
			case "/rpc/OtherService/UpdateUser":
				match++
			case "/rpc/OtherService/DeleteUser":
				match++
			case "/rpc/GetterService/GetA":
				match++
			case "/rpc/GetterService/GetB":
				match++
			case "/rpc/GetterService/GetC":
				match++
			case "/rpc/GetterService/GetD":
				match++
			case "/rpc/GetterService/GetE":
				match++
			case "/rpc/GetterService/GetF":
				match++
			case "/rpc/GetterService/GetG":
				match++
			case "/rpc/GetterService/GetH":
				match++
			case "/rpc/GetterService/GetI":
				match++
			case "/rpc/GetterService/GetJ":
				match++
			case "/rpc/GetterService/GetK":
				match++
			case "/rpc/GetterService/GetL":
				match++
			case "/rpc/GetterService/GetM":
				match++
			case "/rpc/GetterService/GetN":
				match++
			case "/rpc/GetterService/GetO":
				match++
			case "/rpc/GetterService/GetP":
				match++
			case "/rpc/GetterService/GetQ":
				match++
			case "/rpc/GetterService/GetR":
				match++
			case "/rpc/GetterService/GetS":
				match++
			case "/rpc/GetterService/GetT":
				match++
			case "/rpc/GetterService/GetU":
				match++
			case "/rpc/GetterService/GetV":
				match++
			case "/rpc/GetterService/GetW":
				match++
			case "/rpc/GetterService/GetX":
				match++
			case "/rpc/GetterService/GetY":
				match++
			case "/rpc/GetterService/GetZ":
				match++
			}
		}
	}
}
