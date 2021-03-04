// Package job -- входящее задание на обработку
package job

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/prospero78/SolidSoftTest/internal/std"
)

// TMsg -- входящее сообщение на контроль
type TMsg struct {
	ID     int     `json:"id"`
	Replic []*TMsg `json:"replies"`
}

// TJob -- структура для обработки входящего JSON
type TJob struct {
	MsgID     int     `json:"id"`
	MsgBody   string  `json:"body"`
	MsgReplic []*TJob `json:"replies"`
	taskID    string
	job       string
	std       std.IStd
	msg       *TMsg
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
		msg:    &TMsg{},
		std:    std.GetStd(),
	}
	return job, nil
}

// Run -- вполняет всю работу по оформлению и отправке ответа (отдельный поток!)
func (sf *TJob) Run(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		if err := json.Unmarshal([]byte(sf.job), sf.msg); err != nil {
			logrus.WithError(err).Fatalf("TCliApp.unmarshalBody(): unmarshal JSON")
		}
		sf.fillMsg(sf.msg)
		binJSON, err := json.MarshalIndent(sf, "", "\t")
		if err != nil {
			logrus.WithError(err).WithField("taskID", sf.taskID).Fatalf("TJob.Run()")
		}
		// Записать сообщение в выходной поток
		strOut := "\nRESP:" + sf.taskID + "\n"
		strOut += string(binJSON) + "\n"
		strOut += "RESPEND\n"
		if err = sf.std.Write(strOut); err != nil {
			logrus.WithError(err).WithField("taskID", sf.taskID).Fatalf("TJob.Run()")
		}
		return
	}
}

// Ходит на сервис, что-то там забирает
func (sf *TJob) getURL() {
	url := "https://25.ms/posts/" + fmt.Sprint(sf.MsgID)
	resp, err := http.Get(url) //nolint
	if err != nil {
		logrus.WithField("url", url).WithError(err).Fatalf("TJob.getURL()")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		logrus.WithField(
			"code", resp.StatusCode,
		).WithField("msg", resp.Status).WithField("url", url).Fatalf("TJob.getURL()")
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Fatalf("TJob.getURL()")
	}
	sf.MsgBody = string(bytes)
	sf.unmarshalBody(bytes)
}

// Десереализует сообщение после GET-зпроса
func (sf *TJob) unmarshalBody(bytes []byte) {
	if err := json.Unmarshal(bytes, sf); err != nil {
		logrus.WithError(err).Fatalf("TCliApp.unmarshalBody(): unmarshal JSON")
	}
}

// Рекурсивно заполняет сообщения
func (sf *TJob) fillMsg(msg *TMsg) {
	sf.MsgID = msg.ID
	sf.getURL()
	for len(sf.MsgReplic) < len(msg.Replic) {
		strID := fmt.Sprint(msg.Replic[len(sf.MsgReplic)].ID)
		job, _ := New(strID, " ")
		job.MsgID = msg.Replic[len(sf.MsgReplic)].ID
		sf.MsgReplic = append(sf.MsgReplic, job)
	}
	for ind, job := range sf.MsgReplic {
		job.fillMsg(msg.Replic[ind])
	}
}
