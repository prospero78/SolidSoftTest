// Package main -- команда запуска утилиты командной строки на обработку запросов
package main

import (
	"github.com/sirupsen/logrus"

	"github.com/prospero78/SolidSoftTest/cmd/solidtest/cmdarg"
)

func main() {
	cmd := cmdarg.New()
	cmd.Run()

	logrus.Debugf("main(): test of SolidSoft\n")
}
