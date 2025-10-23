package repository

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"taskmanager/model"
	"taskmanager/utils"
)

type CSVRepository struct {
	filepath string
}

func NewCSVRepository() *CSVRepository {
	return &CSVRepository{
		filepath: "store/tasks.csv",
	}
}

func (r *CSVRepository) Load(orderBy string) ([]model.Task, error) {
	file, err := os.Open(r.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	for _, row := range records {
		id, _ := strconv.Atoi(row[0])
		status := model.Status(row[2])
		dueDate, _ := time.Parse(utils.DateLayout, row[3])

		task := model.Task{
			Id:      id,
			Title:   row[1],
			Status:  status,
			DueDate: dueDate,
		}

		tasks = append(tasks, task)
	}

	return utils.SortTasks(tasks, orderBy)
}

func (r *CSVRepository) Save(task model.Task) error {
	file, err := os.OpenFile(r.filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		strconv.Itoa(task.Id),
		task.Title,
		string(task.Status),
		task.DueDate.Format(utils.DateLayout),
	}

	writer.Write(record)
	return nil
}

func (r *CSVRepository) Update(id int, fieldname string, value any) error {
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

func (r *CSVRepository) SaveAll(tasks []model.Task) error {
	file, err := os.Create(r.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.Id),
			task.Title,
			string(task.Status),
			task.DueDate.Format(utils.DateLayout),
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func (r *CSVRepository) Delete(id int) error {
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

func (r *CSVRepository) GenerateId() int {
	tasks, _ := r.Load("Id")
	if len(tasks) == 0 {
		return 0
	}

	return tasks[len(tasks)-1].Id + 1
}
