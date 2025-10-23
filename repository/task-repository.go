package repository

import (
	"taskmanager/model"
)

type TaskRepository interface {
	Load(orderBy string) ([]model.Task, error)
	Save(task model.Task) error
	Update(id int, fieldname string, value any) error
	Delete(id int) error
	GenerateId() int
}

func GetTaskRepository(repoType string) TaskRepository {
	switch repoType {
	case "csv":
		return NewCSVRepository()
	case "json":
		return NewJSONRepository()
	default:
		return nil
	}
}
