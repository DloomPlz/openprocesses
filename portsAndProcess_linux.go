// +build linux,!windows

package core

import (
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

	splits := strings.Split(str, "\n")
	for i := 0; i < len(splits); i++ {
		var rePorts = regexp.MustCompile(`(?m) \d{2,5}`)
		tmpPort := rePorts.FindString(splits[i])
		tmpPort = string.Replace(tmpPort, " ", "", -1)

		if tmpPort == "" {
			continue
		}
		port, err := strconv.Atoi(tmpPort)
		if err != nil {
			return nil, err
		}
		var reProcess = regexp.MustCompile(`(?m)"(.*?[^\\])"`)
		process := reProcess.FindString(splits[i])
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

	cmd := `/usr/bin/ss -n -p -l -A 'tcp' | grep -vE '(127.0.0.1|[::1]|[::]):' | grep -vE 'Local'` // will get all processes list with ip ports and process (needs su for some processes to display)

	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}

	infos, err := parsePortsAndProcess(string(out))
	if err != nil {
		return nil, err
	}

	return infos, nil
}
