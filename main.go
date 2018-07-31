package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/chneau/anecdote/pkg/anecdote"
)

var (
	source string
)

func init() {
	gracefulExit()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	flag.StringVar(&source, "source", "SI", "source can be SI or SCMB")
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
		aa, err = anecdote.SI()
		ce(err, "anecdote")
	default:
		aa, err = anecdote.SCMB()
		ce(err, "anecdote")
	}
	fmt.Println(aa[0].String())
}
