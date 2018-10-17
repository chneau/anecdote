package main

import (
	"os"

	"github.com/chneau/anecdote/pkg/anecdote"
)

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
		if err != nil {
			panic(err)
		}
		println(aa[0].String())
		return
	}
	printUsage()
}
