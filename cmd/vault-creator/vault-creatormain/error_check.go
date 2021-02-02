package vault_creatormain

import (
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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
