// Package cmdarg -- аргументы командной строки
package cmdarg

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/prospero78/SolidSoftTest/internal/cliapp"
)

const (
	BUILD = "build 0010"
)

// TCmdArg -- операции с аргументами командной строки
type TCmdArg struct {
	app     *cli.App        // Обработчик командной строки
	isDebug bool            // Признак отладки
	cli     *cliapp.TCliApp // Потоковый обработчик JSON
}

// New -- возвращает новый *TCmdArg
func New() *TCmdArg {
	ca := &TCmdArg{}

	ca.app = &cli.App{
		Name:      "Приложение командной строки для обработки запросов",
		Version:   BUILD,
		Copyright: "prospero.78.su@gmail.com (a)(c) Moskow 2021",
		Description: `		Приложение позволяет используя опции командной строки обрабатывать запросы
	на заполнение JSON-структур.
	
	Типичный цикл работы:
		1. Запрос на обработку JSON начинается с ключа JSONBEG:NN (NN -- номер запроса в любом виде).
		2. Далее идёт тело JSON-запроса (с новой строки, можно в несколько строк).
		3. Запрос закрывается ключом JSONEND в отдельной строке.
	
	После отработки запроса:
		1. Приложение отвечает выводом в стандартный поток вывода RESP:NN
			(тот же самый номер запроса).
		2. Со следующей строки начинается тело JSON-ответа (можно в несколько строк).
		3. Ответ заканчивается RESPEND в отдельной строке.
	
	Запросы могут поступать в произвольном порядке, ответы также следуют в произвольном порядке.
	Перемешивать как запросы, так и ответы в одних ключах нельзя.`,
		Commands: []*cli.Command{
			{
				Name:    "debug",
				Aliases: []string{"d", "-d", "--debug"},
				Usage:   "включить режим отладки",
				Action:  ca.setDebug,
			},
			{
				Name:    "run",
				Aliases: []string{"r", "-r", "--run"},
				Usage:   "обработка входящего потока",
				Action:  ca.run,
			},
		},
	}
	ca.cli = cliapp.New()
	return ca
}

// Run -- запускает обработчик командной строки
func (sf *TCmdArg) Run() {
	err := sf.app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
	// Если поставить впереди, то до срабатывания обработчика выводиться не будет
	//  (по умолчанию у logrus -- debug отключен)
	logrus.Debugf("TCmdArg.Run()")
}

// Устанавливает режим отладки
func (sf *TCmdArg) setDebug(ctx *cli.Context) error {
	logrus.Infof("Режим отладки включен")
	logrus.SetLevel(logrus.DebugLevel)
	sf.isDebug = true
	return nil
}

// Запускает обработчик командной строки
func (sf *TCmdArg) run(ctx *cli.Context) error {
	return sf.cli.Run()
}

// IsDebug -- возвращает признак отладки
func (sf *TCmdArg) IsDebug() bool {
	return sf.isDebug
}
