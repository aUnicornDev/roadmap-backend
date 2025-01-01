package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"task-tracker.com/cmd"
)

var out io.Writer = os.Stdout

func main() {
	taskTracker()
}
func readTaskID() int64 {
	if len(os.Args) < 3 {
		fmt.Fprint(out, "Please provide task ID")
		os.Exit(1)
	}
	intTaskID, err := strconv.ParseInt(os.Args[2], 10, 8)
	if err != nil {
		fmt.Fprint(out, "Invalid task ID")
		os.Exit(1)
	}
	return intTaskID

}

func taskTracker() {
	command := os.Args[1]
	switch command {
	case "add":
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
		taskID := readTaskID()
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
		taskID := readTaskID()
		deletedTaskID, err := cmd.DeleteTask(taskID)
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		fmt.Fprintf(out, `Task deleted successfully (ID: %d)`, deletedTaskID)
		break
	case "mark-in-progress":
		taskID := readTaskID()
		updatedStatusTaskID, err := cmd.MarkInProgress(taskID)
		if err != nil {
			fmt.Fprint(out, err.Error())
			return
		}
		fmt.Fprintf(out, `Task updated to in-progress successfully (ID: %d)`, updatedStatusTaskID)

		break
	case "mark-done":
		taskID := readTaskID()
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
