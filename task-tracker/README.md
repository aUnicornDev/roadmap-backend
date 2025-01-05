
# Task Tracker CLI

Task Tracker CLI is a command-line tool for managing tasks. You can add, update, delete, and list tasks, as well as mark them as in-progress or done.
[Project Summary](https://roadmap.sh/projects/task-tracker)

## Installation

To install the Task Tracker CLI, clone the repository and build the project:

```sh
git clone https://github.com/yourusername/task-tracker.git
cd task-tracker
go build -o task-cli
```

## Usage

Here are the available commands for the Task Tracker CLI:

### Adding a new task

```sh
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Updating a task

```sh
./task-cli update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully (ID: 1)
```

### Deleting a task

```sh
./task-cli delete 1
# Output: Task deleted successfully (ID: 1)
```

### Marking a task as in-progress

```sh
./task-cli mark-in-progress 1
# Output: Task updated to in-progress successfully (ID: 1)
```

### Marking a task as done

```sh
./task-cli mark-done 1
# Output: Task done successfully (ID: 1)
```

### Listing all tasks

```sh
./task-cli list
# Output: List of all tasks
```

### Listing tasks by status

```sh
./task-cli list done
# Output: List of tasks with status 'done'

./task-cli list todo
# Output: List of tasks with status 'todo'

./task-cli list in-progress
# Output: List of tasks with status 'in-progress'
```

## Error Handling

The CLI will provide appropriate error messages for invalid commands or missing arguments. For example:

```sh
./task-cli add
# Output: Please provide task name

./task-cli update
# Output: Please provide task ID

./task-cli delete abc
# Output: Invalid task ID

./task-cli delete 999
# Output: Task ID (999) not found
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
```
