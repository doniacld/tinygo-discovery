// Simple Hello world program that you can build using go and tinygo compilers
// to see from your own eyes the size difference.
// go build -o bin/helloworld-go hello/hello_world.go
// tinygo build -o bin/helloworld-tinygo hello/hello_world.go
// ll bin/
// total 3888
// -rwxr-xr-x  1 doniacld  staff   1.8M Jul 23 16:40 helloworld-go
// -rwxr-xr-x  1 doniacld  staff   107K Jul 23 16:40 helloworld-tinygo
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
