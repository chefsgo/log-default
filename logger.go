package log_default

import (
	"io"
	"os"
	"sync"

	"github.com/chef-go/chef"
)

type (
	defaultLogDriver struct {
	}
	defaultLogConnect struct {
		config         chef.LogConfig
		stdout, stderr io.Writer
	}
	defaultLogWriter struct {
		lock sync.Mutex
	}
)

func (driver *defaultLogDriver) Connect(config chef.LogConfig) (chef.LogConnect, error) {
	return &defaultLogConnect{
		config: config, stdout: os.Stdout, stderr: os.Stderr,
	}, nil
}

//打开连接
func (connect *defaultLogConnect) Open() error {
	return nil
}

//关闭连接
func (connect *defaultLogConnect) Close() error {
	connect.Flush()
	return nil
}

func (connect *defaultLogConnect) Write(log chef.Log) error {
	msg := log.Format()
	if log.Level <= chef.LogWarning {
		connect.stderr.Write([]byte(msg + "\n"))
	} else {
		connect.stdout.Write([]byte(msg + "\n"))
	}
	return nil
}
func (connect *defaultLogConnect) Flush() {
}
