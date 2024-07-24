package oracle

import (
	"fmt"
	"nvt-server/src/model"
	"nvt-server/src/tracer"
	"strings"
)

func (oracle *Oracle) GetFrame(
	obj *model.Frame,
	tracer tracer.Tracer,
) error {
	config := oracle.Config.WorkerAliases[oracle.Config.Worker].FrameSource

	sbQueryString := strings.Builder{}
	sbQueryString.WriteString(fmt.Sprintf("SELECT %s, %s,\n", config.Line1Alias, config.Line2Alias))
	sbQueryString.WriteString(fmt.Sprintf("       %s, %s,\n", config.Line3Alias, config.Line4Alias))
	sbQueryString.WriteString(fmt.Sprintf("       %s, %s,\n", config.Line5Alias, config.Line6Alias))
	sbQueryString.WriteString(fmt.Sprintf("       %s, %s,\n", config.Line7Alias, config.Line8Alias))
	sbQueryString.WriteString(fmt.Sprintf("       %s, %s,\n", config.ClassAlias, config.DeviceTypeAlias))
	sbQueryString.WriteString(fmt.Sprintf("       %s\n", config.InputFrameAlias))
	sbQueryString.WriteString(fmt.Sprintf("  FROM %s\n", config.Source))
	sbQueryString.WriteString(fmt.Sprintf(" WHERE %s=:FrameName", config.FrameNameAlias))

	stmt, err := oracle.DB.Prepare(sbQueryString.String())
	if err != nil {
		tracer.Error(err)
		return err
	}
	defer stmt.Close()

	dto := oracle.DtoFrame.GetDto()
	dto.SetDto(obj)

	sqlReturn := stmt.QueryRow(dto.FrameName.String)
	if err := sqlReturn.Err(); err != nil {
		tracer.Error(err)
		return err
	}

	err = sqlReturn.Scan(
		&dto.Ln1, &dto.Ln2, &dto.Ln3, &dto.Ln4,
		&dto.Ln5, &dto.Ln6, &dto.Ln7, &dto.Ln8,
		&dto.Class,
		&dto.DeviceType,
		&dto.InputFrame,
	)
	if err != nil {
		tracer.Error(err)
		return err
	}

	dto.SetObject(obj)

	return nil
}
