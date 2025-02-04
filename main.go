package main

import (
	"fmt"
	"os"
)

//Mappa che associa ad ogni comando una funzione
var commands = map[string]func(interface{}) error{
    "add": add,
    "list": list,
    "done": done,
    "delete": delete,
}


/********************************************************************/
// CHECKS
/********************************************************************/
func checkArgsLen(args *[]string) {
    if len(*args) < 2 || len(*args) > 6 {
        fmt.Println("Usage: <command> <flag> <args>")
        os.Exit(1)
    }
}




/********************************************************************/
// MAIN
/********************************************************************/

// args: nessuno.
// return: nessuno.
// desc: Analizza gli argomenti della riga di comando per determinare quale comando eseguire.
//       Utilizza una mappa `commands` che associa i nomi dei comandi alle rispettive. 
//		 Se il comando è valido, esegue la funzione associata al comando con gli argomenti forniti.
//       Se il comando non è valido, stampa un messaggio di errore e termina il programma. 
//       La funzione supporta i seguenti comandi:
//       - "list": elenca tutte le task.
//       - "add": aggiunge una nuova task.
//       - "done": segna una task come completata.
//       - "delete": elimina una task.

// Nota: L'uso di un `interface{}` per gli argomenti dei comandi potrebbe 
//       non essere l'implementazione migliore in un programma più complesso,
//       data la necessità di un'asserzione di tipo ogni volta che si accede agli argomenti. 
//       Tuttavia, l'ho utilizzata qui per scopi di studio.

func main() {
	fmt.Println("=== Welcome to the Task Manager ===")

	checkArgsLen(&os.Args)

	command := os.Args[1]

	if commandFunc, exists := commands[command]; exists {
    	var args interface{}

    	switch command {
    	case "list":
        	args = nil
		default:
			args = os.Args[2:]
    	}

    	if err := commandFunc(args); err != nil {
        	fmt.Printf("Error executing command '%s': %v\n", command, err)
        	os.Exit(1)
    	}
	} else {
    	fmt.Println("Command not found")
    	os.Exit(1)
	}
}