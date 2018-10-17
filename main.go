package main

import (
	"os"
	"sort"
	"strconv"

	"github.com/chneau/limiter"

	"github.com/chneau/anecdote/pkg/anecdote"
)

const line = "############################################################################"

func printUsage() {
	println(`Possible values:`)
	sorted := []string{}
	for k := range anecdote.Sources {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	for i := range sorted {
		println(sorted[i], "<=>", anecdote.Sources[sorted[i]].Desc)
	}
}

func main() {
	source := "SCMB"
	nb := 1
	if len(os.Args) > 1 {
		source = os.Args[1]
	}
	if len(os.Args) > 2 {
		nb, _ = strconv.Atoi(os.Args[2])
		if nb < 1 {
			nb = 1
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
