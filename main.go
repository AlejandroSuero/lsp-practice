package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"

	"lsp/lsp"
	"lsp/rpc"
)

func main() {
	logger := getLogger("/logs/log.txt")
	logger.Println("Logger started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Text()
		method, content, err := rpc.DecodeMessage([]byte(msg))
		if err != nil {
			logger.Printf("[ERROR]: %s", err)
			continue
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Received msg with method: %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("[ERROR]: Couldn't parse this, %s", err)
		}
		logger.Printf(
			"Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version,
		)
	}
}

func getLogger(filename string) *log.Logger {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := pwd + "/logs"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	logfile, err := os.OpenFile(pwd+filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, this is not a good file my guy")
	}
	return log.New(logfile, "[lsp-practice]", log.Ldate|log.Ltime|log.Lshortfile)
}
