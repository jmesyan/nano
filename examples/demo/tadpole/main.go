package main

import (
	"github.com/jmesyan/nano/utils"
	"log"
	"net/http"
	"os"

	"github.com/jmesyan/nano"
	"github.com/jmesyan/nano/examples/demo/tadpole/logic"
	"github.com/jmesyan/nano/serialize/json"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "tadpole"
	app.Author = "github.com/jmesyan/nano authors"
	app.Version = "0.0.1"
	app.Copyright = "github.com/jmesyan/nano authors reserved"
	app.Usage = "tadpole"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: ":23456",
			Usage: "game server address",
		},
	}

	app.Action = serve

	app.Run(os.Args)
}

func serve(ctx *cli.Context) error {
	// register all service
	nano.Register(logic.NewManager())
	nano.Register(logic.NewWorld())
	utils.SetSerializer(json.NewSerializer())

	//nano.EnableDebug()
	log.SetFlags(log.LstdFlags | log.Llongfile)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	nano.SetCheckOriginFunc(func(_ *http.Request) bool { return true })

	addr := ctx.String("addr")
	nano.ListenWS(addr)

	return nil
}
