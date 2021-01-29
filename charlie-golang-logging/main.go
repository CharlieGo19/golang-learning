package main

import(
	"log"
	"os"
	"math/rand"
	"fmt"
	"runtime/trace"
)

type messageType int

	// There is there are three basic error levels - though you can set your own.
	// Information - Tell user something or that something has occured. i.e. network details, os details, etc.
	// Warning - A little more serious, something bad can happen, but program will run fine.
	// Error - More serious! Something you expected to happen, didn't, therefore requires attention.
	// Fatal - Yeah we are screwed, need to close program immediately.

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)


func main() {
	//writeLog(INFO, "This is some information lol.")
	//writeLog(FATAL, "Absolute ball breaking error! Damn I'm done for boys.")
	TraceDemo()
}

func writeLog(messagetype messageType, message string) {
	logFile, err := os.OpenFile("test-log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0740)
	// Close file once enclosing function finishes executing, in this example, writeLog.
	defer logFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	switch messagetype {
	case INFO:
		logger := log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(logFile, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}

func TraceDemo() {
	// To view trace use command: go tool trace $traceFileName
	traceFile, err := os.Create("trace.out")
	
	if err != nil {
		log.Fatalf("Could not create trace file! %v\n", err)
	}
	defer func ()  {
		if err := traceFile.Close(); err != nil {
			log.Fatalf("Failed to close trace file! %v\n", err)
		}
	}()

	if err := trace.Start(traceFile); err != nil {
		log.Fatalf("We failed to start the trace: %d\n", err)
	}

	defer trace.Stop()

	AddRandomNumbers()
}

func AddRandomNumbers() {
	var firstNumber int = rand.Intn(100)
	var secondNumber int = rand.Intn(100)

	fmt.Println(firstNumber + secondNumber)
}