package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"task-tracker.com/cmd"
)

var out io.Writer = os.Stdout

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(out, "Please provide any of the command")
		fmt.Println(`
		# Adding a new task
		task-cli add "Buy groceries"
		# Output: Task added successfully (ID: 1)
		
		# Updating and deleting tasks
		task-cli update 1 "Buy groceries and cook dinner"
		task-cli delete 1
		
		# Marking a task as in progress or done
		task-cli mark-in-progress 1
		task-cli mark-done 1
		
		# Listing all tasks
		task-cli list
		
		# Listing tasks by status
		task-cli list done
		task-cli list todo
		task-cli list in-progress`)
		return
	}
	taskTracker()
}
func readTaskID() (int64, error) {
	if len(os.Args) < 3 {
		return 0, errors.New("Please provide task ID")

	}
	intTaskID, err := strconv.ParseInt(os.Args[2], 10, 8)
	if err != nil {
		return 0, errors.New("Invalid task ID")
	}
	return intTaskID, nil

}

func taskTracker() {
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprint(out, "Please provide task name")
			return
		}
		taskName := os.Args[2]
		taskID, err := cmd.AddTask(taskName)
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		fmt.Fprintf(out, `Task added successfully (ID: %d)`, taskID)
		break
	case "update":
		// check for 3 arguments
		taskID, err := readTaskID()
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		if len(os.Args) < 4 {
			fmt.Fprint(out, "Please provide task name")
			return
		}
		taskDescription := os.Args[3]
		updatedTaskID, err := cmd.UpdateTask(taskDescription, taskID)
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		fmt.Fprintf(out, `Task updated successfully (ID: %d)`, updatedTaskID)

		break
	case "delete":
		// check for 3 arguments
		taskID, err := readTaskID()
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		deletedTaskID, err := cmd.DeleteTask(taskID)
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		fmt.Fprintf(out, `Task deleted successfully (ID: %d)`, deletedTaskID)
		break
	case "mark-in-progress":
		taskID, err := readTaskID()
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		updatedStatusTaskID, err := cmd.MarkInProgress(taskID)
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		fmt.Fprintf(out, `Task updated to in-progress successfully (ID: %d)`, updatedStatusTaskID)

		break
	case "mark-done":
		taskID, err := readTaskID()
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		updatedStatusTaskID, err := cmd.MarkDone(taskID)
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		fmt.Fprintf(out, `Task done successfully (ID: %d)`, updatedStatusTaskID)
		break
	case "list":
		var status string
		if len(os.Args) == 2 {
			status = ""
		} else {
			status = os.Args[2]
			if status != "done" && status != "todo" && status != "in-progress" {
				fmt.Fprint(out, "Invalid status")
				return
			}
		}
		tasks, err := cmd.GetList(status)
		if err != nil {
			fmt.Fprint(out, err.Error())
		}
		for _, task := range tasks {
			cmd.PrintTask(task)
		}
		break
	default:
		fmt.Fprint(out, "Invalid command")

	}

}
