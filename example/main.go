package main

import (
	"fmt"
	"time"

	"github.com/bcho/whitetree"
)

const NumWorkers = 5

var parsers = whitetree.ParserPackage{
	"foo": func(whitetree.TaskContext) (whitetree.HandlerId, error) { return "foo", nil },
}
var handlers = whitetree.HandlerPackage{
	"foo": func(ctx whitetree.TaskContext) error {
		panic(1)
		fmt.Printf("executing: %s\n", string(ctx.Data))
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Printf("execute finished\n")
		return nil
	},
}

func main() {
	wg := whitetree.NewWorkerGroup(
		NumWorkers,
		&parsers,
		&handlers,
	)

	tick := time.NewTicker(time.Duration(1000000) * time.Nanosecond)

	for {
		select {
		case err := <-wg.ErrChan:
			fmt.Printf("error: %v\n", err)
		case <-tick.C:
			// fmt.Printf("received work\n")
			go func() {
				wg.Spawn(&whitetree.TaskContext{[]byte("test")})
				fmt.Printf("pushed new work\n")
			}()
		}
	}
}
