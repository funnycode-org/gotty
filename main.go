package main

import (
	"fmt"
	"github.com/funnycode-org/gotty/server"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"time"
)

var (
	cmdStartServer = cli.Command{
		Name:  "gServer-start",
		Usage: "start server",
		Action: func(c *cli.Context) error {
			server := server.NewServer()
			return server.Start()
		},
	}
	cmdStartClient = cli.Command{
		Name:  "gClient-start",
		Usage: "start client",
		Action: func(c *cli.Context) error {
			//proxy.Start(bootstrap)
			return nil
		},
	}
)

func main() {

	app := cli.NewApp()
	app.Name = "Gotty proxy"
	app.Compiled = time.Now()
	app.Copyright = "(c) " + strconv.Itoa(time.Now().Year()) + " Gotty"
	app.Usage = "Gotty is a network communication framework."

	app.Commands = []cli.Command{
		cmdStartServer,
		cmdStartClient,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(fmt.Sprintf("启动Gotty失败:%v", err))
	}

}
