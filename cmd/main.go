package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/porium/porium"
	"github.com/porium/porium/servers"
)

func main() {
	serverURL := fmt.Sprintf(":%d", backend.Cfg.HTTPPort)

	sv := servers.NewHTTPSv()
	log.Fatal(sv.E.Start(serverURL))
}
