# todo-cli

A simple task tracker CLI written in Go.

## Requirements

- Go 1.22+

## Run

From the project root:

```bash
go run . <command> [args]
```

Examples:

```bash
go run . add "buy milk"
go run . list
go run . mark-in-progress 1
go run . mark-done 1
go run . update 1 "buy milk and eggs"
go run . delete 1
```

## Commands

### Add

```bash
go run . add "task description"
```

Creates a new task with status `todo`.

### Update

```bash
go run . update <id> "new description"
```

Updates the task description.

### Delete

```bash
go run . delete <id>
```

Deletes a task by ID.

### Mark In Progress

```bash
go run . mark-in-progress <id>
```

Sets task status to `in-progress`.

### Mark Done

```bash
go run . mark-done <id>
```

Sets task status to `done`.

### List

List all tasks:

```bash
go run . list
```

List by status:

```bash
go run . list todo
go run . list in-progress
go run . list done
```

## Data Storage

Tasks are persisted in `tasks.json` at the project root.

## Build Binary

```bash
go build -o todo-cli .
./todo-cli list
```
