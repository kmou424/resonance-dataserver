package server

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gookit/goutil/sysutil"
	"github.com/kmou424/resonance-dataserver/server/engine"
	"github.com/kmou424/resonance-dataserver/server/middlewares"
	"github.com/kmou424/resonance-dataserver/server/routers"
)

func Run() {
	middlewares.Register()
	routers.Register()

	serverHost := sysutil.Getenv("SERVER_HOST", "0.0.0.0")
	serverPort := sysutil.Getenv("SERVER_PORT", "8080")

	err := engine.E.Run(fmt.Sprintf("%s:%s", serverHost, serverPort))
	if err != nil {
		log.Fatal("run server failed", "error", err)
		return
	}
}
