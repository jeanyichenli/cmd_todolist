package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := []string{}
	complete := []string{}

	for {
		fmt.Printf("TO-DO type: ")
		if scanner.Scan() {
			t := strings.TrimSpace(scanner.Text())

			switch t {
			case "add":
				fmt.Printf("TODO item: ")
				scanner.Scan()
				item := scanner.Text()
				tasks = append(tasks, item)

			case "remove":
				var id int
				var err error
				for {
					fmt.Printf("Enter task number to remove: ")
					scanner.Scan()
					id, err = strconv.Atoi(scanner.Text())
					if err != nil {
						fmt.Println("Error:", err.Error(), "Please re-input task number.")
					} else if id < 1 || id > len(tasks) {
						fmt.Println("Task number is out of range. Please re-input.")
					} else {
						break
					}
				}
				fmt.Printf("The item you want to remove is %v. %s. \n", id, tasks[id-1])
				fmt.Printf("Ensure you want to delete:[Y/n] ")
				scanner.Scan()

				switch scanner.Text() {
				case "Y":
					id--
					tasks = append(tasks[:id], tasks[id+1:]...)
				case "", "n":
				}

			case "complete":
				var id int
				var err error
				for {
					fmt.Printf("Enter task number to set complete: ")
					scanner.Scan()
					id, err = strconv.Atoi(scanner.Text())
					if err != nil {
						fmt.Println("Error:", err.Error(), "Please re-input task number.")
					} else if id < 1 || id > len(tasks) {
						fmt.Println("Task number is out of range. Please re-input.")
					} else {
						break
					}
				}
				fmt.Printf("Mission [%s] is complete.\n", tasks[id-1])
				id--
				complete = append(complete, tasks[id])
				if id == len(tasks) { // tail element
					tasks = tasks[:id]
				} else {
					tasks = append(tasks[:id], tasks[id+1:]...)
				}

			case "list":
				idx := 1
				for _, t := range tasks {
					fmt.Printf("%v. %s\n", idx, t)
					idx++
				}
			case "complete list":
				idx := 1
				for _, t := range complete {
					fmt.Printf("%v. %s\n", idx, t)
					idx++
				}

			case "list all":
				idx := 1
				for _, t := range tasks {
					fmt.Printf("%v. %s\n", idx, t)
					idx++
				}
				fmt.Println("")
				idx = 1
				for _, t := range complete {
					fmt.Printf("(Complete) %v. %s\n", idx, t)
					idx++
				}
			case "exit":
				return
			default:
				fmt.Println("Unknown type.")
			}
		} else {
			break
		}

	}
}
