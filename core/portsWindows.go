package core

import (
	"bytes"
	"log"
	"os/exec"
)

//get-nettcpconnection | where {($_.State -eq "Listen") -and ($_.RemoteAddress -eq '0.0.0.0')}

func GetWindowsListeningSockets() ([]int, error) {
	// TODO: Exec la commande au dessus
	// Recup et parse l'output dans un tableau d'int
	cmd := exec.Command("get-nettcpconnection | where {($_.State -eq 'Listen') -and ($_.RemoteAddress -eq '0.0.0.0')}")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	// Parse output of command and return array of int Ports + nil error

	return nil, nil
}
