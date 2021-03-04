// Package std -- реализация потокобезопасного ввода/вывода
package std

import (
	"fmt"
	"strings"
	"sync"
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
	chWrite chan string // Выходной канал для терминала
	chErr   chan string // Обратный канал для ошибок
	block   sync.RWMutex
}

var (
	std *tStd // Глобальный объект потока ввода/вывода
)

// GetStd -- возвращает глобальный объект ввода/вывода
func GetStd() IStd {
	return std
}

// Write -- пишет в выходной поток
func (sf *tStd) Write(strWrite string) (err error) {
	defer sf.block.Unlock()
	sf.block.Lock()
	if strWrite == "" {
		return nil
	}
	sf.chWrite <- strWrite
	if strErr := <-sf.chErr; strErr != " " {
		return fmt.Errorf("tStd.Write(): err=%v", strErr)
	}
	return nil
}

// Read -- читает ввод программы
func (sf *tStd) Read() (strIn string, err error) {
	for {
		if _, err = fmt.Scan(&strIn); err != nil {
			if strings.Contains(err.Error(), "expected newline") {
				continue
			}
			return "", fmt.Errorf("tStd.Read(): err=%w", err)
		}
		if strIn == "" {
			continue
		}
		break
	}

	return strIn, nil
}

// В отдельном цикле работает вывод в консоль
func (sf *tStd) run() {
	for str := range sf.chWrite {
		if _, err := fmt.Print(str); err != nil {
			sf.chErr <- fmt.Errorf("tStd.run(): err=%w", err).Error()
		}
		sf.chErr <- " "
	}
}

func init() {
	std = &tStd{
		chWrite: make(chan string, 20),
		chErr:   make(chan string, 20),
	}
	go std.run()
}
