package lib

import "log"

func Info(message string) {
	log.Printf("[INFO] %s\n", message)
}

func Error(message string) {
	log.Println("[ERROR] #s\n", message)
}
