package config

import (
	"log"
	"os"
)

// Logger is a pointer to custom logger object.
var Logger = log.New(os.Stdout, "server: ", log.Lshortfile|log.Ldate|log.Lmicroseconds)
