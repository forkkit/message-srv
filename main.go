package main

import (
	"log"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/message-srv/handler"
	"github.com/micro/message-srv/message"

	"github.com/micro/go-os/kv"
	"github.com/micro/go-os/sync"
	"github.com/micro/go-os/sync/consul"

	proto "github.com/micro/message-srv/proto/message"
)

var (
	SyncAddress = "127.0.0.1:8500"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.message"),

		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),

		micro.Flags(cli.StringFlag{
			Name:   "sync_address",
			EnvVar: "SYNC_ADDRESS",
			Usage:  "Address for the synchronization service e.g. consul",
		}),

		micro.Action(func(c *cli.Context) {
			if addr := c.String("sync_address"); len(addr) > 0 {
				SyncAddress = addr
			}
		}),
	)

	service.Init()

	message.Init(
		service.Server().Options().Broker,
		kv.NewKV(
			kv.Namespace("go.micro.srv.message"),
			kv.Client(service.Client()),
			kv.Server(service.Server()),
		),
		consul.NewSync(sync.Nodes(SyncAddress)),
	)

	proto.RegisterMessageHandler(service.Server(), new(handler.Message))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
