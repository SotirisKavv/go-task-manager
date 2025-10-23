# 📋 CLI Task Manager (Go) - **Intermediate Level**

A modular task management CLI application with persistent storage supporting both CSV and JSON formats. Demonstrates repository pattern, interface design, and clean architecture principles in Go.

---

## 🚀 What is this?

A command-line task manager showcasing clean architecture with pluggable storage backends. Features task CRUD operations, due date management, and demonstrates advanced Go patterns like interfaces and dependency injection.

---

## ✨ Features

- **CRUD Operations:** Add, list, complete, and delete tasks
- **Persistent Storage:** CSV and JSON storage backends
- **Repository Pattern:** Pluggable storage interface design
- **Date Handling:** Due date parsing and sorting
- **Status Management:** Task completion state tracking

---

## 🦄 Go Skills Demonstrated

- **Interface Design:** `TaskRepository` abstraction for storage
- **Design Patterns:** Repository pattern and factory method
- **Multiple Storage:** CSV and JSON serialization/deserialization
- **Time Handling:** Date parsing, formatting, and sorting
- **Struct Methods:** Custom types with behavior
- **Error Handling:** Comprehensive validation and error management

---

## 🏗️ Architecture

```
cli-task-manager/
├── model/              # Task and Status definitions
├── repository/         # Storage interfaces and implementations
├── utils/              # Helper functions and utilities
└── taskmanager.go      # CLI interface and main logic
```

**Repository Pattern:** Abstracts storage logic allowing easy switching between CSV, JSON, or future database backends.

---

## 🛠️ Usage

```sh
# Add tasks
go run taskmanager.go add "Update CV" 2025-09-30
go run taskmanager.go add "Learn Go" 2025-08-15

# List all tasks (sorted by due date)
go run taskmanager.go list

# Mark task as completed
go run taskmanager.go done 1

# Delete task
go run taskmanager.go delete 2
```

---

## 🎯 Learning Objectives

This project demonstrates:
- **Clean Architecture:** Separation of concerns with clear layers
- **Interface Design:** Abstract storage with multiple implementations
- **Data Persistence:** File-based storage with multiple formats
- **CLI Development:** Command routing and argument processing

**Difficulty:** ⭐⭐⭐ Intermediate - Architecture patterns and interface design

---

**Author:** IAmSotiris
