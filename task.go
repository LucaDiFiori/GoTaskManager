package main

import (
	"fmt"
	"time"
	"flag"
	"strconv"
)

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Done        bool      `json:"done"`
    CreatedAt   string `json:"created_at"`
}




/********************************************************************/
// UTILS
/********************************************************************/

//Funzione per leggere generare un nuovo ID
func generateNewID(tasks []Task) int {
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}
	return newID
}

// parseAddArgs analizza gli argomenti per il comando "add".
// args: una slice di stringhe che rappresenta gli argomenti della riga di comando.
// return: restituisce il titolo e la descrizione della task. Restituisce un 
//         errore se il titolo è mancante.
// desc: Analizza gli argomenti della riga di comando per il comando "add" e 
//       restituisce il titolo e la descrizione della task.
//       Se le flag "title" e "desc" non sono fornite, utilizza gli argomenti 
//       posizionali per ottenere il titolo e la descrizione.
func parseAddArgs(args []string) (string, string, error) {
    addCmd := flag.NewFlagSet("add", flag.ExitOnError)
    title := addCmd.String("title", "", "Title of the task")
    description := addCmd.String("desc", "", "Description of the task")
    addCmd.Parse(args)

    if *title == "" {
        remainingArgs := addCmd.Args()
        if len(remainingArgs) < 1 {
            return "", "", fmt.Errorf("either flags or positional arguments for title are required")
        }
        *title = remainingArgs[0]
        if len(remainingArgs) > 1 {
            *description = remainingArgs[1]
        }
    }

    if *title == "" {
        return "", "", fmt.Errorf("title is required")
    }

    return *title, *description, nil
}

// parseIDArgs analizza gli argomenti per i comandi che richiedono un ID.
// args: una slice di stringhe che rappresenta gli argomenti della riga di comando.
// return: restituisce l'ID della task. Restituisce un errore se l'ID è mancante o non valido.
// desc: Analizza gli argomenti della riga di comando per i comandi che richiedono un ID e restituisce l'ID della task.
//       Se la flag "id" non è fornita, utilizza gli argomenti posizionali per ottenere l'ID.
func parseIDArgs(args []string) (int, error) {
    cmd := flag.NewFlagSet("id", flag.ExitOnError)
    id := cmd.Int("id", -1, "ID of the task")
    cmd.Parse(args)

    if *id == -1 {
        remainingArgs := cmd.Args()
        if len(remainingArgs) < 1 {
            return -1, fmt.Errorf("either flags or positional arguments are required")
        }
        parsedID, err := strconv.Atoi(remainingArgs[0])
        if err != nil {
            return -1, fmt.Errorf("invalid ID: %v", err)
        }
        *id = parsedID
    }
    return *id, nil
}





/********************************************************************/
// LIST COMMAND
/********************************************************************/

// list gestisce il comando "list" per elencare tutte le task.
// args: un'interfaccia che rappresenta gli argomenti della riga di comando 
//       (non utilizzati in questa funzione).
// return: restituisce un errore se la lettura delle task fallisce.
// desc: Questa funzione legge le task esistenti dal file e le stampa. 
//       Se non ci sono task, stampa un messaggio d'errore.
func list(args interface{}) error {
	tasks, err := readTasks()
	if err != nil {
		return fmt.Errorf("error reading tasks: %v", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}
	
	// NOTA: "range" restituisce l'indice e il valore di ogni elemento dell'array.
	// Se non si vuole utilizzare l'indice, si può sostituire con un underscore "_"
	for _, task := range tasks {
		doneStatus := ""
		if task.Done {
			doneStatus = "✓"
		}
		fmt.Printf("ID: %d\nTitle: %s\nDescription: %s\nDone: %s\nCreatedAt: %v\n\n",
			task.ID, task.Title, task.Description, doneStatus, task.CreatedAt)
	}
	
	return nil
}





/********************************************************************/
// ADD COMMAND
/********************************************************************/

// addTask aggiunge una nuova task alla lista delle task.
// args: title (string) - il titolo della task, description (string) - la descrizione della task.
// return: restituisce un errore se la lettura o la scrittura delle task fallisce.
// desc: Crea una nuova task con il titolo e la descrizione forniti e la 
//       aggiunge alla lista delle task.
//       La funzione legge le task esistenti dal file, genera un nuovo ID per la nuova task, a
//       ggiunge la nuova task alla lista e salva la lista aggiornata nel file.
func addTask(title, description string) error {
    tasks, err := readTasks()
    if err != nil {
        return err
    }

    newtask := Task{
        ID:          generateNewID(tasks),
        Title:       title,
        Description: description,
        Done:        false,
        CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
    }

    tasks = append(tasks, newtask)

    err = saveTasks(tasks)
    if err != nil {
        return err
    }

    fmt.Println("Task added successfully")
    return nil
}

// add gestisce il comando "add" per aggiungere una nuova task.
// args: un'interfaccia che rappresenta gli argomenti della riga di comando.
// return: restituisce un errore se il parsing degli argomenti o l'aggiunta della task fallisce.
// desc: Questa funzione analizza gli argomenti della riga di comando per il comando "add" e aggiunge 
//       una nuova task con il titolo e la descrizione forniti.
//       Utilizza parseAddArgs per analizzare gli argomenti e addTask per aggiungere la nuova task.
func add(args interface{}) error {
    argList, ok := args.([]string)
    if !ok {
        return fmt.Errorf("invalid arguments for add command")
    }

    title, description, err := parseAddArgs(argList)
    if err != nil {
        return err
    }

    return addTask(title, description)
}





/********************************************************************/
// DONE COMMAND
/********************************************************************/

// markTaskAsDone segna una task come completata.
// args: id (int) - l'ID della task da segnare come completata.
// return: restituisce un errore se la lettura o la scrittura delle task fallisce,
//         o se la task con l'ID specificato non viene trovata.
// desc: Segna la task con l'ID specificato come completata e salva le task aggiornate.
//       La funzione legge le task esistenti dal file, cerca la task con l'ID specificato, 
//       aggiorna il campo "Done" a true, e salva la lista aggiornata nel file. 
//       Se la task è già completata, stampa un messaggio appropriato.
func markTaskAsDone(id int) error {
    tasks, err := readTasks()
    if err != nil {
        return err
    }

    found := false
    alreadyDone := false
    for i, task := range tasks {
        if task.ID == id {
            if tasks[i].Done {
                alreadyDone = true
            } else {
                tasks[i].Done = true
                found = true
            }
            break
        }
    }

    if !found && !alreadyDone {
        return fmt.Errorf("task with ID %d not found", id)
    }

    if found {
        fmt.Printf("Task %d marked as done\n", id)
    } else if alreadyDone {
        fmt.Printf("You've already completed this task, well done :)\n")
    }

    return saveTasks(tasks)
}


// done gestisce il comando "done" per segnare una task come completata.
// args: un'interfaccia che rappresenta gli argomenti della riga di comando.
// return: restituisce un errore se il parsing degli argomenti o il 
//         completamento della task fallisce.
// desc: Questa funzione analizza gli argomenti della riga di comando 
//       per il comando "done" e segna la task con l'ID specificato come completata.
//       Utilizza parseIDArgs per analizzare gli argomenti e markTaskAsDone 
//       per segnare la task come completata.
func done(args interface{}) error {
    argList, ok := args.([]string)
    if !ok {
        return fmt.Errorf("invalid arguments for done command")
    }

    id, err := parseIDArgs(argList)
    if err != nil {
        return err
    }

    return markTaskAsDone(id)
}





/********************************************************************/
// DELETE COMMAND
/********************************************************************/

// deleteTask elimina una task dalla lista delle task.
// args: id (int) - l'ID della task da eliminare.
// return: restituisce un errore se la lettura o la scrittura delle task fallisce, 
//         o se la task con l'ID specificato non viene trovata.
// desc: Elimina la task con l'ID specificato dalla lista delle task e salva le task aggiornate.
//       La funzione legge le task esistenti dal file, cerca la task 
//       con l'ID specificato, rimuove la task dalla lista, e salva la lista aggiornata nel file.
func deleteTask(id int) error {
    tasks, err := readTasks()
    if err != nil {
        return err
    }

    found := false
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            found = true
            break
        }
    }

    if !found {
        return fmt.Errorf("task with ID %d not found", id)
    }

    err = saveTasks(tasks)
    if err != nil {
        return err
    }

    fmt.Printf("Task %d deleted successfully\n", id)
    return nil
}


// delete gestisce il comando "delete" per eliminare una task.
// args: un'interfaccia che rappresenta gli argomenti della riga di comando.
// return: restituisce un errore se il parsing degli argomenti o l'eliminazione della task fallisce.
// desc: Questa funzione analizza gli argomenti della riga di comando per 
//       il comando "delete" e elimina la task con l'ID specificato.
//       Utilizza parseIDArgs per analizzare gli argomenti e deleteTask per eliminare la task.
func delete(args interface{}) error {
    argList, ok := args.([]string)
    if !ok {
        return fmt.Errorf("invalid arguments for delete command")
    }

    id, err := parseIDArgs(argList)
    if err != nil {
        return err
    }

    return deleteTask(id)
}
