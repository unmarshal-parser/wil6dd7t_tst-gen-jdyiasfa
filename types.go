// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"sync"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetPairCreatedEventHash() string {
	return "0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9"
}

type PairCreatedEvent struct {
	EventToken0 string
	EventToken1 string
	EventPair   string
	EventArg3   decimal.Decimal `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:7848ba3d-11c6-4e07-b172-ea17ea532c95,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:7848ba3d-11c6-4e07-b172-ea17ea532c95,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:7848ba3d-11c6-4e07-b172-ea17ea532c95,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetCreatePairMethodHash() string {
	return "c9c65396"
}

type CreatePairMethod struct {
	MethodTokenA string
	MethodTokenB string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:18510ab3-39b6-43ca-b47f-0b17d6b04478,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:18510ab3-39b6-43ca-b47f-0b17d6b04478,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSetFeeToMethodHash() string {
	return "f46901ed"
}

type SetFeeToMethod struct {
	MethodFeeTo string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:bcb6645a-0630-4446-889a-831cc78d888e,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:bcb6645a-0630-4446-889a-831cc78d888e,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSetFeeToSetterMethodHash() string {
	return "a2e74af6"
}

type SetFeeToSetterMethod struct {
	MethodFeeToSetter string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:4a634561-241b-492d-b114-86a6c8158b1f,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:4a634561-241b-492d-b114-86a6c8158b1f,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models
type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
}

var tokenCache = sync.Map{}

// Config
type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
}

type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
	PrometheusConfig  `mapstructure:"prometheus_config"`
	Metrics           *Metrics
}

type PrometheusConfig struct {
	PrometheusNameSpace string //PrometheusNameSpace The prometheus namespace metrics will be reported on. Default -> custom_indexer
	PrometheusSubsystem string //PrometheusSubsystem The prometheus subsystem metrics will be reported on. Default -> parser
	MetricsPort         int    `mapstructure:"metrics_port"`    //MetricsPort port metrics is scraped on. Default -> 9091
	PrometheusHost      string `mapstructure:"prometheus_host"` //PrometheusHost Host for the metrics server. Default -> 0.0.0.0
}

type Metrics struct {
	//List of metrics
	NodeRestarts         *prometheus.CounterVec   //NodeRestarts node restarts specific to this parser
	NodeLatency          *prometheus.HistogramVec //NodeLatency Response latency for node
	UnmarshalAPIRestarts *prometheus.CounterVec   //UnmarshalAPIRestarts api gateway restarts specific to this parser
	UnmarshalAPILatency  *prometheus.HistogramVec //UnmarshalAPILatency Response Latency for unmarshal API calls

}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
	if i.ParallelCalls == 0 {
		i.ParallelCalls = 1
	}
	if i.MetricsPort == 0 {
		i.MetricsPort = 9091
	}
	if i.PrometheusHost == "" {
		i.PrometheusHost = "0.0.0.0"
	}

	i.PrometheusNameSpace = "custom_indexer"
	i.PrometheusSubsystem = "parser"
}
