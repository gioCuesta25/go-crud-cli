package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gioCuesta25/go-cli-crud/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var taskList []tasks.Task

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &taskList)

		if err != nil {
			panic(err)
		}
	} else {
		taskList = []tasks.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "list":
		tasks.ListTasks(taskList)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Cual es tu tarea")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		updatedTasks := tasks.CreateTask(taskList, name)
		tasks.SaveTasks(file, updatedTasks)
		tasks.ListTasks(updatedTasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar un ID para eliminar")
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			panic(err)
		}

		updatedTasks := tasks.DeleteTask(taskList, id)
		tasks.SaveTasks(file, updatedTasks)
		tasks.ListTasks(updatedTasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Debe proporcionar el ID de la tarea que se completará.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un número entero.")
			return
		}
		updatedTasks := tasks.CompleteTask(taskList, id)
		tasks.SaveTasks(file, updatedTasks)
		tasks.ListTasks(updatedTasks)
	}
}

func printUsage() {
	fmt.Println("Uso: go-cli-crud [list|add|complete|delete]")
}
