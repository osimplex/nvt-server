package oracle

import (
	"database/sql"
	"fmt"
	"nvt-server/src/model"
	"nvt-server/src/tracer"
	"strings"
)

func (oracle *Oracle) GetFrameCommand(
	obj *model.FrameCommand,
	tracer tracer.Tracer,
) error {
	config := oracle.Config.WorkerAliases[oracle.Config.Worker].FrameCommandSource

	sbQueryString := strings.Builder{}
	sbQueryString.WriteString(fmt.Sprintf("SELECT %s, %s\n", config.TransactionAlias, config.FrameRefAlias))
	sbQueryString.WriteString(fmt.Sprintf("  FROM %s\n", config.Source))
	sbQueryString.WriteString(fmt.Sprintf(" WHERE %s=:FrameName", config.FrameNameAlias))
	sbQueryString.WriteString(fmt.Sprintf("   AND %s=:Command", config.CommandAlias))

	stmt, err := oracle.DB.Prepare(sbQueryString.String())
	if err != nil {
		tracer.Error(err)
		return err
	}
	defer stmt.Close()

	dto := oracle.DtoFrameCommand.GetDto()
	dto.SetDto(obj)

	sqlReturn := stmt.QueryRow(dto.FrameName.String, dto.Command.String)
	if err := sqlReturn.Err(); err != nil {
		tracer.Error(err)
		return err
	}

	err = sqlReturn.Scan(
		&dto.Transaction,
		&dto.FrameRef,
	)
	switch {
	case err == sql.ErrNoRows:
		dto.Transaction.Valid = false
		dto.FrameRef.Valid = false
	case err != nil:
		tracer.Error(err)
		return err
	}

	dto.SetObject(obj)

	return nil
}
