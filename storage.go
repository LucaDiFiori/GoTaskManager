package main

import (
	"encoding/json"
	"fmt"
	"os"
)


// Nota: I processi di lettura e salvataggio prevedono ogni volta la lettura del file, 
//       il salvataggio del contenuto su un array e la riscrittura del file. 
//       Questo potrebbe non essere l'implementazione più ottimale, 
//       ma per questo esercizio specifico può andare bene.



/********************************************************************/
// READ TASKS
/********************************************************************/

// readTasks legge le task dal file JSON.
// args: nessuno.
// return: restituisce una slice di Task e un errore. Restituisce un errore 
//         se la lettura del file o la deserializzazione fallisce.
// desc: Questa funzione legge le task dal file "tasks.json". Se il file non esiste, 
//       restituisce una slice vuota. 
//       Se il file esiste, legge il contenuto del file, deserializza i dati JSON 
//       in una slice di Task e la restituisce.
func readTasks() ([]Task, error) {
    tasks := []Task{}

    _, err := os.Stat("tasks.json")
   
    if !os.IsNotExist(err) {
        data, err := os.ReadFile("tasks.json")
        if err != nil {
            return nil, fmt.Errorf("error reading tasks file: %v", err)
        }
        
        //se il file non è vuoto
        if len(data) > 0 {
            if err := json.Unmarshal(data, &tasks); err != nil {
                return nil, fmt.Errorf("error unmarshaling tasks: %v", err)
            }
        }
    }
    return tasks, nil
}



/********************************************************************/
// SAVE TASKS
/********************************************************************/

// saveTasks salva le task nel file JSON.
// args: tasks ([]Task) - una slice di Task da salvare.
// return: restituisce un errore se la serializzazione o la scrittura delle task fallisce.
// desc: Questa funzione serializza la slice di Task in formato JSON e la salva nel file "tasks.json".
//       Se la serializzazione fallisce, restituisce un errore. Se la scrittura 
//       del file fallisce, restituisce un errore.
func saveTasks(tasks []Task) error {
    jsonData, err := json.MarshalIndent(tasks, "", "    ")
    if err != nil {
        return fmt.Errorf("error marshaling tasks: %v", err)
    }

    if err := os.WriteFile("tasks.json", jsonData, 0644); err != nil {
        return fmt.Errorf("error writing tasks file: %v", err)
    }
    
    return nil
}