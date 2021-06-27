package main

import (
	"fmt"
	"git.onespace.co.th/project-v/project-v-go-backend/cli"
	"git.onespace.co.th/project-v/project-v-go-backend/setting"
	"github.com/patcharp/golib/server"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"net"
	"os"
	"path/filepath"
)

var (
	// Application identity
	ProgramName string
	Version     string
	BuildTime   string
	BuildCommit string
)

func main() {
	var (
		startArgs = struct {
			host *net.IP
			port *string
		}{}
	)
	a := kingpin.New(filepath.Base(os.Args[0]), fmt.Sprintf("%s %s", ProgramName, Version))
	a.Version(Version)
	a.HelpFlag.Short('h')

	// Start
	startCmd := a.Command("start", "Start server command.")
	startArgs.host = startCmd.Flag("host", "Set server host address.").Envar("SERVER_HOST").Default("0.0.0.0").IP()
	startArgs.port = startCmd.Flag("port", "Set server listen port.").Envar("SERVER_PORT").Default("5000").String()

	// Set configure
	setting.AppName = ProgramName
	setting.AppVersion = Version
	setting.ListeningPort = *startArgs.port
	switch kingpin.MustParse(a.Parse(os.Args[1:])) {
	case startCmd.FullCommand():
		s := server.New(server.Config{
			Host: startArgs.host.String(),
			Port: *startArgs.port,
			Prod: setting.Production,
		})
		if err := cli.StartServer(&s); err != nil {
			logrus.Infoln("Start server error ->", err)
		}
	}
}
