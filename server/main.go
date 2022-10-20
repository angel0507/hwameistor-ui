package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/hwameistor/hwameistor-ui/server/app"
)

func cmdNotFound(c *cli.Context, command string) {
	panic(fmt.Errorf("unrecognized command: %s", command))
}

func onUsageError(c *cli.Context, err error, isSubcommand bool) error {
	panic(fmt.Errorf("usage error, please check your command"))
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	a := cli.NewApp()
	a.Usage = "Hwameistor-UI Manager"

	a.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		if c.GlobalBool("log-json") {
			logrus.SetFormatter(&logrus.JSONFormatter{})
		}
		return nil
	}

	a.Commands = []cli.Command{
		app.DaemonCmd(),
	}
	a.CommandNotFound = cmdNotFound
	a.OnUsageError = onUsageError

	if err := a.Run(os.Args); err != nil {
		logrus.Fatalf("Critical error: %v", err)
	}
}
