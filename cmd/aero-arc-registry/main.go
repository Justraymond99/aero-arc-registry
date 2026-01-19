package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/makinje/aero-arc-registry/internal/registry"
	"github.com/urfave/cli/v3"
)

var registryCmd = cli.Command{
	Usage:  "run the aero arc registry process",
	Action: RunRegistry,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "backend",
			Value: "memory",
			Usage: "the specified backend for the registry.",
		},
		&cli.StringFlag{
			Name:  "grpc-listen-address",
			Usage: "the address the grpc server should listen on",
			Value: "0.0.0.0",
		},
		&cli.IntFlag{
			Name:  "grpc-listen-port",
			Usage: "the port the registry's grpc server will listen on",
			Value: 50051,
		},
		&cli.StringFlag{
			Name:  "tls-key-path",
			Usage: "path to tls key file",
			Value: fmt.Sprintf("~/%s", registry.DebugTLSKeyPath),
		},
		&cli.StringFlag{
			Name:  "tls-cert-path",
			Usage: "path to tls crt file",
			Value: fmt.Sprintf("~/%s", registry.DebugTLSCertPath),
		},
		&cli.DurationFlag{
			Name:  "relay-ttl",
			Usage: "ttl for relay health",
			Value: time.Second * 30,
		},
		&cli.DurationFlag{
			Name:  "agent-ttl",
			Usage: "ttl for agent health",
			Value: time.Second * 30,
		},
		&cli.DurationFlag{
			Name:  "heartbeat-interval",
			Usage: "expected relay heartbeat interval",
			Value: time.Second,
		},
		&cli.StringFlag{
			Name:  "redis-addr",
			Usage: "redis instance address",
			Value: "localhost",
		},
		&cli.IntFlag{
			Name:  "redis-port",
			Usage: "redis instance port",
			Value: 6379,
		},
		&cli.StringFlag{
			Name:  "redis-user",
			Usage: "redis username",
			Value: "default",
		},
		&cli.StringFlag{
			Name:  "redis-password",
			Usage: "redis password",
			Value: "",
		},
		&cli.IntFlag{
			Name:  "redis-db",
			Usage: "specified redis db to use",
			Value: 0,
		},
	},
}

func RunRegistry(ctx context.Context, cmd *cli.Command) error {
	return nil
}

func main() {
	if err := registryCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
