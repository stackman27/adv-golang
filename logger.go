package main

import (
	"log"
	"os"
)

func main_log() {

	//parseEdges()

	log.SetFlags(log.Ldate | log.Lshortfile)
	// log.Println("Sishir")
	//log.Fatal("FATALL")

	// log.Panic("PANICCCC")

	// Output the log
	file, _ := os.Create("file.log")
	log.SetOutput(file)

	log.Println("Hello World, log this in file")
	file.Close()

	log.SetOutput(os.Stdout)
	log.Println("More information")

	// Common Loggers
	flags := log.LstdFlags | log.Lshortfile

	infoLogger := log.New(os.Stdout, "INFO: ", flags)
	warnLogger := log.New(os.Stdout, "WARN: ", flags)
	errorLogger := log.New(os.Stdout, "ERROR: ", flags)

	infoLogger.Println("THis in info log")
	warnLogger.Println("This is warn log")
	errorLogger.Println("This is errror log")

}
