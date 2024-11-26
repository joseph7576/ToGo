package main

import (
	"fmt"
	"time"
)

type task struct {
	content string
	isDone  bool
	created time.Time
}

func (t task) String() string {
	return fmt.Sprintf("|- Content: %s\n|- Is Done: %t\n|- Created: %s",
		t.content,
		t.isDone,
		t.created.Format("Jan 2 2006 @ 3:4 PM"),
	)
}

func main() {
	var tasksList []task

	t1 := task{
		"Test Content Task 1",
		false,
		time.Now(),
	}

	tasksList = append(tasksList, t1)

	fmt.Println(tasksList)
	fmt.Println(t1)
}
