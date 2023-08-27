package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shyam0507/to-do-cmd/db"
	"github.com/shyam0507/to-do-cmd/model"
	"github.com/shyam0507/to-do-cmd/service"
)

func main() {
	log.Print("Start the todo app...")

	ser := service.NewService(db.New())

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		cmd := strings.Split(scanner.Text(), " ")

		// log.Printf("Operation %s", cmd[0])
		switch strings.ToLower(cmd[0]) {
		case "add":
			ser.Add(&model.ToDo{Label: strings.Join(cmd[1:], " ")})
		case "delete":
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Printf("Please provide a valid Id - %v\n", id)
				continue
			}
			if err := ser.Delete(id); err != nil {
				fmt.Printf("Id %d not found\n", id)
			}
		case "update":
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Printf("Please provide a valid Id - %v\n", id)
				continue
			}
			ser.Update(&model.ToDo{Id: id, Label: strings.Join(cmd[2:], " ")})
		case "get":
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Printf("Please provide a valid Id - %v\n", cmd[1])
				continue
			}
			todo := ser.GetById(id)
			if todo == nil {
				fmt.Printf("todo with id :%d not found\n", id)
			} else {
				fmt.Printf("%d    %s\n", todo.Id, todo.Label)
			}
		case "getall":
			todos := ser.GetAll()
			if len(todos) == 0 {
				fmt.Printf("No Todo's available\n")
				continue
			}
			for _, todo := range todos {
				fmt.Printf("%d    %s\n", todo.Id, todo.Label)
			}
		default:
			fmt.Printf("Invalid operation %s", cmd[0])
		}

	}

}
