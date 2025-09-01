package main

import "fmt"

func main() {
	var things string
	var num int
	var queue [5]string

	for {
		i, _ := fmt.Scanln(&things, &num)
		if i == 1 {
			switch {
			case things == "очередь":
				for i := 0; i < 5; i++ {
					if queue[i] != "" {
						fmt.Printf("%d. %s\n", i+1, queue[i])
					} else {
						fmt.Printf("%d. -\n", i+1)
					}
				}

			case things == "количество":
				free := 0
				for i := range queue {
					if queue[i] != "" {
						free++
					}
				}
				fmt.Printf("Осталось свободных мест: %d\n", 5-free)
				fmt.Printf("Всего человек в очереди: %d\n", free)

			case things == "конец":
				for i := 0; i < 5; i++ {
					if queue[i] != "" {
						fmt.Printf("%d. %s\n", i+1, queue[i])
					} else {
						fmt.Printf("%d. -\n", i+1)
					}
				}
				return

			default:
				fmt.Println("некорректный ввод")
			}
		}

		if i == 2 {
			if num < 1 || num > 5 {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)

			} else {
				isFull := true
				for _, v := range queue {
					if v == "" {
						isFull = false
						break
					}
				}

				if isFull {
					fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)

				} else if queue[num-1] != "" {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)

				} else {
					queue[num-1] = things

				}
			}

		}
	}
}
