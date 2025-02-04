# DOCUMENTAZIONE
[Descrizione e utilizzo](#descrizione-e-utilizzo)  
[Strutture dati](#strutture-dati)  
[Struttura Dati: Task](#struttura-dati-task)  
[Struttura Dati: Map](#struttura-dati-map)  
[Funzioni](#funzioni)  

## DESCRIZIONE E UTILIZZO
### Descrizione del Programma
Questo programma simula un Task Manager da linea di comando che permette di gestire una lista di task. 
Supporta i seguenti comandi:
- **add**: Aggiunge una nuova task con un titolo e una descrizione.
- **list**: Elenca tutte le task esistenti.
- **done**: Segna una task come completata utilizzando il suo ID.
- **delete**: Elimina una task utilizzando il suo ID.

Il programma utilizza un file JSON (tasks.json) per memorizzare le task. Ogni task ha un ID univoco, un titolo, una descrizione, uno stato di completamento e una data di creazione.

### Esempio di Utilizzo
Ecco alcuni esempi di come utilizzare il programma dalla riga di comando:

- Aggiungere una nuova task:
```bash
go run main.go add -title "Fare la spesa" -desc "Comprare latte, pane e uova"
#oppure
go run main.go add "Fare la spesa" "Comprare latte, pane e uova"
```

- Elencare tutte le task:
```bash
go run main.go list
```

- Segnare una task come completata:
```bash
go run main.go done -id 1
#oppure
go run main.go done 1
```

- Eliminare una task:
```bash
go run main.go delete -id 1
#oppure
go run main.go delete 1
```

***

## STRUTTURE DATI
### Struttura Dati: Task
La struttura `Task` rappresenta una singola attività nel Task Manager. Ogni task ha un ID univoco, un titolo, una descrizione, uno stato di completamento e una data di creazione.

```go
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Done        bool      `json:"done"`
    CreatedAt   string    `json:"created_at"`
}
```

**Campi**
- `ID` (int): Un identificatore univoco per la task. Viene generato automaticamente quando una nuova task viene aggiunta.
- `Title` (string): Il titolo della task. Questo campo è obbligatorio e descrive brevemente l'attività.
- `Description` (string): Una descrizione dettagliata della task. Questo campo è opzionale.
- `Done` (bool): Indica se la task è stata completata. Il valore predefinito è false.
- `CreatedAt` (string): La data e l'ora in cui la task è stata creata. Viene impostata automaticamente quando la task viene aggiunta.

**JSON tag**
Ogni campo della struttura Task è annotato con un tag JSON che specifica il nome del campo quando la struttura viene serializzata o deserializzata in formato JSON. Ad esempio, il campo ID verrà rappresentato come id nel JSON.



### Struttura Dati: Map
commands è una mappa che associa i nomi dei comandi alle rispettive funzioni.
Viene utilizzata nel punto di ingresso del programma (main) per determinare quale funzione eseguire in base al comando fornito dall'utente.

Ogni funzione accetterà un'interfaccia (interface{}) come argomento e restituirà un errore (error).

```go
var commands = map[string]func(interface{}) error{
    "add": add,
    "list": list,
    "done": done,
    "delete": delete,
}
```
I comandi supportati sono:
- "add": aggiunge una nuova task.
- "list": elenca tutte le task.
- "done": segna una task come completata.
- "delete": elimina una task.

***

## FUNZIONI

#### main.go
```go
// La funzione checkArgsLen verifica che il numero di argomenti passati al programma sia corretto. Se il numero di 
//argomenti non è compreso tra 2 e 6, la funzione stampa un messaggio di utilizzo e termina il programma.
func checkArgsLen(args *[]string)

// Il main analizza gli argomenti della riga di comando per determinare quale comando eseguire. Utilizza una mappa 
//commands che associa i nomi dei comandi alle rispettive funzioni. Se il comando è valido, esegue la funzione associata 
//al comando con gli argomenti forniti. Se il comando non è valido, stampa un messaggio di errore e termina il programma.
func main()
```

#### storage.go
```go
// La funzione readTasks legge le task dal file JSON. Se il file non esiste, restituisce una slice vuota. Se il file esiste, 
//legge il contenuto del file, deserializza i dati JSON in una slice di Task e la restituisce. Restituisce un errore se 
//la lettura del file o la deserializzazione fallisce.
func readTasks() ([]Task, error)

// La funzione saveTasks salva le task nel file JSON. Serializza la slice di Task in formato JSON e la salva nel file 
//"tasks.json". Restituisce un errore se la serializzazione o la scrittura del file fallisce.
func saveTasks(tasks []Task) error
```

#### task.go
```go
// La funzione generateNewID genera un nuovo ID univoco per una task.
func generateNewID(tasks []Task) int 

// parseAddArgs analizza gli argomenti della riga di comando per il comando "add" e restituisce il titolo e la descrizione della task. 
// Se le flag "title" e "desc" non sono fornite, utilizza gli argomenti posizionali per ottenere il titolo e la descrizione.
func parseAddArgs(args []string) (string, string, error)

// Questa funzione analizza gli argomenti della riga di comando per i comandi che richiedono un ID erestituisce l'ID della task.
// Se la flag "id" non è fornita, utilizza gli argomenti posizionali per ottenere l'ID.
func parseIDArgs(args []string) (int, error)

// Questa funzione legge le task esistenti dal file e le stampa. Se non ci sono task, stampa un messaggio appropriato.
func list(args interface{}) error

// addTask crea una nuova task con il titolo e la descrizione forniti e la aggiunge alla lista delle task.La funzione 
// legge le task esistenti dal file, genera un nuovo ID per la nuova task, aggiunge la nuova task alla lista e salva la lista aggiornata nel file.
func addTask(title, description string) error

// Questa funzione analizza gli argomenti della riga di comando per il comando "add" e aggiunge una nuova task con il titolo 
// e la descrizione forniti. La funzione utilizza parseAddArgs per analizzare gli argomenti e addTask per aggiungere la nuova task.
func add(args interface{}) error

// Questa funzione segna la task con l'ID specificato come completata e salva le task aggiornate.La funzione legge le 
// task esistenti dal file, cerca la task con l'ID specificato, aggiorna il campo "Done" a true,e salva la lista aggiornata nel file. 
// Se la task è già completata, stampa un messaggio appropriato.
func markTaskAsDone(id int) error

// Questa funzione analizza gli argomenti della riga di comando per il comando "done" e segna la task con l'ID specificato come completata. 
// La funzione utilizza parseIDArgs per analizzare gli argomenti e markTaskAsDone per segnare la task come completata.
func done(args interface{}) error

// Questa funzione elimina la task con l'ID specificato dalla lista delle task e salva le task aggiornate.Legge le task esistenti dal file, 
// cerca la task con l'ID specificato, rimuove la task dalla lista,e salva la lista aggiornata nel file.
func deleteTask(id int) error

// Questa funzione analizza gli argomenti della riga di comando per il comando "delete" e elimina la task con l'ID specificato.
// Utilizza parseIDArgs per analizzare gli argomenti e deleteTask per eliminare la task.
func delete(args interface{}) error
```

***
***

# CONCETTI UTILI
- Subject [#Go_CLI_Task_Manager_-_Getting_Started_with_Go]
- Funzioni importanti [#FUNZIONI_IMPORTANTI]
    - os.Stat [##OS.STAT]
    - os.IsNotExist [##OS.ISNOTEXIST]
    - os.ReadFile [##OS.READFILE]
    - json.Unmarshal [##JSON.UNMARSHAL]
    - json.Marshal [##JSON.MARSHAL]
    - json.MarshalIndent [##JSON.MARSHALINDENT]
- Interface [#INTERFACE]
    - interface{} [##INTERFACE{}]
    - type assertion [##TYPE_ASSERTION]

***
***

# Go CLI Task Manager - Getting Started with Go

## Project Overview
You'll be building a simple command-line task manager in Go. This project will help you learn the fundamentals of Go programming while creating something practical.

## Learning Objectives
- Understanding basic Go syntax and types
- Working with structs and slices
- Basic file I/O operations
- Command-line argument parsing
- Error handling in Go
- Writing modular Go code

## Project Requirements

### Core Features
1. Add a new task with a title and optional description
2. List all tasks
3. Mark a task as complete
4. Delete a task
5. Save tasks to a JSON file
6. Load tasks from a JSON file

### Data Structure
```go
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Done        bool      `json:"done"`
    CreatedAt   time.Time `json:"created_at"`
}
```

### Command-Line Interface
The program should support the following commands:
```bash
# Add a new task
./task-manager add "Buy groceries" "Milk, bread, and eggs"

# List all tasks
./task-manager list

# Mark task as done (using task ID)
./task-manager done 1

# Delete a task (using task ID)
./task-manager delete 1
```

## Step-by-Step Implementation Guide

### 1. Project Setup
Create a new directory and initialize your Go module:
```bash
mkdir task-manager
cd task-manager
go mod init task-manager
```

### 2. Project Structure
```
task-manager/
├── main.go       # Entry point
├── task.go       # Task type and methods
├── storage.go    # File operations
└── tasks.json    # Data storage file
```

### 3. Implementation Steps - Up to you !

## Testing Your Implementation

Test your implementation with these commands:
```bash
# Add some tasks
go run . add -title "Learn Go basics" -desc "Complete the CLI project"
go run . add -title "Read Go documentation" -desc "Focus on slices and structs"

# List all tasks
go run . list

# Mark the first task as done
go run . done -id 1

# List again to see the changes
go run . list

# Delete a task
go run . delete -id 2
```

## Bonus Challenges
Once you've completed the basic implementation, try these extensions:
1. Add due dates to tasks
2. Add priority levels (High, Medium, Low)
3. Add a search function to find tasks by title
4. Add task categories or tags

## Go Concepts Used
- Structs
- Slices
- JSON marshaling/unmarshaling
- File I/O
- Command-line flags
- Error handling
- Time handling
- Basic control structures (if, for, switch)

## Resources
- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)

Remember: The goal is to learn Go basics, so focus on understanding each concept as you implement it. Don't hesitate to consult the Go documentation and experiment with the code!



***
***



# FUNZIONI IMPORTANTI

## OS.STAT

```go
func Stat(name string) (FileInfo, error)
```

In Go, la funzione os.Stat viene utilizzata per ottenere informazioni su un file o una directory. Restituisce una struttura di tipo os.FileInfo che contiene dettagli come dimensione, permessi, timestamp di modifica e se il percorso specificato è un file o una directory.

Esempio:
```go
import (
    "fmt"
    "os"
)

func main() {
    fileInfo, err := os.Stat("nomefile.txt")
    if err != nil {
        fmt.Println("Errore:", err)
        return
    }

    fmt.Println("Nome:", fileInfo.Name())
    fmt.Println("Dimensione:", fileInfo.Size(), "byte")
    fmt.Println("Permessi:", fileInfo.Mode())
    fmt.Println("Ultima modifica:", fileInfo.ModTime())
    fmt.Println("È una directory?", fileInfo.IsDir())
}
```

**Gestione degli errori**
Questa funzione può restituire diversi errori: se il file non esiste, se il file è corrotto.

Se il file o la directory non esistono, possiamo verificarlo  usando os.IsNotExist(err):
```go
if os.IsNotExist(err) {
    fmt.Println("Il file non esiste")
}
```

***

## OS.ISNOTEXIST

```go
func IsNotExist(err error) bool
```

La funzione os.IsNotExist in Go viene utilizzata per verificare se un errore restituito da una funzione del pacchetto os è dovuto al fatto che un file o una directory non esistono.

Ad esempio può prendere in input l'errore restituito da os.Stat e restituire *true* solo se l'errore indica che il file non esiste


***

## OS.READFILE

```go
func ReadFile(name string) ([]byte, error)
```

La funzione os.ReadFile viene utilizzata per leggere il contenuto di un file e restituirlo come una slice di byte ([]byte). È una funzione di utilità che combina l'apertura, la lettura e la chiusura del file in un'unica operazione.

Fra gli errori che può riportare c'è l'assenza del file, problemi di permessi etc (come per os.Stat)

NOTA: Non va usato per file di grandi dimensioni

**Alternative a os.ReadFile: bufio.Scanner**
Se vuoi leggere un file riga per riga invece di caricare tutto in memoria, puoi usare bufio.Scanner:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Errore:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text()) // Stampa una riga alla volta
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Errore nella lettura:", err)
    }
}
```


***

## JSON.UNMARSHAL

La funzione json.Unmarshal in Go trasforma (deserializza) una stringa JSON in una variabile Go, come una struct, una mappa, o un array.

Esempio:
- **JSON (stringa)**: {"nome": "Luca", "età": 30}
- **Go struct**: Persona{Nome: "Luca", Età: 30}

**Fondamentalmente, Unmarshal legge il JSON e riempie la variabile passata.**

**Firma**
```go
func Unmarshal(data []byte, v any) error
```
- `data []byte` → un array di byte contenente le informazioni del JSON da decodificare.
- `v any` → Il puntatore alla variabile in cui salvare il risultato.
- `Restituisce error` → nil se il parsing è riuscito, altrimenti un errore.



**Esempio**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Persona struct {
    Nome  string `json:"nome"`
    Età   int    `json:"età"`
}

func main() {
    jsonData := []byte(`{"nome": "Luca", "età": 30}`)

    var p Persona
    err := json.Unmarshal(jsonData, &p)
    if err != nil {
        fmt.Println("Errore:", err)
        return
    }

    fmt.Println("Nome:", p.Nome)
    fmt.Println("Età:", p.Età)
}
```
🔹 **Nota**: Devi passare il puntatore alla variabile (&p), altrimenti Unmarshal non può modificarne il valore.


### Quando Unmarshal restituisce errore
json.Unmarshal restituisce un errore (error) nei seguenti casi:

#### 1. Il JSON non è valido (errore di sintassi)
Esempio di JSON non valido (manca una virgola):
```go
jsonData := []byte(`{"nome": "Luca" "età": 30}`) // ❌ ERRORE
```

Errore restituito:
```go
invalid character '"' after object key:value pair
```


#### 2. Il JSON non corrisponde alla struttura Go
Se il JSON ha dati che non possono essere convertiti nel tipo Go specificato.
Esempio: il JSON ha età come stringa, ma la struct lo aspetta come int.

```go
type Persona struct {
    Nome string `json:"nome"`
    Età  int    `json:"età"`  // ❌ Deve essere un numero
}

jsonData := []byte(`{"nome": "Luca", "età": "trenta"}`)
err := json.Unmarshal(jsonData, &p)
```

Errore restituito
```go
json: cannot unmarshal string into Go struct field Persona.età of type int
```

#### Altri errori
- **Non passi un puntatore alla variabile**
- **Il JSON ha tipi diversi da quelli attes**



***

## JSON.MARSHAL
La funzione json.Marshal in Go è usata per convertire (serializzare) una variabile Go in JSON.

- Prende una struct, una mappa, un array, ecc.
- Restituisce una slice di byte ([]byte) contenente il JSON.

**Frima**
```go
func Marshal(v any) ([]byte, error)
```
- **v any** → L'oggetto Go da convertire in JSON.
- **Restituisce**:
    - []byte → Il JSON generato.
    - error → nil se tutto va bene, altrimenti un errore.


Esempio
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Persona struct {
    Nome  string `json:"nome"`
    Età   int    `json:"età"`
}

func main() {
    p := Persona{Nome: "Luca", Età: 30}

    jsonData, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Errore:", err)
        return
    }

    fmt.Println(string(jsonData)) // Output: {"nome":"Luca","età":30}
}
```
🔹 **Nota**: json.Marshal restituisce []byte, quindi usiamo string(jsonData) per stamparlo.


### Quando Unmarshal restituisce errore
#### Strutture con campi non serializzabili
Se la struct ha campi non esportati (iniziano con una lettera minuscola), non vengono inclusi nel JSON.
```go
type Persona struct {
    Nome  string `json:"nome"`
    età   int    // ❌ Non esportato, non verrà incluso
}
```

In Go, infatti, la visibilità dei campi di una struct segue le regole del modificatore di accesso implicito:

- **Lettera maiuscola** → Il campo è esportato (visibile all'esterno del pacchetto).
- **Lettera minuscola** → Il campo è non esportato (privato, visibile solo dentro il pacchetto).

La libreria encoding/json può accedere solo ai campi esportati, quindi:

- Se un campo ha una lettera maiuscola, json.Marshal può leggerlo e convertirlo in JSON.
- Se un campo ha una lettera minuscola, json.Marshal lo ignora.


***


## JSON.MARSHALINDENT

In Go, la funzione json.MarshalIndent viene utilizzata per serializzare una struttura dati in JSON formattato con rientri, rendendolo più leggibile. È simile a json.Marshal, ma aggiunge un'indentazione personalizzabile per migliorare la leggibilità.

**Sintassi**
```go
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```

- **v**: L'oggetto da serializzare in JSON.
- **prefix**: Una stringa opzionale da aggiungere all'inizio di ogni riga (di solito lasciata vuota "").
- **indent**: La stringa utilizzata per l'indentazione (spesso " " per due spazi o "\t" per un tab) rispetto alla parentesi graffa di apertura del file JSON 

**Ritorno**
- Un []byte contenente il JSON formattato.
- Un error se la serializzazione fallisce.



***
***

# INTERFACE
Un'interfaccia è un insieme di metodi che un tipo deve implementare per soddisfare l'interfaccia stessa. Ecco una spiegazione dettagliata delle interfacce in Go:

### Definizione di un'interfaccia
Un'interfaccia è definita utilizzando la parola chiave interface. Ecco un esempio di un'interfaccia chiamata Speaker che richiede un metodo Speak:

```go
type Speaker interface {
    Speak() string
}
```

### Implementazione di un'interfaccia
Qualsiasi tipo che implementa tutti i metodi definiti nell'interfaccia soddisfa automaticamente quell'interfaccia. Non è necessario dichiarare esplicitamente che un tipo implementa un'interfaccia; basta fornire i metodi richiesti. Ecco un esempio:

```go
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}
```
In questo esempio, sia Dog che Cat implementano l'interfaccia Speaker poiché entrambi forniscono un metodo Speak.


### Utilizzo delle interfacce
Le interfacce possono essere utilizzate come tipi di variabili, consentendo di scrivere codice generico che può funzionare con diversi tipi che implementano l'interfaccia. Ecco un esempio di utilizzo:

```go
package main

import (
    "fmt"
)

// Definizione dell'interfaccia Speaker
type Speaker interface {
    Speak() string
}

// Definizione della struct Dog
type Dog struct{}

// Implementazione del metodo Speak per Dog
func (d Dog) Speak() string {
    return "Woof!"
}

// Definizione della struct Cat
type Cat struct{}

// Implementazione del metodo Speak per Cat
func (c Cat) Speak() string {
    return "Meow!"
}

// Funzione MakeSound che accetta un tipo Speaker
func MakeSound(s Speaker) {
    fmt.Println(s.Speak())
}

// Funzione main
func main() {
    dog := Dog{} // Creazione di un'istanza di Dog
    cat := Cat{} // Creazione di un'istanza di Cat

    MakeSound(dog) // Chiamata a MakeSound con l'istanza di Dog
    MakeSound(cat) // Chiamata a MakeSound con l'istanza di Cat
}
```

***

# INTERFACE{}
In Go, interface{} è un tipo di interfaccia vuota che può contenere un valore di qualsiasi tipo. È un modo per definire variabili che possono contenere qualsiasi valore, rendendo interface{} un tipo molto flessibile e versatile.

#### Caratteristiche di interface{}
- **Tipi dinamici**: Poiché interface{} può contenere valori di qualsiasi tipo, è utile per scrivere funzioni generiche che devono accettare input di diversi tipi.

- **Nessun metodo richiesto**: Non è necessario implementare alcun metodo specifico per utilizzare interface{}, il che significa che puoi passare oggetti di qualsiasi tipo senza doverli adattare a un'interfaccia specifica.

- **Type Assertion**: Quando utilizzi interface{}, per accedere al valore effettivo contenuto nell'interfaccia, puoi utilizzare l'asserzione di tipo per ottenere il valore specifico.


**Esempio**
```go
package main

import (
	"fmt"
)

func printValue(v interface{}) {
	// Utilizziamo l'asserzione di tipo per identificare il tipo specifico
	switch value := v.(type) {
	case int:
		fmt.Printf("Il valore è un intero: %d\n", value)
	case string:
		fmt.Printf("Il valore è una stringa: %s\n", value)
	case float64:
		fmt.Printf("Il valore è un float: %f\n", value)
	default:
		fmt.Println("Tipo non supportato")
	}
}

func main() {
	printValue(42)               // Passa un intero
	printValue("Hello, Go!")     // Passa una stringa
	printValue(3.14)             // Passa un float
	printValue(true)             // Passa un booleano
}
```

#### Utilizzo pratico
interface{} è molto comune in:

- Funzioni generiche: Per scrivere funzioni che possono accettare diversi tipi di dati.
- Ritorni da funzioni: Quando non è possibile determinare in anticipo il tipo di valore da restituire.
- Strutture dati generiche: Come stack o code, che possono contenere valori di vari tipi.

#### Limitazioni
- Controllo dei tipi: Poiché interface{} non offre alcun controllo di tipo durante la compilazione, potresti ricevere errori di runtime se non gestisci correttamente le asserzioni di tipo.
- Prestazioni: L'uso eccessivo di interface{} può avere un impatto sulle prestazioni a causa della necessità di effettuare asserzioni di tipo e di gestire la memoria in modo dinamico.


## TYPE ASSERTION
In Go, la type assertion (asserzione di tipo) è un meccanismo che consente di ottenere il valore effettivo di un'interfaccia e di verificarne il tipo sottostante. È particolarmente utile quando si lavora con l'interfaccia vuota interface{}, poiché non si conosce in anticipo quale tipo di valore sarà contenuto.

```go
value := x.(T)
```

Dove:
- x è l'interfaccia da cui si desidera estrarre il valore.
- T è il tipo a cui si desidera fare l'asserzione.
- value sarà del tipo T se l'asserzione ha successo.

#### Type Assertion con Gestione degli Errori
Per gestire gli errori di asserzione, puoi utilizzare una forma di asserzione di tipo che restituisce anche un secondo valore booleano:

```go
value, ok := x.(T)
```
- value conterrà il valore estratto se l'asserzione ha successo.
- ok sarà true se l'asserzione ha avuto successo, altrimenti sarà false.

