package main

import (
	"reactgo/models"
)

func main() {
	models.Setup()

	models.GetDB().AutoMigrate(&models.Task{})

	tasks := []models.Task{
		models.Task{Text: "School Meeting", Day: "6th Feb 08:00", Reminder: true},
		models.Task{Text: "Dog to Petshop", Day: "8th Feb 14:00", Reminder: false},
		models.Task{Text: "Date next week", Day: "12th Feb 20:00", Reminder: true},
		models.Task{Text: "Buy a guitar", Day: "13th Feb 13:00", Reminder: true},
	}

	for _, task := range tasks {
		models.AddTask(&task)
	}

}
