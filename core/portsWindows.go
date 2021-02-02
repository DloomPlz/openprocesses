package core

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//get-nettcpconnection | where {($_.State -eq "Listen") -and ($_.RemoteAddress -eq '0.0.0.0')}

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

func GetWindowsListeningSockets() (WindowsInfos, error) {
	// Get localAdress + ports : get-nettcpconnection | where {($_.State -eq 'Listen') -and ($_.RemoteAddress -eq '0.0.0.0') -and ($_.LocalAddress -ne '127.0.0.1')} | select LocalAddress,LocalPort

	powershell := New()
	stdOut, _, err := powershell.execute("get-nettcpconnection | where {($_.State -eq 'Listen') -and ($_.RemoteAddress -eq '0.0.0.0') -and ($_.LocalAddress -ne '127.0.0.1')} | select LocalPort,@{Name='Process';Expression={(Get-Process -Id $_.OwningProcess).ProcessName}}")
	if err != nil {
		return WindowsInfos{}, err
	}
	fmt.Print(stdOut)

	infos, err := parseInfos(stdOut)
	if err != nil {
		return WindowsInfos{}, err
	}

	return infos, nil
}

type WindowsInfos struct {
	MachineInfo []Infos `json:"infos"`
}

type Infos struct {
	port    int    `json:"port"`
	process string `json:"process"`
}

func parseInfos(out string) (WindowsInfos, error) {
	var output = WindowsInfos{}
	lines := strings.Split(strings.Replace(out, "\r\n", "\n", -1), "\n")
	for _, l := range lines[3 : len(lines)-3] {
		space := regexp.MustCompile(`\s+`)
		infos := strings.Split(space.ReplaceAllString(l, " "), " ")
		port, err := strconv.Atoi(infos[1])
		if err != nil {
			return WindowsInfos{}, err
		}
		i := Infos{port: port, process: infos[2]}
		output.MachineInfo = append(output.MachineInfo, i)
	}

	return output, nil
}
