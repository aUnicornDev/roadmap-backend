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
		{"Deleted Task", []string{"cmd", "delete", "4"}, "Task deleted successfully (ID: 4)"},
		{"mark-in-progress", []string{"cmd", "mark-in-progress", "1"}, "Task updated to in-progress successfully (ID: 1)"},
		{"mark-done", []string{"cmd", "mark-done", "2"}, "Task done successfully (ID: 2)"},
		{"list-all", []string{"cmd", "list"}, ""},
		{"list-done", []string{"cmd", "list", "done"}, ""},
		{"list-todo", []string{"cmd", "list", "todo"}, ""},
		{"list-in-progress", []string{"cmd", "list", "in-progress"}, ""},
		{"Update Non-existent Task", []string{"cmd", "update", "999", "Non-existent Task"}, "Task ID (999) not found"},
		{"Delete Non-existent Task", []string{"cmd", "delete", "999"}, "Task ID (999) not found"},
		{"Mark In Progress Non-existent Task", []string{"cmd", "mark-in-progress", "999"}, "Task ID (999) not found"},
		{"Mark Done Non-existent Task", []string{"cmd", "mark-done", "999"}, "Task ID (999) not found"},
	}

	//setup()
	defer os.Remove(testFileName)
	cmd.FileName = testFileName
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			out = bytes.NewBuffer(nil)

			os.Args = tt.osArgs
			main()
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

		{"Invalid Command", []string{"cmd", "adding", `First Task`}, "Invalid command"},
		{"Missing Task ID for Update", []string{"cmd", "update"}, "Please provide task ID"},
		{"Invalid Task ID for Update", []string{"cmd", "update", "abc", "New Description"}, "Invalid task ID"},
		{"Missing Task ID for Delete", []string{"cmd", "delete"}, "Please provide task ID"},
		{"Invalid Task ID for Delete", []string{"cmd", "delete", "abc"}, "Invalid task ID"},
		{"Missing Task ID for Mark In Progress", []string{"cmd", "mark-in-progress"}, "Please provide task ID"},
		{"Invalid Task ID for Mark In Progress", []string{"cmd", "mark-in-progress", "abc"}, "Invalid task ID"},
		{"Missing Task ID for Mark Done", []string{"cmd", "mark-done"}, "Please provide task ID"},
		{"Invalid Task ID for Mark Done", []string{"cmd", "mark-done", "abc"}, "Invalid task ID"},
		{"Invalid List Status", []string{"cmd", "list", "unknown"}, "Invalid status"},
		{"Missing Task Name for Add", []string{"cmd", "add"}, "Please provide task name"},
		{"Missing Task Name for Update", []string{"cmd", "update", "1"}, "Please provide task name"},
		{"No command", []string{"cmd"}, "Please provide any of the command"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out = bytes.NewBuffer(nil)
			os.Args = tt.osArgs
			main()
			if actual := out.(*bytes.Buffer).String(); actual != tt.output {
				t.Errorf("expected '%s', but got '%s'", tt.output, actual)
			}
		})
	}
}
