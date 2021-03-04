package job

/*
	Исходник предоставляет тест для проверки работы TJob.

	Он недописан, потому что для тестового задания потратить
	целый день -- хотелось бы увидеть тестовую зарплату (уж извините).
*/

import (
	"context"
	"testing"
)

var (
	job *TJob
	err error
	ctx context.Context
)

func TestJob(test *testing.T) {
	ctx = context.Background()
	createBad(test)
	create(test)
	runBadJob(test)
	runJob(test)
}

func runJob(test *testing.T) {
	test.Logf("runJob()\n")
	defer func() {
		if _panic := recover(); _panic != nil {
			test.Errorf("runJob(): generate panic, panic=%v", _panic)
		}
	}()
	job.job = strJob
	job.Run(ctx)
}

func runBadJob(test *testing.T) {
	test.Logf("runBadJob()\n")
	defer func() {
		if _panic := recover(); _panic == nil {
			test.Error("runBadJob(): empty panic")
		}
	}()
	job.job = "fdghdfgh"
	job.Run(ctx)
}

func createBad(test *testing.T) {
	test.Logf("createBad()\n")
	{ // Нет taskID
		job, err := New("", "dfghdfgh")
		if err == nil {
			test.Errorf("createBad(): err==nil")
		}
		if job != nil {
			test.Errorf("createBad(): job!=nil")
		}
	}
	{ // Нет strJob
		job, err := New("amd", "")
		if err == nil {
			test.Errorf("createBad(): err==nil")
		}
		if job != nil {
			test.Errorf("createBad(): job!=nil")
		}
	}
}

func create(test *testing.T) {
	test.Logf("create()\n")
	{ // Нормальное создание
		job, err = New("amd", strJob)
		if err != nil {
			test.Errorf("create(): err=%v", err)
		}
		if job == nil {
			test.Errorf("create(): job==nil")
		}
	}
}

const (
	strJob = `{
		"id": 1,
		"replies": [
			{
				"id": 2,
				"replies": []
			},
			{
				"id": 3,
				"replies": [
					{
						"id": 4,
						"replies": []
					},
					{
						"id": 5,
						"replies": []
					}
				]
			}
		]
	}`
)
