package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go list.txt")
		return
	}

	lists := args[0]
	file, err := os.Open(lists)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		site := scanner.Text()
		if strings.Contains(checkConfigurationFile(site), "DB_DATABASE") {
			fmt.Printf("\033[32m[+] FOUND: %s/.env\n", site)
		} else {
			fmt.Printf("\033[31m[-] NOT FOUND: %s\n", site)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func checkConfigurationFile(site string) string {
	url := site + "/.env"
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
