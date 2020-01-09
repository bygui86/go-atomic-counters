package atomics

import (
	"errors"
	"fmt"

	"github.com/robfig/cron/v3"
)

type JobLogger struct {
	interval     int
	intervalUnit string
	cronExpr     string
	cronJob      *cron.Cron
	cronEntryId  cron.EntryID
}

func InitJobLogger(interval int, intervalUnit string) (*JobLogger, error) {
	cronJob := cron.New()
	cronExpr, cronErr := getCronExpr(interval, intervalUnit)
	if cronErr != nil {
		return nil, cronErr
	}
	entryId, funcErr := cronJob.AddFunc(cronExpr, jobLogger)
	if funcErr != nil {
		return nil, funcErr
	}
	return &JobLogger{
		interval:    interval,
		cronExpr:    cronExpr,
		cronJob:     cronJob,
		cronEntryId: entryId,
	}, nil
}

func (jl *JobLogger) Start() {
	jl.cronJob.Start()
}

func (jl *JobLogger) Stop() {
	jl.cronJob.Stop()
}

func jobLogger() {
	tempCounts := TotalCounts.Get()
	TotalCounts.Reset()
	fmt.Println("Counts:", tempCounts)
}

func getCronExpr(interval int, unit string) (string, error) {
	if interval > 0 {
		return fmt.Sprintf("@every %d%s", interval, unit), nil
	}
	return "", errors.New("interval must be positive")
}
