package main

import (
	"os"
	"strconv"

	"github.com/chneau/limiter"

	"github.com/chneau/anecdote/pkg/anecdote"
)

const line = "############################################################################"

func printUsage() {
	println(`Possible values:`)
	for k, v := range anecdote.Sources {
		println(k, "<=>", v.Desc)
	}
}

func main() {
	source := "SCMB"
	nb := 1
	if len(os.Args) > 1 {
		source = os.Args[1]
	}
	if len(os.Args) > 2 {
		var err error
		nb, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}
	if _, exist := anecdote.Sources[source]; !exist {
		printUsage()
		return
	}
	limit := limiter.New(nb)
	for i := 0; i < nb; i++ {
		limit.Execute(func() {
			if v, exist := anecdote.Sources[source]; exist {
				aa, err := v.Anecdotes()
				if err != nil {
					panic(err)
				}
				println(line + "\n" + aa[0].String())
				return
			}
		})
	}
	limit.Wait()
	println(line)
}
