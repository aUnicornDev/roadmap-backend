package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID     int
	Name   string
	Status string
}

const fileName = "tasks.json"

func fileExists(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddTask(task string) int {
	taskList := []Task{}
	var newID int
	if fileExists(fileName) != nil {
		newID = 1
	} else {
		file, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Errorf("Erro reading from File!")
		}
		json.Unmarshal(file, &taskList)

		// fmt.Print(taskList)
		if len(taskList) == 0 {
			newID = 1
		} else {
			newID = taskList[len(taskList)-1].ID + 1
		}
	}

	taskJSON := Task{}
	taskJSON.ID = newID
	taskJSON.Name = task
	taskJSON.Status = "todo"

	taskList = append(taskList, taskJSON)

	dataBytes, err := json.Marshal(taskList)
	if err != nil {
		fmt.Printf("format string %s", err.Error())
	}
	err = os.WriteFile("tasks.json", dataBytes, 0644)
	if err != nil {
		fmt.Printf("format string %s", err.Error())
	}

	return 1
}

func UpdateTask(task string, taskID string) int {
	// update task in the file
	//

	return 1
}

func DeleteTask(task string) int {
	// delete task in the file
	//
	return 1
}

func MarkInProgress(task string) int {
	// make task in progress
	//
	return 1
}

func MarkDone(task string) int {
	// mark task as done
	//
	return 1
}
