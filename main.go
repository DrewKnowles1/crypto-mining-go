package main

import (
	"log"
	"os"
)

//Creating a struct of type application, so i can info/error log to a single object
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	app.getMiner()
}
