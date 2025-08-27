package main

import (
	"fmt"
	foo "net/http" // named imports [http -> go]
)

func main() {
	fmt.Println("Hello, Go standard library")
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Http Response status: ", resp.Status)

}

/*
Tree shaking in Go refers to the process of removing unused code (functions, variables, packages) during compilation so the final binary only contains what’s actually referenced.
Unlike JavaScript where bundlers perform tree shaking, Go’s compiler and linker do it automatically:
The compiler builds everything reachable from main.
Unused packages, functions, or symbols that aren’t referenced are dropped.
This keeps binaries smaller and more efficient without extra tooling.
In short: Go’s build system inherently does tree shaking via dead-code elimination at compile/link time.
*/
