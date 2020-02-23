package main

import (
	"consensus-demo/node"
	"flag"
)

func main() {
	name := flag.String("n", "node001", "node name")
	difficulty := flag.Uint("d", 1, "mine difficulty")
	cs := flag.String("c", "pos", "consensus, pow or pos")
	flag.Parse()

	conf := node.Config{
		Name:       *name,
		Difficulty: *difficulty,
		Consensus:  *cs,
	}
	n, err := node.New(conf)
	if err != nil {
		panic(err)
	}
	// start node
	n.Start()

	// sigc := make(chan os.Signal)
	// signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	// defer signal.Stop(sigc)
	// <-sigc
	// fmt.Println("Got interrupt, shutting down...")
	// n.Stop()
}
