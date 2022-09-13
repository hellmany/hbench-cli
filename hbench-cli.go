package main

import (
	"flag"

	hbench "github.com/hellmany/hbench"
)

func main() {

	c := hbench.ConfData{}

	var action string
	flag.StringVar(&action, "a", "", "Action gen|bench - required")

	flag.IntVar(&c.Threads, "t", 100, "threads to read/write")
	flag.IntVar(&c.Max, "m", 1000, "max files to read/write, 0 - to read all")
	flag.IntVar(&c.LimitMax, "l", 1000, "limit max files to read before rand, 0 - to read all")
	flag.IntVar(&c.Inter, "i", 1, "interations to read")
	flag.IntVar(&c.Size, "s", 1, "size to read/write in Mb, 0 - full file to read")
	flag.IntVar(&c.RandSize, "r", 3, "add rand mb to size writing")
	flag.StringVar(&c.Path, "p", "/stor/tmp", "path")
	flag.Parse()

	if action == "gen" {
		hbench.Gen(c)
	} else if action == "bench" {
		hbench.Bench(c)
	} else {
		flag.PrintDefaults()
	}

}
