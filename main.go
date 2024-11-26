package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	TaskStringerFormat = "|- Content: %s\n|- Is Done: %t\n|- Created: %s"
	TaskTimeFormat     = "Jan 2 2006 @ 3:4 PM"
)

type task struct {
	content string
	isDone  bool
	created time.Time
}

func (t task) String() string {
	return fmt.Sprintf(TaskStringerFormat,
		t.content,
		t.isDone,
		t.created.Format(TaskTimeFormat),
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
		content: scanner.Text(),
		isDone:  false,
		created: time.Now(),
	}

	tasksList = append(tasksList, newTask)
}
