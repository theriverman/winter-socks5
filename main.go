package main

import (
	"fmt"
	"log"
	"os"

	socks5 "github.com/theriverman/go-socks5"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func main() {
	var listen_ip, listen_port, version string
	app := &cli.App{
		Name:    "SOCKS5 CLI",
		Usage:   "SOCKS5 CLI | A tiny CLI wrapper around github.com/theriverman/go-socks5 (forked from github.com/armon/go-socks5)",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Usage:       "Address to listen on for incoming SOCKS5 requests. To listen on all addresses, set to 0.0.0.0",
				Value:       "127.0.0.1",
				Destination: &listen_ip,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "port",
				Usage:       "Port to listen on for incoming SOCKS5 requests. Defaults to 1080",
				Value:       "1080",
				Destination: &listen_port,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			var listen_addr string = fmt.Sprintf("%s:%s", listen_ip, listen_port)
			// Create a SOCKS5 server
			conf := &socks5.Config{}
			fmt.Println("SOCKS5 listening on:", listen_addr)
			server, err := socks5.New(conf)
			if err != nil {
				panic(err)
			}

			// Create SOCKS5 proxy on localhost port 8000
			if err := server.ListenAndServe("tcp", listen_addr); err != nil {
				panic(err)
			}

			return nil
		},
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "Prints the version of go-socks5-cli",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
