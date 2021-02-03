// +build linux,!windows

package core

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

/* EXAMPLE OUTPUT of : ss -n -p -l -A 'tcp' | grep -vE '(127.0.0.1|[::1]|[::]):' | grep -vE 'Local'
`LISTEN 0      4096               *:4320            :    users:(("process1",pid=9803,fd=3))
	LISTEN 0      80                 *:3306            :    users:(("process2",pid=9534,fd=21))`
*/

func parsePortsAndProcess(str string) ([]PortsAndProcessesInformations, error) {
	var output []PortsAndProcessesInformations

	splits := strings.Split(strings.Replace(str, " ", "", -1), "\n")
	for i := 0; i < len(splits); i++ {
		tmpPort := strings.Split(splits[i], ":")
		port, err := strconv.Atoi(tmpPort[1])
		if err != nil {
			return PortsAndProcessesInformations{}, err
		}
		var re = regexp.MustCompile(`(?m)"(.*?[^\\])"`)
		process := re.FindString(splits[i])
		if process != "" {
			process = process[1 : len(process)-1] // delete quotes
		} else {
			process = ""
		}

		i := PortsAndProcessesInformations{Port: port, Process: process}
		output = append(output, i)
	}
	return output, nil
}

func GetListeningSockets() ([]PortsAndProcessesInformations, error) {

	cmd := exec.Command("ss -n -p -l -A 'tcp' | grep -vE '(127.0.0.1|[::1]|[::]):' | grep -vE 'Local'") // will get all processes list with ip ports and process (needs su for some processes to display)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return PortsAndProcessesInformations{}, err
	}

	infos, err := parsePortsAndProcess(out.String())
	if err != nil {
		return PortsAndProcessesInformations{}, err
	}

	return infos, nil
}
