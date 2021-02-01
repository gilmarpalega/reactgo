package handlers

import (
	"fmt"
	"reactgo/models"
	"runtime/debug"
	"strconv"

	"github.com/gofiber/fiber"
)

// CheckError shortcut to deal errors
func CheckError(c *fiber.Ctx, err error) bool {
	if err != nil {
		c.SendStatus(500) // Não encontrado
		debug.PrintStack()
		return true
	}
	return false
}

// GetId from URL
func GetId(c *fiber.Ctx) int {
	id, err := strconv.Atoi(c.Params("id"))

	if CheckError(c, err) {
		id = 0
		debug.PrintStack()
	}
	return id
}

// GetTask by id
func GetTask(c *fiber.Ctx) error {
	id := GetId(c)
	task, err := models.GetTask(id)

	if CheckError(c, err) {
		return err
	}

	c.JSON(task)
	return nil
}

// GetTasks get all tasks
func GetTasks(c *fiber.Ctx) error {
	tasks, err := models.GetTasks()
	if CheckError(c, err) {
		return err
	}
	c.JSON(tasks)
	return nil
}

// AddTask adds a task
func AddTask(c *fiber.Ctx) error {
	task := new(models.Task)
	err := c.BodyParser(task)

	if CheckError(c, err) {
		return err
	}
	err = models.AddTask(task)

	if CheckError(c, err) {
		return err
	}
	c.JSON(task)
	return nil
}

// UpdateTask updates a task
func UpdateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	err := c.BodyParser(task)
	id := GetId(c)

	task.Id = id

	if CheckError(c, err) {
		return err
	}
	err = models.UpdateTask(task)

	if CheckError(c, err) {
		return err
	}
	c.JSON(task)
	return nil
}

// DeleteTask removes a task
func DeleteTask(c *fiber.Ctx) error {
	id := GetId(c)
	err := models.DeleteTask(id)

	if CheckError(c, err) {
		return err
	}
	c.Send([]byte(fmt.Sprintf("Task id %v deleted", id)))
	return nil
}
