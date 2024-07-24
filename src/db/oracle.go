package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	go_ora "github.com/sijms/go-ora/v2"
)

var (
	ErrDatabaseSetup    = errors.New("database setup error")
	ErrTnsNamesReadFile = errors.New("couldn't read TNS names file")
)

type oracleStructure struct {
	tnsNamesFile    string
	user            string
	pass            string
	alias           string
	maxIdleConns    int
	maxOpenConns    int
	connMaxIdleTime int
}

func GetOracleStructure(config map[string]string) DatabaseStructure {
	maxIdleConns, _ := strconv.Atoi(config["MaxIdleConns"])
	maxOpenConns, _ := strconv.Atoi(config["MaxOpenConns"])
	connMaxIdleTime, _ := strconv.Atoi(config["ConnMaxIdleMin"])

	return &oracleStructure{
		tnsNamesFile:    filepath.Clean(config["TnsNames"]),
		user:            config["User"],
		pass:            config["Password"],
		alias:           config["Alias"],
		maxIdleConns:    maxIdleConns,
		maxOpenConns:    maxOpenConns,
		connMaxIdleTime: connMaxIdleTime,
	}
}

func (oracle *oracleStructure) readTnsNames(connName string) (string, error) {
	file, err := os.ReadFile(oracle.tnsNamesFile)
	if err != nil {
		return "", ErrTnsNamesReadFile
	}

	tns := string(file)
	_, tns, _ = strings.Cut(tns, connName+" =")
	tns, _, _ = strings.Cut(tns, "\r\n\r\n")

	return tns, nil
}

func (oracle *oracleStructure) GetDatabasePoll() (*sql.DB, error) {
	connString, err := oracle.readTnsNames(oracle.alias)
	if err != nil {
		return nil, err
	}

	connUrl := go_ora.BuildJDBC(
		oracle.user,
		oracle.pass,
		connString,
		nil,
	)

	poll, err := sql.Open("oracle", connUrl)
	if err != nil {
		return nil, ErrDatabaseSetup
	}

	maxIdleTime, _ := time.ParseDuration(fmt.Sprintf("%dm", oracle.connMaxIdleTime))

	poll.SetConnMaxIdleTime(maxIdleTime)
	poll.SetMaxIdleConns(oracle.maxIdleConns)
	poll.SetMaxOpenConns(oracle.maxOpenConns)

	return poll, nil
}
