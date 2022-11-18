// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"encoding/json"
	sdkTransactionTypes "github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
	"github.com/onrik/ethrpc"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = ethrpc.Transaction{}
	_ = time.Time{}
	_ = strings.Builder{}
	_ = sdkTransactionTypes.RawTransaction{}
)

func getJSONFromInterface(data interface{}) datatypes.JSON {
	var (
		err  error
		temp datatypes.JSON
	)
	temp, err = json.Marshal(data)
	if err != nil {
		log.Error("Error Marshalling Data: " + err.Error())
	}
	return temp
}

func convertToPairCreatedEvent(abiEvent *ABIPairCreatedEvent,
	transaction sdkTransactionTypes.TxnByID, chainID string) PairCreatedEvent {
	return PairCreatedEvent{
		EventToken0: strings.ToLower(abiEvent.Token0.String()),
		EventToken1: strings.ToLower(abiEvent.Token1.String()),
		EventPair:   strings.ToLower(abiEvent.Pair.String()),
		EventArg3:   decimal.NewFromBigInt(abiEvent.Arg3, 0),

		BlockHash:       strings.ToLower(abiEvent.Raw.BlockHash),
		BlockNumber:     uint64(abiEvent.Raw.BlockNumber),
		BlockTime:       time.Unix(int64(transaction.Date), 0),
		ChainID:         chainID,
		ContractAddress: strings.ToLower(abiEvent.Raw.Address),
		Gas:             getDecimalFromString(transaction.Fee),
		GasPrice:        getDecimalFromString(transaction.GasPrice),
		Index:           uint(abiEvent.Raw.LogIndex),
		TxFrom:          strings.ToLower(transaction.From),
		TxHash:          strings.ToLower(abiEvent.Raw.TransactionHash),
		TxIndex:         uint(abiEvent.Raw.TransactionIndex),
		TxTo:            strings.ToLower(transaction.To),
		TxValue:         getDecimalFromString(transaction.Value),
	}
}

func convertToCreatePairMethod(abiMethod *ABICreatePairMethod, chainID string) CreatePairMethod {
	return CreatePairMethod{
		MethodTokenA: strings.ToLower(abiMethod.TokenA.String()),
		MethodTokenB: strings.ToLower(abiMethod.TokenB.String()),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToSetFeeToMethod(abiMethod *ABISetFeeToMethod, chainID string) SetFeeToMethod {
	return SetFeeToMethod{
		MethodFeeTo: strings.ToLower(abiMethod.FeeTo.String()),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToSetFeeToSetterMethod(abiMethod *ABISetFeeToSetterMethod, chainID string) SetFeeToSetterMethod {
	return SetFeeToSetterMethod{
		MethodFeeToSetter: strings.ToLower(abiMethod.FeeToSetter.String()),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func getUint64FromString(numberString string) uint64 {
	number, err := strconv.ParseUint(numberString, 10, 64)
	if err != nil {
		return 0
	}
	return number
}

func getDecimalFromString(numberString string) decimal.Decimal {
	number, err := decimal.NewFromString(numberString)
	if err != nil {
		return decimal.NewFromInt(0)
	}
	return number
}
