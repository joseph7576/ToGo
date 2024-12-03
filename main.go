package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	TaskStringerFormat = "|- Content: %s\n|- Is Done: %t\n|- Created: %s\n"
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
	loadTasks()

	fmt.Println(tasksList)

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

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("X Can't clost the file -> ", err)
		}
	}()

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

func loadTasks() {

	file, err := os.Open("taskDB.txt")
	if err != nil {
		fmt.Println("X Can't open or create file -> ", err)

		return
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("X Can't clost the file -> ", err)

			return
		}
	}()

	dataBuffer := make([]byte, 8*1024)

	_, err = file.Read(dataBuffer)
	if err != nil {
		fmt.Println("X Can't read the file -> ", err)

		return
	}

	dataString := string(dataBuffer)

	taskSlice := strings.Split(dataString, "\n")
	for _, t := range taskSlice {
		taskStruct := task{}

		if t[0] != '{' && t[len(t)-1] != '}' {
			continue
		}

		err = json.Unmarshal([]byte(t), &taskStruct)
		if err != nil {
			fmt.Println("X Cna't unmarshal task record to task struct -> ", err)

			return
		}

		tasksList = append(tasksList, taskStruct)
	}

}
