package main

import (
	"bytes"
	"os"
	"task-tracker.com/cmd"
	"testing"
)

var testFileName = "tasks_test.json"

func setup() {
	// Create or truncate the test file before each test
	file, err := os.CreateTemp("", testFileName)
	if err != nil {
		panic(err)
	}
	file.Close()
}
func Test_GetTaskTrackerCLI(t *testing.T) {

	tests := []struct {
		name   string
		osArgs []string
		output string
	}{
		{"Add First Task", []string{"cmd", "add", `First Task`}, "Task added successfully (ID: 1)"},
		{"First Added Task", []string{"cmd", "add", `First Added Task`}, "Task added successfully (ID: 2)"},
		{"First Non Updated Task", []string{"cmd", "add", `First Non Updated Task`}, "Task added successfully (ID: 3)"},
		{"First Non Deleted Task", []string{"cmd", "add", `First Non Deleted Task`}, "Task added successfully (ID: 4)"},
		{"First Test Task", []string{"cmd", "add", `First Test Task`}, "Task added successfully (ID: 5)"},
		{"First Updated Task", []string{"cmd", "update", "3", `First Updated Task`}, "Task updated successfully (ID: 3)"},
		{"Deletd Task", []string{"cmd", "delete", "4"}, "Task deleted successfully (ID: 4)"},
		{"mark-in-progress", []string{"cmd", "mark-in-progress", "1"}, "Task updated to in-progress successfully (ID: 1)"},
		{"mark-done", []string{"cmd", "mark-done", "2"}, "Task done successfully (ID: 2)"},
		{"list-all", []string{"cmd", "list"}, ""},
		{"list-done", []string{"cmd", "list", "done"}, ""},
		{"list-todo", []string{"cmd", "list", "todo"}, ""},
		{"list-in-progress", []string{"cmd", "list", "in-progress"}, ""},
	}

	//setup()
	defer os.Remove(testFileName)
	cmd.FileName = testFileName
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			out = bytes.NewBuffer(nil)

			os.Args = tt.osArgs
			taskTracker()
			if actual := out.(*bytes.Buffer).String(); actual != tt.output {
				t.Errorf("expected '%s', but got '%s'", tt.output, actual)
			}
		})
	}
}

func Test_GetIncorrectInputs(t *testing.T) {

	tests := []struct {
		name   string
		osArgs []string
		output string
	}{

		{"Wrong Argument --adding", []string{"cmd", "adding", `First Task`}, "Invalid command"},
		{"Wrong Argument --updating", []string{"cmd", "updating", `First Task`}, "Invalid command"},
		{"Wrong Argument --deleting", []string{"cmd", "deleting", `First Task`}, "Invalid command"},
		//{"Wrong Argument --adding", []string{"cmd", "add"}, "Please provide task ID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out = bytes.NewBuffer(nil)
			os.Args = tt.osArgs
			taskTracker()
			if actual := out.(*bytes.Buffer).String(); actual != tt.output {
				t.Errorf("expected '%s', but got '%s'", tt.output, actual)
			}
		})
	}
}

// func Suite
