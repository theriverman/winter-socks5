package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/theriverman/go-socks5"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

// Address config
var listen_ip, listen_port string

// Build details
var app_title string = "SOCKS5 CLI"
var app_built_date string = "" // currentTime.Format("Mon Jan 2 15:04:05 2006")
var app_build_type string = "unreleased/internal"
var app_sem_version, git_commit string

func init() {
	if app_built_date == "" {
		var currentTime time.Time = time.Now()
		app_built_date = currentTime.Format("Mon Jan 2 15:04:05 2006")
	}
}

func main() {
	app := &cli.App{
		Name:    app_title,
		Usage:   fmt.Sprintf("%s | A tiny CLI wrapper around github.com/theriverman/go-socks5 (forked from github.com/armon/go-socks5)", app_title),
		Version: app_sem_version,
		Flags:   []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Starts the SOCKS5 server",
				Action: func(c *cli.Context) error {
					var listen_addr string = fmt.Sprintf("%s:%s", listen_ip, listen_port)
					// Create a SOCKS5 server
					conf := &socks5.Config{}
					fmt.Println("SOCKS5 listening on:", listen_addr)
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
					fmt.Println("  Version:", app_sem_version)
					fmt.Println("  Go version:", runtime.Version())
					fmt.Println("  Git commit:", git_commit)
					fmt.Println("  Built:", app_built_date)
					fmt.Printf("  OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
					fmt.Println("  Build type:", app_build_type)
					return nil
				},
			},
		},
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "Prints version information of go-socks5-cli and quit",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
