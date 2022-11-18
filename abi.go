// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	sdkTransactionTypes "github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
	"github.com/onrik/ethrpc"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = sdkTransactionTypes.RawTransaction{}
	_ = errors.New("")
	_ = hexutil.Big{}
	_ = ethrpc.Transaction{}
	_ = abi.ABI{}
	_ = abi.ABI{}
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

const MainABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

type MainFilterer struct {
	abi      *abi.ABI
	contract *bind.BoundContract
}

type ABIPairCreatedEvent struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    ethrpc.Log
}

func (_Main *MainFilterer) ParseABIPairCreatedEvent(log types.Log, parsedLog ethrpc.Log) (*ABIPairCreatedEvent, error) {
	e := new(ABIPairCreatedEvent)
	if err := _Main.contract.UnpackLog(e, "PairCreated", log); err != nil {
		return nil, err
	}
	e.Raw = parsedLog
	return e, nil
}

type ABICreatePairMethod struct {
	TokenA         common.Address
	TokenB         common.Address
	RawTransaction sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABICreatePairMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABICreatePairMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABICreatePairMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "createPair", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABISetFeeToMethod struct {
	FeeTo          common.Address
	RawTransaction sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABISetFeeToMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABISetFeeToMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABISetFeeToMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "setFeeTo", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABISetFeeToSetterMethod struct {
	FeeToSetter    common.Address
	RawTransaction sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABISetFeeToSetterMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABISetFeeToSetterMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABISetFeeToSetterMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "setFeeToSetter", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, contractAbi, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract, abi: contractAbi}, nil
}

func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, *abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), &parsed, nil
}

func (_Main MainFilterer) UnpackMethodIntoInterface(v interface{}, name string, data []byte) error {
	args := _Main.abi.Methods[name].Inputs
	unpacked, err := args.Unpack(data[4:])
	if err != nil {
		return err
	}
	return args.Copy(v, unpacked)
}
