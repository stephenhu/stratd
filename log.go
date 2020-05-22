package main

import (
	"log"
)

const (
	INFO 				= "INFO"
	WARNING			= "WARNING"
	ERROR       = "ERROR"
)


func logf(sev string, fname, msg string) {
	log.Printf("[%s] %s(): %s", sev, fname, msg)
} // logf
