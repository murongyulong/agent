package main

import (
	"flag"
	"fmt"
	"github.com/smartcaas/agent/cron"
	"github.com/smartcaas/agent/g"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	g.InitRpcClient()

	cron.Heartbeat()
}
