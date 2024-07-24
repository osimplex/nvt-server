package worker

import (
	"context"
	"errors"
	"nvt-server/src/common"
	"nvt-server/src/db"
	sql_dto "nvt-server/src/dto/sql"
	oracle_worker "nvt-server/src/worker/oracle"
	"time"
)

var (
	ErrMissingWorkerOpts = errors.New("required worker opts not found for current worker")
	ErrPing              = errors.New("worker server unrecheable")
	ErrUnknownWorker     = errors.New("unknown worker requested")
)

func Factory(config *common.Config) (Worker, error) {
	switch config.Worker {
	case "oracle":
		if config.WorkerOpts[config.Worker].Opts == nil {
			return nil, ErrMissingWorkerOpts
		}

		dbStructure := db.GetOracleStructure(
			config.WorkerOpts[config.Worker].Opts,
		)

		dbPoll, err := dbStructure.GetDatabasePoll()
		if err != nil {
			return nil, err
		}

		pingCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := dbPoll.PingContext(pingCtx); err != nil {
			return nil, ErrPing
		}

		return &oracle_worker.Oracle{
			DtoFrame:        &sql_dto.Frame{},
			DtoFrameCommand: &sql_dto.FrameCommand{},
			DtoSession:      &sql_dto.Session{},
			Config:          config,
			DB:              dbPoll,
		}, nil
	}

	return nil, ErrUnknownWorker
}
