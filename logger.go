package main

import (
	"log"
	"os"
)

func initLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetPrefix("INFO ")
	log.SetOutput(os.Stdout)
}
