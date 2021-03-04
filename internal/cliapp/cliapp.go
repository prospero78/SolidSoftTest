// Package cliapp -- приложение командной строки по тестовому заданию
package cliapp

import (
	"fmt"

	"github.com/clover87/valid"
	"github.com/sirupsen/logrus"
)

// TcliApp -- операции с приложением командной строки
type TCliApp struct {
}

// New -- возвращает новый *TCliApp
func New() *TCliApp {
	return &TCliApp{}
}

// Run -- запускает приложение в бесконечном цикле на чтение входного потока
func (sf *TCliApp) Run() error {
	logrus.Debugf("TCliApp.Run()")
	if err := valid.Validate("123"); err != nil {
		return fmt.Errorf("TCliApp.Run(): error in validate")
	}
	return nil
}
