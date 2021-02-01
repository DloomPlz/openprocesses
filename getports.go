package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

type Ports struct {
	mux sync.Mutex
	out []int
}

func (p *Ports) Add(i int) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.out = append(p.out, i)
}

const (
	minTCPPort         = 0
	maxTCPPort         = 65535
	maxReservedTCPPort = 1024
	maxRandTCPPort     = maxTCPPort - (maxReservedTCPPort + 1)
)

func IsTCPPortAvailable(ip string, port int) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func getPortsByInterfaceIP(ip string) []int {
	ports := Ports{}
	wg := new(sync.WaitGroup)

	for i := 0; i < maxTCPPort; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			open := IsTCPPortAvailable(ip, i)
			if open {
				ports.Add(i)
			}
		}(i)
	}
	wg.Wait()
	return ports.out
}

func main() {
	ports := getPortsByInterfaceIP("127.0.0.1")
	sort.Ints(ports)
	fmt.Println(ports)
}
