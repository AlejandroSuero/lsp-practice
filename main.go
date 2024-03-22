package main

import (
	"bufio"
	"log"
	"lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("/logs/log.txt")
	logger.Println("Logger started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {}

func getLogger(filename string) *log.Logger {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logfile, err := os.OpenFile(pwd+filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, this is not a good file my guy")
	}
	return log.New(logfile, "[lsp-practice]", log.Ldate|log.Ltime|log.Lshortfile)
}
