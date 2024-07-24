package oracle

import (
	"database/sql"
	"fmt"
	sql_dto "nvt-server/src/dto/sql"
	"nvt-server/src/model"
	"nvt-server/src/tracer"
	"strings"

	go_ora "github.com/sijms/go-ora/v2"
)

func (oracle *Oracle) ProcedureCall(
	obj *model.Session,
	tracer tracer.Tracer,
) error {
	stmt, err := oracle.DB.Prepare(oracle.buildQuery())
	if err != nil {
		tracer.Errorf("transaction prepared statement error: %q", err)
		return err
	}
	defer stmt.Close()

	dto := oracle.DtoSession.GetDto()
	dto.SetDto(obj)

	setValid(&dto)
	err = execQuery(stmt, &dto)
	if err != nil {
		tracer.Errorf("transaction execution error: %q", err)
		return err
	}

	dto.SetObject(obj)

	return nil
}

func (oracle *Oracle) buildQuery() string {
	sbQueryString := strings.Builder{}
	config := oracle.Config.WorkerAliases[oracle.Config.Worker].ProcedureSource

	sbQueryString.WriteString("BEGIN\n")

	sbQueryString.WriteString("	 procmlo_rf_mostratela(\n")
	sbQueryString.WriteString(fmt.Sprintf("    %s => :ClientAddr,\n", config.ClientAddrAlias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Transaction,\n", config.TransactionAlias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Input,\n", config.InputAlias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :SearchType,\n", config.SearchTypeAlias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Beep,\n", config.BeepAlias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :FrameName,\n", config.FrameNameAlias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :PreviousFrameName,\n", config.PreviousFrameNameAlias))

	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv01,\n", config.Rv01Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv02,\n", config.Rv02Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv03,\n", config.Rv03Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv04,\n", config.Rv04Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv05,\n", config.Rv05Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv06,\n", config.Rv06Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv07,\n", config.Rv07Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv08,\n", config.Rv08Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv09,\n", config.Rv09Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv10,\n", config.Rv10Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv11,\n", config.Rv11Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv12,\n", config.Rv12Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv13,\n", config.Rv13Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv14,\n", config.Rv14Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv15,\n", config.Rv15Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv16,\n", config.Rv16Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv17,\n", config.Rv17Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv18,\n", config.Rv18Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv19,\n", config.Rv19Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv20,\n", config.Rv20Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv21,\n", config.Rv21Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv22,\n", config.Rv22Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv23,\n", config.Rv23Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv24,\n", config.Rv24Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv25,\n", config.Rv25Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv26,\n", config.Rv26Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv27,\n", config.Rv27Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv28,\n", config.Rv28Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv29,\n", config.Rv29Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv30,\n", config.Rv30Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv31,\n", config.Rv31Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rv32,\n", config.Rv32Alias))

	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rl1,\n", config.Rl1Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rl2,\n", config.Rl2Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rl3,\n", config.Rl3Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rl4,\n", config.Rl4Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rl5,\n", config.Rl5Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rl6,\n", config.Rl6Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rl7,\n", config.Rl7Alias))
	sbQueryString.WriteString(fmt.Sprintf("    %s => :Rl8);\n", config.Rl8Alias))

	sbQueryString.WriteString("END;")

	return sbQueryString.String()
}

func execQuery(stmt *sql.Stmt, session *sql_dto.Session) error {
	_, err := stmt.Exec(
		&go_ora.Out{Dest: &session.ClientAddr.NullString, In: true, Size: 50},
		&go_ora.Out{Dest: &session.Transaction.NullString, In: true, Size: 20},
		&go_ora.Out{Dest: &session.Input.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.SearchType.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Beep.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.FrameName.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.PreviousFrameName.NullString, In: true, Size: 255},

		&go_ora.Out{Dest: &session.Register.Rv01.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv02.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv03.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv04.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv05.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv06.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv07.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv08.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv09.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv10.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv11.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv12.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv13.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv14.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv15.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv16.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv17.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv18.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv19.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv20.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv21.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv22.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv23.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv24.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv25.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv26.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv27.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv28.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv29.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv30.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv31.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.Register.Rv32.NullString, In: true, Size: 255},

		&go_ora.Out{Dest: &session.FrameLines.Rl1.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.FrameLines.Rl2.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.FrameLines.Rl3.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.FrameLines.Rl4.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.FrameLines.Rl5.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.FrameLines.Rl6.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.FrameLines.Rl7.NullString, In: true, Size: 255},
		&go_ora.Out{Dest: &session.FrameLines.Rl8.NullString, In: true, Size: 255},
	)

	return err
}

func setValid(session *sql_dto.Session) {
	session.ClientAddr.NullString.Valid = true
	session.Transaction.NullString.Valid = true
	session.Input.NullString.Valid = true
	session.SearchType.NullString.Valid = true
	session.Beep.NullString.Valid = true
	session.FrameName.NullString.Valid = true
	session.PreviousFrameName.NullString.Valid = true

	session.Register.Rv01.NullString.Valid = true
	session.Register.Rv02.NullString.Valid = true
	session.Register.Rv03.NullString.Valid = true
	session.Register.Rv04.NullString.Valid = true
	session.Register.Rv05.NullString.Valid = true
	session.Register.Rv06.NullString.Valid = true
	session.Register.Rv07.NullString.Valid = true
	session.Register.Rv08.NullString.Valid = true
	session.Register.Rv09.NullString.Valid = true
	session.Register.Rv10.NullString.Valid = true
	session.Register.Rv11.NullString.Valid = true
	session.Register.Rv12.NullString.Valid = true
	session.Register.Rv13.NullString.Valid = true
	session.Register.Rv14.NullString.Valid = true
	session.Register.Rv15.NullString.Valid = true
	session.Register.Rv16.NullString.Valid = true
	session.Register.Rv17.NullString.Valid = true
	session.Register.Rv18.NullString.Valid = true
	session.Register.Rv19.NullString.Valid = true
	session.Register.Rv20.NullString.Valid = true
	session.Register.Rv21.NullString.Valid = true
	session.Register.Rv22.NullString.Valid = true
	session.Register.Rv23.NullString.Valid = true
	session.Register.Rv24.NullString.Valid = true
	session.Register.Rv25.NullString.Valid = true
	session.Register.Rv26.NullString.Valid = true
	session.Register.Rv27.NullString.Valid = true
	session.Register.Rv28.NullString.Valid = true
	session.Register.Rv29.NullString.Valid = true
	session.Register.Rv30.NullString.Valid = true
	session.Register.Rv31.NullString.Valid = true
	session.Register.Rv32.NullString.Valid = true

	session.FrameLines.Rl1.NullString.Valid = true
	session.FrameLines.Rl2.NullString.Valid = true
	session.FrameLines.Rl3.NullString.Valid = true
	session.FrameLines.Rl4.NullString.Valid = true
	session.FrameLines.Rl5.NullString.Valid = true
	session.FrameLines.Rl6.NullString.Valid = true
	session.FrameLines.Rl7.NullString.Valid = true
	session.FrameLines.Rl8.NullString.Valid = true
}
