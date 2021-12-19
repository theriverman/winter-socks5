//go:build txthinking

package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/txthinking/socks5"
	"github.com/urfave/cli/v2"
)

var username, password string
var tcpTimeout, udpTimeout int

var proxyBackendUsage string = fmt.Sprintf("%s | A tiny CLI wrapper around github.com/txthinking/socks5", app_title)

var proxyBackendCommands []*cli.Command = []*cli.Command{
	{
		Name:  "server",
		Usage: "Starts the SOCKS5 server",
		Action: func(c *cli.Context) (err error) {
			var listen_addr string = fmt.Sprintf("%s:%s", listen_ip, listen_port)
			// Create a SOCKS5 server
			var server *socks5.Server
			server, err = socks5.NewClassicServer(listen_addr, listen_ip, username, password, tcpTimeout, udpTimeout)
			if err != nil {
				return err
			}
			log.Printf("%s | Listening at %s\n", app_title, listen_addr)
			if username != "" || password != "" {
				log.Println("Attention! The proxy server is protected by username and password")
			}
			// start listening on the address
			err = server.ListenAndServe(nil)
			return err
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
			&cli.StringFlag{
				Name:        "username",
				Usage:       "Username to authenticate clients to this server",
				Value:       "",
				Destination: &username,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "password",
				Usage:       "Password to authenticate clients to this server",
				Value:       "",
				Destination: &password,
				Required:    false,
			},
			&cli.IntFlag{
				Name:        "tcp-timeout",
				Usage:       "Define TCP timeout in seconds",
				Value:       0,
				Destination: &tcpTimeout,
				Required:    false,
			},
			&cli.IntFlag{
				Name:        "udp-timeout",
				Usage:       "Define UDP timeout in seconds",
				Value:       0,
				Destination: &udpTimeout,
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
