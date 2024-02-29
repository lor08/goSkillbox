package main

import (
	"flag"
	"github.com/lor08/goSkillbox/final/internal/config"
	"github.com/lor08/goSkillbox/final/middleware/httpServer"
)

func init() {
	var path string
	flag.StringVar(&path, "path", "internal/config/config.json", "path to app configuration file")
	flag.Parse()
	config.LoadConfig(path)
}

func main() {
	httpServer.Run(config.Conf.ServerURL)
}
