package main

import (
	"fmt"
	"log"
	"os"
	"time"

	hello "github.com/denkhaus/microservices/greeter/proto/hello"
	"github.com/micro/cli"
	"github.com/micro/go-micro"

	"golang.org/x/net/context"
)

var (
	GitCommit = "undefined"
	Version   = "undefined"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Flags(
			cli.BoolFlag{
				Name:  "version",
				Usage: "Show version info",
			},
			cli.BoolFlag{
				Name:  "revision",
				Usage: "Show revision info",
			},
		),

		micro.Version(Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("version") {
				fmt.Printf("version: %s", Version)
				os.Exit(0)
			}
			if c.Bool("revision") {
				fmt.Printf("revision: %s", GitCommit)
				os.Exit(0)
			}
		}),
	)

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
