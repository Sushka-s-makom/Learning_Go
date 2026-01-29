package http

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func Client() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		fmt.Println("Ошибка при запросе:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//читаем первые пять строк для примера
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
