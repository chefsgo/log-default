package log_default

import (
	"github.com/chefsgo/log"
)

func Driver() log.Driver {
	return &defaultDriver{}
}

func init() {
	log.Register("default", Driver())
}
