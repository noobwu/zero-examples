package main

import (
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

func foo() {
	logx.WithDuration(time.Second).Error("world")
}

func main() {
	c := logx.LogConf{
		Mode: "console", //console|file|volume
		Path: "logs",
	}
	logx.MustSetup(c)  // logx 根据配置初始化
	defer logx.Close() //关闭日志输出

	logx.Info("info")
	logx.Error("error")
	logx.ErrorStack("hello")
	logx.Errorf("%s and %s", "hello", "world")
	logx.Severef("%s severe %s", "hello", "world")
	logx.Slowf("%s slow %s", "hello", "world")
	logx.Statf("%s stat %s", "hello", "world")
	logx.WithDuration(time.Minute + time.Second).Info("hello")
	logx.WithDuration(time.Minute + time.Second).Error("hello")

	foo()
}
