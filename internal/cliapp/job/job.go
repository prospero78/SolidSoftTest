// Package job -- входящее задание на обработку
package job

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/prospero78/SolidSoftTest/internal/std"
)

// TJob -- структура для обработки входящего JSON
type TJob struct {
	MsgID     string  `json:"id"`
	MsgBody   string  `json:"body"`
	MsgReplic []*TJob `ison:"replies"`
	taskID    string
	job       string
	std       std.IStd
}

// New -- возвращает новый *TJob
func New(taskID, strJob string) (job *TJob, err error) {
	{ // Предусловия
		if taskID == "" {
			return nil, fmt.Errorf("job.go/New(): taskID is empty")
		}
		if strJob == "" {
			return nil, fmt.Errorf("job.go/New(): strJob is empty")
		}
	}
	job = &TJob{
		taskID: taskID,
		job:    strJob,
	}
	return job, nil
}

// Run -- вполняет всю работу по оформлению и топравке ответа
func (sf *TJob) Run(ctx context.Context) {
	select {
	case ctx.Done():
	default:
		if err := json.Unmarshal([]byte(sf.job), sf); err != nil {
			logrus.WithError(err).Fatalf("TCliApp.Run(): unmarshal JSON")
		}
		sf.fillMsg()
		outJSON, err := json.Marshal(sf)
		if err != nil {
			logrus.WithError(err).WithField("taskID", sf.taskID).Fatalf("TJob.Run()")
		}
		if err = sf.std.Write(string(outJSON)); err != nil {
			logrus.WithError(err).WithField("taskID", sf.taskID).Fatalf("TJob.Run()")
		}
	}
}

// Рекурсивно заполняет сообщения
func (sf *TJob) fillMsg() {
	for _, msg := range sf.MsgReplic {
		sf.MsgBody = fmt.Sprintf("id:%v msg:%q", sf.MsgID, sf.taskID)
		msg.fillMsg()
	}
}
