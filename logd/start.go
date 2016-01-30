package main




import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/Sirupsen/logrus"
)



var startCommand = cli.Command{
	Name: "start",
	Usage: "start logd and output to stdin",
	Action: func(context *cli.Context) {
		fmt.Println("logd start")
		logrus.SetLevel(logrus.DebugLevel)
		logrus.WithFields(logrus.Fields{
			"animal": "walrus",
		}).Info("A walrus appears")

		logrus.Debug("Useful debugging information")
		logrus.Info("Something noteworthy happened!")
		logrus.Warn("You should probably take a look at this.")
		logrus.Error("Something failed but I'm not quitting.")
	},
}


