package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var FileName = "tasks.json"

func fileExists(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

func getTaskList(createFileIfNotExists bool) (*[]Task, error) {
	var taskList []Task

	err := fileExists(FileName)
	if err != nil && createFileIfNotExists {
		_, err := os.Create(FileName)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, errors.New("file does not exist. Please add a new task.")
	} else {
		file, err := os.ReadFile(FileName)
		if err != nil {
			return nil, fmt.Errorf("error reading from File, %s", err.Error())
		}
		err = json.Unmarshal(file, &taskList)
		if err != nil {
			return nil, fmt.Errorf("error marshalling from File ,%s", err.Error())
		}
	}

	return &taskList, nil
}

func writeTaskList(taskList *[]Task) error {

	dataBytes, err := json.Marshal(*taskList)
	if err != nil {
		return fmt.Errorf("format string %s", err.Error())
	}
	err = os.WriteFile(FileName, dataBytes, 0644)
	if err != nil {
		return fmt.Errorf("format string %s", err.Error())
	}
	return nil
}
func AddTask(task string) (int64, error) {
	var newID int64
	taskList, err := getTaskList(true)
	if err != nil {
		return 0, err
	}
	if len(*taskList) == 0 {
		newID = 1
	} else {
		newID = (*taskList)[len(*taskList)-1].ID + 1
	}

	taskJSON := Task{}
	taskJSON.ID = newID
	taskJSON.Description = task
	taskJSON.Status = "todo"
	taskJSON.CreatedAt = time.Now().UTC()
	taskJSON.UpdatedAt = time.Now().UTC()

	*taskList = append(*taskList, taskJSON)
	err = writeTaskList(taskList)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func UpdateTask(description string, taskID int64) (int, error) {
	// update task in the file
	taskList, err := getTaskList(false)
	if err != nil {
		return 0, err
	}
	updated := false
	for index, _ := range *taskList {
		if (*taskList)[index].ID == taskID {
			(*taskList)[index].Description = description
			(*taskList)[index].UpdatedAt = time.Now().UTC()
			updated = true
			break
		}
	}

	if !updated {
		return 0, fmt.Errorf("Task ID (%d) not found", taskID)
	}

	err = writeTaskList(taskList)
	if err != nil {
		return 0, err
	}

	return int(taskID), nil
}

func DeleteTask(taskID int64) (int, error) {
	// delete task in the file
	newTaskList := []Task{}
	taskList, err := getTaskList(false)
	if err != nil {
		return 0, err
	}
	deleted := false
	for index, task := range *taskList {
		if (*taskList)[index].ID != taskID {
			newTaskList = append(newTaskList, task)
		} else {
			deleted = true
		}
	}
	if !deleted {
		return 0, fmt.Errorf("Task ID (%d) not found", taskID)
	}
	err = writeTaskList(&newTaskList)
	if err != nil {
		return 0, err
	}

	return int(taskID), nil
}

func MarkInProgress(taskID int64) (int, error) {
	// make task in progress
	taskList, err := getTaskList(false)
	if err != nil {
		return 0, err
	}
	marked := false
	for index, _ := range *taskList {
		if (*taskList)[index].ID == taskID {
			(*taskList)[index].Status = "in-progress"
			(*taskList)[index].UpdatedAt = time.Now().UTC()
			marked = true
			break
		}
	}
	if !marked {
		return 0, fmt.Errorf("Task ID (%d) not found", taskID)
	}

	err = writeTaskList(taskList)
	if err != nil {
		return 0, err
	}
	return int(taskID), nil
}

func MarkDone(taskID int64) (int64, error) {
	// mark task as done
	taskList, err := getTaskList(false)
	if err != nil {
		return 0, err
	}
	marked := false
	for index, _ := range *taskList {
		if (*taskList)[index].ID == taskID {
			(*taskList)[index].Status = "done"
			(*taskList)[index].UpdatedAt = time.Now().UTC()
			marked = true
			break
		}
	}

	if !marked {
		return 0, fmt.Errorf("Task ID (%d) not found", taskID)
	}
	err = writeTaskList(taskList)
	if err != nil {
		return 0, err
	}
	return taskID, nil
}

func GetList(status string) ([]Task, error) {
	newTaskList := []Task{}
	taskList, err := getTaskList(false)
	if err != nil {
		return nil, err
	}

	for index, task := range *taskList {
		if status == "" {
			newTaskList = append(newTaskList, task)
		} else if (*taskList)[index].Status == status {
			newTaskList = append(newTaskList, task)
		}

	}
	return newTaskList, nil

}
func PrintTask(task Task) {

	fmt.Printf("ID : %d,Task : %s, Status : %s, Created : %s, Updated :%s", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println()
}
