package main

import (
	"flag"
	"log"

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
	flag.BoolVar(&c.DebugInfo, "d", true, "Debug info")
	flag.Parse()

	if action == "gen" {
		r, err := hbench.Gen(c)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("r: %+v", r)

	} else if action == "bench" {
		r, err := hbench.Bench(c)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("r: %+v", r)
	} else {
		flag.PrintDefaults()
	}

}
