package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/realChainLife/porium"
	"github.com/realChainLife/porium/servers"
)

func main() {
	serverURL := fmt.Sprintf(":%d", porium.Cfg.HTTPPort)

	sv := servers.NewHTTPSv()
	log.Fatal(sv.E.Start(serverURL))
}
