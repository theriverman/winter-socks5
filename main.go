package main

import (
	"log"
	"os"
	"time"

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
		Name:     app_title,
		Usage:    proxyBackendUsage,
		Version:  app_sem_version,
		Flags:    []cli.Flag{},
		Commands: proxyBackendCommands,
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
