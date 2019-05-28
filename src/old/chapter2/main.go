package main

import (
	_ "chapter2/matchers" // not used in this package therfore (_). This is still needed to be imported to handle the init function in the matchers (rss.go) package.
	"chapter2/search"
	"log"
	"os"
	// these paths must be relative to GOPATH and GOROOT env variables
)

//init is called prior to main
func init() {
	//SetOutput sets the logging destination.
	//Print[f|l|n], Fatal[f|l|n] and Panic[f|l|n] are standard output commands
	//Fmt can be used to log in console.
	// log will create a Logger type
	//stdout, stdin, stderr -> open standard I / O and error files.
	// default logger file is stderr
	// change defalut logging device
	log.SetOutput(os.Stdout)

}

func main() {
	// search the passed in term
	search.Run("President")
}
