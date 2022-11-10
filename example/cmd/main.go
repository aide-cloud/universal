package main

import (
	"flag"
	"github.com/aide-cloud/universal/example/internal/conf"
	"github.com/aide-cloud/universal/example/internal/server"
	"github.com/aide-cloud/universal/executor"
)

const (
	appName = "aide-family-layout"
	cmdName = "main"
	desc    = "aide-family-layout is a service for layout"
	author  = "aide-cloud"
)

//Version go build -ldflags "-X main.server.Version=x.y.z"
var Version string
var filePath = flag.String("f", conf.DefaultPath, "config file path")

func init() {
	flag.Parse()
	conf.LoadConfig(*filePath)
}

func main() {
	executor.ExecMulSerProgram(
		executor.NewLierCmd(
			executor.WithServices(server.NewHttpServer(server.GetGlobalLog())),
			executor.AddProperty("appName", appName),
			executor.AddProperty("cmdName", cmdName),
			executor.AddProperty("desc   ", desc),
			executor.AddProperty("author ", author),
			executor.AddProperty("version", Version),
			executor.WithLogger(server.GetGlobalLog()),
		),
	)
}
