package main

import (
	"fmt"
	"openprocesses/core"
)

func main() {
	ports := core.GetLinuxListeningSockets()
	fmt.Printf("translated phrase: %q\n", ports)
}
