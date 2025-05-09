package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting in a few...... ")
}

type Task struct {
	id            int32
	title         string
	completed     bool
	creation_time time.Time
}

func ListTask(status status) error {
	tasks, err := ReadTaskFromList()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No Tasks Found!!!!")
		return nil
	}

	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks == tasks
	case STATUS_TODO:
		for _, task := range tasks {
			if task.Status == STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case STATUS_DONE:
		for _, task := range tasks {
			if task.Status == STATUS_DONE {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}
	fmt.Println("Tasks %s", tasks)
	return nil
}

func AddTasks(description string) error {
	tasks, err := ReadTaskFromFile()
	if err != nil {
		return err
	}

	var newTaskID int32
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskID = lastTask.ID + 1
	} else {
		newTaskID = 0
	}

	task := NewTask(newTaskID, description)
	tasks = append(tasks, *task)
	fmt.Println("Task added succesfully %s", task)
	return writeTaskToFile(task)
}
