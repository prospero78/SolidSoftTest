// Package std -- реализация потокобезопасного ввода/вывода
package std

import (
	"fmt"
)

// IStd -- интерфейс к потокобезопасному вводу/выводу
type IStd interface {
	// Read -- читает входной поток
	Read() (string, error)
	// Write -- пишет выходной поток
	Write(string) error
}

// tStd -- потокобезопасные операции с потоком ввода/вывода
type tStd struct {
	chOut chan string // Выходной канал для терминала
	chErr chan string // Обратный канал для ошибок
}

var (
	std *tStd // Глобальный объект потока ввода/вывода
)

// GetStd -- возвращает глобальный объект ввода/вывода
func GetStd() IStd {
	return std
}

// Write -- пишет в выходной поток
func (sf *tStd) Write(strOut string) (err error) {
	if strOut == "" {
		return nil
	}
	sf.chOut <- strOut
	if strErr := <-sf.chErr; strErr != "" {
		return fmt.Errorf("tStd.Write(): err=%v", strErr)
	}
	return nil
}

// Read -- читает ввод программы
func (sf *tStd) Read() (strIn string, err error) {
	if _, err = fmt.Scanln(&strIn); err != nil {
		return "", fmt.Errorf("tStd.Read(): err=%w", err)
	}
	return strIn, nil
}

// В отдельном цикле работает вывод в консоль
func (sf *tStd) run() {
	for str := range sf.chOut {
		if _, err := fmt.Print(str); err != nil {
			sf.chErr <- fmt.Errorf("tStd.run(): err=%w", err).Error()
		}
	}
}

func init() {
	std = &tStd{
		chOut: make(chan string, 20),
		chErr: make(chan string, 20),
	}
	go std.run()
}
