//go:build armon

package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/theriverman/go-socks5"
	"github.com/urfave/cli/v2"
)

var proxyBackend string = "github.com/armon/go-socks5"
var proxyBackendUsage string = fmt.Sprintf("%s | A tiny CLI wrapper around github.com/theriverman/go-socks5 (forked from github.com/armon/go-socks5)", app_title)
var proxyBackendCommands []*cli.Command = []*cli.Command{
	{
		Name:  "server",
		Usage: "Starts the SOCKS5 server",
		Action: func(c *cli.Context) error {
			var listen_addr string = fmt.Sprintf("%s:%s", listen_ip, listen_port)
			// Create a SOCKS5 server
			conf := &socks5.Config{}
			log.Println("SOCKS5-CLI listening on:", listen_addr)
			server, err := socks5.New(conf)
			if err != nil {
				return err
			}
			// Create SOCKS5 proxy on `listen_addr`
			if err := server.ListenAndServe("tcp", listen_addr); err != nil {
				return err
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Usage:       "Address to listen on for incoming SOCKS5 requests. To listen on all addresses, set to 0.0.0.0",
				Value:       "127.0.0.1",
				Destination: &listen_ip,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "port",
				Usage:       "Port to listen on for incoming SOCKS5 requests",
				Value:       "1080",
				Destination: &listen_port,
				Required:    false,
			},
		},
	},
	{
		Name:  "version",
		Usage: fmt.Sprintf("Show the %s version information (detailed)", app_title),
		Action: func(c *cli.Context) error {
			fmt.Println(app_title + ":")
			fmt.Printf("  Version: %s\n", app_sem_version)
			fmt.Printf("  Go version: %s\n", runtime.Version())
			fmt.Printf("  Git commit: %s\n", git_commit)
			fmt.Printf("  Built: %s\n", app_built_date)
			fmt.Printf("  OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
			fmt.Printf("  Build type: %s\n", app_build_type)
			fmt.Printf("  Proxy Backend: %s\n", proxyBackend)
			return nil
		},
	},
}
