package main

//ss -n -l -A 'tcp' | grep -vE "(127.0.0.1|[::1]|[::]):" | tr -s ' ' | cut -d ' ' -f 4 | cut -d ':' -f 2 | grep -vE "Local"

func main() {
	// TODO: Exec la commande au dessus
	// Recup et parse l'output dans un tableau d'int

}
