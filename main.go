package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var a int

func init() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02/01/2006 15:04:05",
	})
}

func main() {

	file, err := os.OpenFile("text.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0666)
	log.SetOutput(file)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	for i := 0; i <= 10; i++ {
		a += i * i
		log.WithFields(log.Fields{
			"i": i,
			"a": a,
		}).Info("")
	}
	fmt.Println("done")

}
