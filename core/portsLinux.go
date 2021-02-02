package core

import (
	"bytes"
	"log"
	"os/exec"
)

//ss -n -l -A 'tcp' | grep -vE "(127.0.0.1|[::1]|[::]):" | tr -s ' ' | cut -d ' ' -f 4 | cut -d ':' -f 2 | grep -vE "Local"

func GetLinuxListeningSockets() ([]int, error) {
	// TODO: Exec la commande au dessus
	// Recup et parse l'output dans un tableau d'int
	cmd := exec.Command("ss -n -l -A 'tcp' | grep -vE '(127.0.0.1|[::1]|[::]):' | tr -s ' ' | cut -d ' ' -f 4 | cut -d ':' -f 2 | grep -vE 'Local'")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	// Parse output of command and return array of int Ports + nil error

	return nil, nil
}
