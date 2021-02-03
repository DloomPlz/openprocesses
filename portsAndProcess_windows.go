// +build windows,!linux

package core

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//get-nettcpconnection | where {($_.State -eq 'Listen') -and ($_.RemoteAddress -eq '0.0.0.0') -and ($_.LocalAddress -ne '127.0.0.1')} | select LocalPort,@{Name='Process';Expression={(Get-Process -Id $_.OwningProcess).ProcessName}}

/* EXAMPLE OUTPUT
LocalPort Process
--------- -------
    49680 services
    49668 spoolsv
    49667 svchost
    49666 svchost
    49665 wininit
    49664 lsass
    45769 DiscSoftBusServiceLite
    17500 Dropbox
     5040 svchost
     2179 vmms
      139 System
      139 System
      139 System
      135 svchost
*/

// PowerShell struct
type PowerShell struct {
	powerShell string
}

// New create new session
func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

func GetListeningSockets() ([]PortsAndProcessesInformations, error) {
	// Get localAdress + ports : get-nettcpconnection | where {($_.State -eq 'Listen') -and ($_.RemoteAddress -eq '0.0.0.0') -and ($_.LocalAddress -ne '127.0.0.1')} | select LocalAddress,LocalPort

	powershell := New()
	stdOut, _, err := powershell.execute("get-nettcpconnection | where {($_.State -eq 'Listen') -and ($_.RemoteAddress -eq '0.0.0.0') -and ($_.LocalAddress -ne '127.0.0.1')} | select LocalPort,@{Name='Process';Expression={(Get-Process -Id $_.OwningProcess).ProcessName}}")
	if err != nil {
		return nil, err
	}
	fmt.Print(stdOut)

	infos, err := parsePortsAndProcess(stdOut)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func parsePortsAndProcess(out string) ([]PortsAndProcessesInformations, error) {
	var output []PortsAndProcessesInformations
	lines := strings.Split(strings.Replace(out, "\r\n", "\n", -1), "\n")
	for _, l := range lines[3 : len(lines)-3] {
		space := regexp.MustCompile(`\s+`)
		infos := strings.Split(space.ReplaceAllString(l, " "), " ")
		port, err := strconv.Atoi(infos[1])
		if err != nil {
			return nil, err
		}
		i := PortsAndProcessesInformations{Port: port, Process: infos[2]}
		output = append(output, i)
	}

	return output, nil
}
