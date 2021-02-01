package models

import (
	"errors"
	"fmt"
)

// Task DB Model
type Task struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	Day      string `json:"day"`
	Reminder bool   `json:"reminder"`
}

// Tablename for task
func (Task) Tablename() string {
	return "task"
}

// GetTask by Id
func GetTask(id int) (t *Task, err error) {
	if id == 0 {
		return t, errors.New("Id not informed")
	}
	t = &Task{Id: id}
	err = DB.Model(Task{Id: id}).Take(t).Error
	fmt.Println(err)
	return
}

//GetTasks get all tasks
func GetTasks() (t []Task, err error) {
	err = DB.Find(&t).Error
	return
}

// AddTask Add a task
func AddTask(t *Task) (err error) {
	if len(t.Text) == 0 {
		return errors.New("Text not informed")
	}
	err = DB.Model(Task{}).Save(t).Error
	return
}

// UpdateTask update a task
func UpdateTask(t *Task) (err error) {
	if len(t.Text) == 0 {
		return errors.New("Text not informed")
	}
	err = DB.Save(t).Error
	return
}

// DeleteTask delete a task
func DeleteTask(id int) (err error) {
	if id == 0 {
		return errors.New("Id not informed")
	}
	err = DB.Delete(Task{Id: id}).Error
	return
}
