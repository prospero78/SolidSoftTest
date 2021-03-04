// Package cliapp -- приложение командной строки по тестовому заданию
package cliapp

import (
	"context"
	"fmt"
	"time"

	"github.com/clover87/valid"
	"github.com/sirupsen/logrus"

	"github.com/prospero78/SolidSoftTest/internal/cliapp/job"
	"github.com/prospero78/SolidSoftTest/internal/std"
)

// TcliApp -- операции с приложением командной строки
type TCliApp struct {
	std std.IStd // Глобальный объект стандартного ввода/вывода
	ctx context.Context
}

// New -- возвращает новый *TCliApp
func New() *TCliApp {
	return &TCliApp{
		std: std.GetStd(),
		ctx: context.Background(),
	}
}

// Run -- запускает приложение в бесконечном цикле на чтение входного потока
func (sf *TCliApp) Run() (err error) {
	logrus.Debugf("TCliApp.Run()")
	// var id string
	for taskID, err := sf.readTag(); err == nil; {
		strJson := ""
		for { // Фомирование тела задачи
			strPartJSON, err := sf.std.Read()
			if err != nil {
				logrus.WithError(err).Fatalf("TCliApp.Run(): in get body JSON")
			}
			if valid.JsonEnd(strPartJSON) {
				break
			}
			strJson += strPartJSON
		}

		job, err := job.New(taskID, strJson)
		if err == nil {
			break
		}
		ctx, cancel := context.WithTimeout(sf.ctx, time.Duration(time.Second*1))
		defer cancel()
		go job.Run(ctx)
	}
	return fmt.Errorf("TCliaApp.Run(): err=%w", err)
}

// Читает открывающий тег на обработку
func (sf *TCliApp) readTag() (numJob string, err error) {
	strIn, err := sf.std.Read()
	if err != nil {
		return "", fmt.Errorf("TCliaApp.readTag(): err=%w", err)
	}
	id, err := valid.JsonBeg(strIn)
	if err != nil {
		return "", fmt.Errorf("TCliApp.readTag(): error in validate")
	}
	return id, nil
}
