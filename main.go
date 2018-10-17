package main

import (
	"os"
	"os/signal"

	"github.com/chneau/anecdote/pkg/anecdote"
)

func init() {
	gracefulExit()
}

func gracefulExit() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		os.Exit(0)
	}()
}

// checkError
func ce(err error) {
	if err != nil {
		panic(err)
	}
}

func printUsage() {
	println(`Possible values:`)
	for k, v := range anecdote.Sources {
		println(k, "<=>", v.Desc)
	}
}

func main() {
	source := "SCMB"
	if len(os.Args) > 1 {
		source = os.Args[1]
	}
	if v, exist := anecdote.Sources[source]; exist {
		aa, err := v.Anecdotes()
		ce(err)
		println(aa[0].String())
		return
	}
	printUsage()
}
