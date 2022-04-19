package log_default

import (
	"github.com/chefsgo/chef"
)

func Driver() chef.LogDriver {
	return &defaultLogDriver{}
}

func init() {
	chef.Register("default", Driver())
	chef.Register("console", Driver())
}
