package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	TaskStringerFormat = "|- Content: %s\n|- Is Done: %t\n|- Created: %s"
	TaskTimeFormat     = "Jan 2 2006 @ 3:4 PM"
)

type task struct {
	Content string    `json:"content"`
	IsDone  bool      `json:"isDone"`
	Created time.Time `json:"created"`
}

func (t task) String() string {
	return fmt.Sprintf(TaskStringerFormat,
		t.Content,
		t.IsDone,
		t.Created.Format(TaskTimeFormat),
	)
}

var tasksList []task

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("> Please enter your task:")
	scanner.Scan()

	createTask(scanner)

	fmt.Println("--- User Tasks List ---")
	for _, t := range tasksList {
		fmt.Println(t)
		fmt.Println("---")
	}

}

func createTask(scanner *bufio.Scanner) {
	newTask := task{
		Content: scanner.Text(),
		IsDone:  false,
		Created: time.Now(),
	}

	file, err := os.OpenFile("taskDB.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("X Can't open or create file -> ", err)
		return

	}

	defer file.Close()

	data, err := json.Marshal(newTask)
	if err != nil {
		fmt.Println("X Can't marshal newTask -> ", err)
	}

	data = append(data, []byte("\n")...)

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("X Can't write to file -> ", err)
	}

	tasksList = append(tasksList, newTask)
}
