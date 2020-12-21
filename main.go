package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {

	var host string
	var port string
	app := &cli.App{
		Name:     "gomap",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Lewis Baston",
				Email: "lewis@baston.dev",
			},
		},
		HelpName:  "Menu",
		Usage:     "[ip] [args]",
		UsageText: "--ip 10.10.10.10 [args]",
		ArgsUsage: "[args and such]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Aliases:     []string{"ip"},
				Usage:       "IP adress of the host",
				Destination: &host,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "Port you want to scan",
				Destination: &port,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Running GoMap")
			dialer(port, host)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func dialer(port, ip string) {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	conn, err := d.DialContext(ctx, "tcp", ip+":"+port)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connected")
	}
}
