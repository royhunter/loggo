// logd project main.go
package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Logd"
	app.Usage = "Fcprobe log process tool"

	app.Commands = []cli.Command{
		startCommand,
		outputFile,
	}
/*
	app.Action = func(c *cli.Context) {
		fmt.Println("main process")
		
	}
*/
	app.Run(os.Args)
}
