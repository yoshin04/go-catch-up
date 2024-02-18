package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	flags := log.Lshortfile
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", flags)
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)

	warnLogger.Println("warning a")
	// ログを出力した後に強制終了
	errorLogger.Fatalln("critical error")
}
