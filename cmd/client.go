package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число: ")

	for scanner.Scan() {

		input := scanner.Text()

		firstNum, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if firstNum > 2147483547 {
			log.Fatal("Слишком большое число")
		}

		rand.Seed(time.Now().UnixNano())
		secondNum := rand.Intn(100)

		fmt.Println("Пусть второе число будет... Например", secondNum)

		baseURL := "http://127.0.0.1:8080/"
		params := url.Values{}
		params.Add("a", strconv.Itoa(firstNum))
		params.Add("b", strconv.Itoa(secondNum))
		reqURL := baseURL + "?" + params.Encode()

		resp, err := http.Get(reqURL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		fmt.Println("Отправлен GET запрос:", reqURL)

		buf := make([]byte, 32)

		resp.Body.Read(buf)

		if len(buf) > 0 {
			fmt.Println("Получен ответ:", resp.Status, ":", string(buf))
		} else {
			fmt.Println("Ответ не был получен.")
		}
		fmt.Print("Введите новое число: ")
	}
}
