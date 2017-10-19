package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io"
	"regexp"
	"time"
	"github.com/aukhatov/MqttService/mqtt"
)

const defaultHttpPort = ":80"

func main()  {
	loggerInitialize()
	log.Println("Connection test URL on tcp://192.168.1.133:1883")
	mqtt.Test("tcp://192.168.1.133:1883")
}

func loggerInitialize() {
	timeSuffix := time.Now().Format("2006-01-02")
	logFileName := "access-" + timeSuffix + ".log"
	accessLogFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, accessLogFile)
	log.SetOutput(multiWriter)
}

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("%v %v %v", request.Proto, request.Method, request.RequestURI)
	fmt.Fprint(writer, "The web-service is working by Go!")
}

func readCommandLineArguments() []string {
	args := os.Args[1:]
	return args
}

func parseHttpPort(args []string) string {
	var httpPort = defaultHttpPort
	var argumentPattern = "--http.port="
	var valueStartIndex = 12

	for _, entry := range args {
		found, _ := regexp.MatchString(argumentPattern, entry)
		if found {
			httpPort = ":" + entry[valueStartIndex:]
		}

	}

	return httpPort
}

func Travis(first int, second int) int {
	return first + second
}