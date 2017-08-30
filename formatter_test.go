package consolus

import "testing"
import log "github.com/sirupsen/logrus"

func TestInit(t *testing.T) {
	log.SetFormatter(&ConsoleFormatter{})
	log.Info("it's merely a flesh wound!")

	"TIMES"
}