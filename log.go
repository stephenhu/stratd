package main

import (
	"log"
	"runtime"
)

const (
	INFO 				= "INFO"
	WARNING			= "WARNING"
	ERROR       = "ERROR"
)


func logf(sev string, msg string) {

	pc := make([]uintptr, 10)

	runtime.Callers(2, pc)

	f := runtime.FuncForPC(pc[0])

	log.Printf("[%s] %s(): %s", sev, f.Name(), msg)

} // logf
