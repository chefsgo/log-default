package log_default

import (
	"io"
	"os"

	"github.com/chefsgo/log"
)

type (
	defaultDriver struct {
	}
	defaultConnect struct {
		config         log.Config
		stdout, stderr io.Writer
	}
)

func (driver *defaultDriver) Connect(config log.Config) (log.Connect, error) {
	return &defaultConnect{
		config: config, stdout: os.Stdout, stderr: os.Stderr,
	}, nil
}

//打开连接
func (connect *defaultConnect) Open() error {
	return nil
}

//关闭连接
func (connect *defaultConnect) Close() error {
	connect.Flush()
	return nil
}

func (connect *defaultConnect) Write(msg *log.Log) error {
	body := msg.Format()
	if msg.Level <= log.LevelWarning {
		connect.stderr.Write([]byte(body + "\n"))
	} else {
		connect.stdout.Write([]byte(body + "\n"))
	}
	return nil
}
func (connect *defaultConnect) Flush() {
}
