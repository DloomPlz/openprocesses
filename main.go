package main

import (
	"fmt"
	"log"
	"openprocesses/core"
)

func main() {
	/*
		portsLinux, err := core.GetLinuxListeningSockets()
		if err != nil {
			log.Fatal(err)
		}*/
	machineInfos, err := core.GetWindowsListeningSockets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Machine infos: %q\n", machineInfos)
}
