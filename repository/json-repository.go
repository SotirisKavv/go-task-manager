package repository

import (
	"encoding/json"
	"os"

	"taskmanager/model"
	"taskmanager/utils"
)

type JSONRepository struct {
	filepath string
}

func NewJSONRepository() *JSONRepository {
	return &JSONRepository{
		filepath: "store/tasks.json",
	}
}

func (r *JSONRepository) Load(orderBy string) ([]model.Task, error) {
	file, err := os.Open(r.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []model.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}

	return utils.SortTasks(tasks, orderBy)
}

func (r *JSONRepository) Save(task model.Task) error {
	tasks, err := r.Load("Id")
	if err != nil {
		return err
	}

	tasks = append(tasks, task)
	return r.SaveAll(tasks)
}

func (r *JSONRepository) Update(id int, fieldname string, value any) error {
	tasks, err := r.Load("Id")
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if id == task.Id {
			utils.SetField(&task, fieldname, value)
			tasks[i] = task
			break
		}
	}

	r.SaveAll(tasks)
	return nil
}

func (r *JSONRepository) SaveAll(tasks []model.Task) error {
	file, err := os.Create(r.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

func (r *JSONRepository) Delete(id int) error {
	tasks, err := r.Load("Id")
	if err != nil {
		return err
	}

	var updatedTasks []model.Task
	for _, task := range tasks {
		if task.Id != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	r.SaveAll(updatedTasks)
	return nil
}

func (r *JSONRepository) GenerateId() int {
	tasks, _ := r.Load("Id")
	if len(tasks) == 0 {
		return 0
	}

	return tasks[len(tasks)-1].Id + 1
}
