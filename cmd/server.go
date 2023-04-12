package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	a, err := strconv.Atoi(r.FormValue("a"))
	writer := bufio.NewWriter(w)
	if err != nil {
		writer.WriteString("Неверный параметр a")
		if writer.Flush() != nil {
			fmt.Println("Не получилось записать ответ")
		}

		return
	}
	b, err := strconv.Atoi(r.FormValue("b"))
	if err != nil {
		writer.WriteString("Неверный параметр b")
		if writer.Flush() != nil {
			fmt.Println("Не получилось записать ответ")
		}

		return
	}
	sum := a + b
	fmt.Println(sum)
	writer.WriteString("Сумма: " + strconv.Itoa(sum))
	if writer.Flush() != nil {
		fmt.Println("Не получилось записать ответ")
	}
}

func main() {
	fmt.Println("Сервер работает...")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}