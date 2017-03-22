package wbt

import (
	"fmt"
	"runtime"

	"github.com/kuntao/widebullet/config"
	"github.com/kuntao/widebullet/wlog"
)

const (
	Version = "0.3.1"
)

var (
	Config config.Config
	AL     wlog.Logger
	EL     wlog.Logger
)

func ServerHeader() string {
	return fmt.Sprintf("WideBullet %s", Version)
}

func PrintVersion() {
	fmt.Printf(`wbt %s
Compiler: %s %s
Copyright (C) 2016 Mercari, Inc.
`,
		Version,
		runtime.Compiler,
		runtime.Version())
}
