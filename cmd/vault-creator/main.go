package main

import (
	"github.com/GrolimundSolutions/vault-creator/cmd/vault-creator/vault-creatormain"
	log "github.com/sirupsen/logrus"
)

func init() {
	// set default log config
	log.SetLevel(log.WarnLevel)
	log.SetReportCaller(false)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {

	// Version is provided by ldflags
	var Version = "unspecified"

	// Build is provided by ldflags
	var Build = "unspecified"

	// Show build information
	log.Infof("Version: %s, Build: %s", Version, Build)

	vault_creatormain.Run()
}
