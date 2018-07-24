package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/chneau/anecdote/pkg/anecdote"
)

var (
	source string
)

func init() {
	gracefulExit()
	flag.StringVar(&source, "source", "SI", "source")
	flag.Parse()
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
func ce(err error, msg string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}

func main() {
	aa := []anecdote.Anecdote{}
	var err error
	switch source {
	case "SCMB":
		aa, err = anecdote.SCMB()
		ce(err, "anecdote")
	case "SI":
		fallthrough
	default:
		aa, err = anecdote.SI()
		ce(err, "anecdote")
	}
	fmt.Println(aa[0].String())
}
