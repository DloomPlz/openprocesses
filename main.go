package main

import (
	"fmt"
	"log"
	"openprocesses/core"
)

func main() {

	portsLinux, err := core.GetLinuxListeningSockets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(portsLinux)
	/*machineInfos, err := core.GetWindowsListeningSockets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(machineInfos)*/
}
