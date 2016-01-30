package main



import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
	"github.com/Sirupsen/logrus"
)


var outputFile = cli.Command{
	Name: "ofile",
	Usage: "start logd and output to indicate file",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "file, f",
			Value: "",
			Usage: "filename and path to the log file",
		},
	},
	Action: func(context *cli.Context) {
		fmt.Println("logd file")
		file := context.String("file")
		if file != "" {
			fmt.Println("log filename:", file)
			f, err := os.Create(file)
			if err != nil {
				return
			}
			logrus.SetOutput(f)
		} else {
			fmt.Println("please input log filename")
		}

		logrus.Debug("Useful debugging information")
		logrus.Info("Something noteworthy happened!")
		logrus.Warn("You should probably take a look at this.")
		logrus.Error("Something failed but I'm not quitting.")
	},
}
