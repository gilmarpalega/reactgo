package main

import (
	"fmt"
	"reactgo/models"
)

func main() {
	models.Setup()

	tasks, err := models.GetTasks()

	if err != nil {
		panic(err.Error())
	}
	for _, task := range tasks {
		fmt.Printf("%+v\n", task)
	}

}
