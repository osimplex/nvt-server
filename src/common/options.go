package common

import (
	"errors"
	"sync"

	"github.com/BurntSushi/toml"
	val "github.com/go-playground/validator/v10"
)

var (
	ErrConfigTomlParsing = errors.New("couldn't parse the config file")
	ErrConfigValidation  = errors.New("invalid config file")

	validator *val.Validate
)

type Config struct {
	Mu            sync.Mutex
	Handler       string                   `validate:"required,alpha"`
	Worker        string                   `validate:"omitempty,alpha"`
	Server        string                   `validate:"required,alpha"`
	LogMode       int                      `validate:"required,gte=0,lte=3"`
	ServerOpts    map[string]ServerOpts    `validate:"required,dive"`
	HandlerOpts   map[string]HandlerOpts   `validate:"dive"`
	WorkerOpts    map[string]WorkerOpts    `validate:"dive"`
	WorkerAliases map[string]WorkerAliases `validate:"dive"`
}

func LoadConfig(configFile string) (*Config, error) {
	validator = val.New(val.WithRequiredStructEnabled())
	config := &Config{}

	config.Mu.Lock()

	_, err := toml.DecodeFile(configFile, config)
	if err != nil {
		return nil, ErrConfigTomlParsing
	}

	if err := validator.Struct(config); err != nil {
		return nil, ErrConfigValidation
	}

	return config, nil
}

type ServerOpts struct {
	Addr string            `validate:"required,hostname_port"`
	Opts map[string]string `validate:"dive,keys,endkeys,alpha"`
}

type HandlerOpts struct {
	Opts map[string]string `validate:"dive,keys,endkeys,printascii"`
}

type WorkerAliases struct {
	FrameSource        Frame
	FrameCommandSource FrameCommand
	ProcedureSource    Procedure
}

type WorkerOpts struct {
	Opts map[string]string `validate:"dive,keys,endkeys,printascii"`
}

type Frame struct {
	Source string `validate:"printascii"`

	FrameNameAlias  string `validate:"printascii"`
	ClassAlias      string `validate:"printascii"`
	DeviceTypeAlias string `validate:"printascii"`
	InputFrameAlias string `validate:"printascii"`

	Line1Alias string `validate:"printascii"`
	Line2Alias string `validate:"printascii"`
	Line3Alias string `validate:"printascii"`
	Line4Alias string `validate:"printascii"`
	Line5Alias string `validate:"printascii"`
	Line6Alias string `validate:"printascii"`
	Line7Alias string `validate:"printascii"`
	Line8Alias string `validate:"printascii"`
}

type FrameCommand struct {
	Source string `validate:"printascii"`

	FrameNameAlias   string `validate:"printascii"`
	CommandAlias     string `validate:"printascii"`
	TransactionAlias string `validate:"printascii"`
	FrameRefAlias    string `validate:"printascii"`
}

type Procedure struct {
	Source string `validate:"printascii"`

	ClientAddrAlias        string `validate:"printascii"`
	TransactionAlias       string `validate:"printascii"`
	InputAlias             string `validate:"printascii"`
	SearchTypeAlias        string `validate:"printascii"`
	BeepAlias              string `validate:"printascii"`
	FrameNameAlias         string `validate:"printascii"`
	PreviousFrameNameAlias string `validate:"printascii"`

	Rv01Alias string `validate:"required,printascii"`
	Rv02Alias string `validate:"required,printascii"`
	Rv03Alias string `validate:"printascii"`
	Rv04Alias string `validate:"printascii"`
	Rv05Alias string `validate:"printascii"`
	Rv06Alias string `validate:"printascii"`
	Rv07Alias string `validate:"printascii"`
	Rv08Alias string `validate:"printascii"`
	Rv09Alias string `validate:"printascii"`
	Rv10Alias string `validate:"printascii"`
	Rv11Alias string `validate:"printascii"`
	Rv12Alias string `validate:"printascii"`
	Rv13Alias string `validate:"printascii"`
	Rv14Alias string `validate:"printascii"`
	Rv15Alias string `validate:"printascii"`
	Rv16Alias string `validate:"printascii"`
	Rv17Alias string `validate:"printascii"`
	Rv18Alias string `validate:"printascii"`
	Rv19Alias string `validate:"printascii"`
	Rv20Alias string `validate:"printascii"`
	Rv21Alias string `validate:"printascii"`
	Rv22Alias string `validate:"printascii"`
	Rv23Alias string `validate:"printascii"`
	Rv24Alias string `validate:"printascii"`
	Rv25Alias string `validate:"printascii"`
	Rv26Alias string `validate:"printascii"`
	Rv27Alias string `validate:"printascii"`
	Rv28Alias string `validate:"printascii"`
	Rv29Alias string `validate:"printascii"`
	Rv30Alias string `validate:"printascii"`
	Rv31Alias string `validate:"printascii"`
	Rv32Alias string `validate:"printascii"`

	Rl1Alias string `validate:"printascii"`
	Rl2Alias string `validate:"printascii"`
	Rl3Alias string `validate:"printascii"`
	Rl4Alias string `validate:"printascii"`
	Rl5Alias string `validate:"printascii"`
	Rl6Alias string `validate:"printascii"`
	Rl7Alias string `validate:"printascii"`
	Rl8Alias string `validate:"printascii"`
}
