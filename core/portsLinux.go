package core

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetLinuxListeningSockets() ([]int, error) {

	cmd := exec.Command("ss -n -l -A 'tcp' | grep -vE '(127.0.0.1|[::1]|[::]):' | tr -s ' ' | cut -d ' ' -f 4 | cut -d ':' -f 2 | grep -vE 'Local'")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	splits := strings.Split(strings.Replace(out.String(), "\n\t", " ", -1), " ")
	var ints []int
	for i := 0; i < len(splits); i++ {
		y, err := strconv.Atoi(splits[i])
		if err != nil {
			return nil, err
		}
		ints = append(ints, y)
	}
	// Parse output of command and return array of int Ports + nil error

	return ints, nil
}
