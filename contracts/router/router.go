// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CosmoswapLibraryMetaData contains all meta data concerning the CosmoswapLibrary contract.
var CosmoswapLibraryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122004289597ffa66f8f0164cf0ca7dcd50332f7c52dc58088a47c1e962db6ca5b6164736f6c63430006060033",
}

// CosmoswapLibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmoswapLibraryMetaData.ABI instead.
var CosmoswapLibraryABI = CosmoswapLibraryMetaData.ABI

// CosmoswapLibraryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmoswapLibraryMetaData.Bin instead.
var CosmoswapLibraryBin = CosmoswapLibraryMetaData.Bin

// DeployCosmoswapLibrary deploys a new Ethereum contract, binding an instance of CosmoswapLibrary to it.
func DeployCosmoswapLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CosmoswapLibrary, error) {
	parsed, err := CosmoswapLibraryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmoswapLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmoswapLibrary{CosmoswapLibraryCaller: CosmoswapLibraryCaller{contract: contract}, CosmoswapLibraryTransactor: CosmoswapLibraryTransactor{contract: contract}, CosmoswapLibraryFilterer: CosmoswapLibraryFilterer{contract: contract}}, nil
}

// CosmoswapLibrary is an auto generated Go binding around an Ethereum contract.
type CosmoswapLibrary struct {
	CosmoswapLibraryCaller     // Read-only binding to the contract
	CosmoswapLibraryTransactor // Write-only binding to the contract
	CosmoswapLibraryFilterer   // Log filterer for contract events
}

// CosmoswapLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CosmoswapLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmoswapLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmoswapLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmoswapLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmoswapLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmoswapLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmoswapLibrarySession struct {
	Contract     *CosmoswapLibrary // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosmoswapLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmoswapLibraryCallerSession struct {
	Contract *CosmoswapLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CosmoswapLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmoswapLibraryTransactorSession struct {
	Contract     *CosmoswapLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CosmoswapLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CosmoswapLibraryRaw struct {
	Contract *CosmoswapLibrary // Generic contract binding to access the raw methods on
}

// CosmoswapLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmoswapLibraryCallerRaw struct {
	Contract *CosmoswapLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// CosmoswapLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmoswapLibraryTransactorRaw struct {
	Contract *CosmoswapLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmoswapLibrary creates a new instance of CosmoswapLibrary, bound to a specific deployed contract.
func NewCosmoswapLibrary(address common.Address, backend bind.ContractBackend) (*CosmoswapLibrary, error) {
	contract, err := bindCosmoswapLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmoswapLibrary{CosmoswapLibraryCaller: CosmoswapLibraryCaller{contract: contract}, CosmoswapLibraryTransactor: CosmoswapLibraryTransactor{contract: contract}, CosmoswapLibraryFilterer: CosmoswapLibraryFilterer{contract: contract}}, nil
}

// NewCosmoswapLibraryCaller creates a new read-only instance of CosmoswapLibrary, bound to a specific deployed contract.
func NewCosmoswapLibraryCaller(address common.Address, caller bind.ContractCaller) (*CosmoswapLibraryCaller, error) {
	contract, err := bindCosmoswapLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmoswapLibraryCaller{contract: contract}, nil
}

// NewCosmoswapLibraryTransactor creates a new write-only instance of CosmoswapLibrary, bound to a specific deployed contract.
func NewCosmoswapLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*CosmoswapLibraryTransactor, error) {
	contract, err := bindCosmoswapLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmoswapLibraryTransactor{contract: contract}, nil
}

// NewCosmoswapLibraryFilterer creates a new log filterer instance of CosmoswapLibrary, bound to a specific deployed contract.
func NewCosmoswapLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*CosmoswapLibraryFilterer, error) {
	contract, err := bindCosmoswapLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmoswapLibraryFilterer{contract: contract}, nil
}

// bindCosmoswapLibrary binds a generic wrapper to an already deployed contract.
func bindCosmoswapLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmoswapLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmoswapLibrary *CosmoswapLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmoswapLibrary.Contract.CosmoswapLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmoswapLibrary *CosmoswapLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmoswapLibrary.Contract.CosmoswapLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmoswapLibrary *CosmoswapLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmoswapLibrary.Contract.CosmoswapLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmoswapLibrary *CosmoswapLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmoswapLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmoswapLibrary *CosmoswapLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmoswapLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmoswapLibrary *CosmoswapLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmoswapLibrary.Contract.contract.Transact(opts, method, params...)
}

// CosmoswapRouter02MetaData contains all meta data concerning the CosmoswapRouter02 contract.
var CosmoswapRouter02MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WPLUG\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WPLUG\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityPLUG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityPLUG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityPLUGSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityPLUGWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactPLUGForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactPLUGForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForPLUG\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForPLUGSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapPLUGForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactPLUG\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"faa5e5e2": "WPLUG()",
		"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
		"206de5b9": "addLiquidityPLUG(address,uint256,uint256,uint256,address,uint256)",
		"c45a0155": "factory()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
		"c3700e55": "removeLiquidityPLUG(address,uint256,uint256,uint256,address,uint256)",
		"c5b21050": "removeLiquidityPLUGSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)",
		"aae97507": "removeLiquidityPLUGWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"89fe6c1f": "removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"e794b31b": "swapExactPLUGForTokens(uint256,address[],address,uint256)",
		"39940723": "swapExactPLUGForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256)",
		"03a618a9": "swapExactTokensForPLUG(uint256,uint256,address[],address,uint256)",
		"9d5148f6": "swapExactTokensForPLUGSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
		"5c11d795": "swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"ec301ba2": "swapPLUGForExactTokens(uint256,address[],address,uint256)",
		"9548ea81": "swapTokensForExactPLUG(uint256,uint256,address[],address,uint256)",
		"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
	},
	Bin: "0x60c060405234801561001057600080fd5b5060405162004769380380620047698339818101604052604081101561003557600080fd5b5080516020909101516001600160601b0319606092831b8116608052911b1660a05260805160601c60a05160601c6145e4620001856000398061015f5280610ce75280610f05528061105e52806110b452806110e8528061115c528061153f52806115d452806116485280611c4a5280611d7e5280611f08528061200252806120b8528061216a52806124f0528061252b528061262d52806126dc52806127d7528061291d52806129a55280612c665280612da95280612e315280612fa2525080610d755280610e4c5280610fe0528061109252806112be52806113f9528061167a528061198b5280611b725280611c285280611e0c528061214852806122ad52806125bf5280612769528061286a52806129d75280612b455280612cf95280612e6352806133dc528061341f52806138c4528061397252806139f25280613b5e5280613cdd52506145e46000f3fe60806040526004361061014f5760003560e01c80639d5148f6116100b6578063c5b210501161006f578063c5b2105014610a10578063d06ca61f14610a63578063e794b31b14610b18578063e8e3370014610b9c578063ec301ba214610bfe578063faa5e5e214610c8257610188565b80639d5148f6146107f0578063aae9750714610886578063ad615dec146108f9578063baa2abde1461092f578063c3700e551461098c578063c45a0155146109df57610188565b8063399407231161010857806339940723146105015780635c11d7951461058557806385f8c2591461061b5780638803dbee1461065157806389fe6c1f146106e75780639548ea811461075a57610188565b806303a618a91461018d578063054d50d4146102735780631f00ca74146102bb578063206de5b9146103705780632195995c146103d457806338ed17391461046b57610188565b3661018857336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461018657fe5b005b600080fd5b34801561019957600080fd5b50610223600480360360a08110156101b057600080fd5b813591602081013591810190606081016040820135600160201b8111156101d657600080fd5b8201836020820111156101e857600080fd5b803590602001918460208302840111600160201b8311171561020957600080fd5b91935091506001600160a01b038135169060200135610c97565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561025f578181015183820152602001610247565b505050509050019250505060405180910390f35b34801561027f57600080fd5b506102a96004803603606081101561029657600080fd5b5080359060208101359060400135610fc4565b60408051918252519081900360200190f35b3480156102c757600080fd5b50610223600480360360408110156102de57600080fd5b81359190810190604081016020820135600160201b8111156102ff57600080fd5b82018360208201111561031157600080fd5b803590602001918460208302840111600160201b8311171561033257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610fd9945050505050565b6103b6600480360360c081101561038657600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a0013561100f565b60408051938452602084019290925282820152519081900360600190f35b3480156103e057600080fd5b5061045260048036036101608110156103f857600080fd5b506001600160a01b038135811691602081013582169160408201359160608101359160808201359160a08101359091169060c08101359060e081013515159060ff61010082013516906101208101359061014001356112b4565b6040805192835260208301919091528051918290030190f35b34801561047757600080fd5b50610223600480360360a081101561048e57600080fd5b813591602081013591810190606081016040820135600160201b8111156104b457600080fd5b8201836020820111156104c657600080fd5b803590602001918460208302840111600160201b831117156104e757600080fd5b91935091506001600160a01b0381351690602001356113ae565b6101866004803603608081101561051757600080fd5b81359190810190604081016020820135600160201b81111561053857600080fd5b82018360208201111561054a57600080fd5b803590602001918460208302840111600160201b8311171561056b57600080fd5b91935091506001600160a01b0381351690602001356114f9565b34801561059157600080fd5b50610186600480360360a08110156105a857600080fd5b813591602081013591810190606081016040820135600160201b8111156105ce57600080fd5b8201836020820111156105e057600080fd5b803590602001918460208302840111600160201b8311171561060157600080fd5b91935091506001600160a01b038135169060200135611921565b34801561062757600080fd5b506102a96004803603606081101561063e57600080fd5b5080359060208101359060400135611b1a565b34801561065d57600080fd5b50610223600480360360a081101561067457600080fd5b813591602081013591810190606081016040820135600160201b81111561069a57600080fd5b8201836020820111156106ac57600080fd5b803590602001918460208302840111600160201b831117156106cd57600080fd5b91935091506001600160a01b038135169060200135611b27565b3480156106f357600080fd5b506102a9600480360361014081101561070b57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a08101359060c081013515159060ff60e08201351690610100810135906101200135611c20565b34801561076657600080fd5b50610223600480360360a081101561077d57600080fd5b813591602081013591810190606081016040820135600160201b8111156107a357600080fd5b8201836020820111156107b557600080fd5b803590602001918460208302840111600160201b831117156107d657600080fd5b91935091506001600160a01b038135169060200135611d2e565b3480156107fc57600080fd5b50610186600480360360a081101561081357600080fd5b813591602081013591810190606081016040820135600160201b81111561083957600080fd5b82018360208201111561084b57600080fd5b803590602001918460208302840111600160201b8311171561086c57600080fd5b91935091506001600160a01b038135169060200135611eba565b34801561089257600080fd5b5061045260048036036101408110156108aa57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a08101359060c081013515159060ff60e0820135169061010081013590610120013561213e565b34801561090557600080fd5b506102a96004803603606081101561091c57600080fd5b5080359060208101359060400135612252565b34801561093b57600080fd5b50610452600480360360e081101561095257600080fd5b506001600160a01b038135811691602081013582169160408201359160608101359160808201359160a08101359091169060c0013561225f565b34801561099857600080fd5b50610452600480360360c08110156109af57600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a001356124a3565b3480156109eb57600080fd5b506109f46125bd565b604080516001600160a01b039092168252519081900360200190f35b348015610a1c57600080fd5b506102a9600480360360c0811015610a3357600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a001356125e1565b348015610a6f57600080fd5b5061022360048036036040811015610a8657600080fd5b81359190810190604081016020820135600160201b811115610aa757600080fd5b820183602082011115610ab957600080fd5b803590602001918460208302840111600160201b83111715610ada57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550612762945050505050565b61022360048036036080811015610b2e57600080fd5b81359190810190604081016020820135600160201b811115610b4f57600080fd5b820183602082011115610b6157600080fd5b803590602001918460208302840111600160201b83111715610b8257600080fd5b91935091506001600160a01b03813516906020013561278f565b348015610ba857600080fd5b506103b66004803603610100811015610bc057600080fd5b506001600160a01b038135811691602081013582169160408201359160608101359160808201359160a08101359160c0820135169060e00135612ae2565b61022360048036036080811015610c1457600080fd5b81359190810190604081016020820135600160201b811115610c3557600080fd5b820183602082011115610c4757600080fd5b803590602001918460208302840111600160201b83111715610c6857600080fd5b91935091506001600160a01b038135169060200135612c1e565b348015610c8e57600080fd5b506109f4612fa0565b60608142811015610cdd576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001686866000198101818110610d1757fe5b905060200201356001600160a01b03166001600160a01b031614610d70576040805162461bcd60e51b815260206004820152601d602482015260008051602061447c833981519152604482015290519081900360640190fd5b610dce7f000000000000000000000000000000000000000000000000000000000000000089888880806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250612fc492505050565b91508682600184510381518110610de157fe5b60200260200101511015610e265760405162461bcd60e51b815260040180806020018281038252602b815260200180614584602b913960400191505060405180910390fd5b610ec486866000818110610e3657fe5b905060200201356001600160a01b031633610eaa7f00000000000000000000000000000000000000000000000000000000000000008a8a6000818110610e7857fe5b905060200201356001600160a01b03168b8b6001818110610e9557fe5b905060200201356001600160a01b0316613110565b85600081518110610eb757fe5b60200260200101516131d0565b610f038287878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525030925061332d915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d83600185510381518110610f4257fe5b60200260200101516040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015610f8057600080fd5b505af1158015610f94573d6000803e3d6000fd5b50505050610fb98483600185510381518110610fac57fe5b6020026020010151613573565b509695505050505050565b6000610fd184848461366b565b949350505050565b60606110067f0000000000000000000000000000000000000000000000000000000000000000848461375b565b90505b92915050565b60008060008342811015611058576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b6110868a7f00000000000000000000000000000000000000000000000000000000000000008b348c8c613893565b909450925060006110d87f00000000000000000000000000000000000000000000000000000000000000008c7f0000000000000000000000000000000000000000000000000000000000000000613110565b90506110e68b3383886131d0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561114157600080fd5b505af1158015611155573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb82866040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156111da57600080fd5b505af11580156111ee573d6000803e3d6000fd5b505050506040513d602081101561120457600080fd5b505161120c57fe5b806001600160a01b0316636a627842886040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050602060405180830381600087803b15801561126457600080fd5b505af1158015611278573d6000803e3d6000fd5b505050506040513d602081101561128e57600080fd5b50519250348410156112a6576112a633853403613573565b505096509650969350505050565b60008060006112e47f00000000000000000000000000000000000000000000000000000000000000008f8f613110565b90506000876112f3578c6112f7565b6000195b6040805163d505accf60e01b815233600482015230602482015260448101839052606481018c905260ff8a16608482015260a4810189905260c4810188905290519192506001600160a01b0384169163d505accf9160e48082019260009290919082900301818387803b15801561136d57600080fd5b505af1158015611381573d6000803e3d6000fd5b505050506113948f8f8f8f8f8f8f61225f565b809450819550505050509b509b9950505050505050505050565b606081428110156113f4576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b6114527f000000000000000000000000000000000000000000000000000000000000000089888880806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250612fc492505050565b9150868260018451038151811061146557fe5b602002602001015110156114aa5760405162461bcd60e51b815260040180806020018281038252602b815260200180614584602b913960400191505060405180910390fd5b6114ba86866000818110610e3657fe5b610fb98287878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525089925061332d915050565b804281101561153d576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168585600081811061157457fe5b905060200201356001600160a01b03166001600160a01b0316146115cd576040805162461bcd60e51b815260206004820152601d602482015260008051602061447c833981519152604482015290519081900360640190fd5b60003490507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561162d57600080fd5b505af1158015611641573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb6116c37f0000000000000000000000000000000000000000000000000000000000000000898960008181106116a657fe5b905060200201356001600160a01b03168a8a6001818110610e9557fe5b836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561171357600080fd5b505af1158015611727573d6000803e3d6000fd5b505050506040513d602081101561173d57600080fd5b505161174557fe5b60008686600019810181811061175757fe5b905060200201356001600160a01b03166001600160a01b03166370a08231866040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156117bc57600080fd5b505afa1580156117d0573d6000803e3d6000fd5b505050506040513d60208110156117e657600080fd5b505160408051602089810282810182019093528982529293506118289290918a918a918291850190849080828437600092019190915250899250613b07915050565b876118da828989600019810181811061183d57fe5b905060200201356001600160a01b03166001600160a01b03166370a08231896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156118a257600080fd5b505afa1580156118b6573d6000803e3d6000fd5b505050506040513d60208110156118cc57600080fd5b50519063ffffffff613e1216565b10156119175760405162461bcd60e51b815260040180806020018281038252602b815260200180614584602b913960400191505060405180910390fd5b5050505050505050565b8042811015611965576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b6119bd8585600081811061197557fe5b905060200201356001600160a01b0316336119b77f0000000000000000000000000000000000000000000000000000000000000000898960008181106116a657fe5b8a6131d0565b6000858560001981018181106119cf57fe5b905060200201356001600160a01b03166001600160a01b03166370a08231856040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015611a3457600080fd5b505afa158015611a48573d6000803e3d6000fd5b505050506040513d6020811015611a5e57600080fd5b50516040805160208881028281018201909352888252929350611aa0929091899189918291850190849080828437600092019190915250889250613b07915050565b866118da8288886000198101818110611ab557fe5b905060200201356001600160a01b03166001600160a01b03166370a08231886040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156118a257600080fd5b6000610fd1848484613e62565b60608142811015611b6d576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b611bcb7f00000000000000000000000000000000000000000000000000000000000000008988888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061375b92505050565b91508682600081518110611bdb57fe5b602002602001015111156114aa5760405162461bcd60e51b81526004018080602001828103825260278152602001806144ed6027913960400191505060405180910390fd5b600080611c6e7f00000000000000000000000000000000000000000000000000000000000000008d7f0000000000000000000000000000000000000000000000000000000000000000613110565b9050600086611c7d578b611c81565b6000195b6040805163d505accf60e01b815233600482015230602482015260448101839052606481018b905260ff8916608482015260a4810188905260c4810187905290519192506001600160a01b0384169163d505accf9160e48082019260009290919082900301818387803b158015611cf757600080fd5b505af1158015611d0b573d6000803e3d6000fd5b50505050611d1d8d8d8d8d8d8d6125e1565b9d9c50505050505050505050505050565b60608142811015611d74576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001686866000198101818110611dae57fe5b905060200201356001600160a01b03166001600160a01b031614611e07576040805162461bcd60e51b815260206004820152601d602482015260008051602061447c833981519152604482015290519081900360640190fd5b611e657f00000000000000000000000000000000000000000000000000000000000000008988888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061375b92505050565b91508682600081518110611e7557fe5b60200260200101511115610e265760405162461bcd60e51b81526004018080602001828103825260278152602001806144ed6027913960400191505060405180910390fd5b8042811015611efe576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001685856000198101818110611f3857fe5b905060200201356001600160a01b03166001600160a01b031614611f91576040805162461bcd60e51b815260206004820152601d602482015260008051602061447c833981519152604482015290519081900360640190fd5b611fa18585600081811061197557fe5b611fdf858580806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250309250613b07915050565b604080516370a0823160e01b815230600482015290516000916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016916370a0823191602480820192602092909190829003018186803b15801561204957600080fd5b505afa15801561205d573d6000803e3d6000fd5b505050506040513d602081101561207357600080fd5b50519050868110156120b65760405162461bcd60e51b815260040180806020018281038252602b815260200180614584602b913960400191505060405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561211c57600080fd5b505af1158015612130573d6000803e3d6000fd5b505050506119178482613573565b600080600061218e7f00000000000000000000000000000000000000000000000000000000000000008e7f0000000000000000000000000000000000000000000000000000000000000000613110565b905060008761219d578c6121a1565b6000195b6040805163d505accf60e01b815233600482015230602482015260448101839052606481018c905260ff8a16608482015260a4810189905260c4810188905290519192506001600160a01b0384169163d505accf9160e48082019260009290919082900301818387803b15801561221757600080fd5b505af115801561222b573d6000803e3d6000fd5b5050505061223d8e8e8e8e8e8e6124a3565b909f909e509c50505050505050505050505050565b6000610fd1848484613f52565b60008082428110156122a6576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b60006122d37f00000000000000000000000000000000000000000000000000000000000000008c8c613110565b604080516323b872dd60e01b81523360048201526001600160a01b03831660248201819052604482018d9052915192935090916323b872dd916064808201926020929091908290030181600087803b15801561232e57600080fd5b505af1158015612342573d6000803e3d6000fd5b505050506040513d602081101561235857600080fd5b50506040805163226bf2d160e21b81526001600160a01b03888116600483015282516000938493928616926389afcb44926024808301939282900301818787803b1580156123a557600080fd5b505af11580156123b9573d6000803e3d6000fd5b505050506040513d60408110156123cf57600080fd5b508051602090910151909250905060006123e98e8e613ffe565b509050806001600160a01b03168e6001600160a01b03161461240c57818361240f565b82825b90975095508a8710156124535760405162461bcd60e51b81526004018080602001828103825260268152602001806144566026913960400191505060405180910390fd5b898610156124925760405162461bcd60e51b81526004018080602001828103825260268152602001806144c76026913960400191505060405180910390fd5b505050505097509795505050505050565b60008082428110156124ea576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b612519897f00000000000000000000000000000000000000000000000000000000000000008a8a8a308a61225f565b90935091506125298986856140dc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d836040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561258f57600080fd5b505af11580156125a3573d6000803e3d6000fd5b505050506125b18583613573565b50965096945050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008142811015612627576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b612656887f0000000000000000000000000000000000000000000000000000000000000000898989308961225f565b604080516370a0823160e01b815230600482015290519194506126da92508a9187916001600160a01b038416916370a0823191602480820192602092909190829003018186803b1580156126a957600080fd5b505afa1580156126bd573d6000803e3d6000fd5b505050506040513d60208110156126d357600080fd5b50516140dc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d836040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561274057600080fd5b505af1158015612754573d6000803e3d6000fd5b50505050610fb98483613573565b60606110067f00000000000000000000000000000000000000000000000000000000000000008484612fc4565b606081428110156127d5576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168686600081811061280c57fe5b905060200201356001600160a01b03166001600160a01b031614612865576040805162461bcd60e51b815260206004820152601d602482015260008051602061447c833981519152604482015290519081900360640190fd5b6128c37f000000000000000000000000000000000000000000000000000000000000000034888880806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250612fc492505050565b915086826001845103815181106128d657fe5b6020026020010151101561291b5760405162461bcd60e51b815260040180806020018281038252602b815260200180614584602b913960400191505060405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db08360008151811061295757fe5b60200260200101516040518263ffffffff1660e01b81526004016000604051808303818588803b15801561298a57600080fd5b505af115801561299e573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb612a037f0000000000000000000000000000000000000000000000000000000000000000898960008181106116a657fe5b84600081518110612a1057fe5b60200260200101516040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015612a6757600080fd5b505af1158015612a7b573d6000803e3d6000fd5b505050506040513d6020811015612a9157600080fd5b5051612a9957fe5b612ad88287878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525089925061332d915050565b5095945050505050565b60008060008342811015612b2b576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b612b398c8c8c8c8c8c613893565b90945092506000612b6b7f00000000000000000000000000000000000000000000000000000000000000008e8e613110565b9050612b798d3383886131d0565b612b858c3383876131d0565b806001600160a01b0316636a627842886040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050602060405180830381600087803b158015612bdd57600080fd5b505af1158015612bf1573d6000803e3d6000fd5b505050506040513d6020811015612c0757600080fd5b5051949d939c50939a509198505050505050505050565b60608142811015612c64576040805162461bcd60e51b81526020600482015260186024820152600080516020614436833981519152604482015290519081900360640190fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031686866000818110612c9b57fe5b905060200201356001600160a01b03166001600160a01b031614612cf4576040805162461bcd60e51b815260206004820152601d602482015260008051602061447c833981519152604482015290519081900360640190fd5b612d527f00000000000000000000000000000000000000000000000000000000000000008888888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061375b92505050565b91503482600081518110612d6257fe5b60200260200101511115612da75760405162461bcd60e51b81526004018080602001828103825260278152602001806144ed6027913960400191505060405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db083600081518110612de357fe5b60200260200101516040518263ffffffff1660e01b81526004016000604051808303818588803b158015612e1657600080fd5b505af1158015612e2a573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb612e8f7f0000000000000000000000000000000000000000000000000000000000000000898960008181106116a657fe5b84600081518110612e9c57fe5b60200260200101516040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015612ef357600080fd5b505af1158015612f07573d6000803e3d6000fd5b505050506040513d6020811015612f1d57600080fd5b5051612f2557fe5b612f648287878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525089925061332d915050565b81600081518110612f7157fe5b6020026020010151341115612ad857612ad83383600081518110612f9157fe5b60200260200101513403613573565b7f000000000000000000000000000000000000000000000000000000000000000081565b606060028251101561301d576040805162461bcd60e51b815260206004820152601e60248201527f436f736d6f737761704c6962726172793a20494e56414c49445f504154480000604482015290519081900360640190fd5b815167ffffffffffffffff8111801561303557600080fd5b5060405190808252806020026020018201604052801561305f578160200160208202803683370190505b509050828160008151811061307057fe5b60200260200101818152505060005b6001835103811015613108576000806130c28786858151811061309e57fe5b60200260200101518786600101815181106130b557fe5b6020026020010151614246565b915091506130e48484815181106130d557fe5b6020026020010151838361366b565b8484600101815181106130f357fe5b6020908102919091010152505060010161307f565b509392505050565b600080600061311f8585613ffe565b604080516bffffffffffffffffffffffff19606094851b811660208084019190915293851b81166034830152825160288184030181526048830184528051908501206001600160f81b031960688401529a90941b9093166069840152607d8301989098527fccb93815acc1b2c4aeeb3c0f3ae3deda878867bc4c65d41707a525d9e1acedaa609d808401919091528851808403909101815260bd909201909752805196019590952095945050505050565b604080516001600160a01b0385811660248301528481166044830152606480830185905283518084039091018152608490920183526020820180516001600160e01b03166323b872dd60e01b17815292518251600094606094938a169392918291908083835b602083106132555780518252601f199092019160209182019101613236565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146132b7576040519150601f19603f3d011682016040523d82523d6000602084013e6132bc565b606091505b50915091508180156132ea5750805115806132ea57508080602001905160208110156132e757600080fd5b50515b6133255760405162461bcd60e51b81526004018080602001828103825260248152602001806145146024913960400191505060405180910390fd5b505050505050565b60005b600183510381101561356d5760008084838151811061334b57fe5b602002602001015185846001018151811061336257fe5b602002602001015191509150600061337a8383613ffe565b509050600087856001018151811061338e57fe5b60200260200101519050600080836001600160a01b0316866001600160a01b0316146133bc578260006133c0565b6000835b91509150600060028a510388106133d75788613418565b6134187f0000000000000000000000000000000000000000000000000000000000000000878c8b6002018151811061340b57fe5b6020026020010151613110565b90506134457f00000000000000000000000000000000000000000000000000000000000000008888613110565b6001600160a01b031663022c0d9f84848460006040519080825280601f01601f191660200182016040528015613482576020820181803683370190505b506040518563ffffffff1660e01b815260040180858152602001848152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156134f35781810151838201526020016134db565b50505050905090810190601f1680156135205780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561354257600080fd5b505af1158015613556573d6000803e3d6000fd5b505060019099019850613330975050505050505050565b50505050565b604080516000808252602082019092526001600160a01b0384169083906040518082805190602001908083835b602083106135bf5780518252601f1990920191602091820191016135a0565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114613621576040519150601f19603f3d011682016040523d82523d6000602084013e613626565b606091505b50509050806136665760405162461bcd60e51b81526004018080602001828103825260248152602001806145606024913960400191505060405180910390fd5b505050565b60008084116136ab5760405162461bcd60e51b815260040180806020018281038252602b81526020018061449c602b913960400191505060405180910390fd5b6000831180156136bb5750600082115b6136f65760405162461bcd60e51b81526004018080602001828103825260288152602001806145386028913960400191505060405180910390fd5b600061370a856103e563ffffffff61430d16565b9050600061371e828563ffffffff61430d16565b9050600061374483613738886103e863ffffffff61430d16565b9063ffffffff61437016565b905080828161374f57fe5b04979650505050505050565b60606002825110156137b4576040805162461bcd60e51b815260206004820152601e60248201527f436f736d6f737761704c6962726172793a20494e56414c49445f504154480000604482015290519081900360640190fd5b815167ffffffffffffffff811180156137cc57600080fd5b506040519080825280602002602001820160405280156137f6578160200160208202803683370190505b509050828160018351038151811061380a57fe5b60209081029190910101528151600019015b80156131085760008061384c8786600186038151811061383857fe5b60200260200101518786815181106130b557fe5b9150915061386e84848151811061385f57fe5b60200260200101518383613e62565b84600185038151811061387d57fe5b602090810291909101015250506000190161381c565b6040805163e6a4390560e01b81526001600160a01b03888116600483015287811660248301529151600092839283927f00000000000000000000000000000000000000000000000000000000000000009092169163e6a4390591604480820192602092909190829003018186803b15801561390d57600080fd5b505afa158015613921573d6000803e3d6000fd5b505050506040513d602081101561393757600080fd5b50516001600160a01b031614156139ea57604080516364e329cb60e11b81526001600160a01b038a81166004830152898116602483015291517f00000000000000000000000000000000000000000000000000000000000000009092169163c9c65396916044808201926020929091908290030181600087803b1580156139bd57600080fd5b505af11580156139d1573d6000803e3d6000fd5b505050506040513d60208110156139e757600080fd5b50505b600080613a187f00000000000000000000000000000000000000000000000000000000000000008b8b614246565b91509150816000148015613a2a575080155b15613a3a57879350869250613afa565b6000613a47898484613f52565b9050878111613a9a5785811015613a8f5760405162461bcd60e51b81526004018080602001828103825260268152602001806144c76026913960400191505060405180910390fd5b889450925082613af8565b6000613aa7898486613f52565b905089811115613ab357fe5b87811015613af25760405162461bcd60e51b81526004018080602001828103825260268152602001806144566026913960400191505060405180910390fd5b94508793505b505b5050965096945050505050565b60005b600183510381101561366657600080848381518110613b2557fe5b6020026020010151858460010181518110613b3c57fe5b6020026020010151915091506000613b548383613ffe565b5090506000613b847f00000000000000000000000000000000000000000000000000000000000000008585613110565b9050600080600080846001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015613bc557600080fd5b505afa158015613bd9573d6000803e3d6000fd5b505050506040513d6060811015613bef57600080fd5b5080516020909101516001600160701b0391821693501690506000806001600160a01b038a811690891614613c25578284613c28565b83835b91509150613c86828b6001600160a01b03166370a082318a6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156118a257600080fd5b9550613c9386838361366b565b945050505050600080856001600160a01b0316886001600160a01b031614613cbd57826000613cc1565b6000835b91509150600060028c51038a10613cd8578a613d0c565b613d0c7f0000000000000000000000000000000000000000000000000000000000000000898e8d6002018151811061340b57fe5b604080516000808252602082019283905263022c0d9f60e01b835260248201878152604483018790526001600160a01b038086166064850152608060848501908152845160a48601819052969750908c169563022c0d9f958a958a958a9591949193919260c486019290918190849084905b83811015613d96578181015183820152602001613d7e565b50505050905090810190601f168015613dc35780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015613de557600080fd5b505af1158015613df9573d6000803e3d6000fd5b50506001909b019a50613b0a9950505050505050505050565b80820382811115611009576040805162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b604482015290519081900360640190fd5b6000808411613ea25760405162461bcd60e51b815260040180806020018281038252602c8152602001806143e5602c913960400191505060405180910390fd5b600083118015613eb25750600082115b613eed5760405162461bcd60e51b81526004018080602001828103825260288152602001806145386028913960400191505060405180910390fd5b6000613f116103e8613f05868863ffffffff61430d16565b9063ffffffff61430d16565b90506000613f2b6103e5613f05868963ffffffff613e1216565b9050613f486001828481613f3b57fe5b049063ffffffff61437016565b9695505050505050565b6000808411613f925760405162461bcd60e51b81526004018080602001828103825260258152602001806143c06025913960400191505060405180910390fd5b600083118015613fa25750600082115b613fdd5760405162461bcd60e51b81526004018080602001828103825260288152602001806145386028913960400191505060405180910390fd5b82613fee858463ffffffff61430d16565b81613ff557fe5b04949350505050565b600080826001600160a01b0316846001600160a01b031614156140525760405162461bcd60e51b81526004018080602001828103825260258152602001806144116025913960400191505060405180910390fd5b826001600160a01b0316846001600160a01b031610614072578284614075565b83835b90925090506001600160a01b0382166140d5576040805162461bcd60e51b815260206004820152601e60248201527f436f736d6f737761704c6962726172793a205a45524f5f414444524553530000604482015290519081900360640190fd5b9250929050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b178152925182516000946060949389169392918291908083835b602083106141595780518252601f19909201916020918201910161413a565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146141bb576040519150601f19603f3d011682016040523d82523d6000602084013e6141c0565b606091505b50915091508180156141ee5750805115806141ee57508080602001905160208110156141eb57600080fd5b50515b61423f576040805162461bcd60e51b815260206004820152601f60248201527f5472616e7366657248656c7065723a205452414e534645525f4641494c454400604482015290519081900360640190fd5b5050505050565b60008060006142558585613ffe565b509050600080614266888888613110565b6001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561429e57600080fd5b505afa1580156142b2573d6000803e3d6000fd5b505050506040513d60608110156142c857600080fd5b5080516020909101516001600160701b0391821693501690506001600160a01b03878116908416146142fb5780826142fe565b81815b90999098509650505050505050565b60008115806143285750508082028282828161432557fe5b04145b611009576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b604482015290519081900360640190fd5b80820182811015611009576040805162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b604482015290519081900360640190fdfe436f736d6f737761704c6962726172793a20494e53554646494349454e545f414d4f554e54436f736d6f737761704c6962726172793a20494e53554646494349454e545f4f55545055545f414d4f554e54436f736d6f737761704c6962726172793a204944454e544943414c5f414444524553534553436f736d6f73776170526f757465723a20455850495245440000000000000000436f736d6f73776170526f757465723a20494e53554646494349454e545f415f414d4f554e54436f736d6f73776170526f757465723a20494e56414c49445f50415448000000436f736d6f737761704c6962726172793a20494e53554646494349454e545f494e5055545f414d4f554e54436f736d6f73776170526f757465723a20494e53554646494349454e545f425f414d4f554e54436f736d6f73776170526f757465723a204558434553534956455f494e5055545f414d4f554e545472616e7366657248656c7065723a205452414e534645525f46524f4d5f4641494c4544436f736d6f737761704c6962726172793a20494e53554646494349454e545f4c49515549444954595472616e7366657248656c7065723a20504c55475f5452414e534645525f4641494c4544436f736d6f73776170526f757465723a20494e53554646494349454e545f4f55545055545f414d4f554e54a26469706673582212207238c2e25c4310da16cd7991963d4f397a25eec52971eee60062cfb7adadd0ab64736f6c63430006060033",
}

// CosmoswapRouter02ABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmoswapRouter02MetaData.ABI instead.
var CosmoswapRouter02ABI = CosmoswapRouter02MetaData.ABI

// Deprecated: Use CosmoswapRouter02MetaData.Sigs instead.
// CosmoswapRouter02FuncSigs maps the 4-byte function signature to its string representation.
var CosmoswapRouter02FuncSigs = CosmoswapRouter02MetaData.Sigs

// CosmoswapRouter02Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmoswapRouter02MetaData.Bin instead.
var CosmoswapRouter02Bin = CosmoswapRouter02MetaData.Bin

// DeployCosmoswapRouter02 deploys a new Ethereum contract, binding an instance of CosmoswapRouter02 to it.
func DeployCosmoswapRouter02(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _WPLUG common.Address) (common.Address, *types.Transaction, *CosmoswapRouter02, error) {
	parsed, err := CosmoswapRouter02MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmoswapRouter02Bin), backend, _factory, _WPLUG)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmoswapRouter02{CosmoswapRouter02Caller: CosmoswapRouter02Caller{contract: contract}, CosmoswapRouter02Transactor: CosmoswapRouter02Transactor{contract: contract}, CosmoswapRouter02Filterer: CosmoswapRouter02Filterer{contract: contract}}, nil
}

// CosmoswapRouter02 is an auto generated Go binding around an Ethereum contract.
type CosmoswapRouter02 struct {
	CosmoswapRouter02Caller     // Read-only binding to the contract
	CosmoswapRouter02Transactor // Write-only binding to the contract
	CosmoswapRouter02Filterer   // Log filterer for contract events
}

// CosmoswapRouter02Caller is an auto generated read-only Go binding around an Ethereum contract.
type CosmoswapRouter02Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmoswapRouter02Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmoswapRouter02Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmoswapRouter02Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmoswapRouter02Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmoswapRouter02Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmoswapRouter02Session struct {
	Contract     *CosmoswapRouter02 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CosmoswapRouter02CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmoswapRouter02CallerSession struct {
	Contract *CosmoswapRouter02Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CosmoswapRouter02TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmoswapRouter02TransactorSession struct {
	Contract     *CosmoswapRouter02Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CosmoswapRouter02Raw is an auto generated low-level Go binding around an Ethereum contract.
type CosmoswapRouter02Raw struct {
	Contract *CosmoswapRouter02 // Generic contract binding to access the raw methods on
}

// CosmoswapRouter02CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmoswapRouter02CallerRaw struct {
	Contract *CosmoswapRouter02Caller // Generic read-only contract binding to access the raw methods on
}

// CosmoswapRouter02TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmoswapRouter02TransactorRaw struct {
	Contract *CosmoswapRouter02Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmoswapRouter02 creates a new instance of CosmoswapRouter02, bound to a specific deployed contract.
func NewCosmoswapRouter02(address common.Address, backend bind.ContractBackend) (*CosmoswapRouter02, error) {
	contract, err := bindCosmoswapRouter02(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmoswapRouter02{CosmoswapRouter02Caller: CosmoswapRouter02Caller{contract: contract}, CosmoswapRouter02Transactor: CosmoswapRouter02Transactor{contract: contract}, CosmoswapRouter02Filterer: CosmoswapRouter02Filterer{contract: contract}}, nil
}

// NewCosmoswapRouter02Caller creates a new read-only instance of CosmoswapRouter02, bound to a specific deployed contract.
func NewCosmoswapRouter02Caller(address common.Address, caller bind.ContractCaller) (*CosmoswapRouter02Caller, error) {
	contract, err := bindCosmoswapRouter02(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmoswapRouter02Caller{contract: contract}, nil
}

// NewCosmoswapRouter02Transactor creates a new write-only instance of CosmoswapRouter02, bound to a specific deployed contract.
func NewCosmoswapRouter02Transactor(address common.Address, transactor bind.ContractTransactor) (*CosmoswapRouter02Transactor, error) {
	contract, err := bindCosmoswapRouter02(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmoswapRouter02Transactor{contract: contract}, nil
}

// NewCosmoswapRouter02Filterer creates a new log filterer instance of CosmoswapRouter02, bound to a specific deployed contract.
func NewCosmoswapRouter02Filterer(address common.Address, filterer bind.ContractFilterer) (*CosmoswapRouter02Filterer, error) {
	contract, err := bindCosmoswapRouter02(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmoswapRouter02Filterer{contract: contract}, nil
}

// bindCosmoswapRouter02 binds a generic wrapper to an already deployed contract.
func bindCosmoswapRouter02(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmoswapRouter02ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmoswapRouter02 *CosmoswapRouter02Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmoswapRouter02.Contract.CosmoswapRouter02Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmoswapRouter02 *CosmoswapRouter02Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.CosmoswapRouter02Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmoswapRouter02 *CosmoswapRouter02Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.CosmoswapRouter02Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmoswapRouter02 *CosmoswapRouter02CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmoswapRouter02.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.contract.Transact(opts, method, params...)
}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() view returns(address)
func (_CosmoswapRouter02 *CosmoswapRouter02Caller) WPLUG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmoswapRouter02.contract.Call(opts, &out, "WPLUG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() view returns(address)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) WPLUG() (common.Address, error) {
	return _CosmoswapRouter02.Contract.WPLUG(&_CosmoswapRouter02.CallOpts)
}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() view returns(address)
func (_CosmoswapRouter02 *CosmoswapRouter02CallerSession) WPLUG() (common.Address, error) {
	return _CosmoswapRouter02.Contract.WPLUG(&_CosmoswapRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CosmoswapRouter02 *CosmoswapRouter02Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmoswapRouter02.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) Factory() (common.Address, error) {
	return _CosmoswapRouter02.Contract.Factory(&_CosmoswapRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CosmoswapRouter02 *CosmoswapRouter02CallerSession) Factory() (common.Address, error) {
	return _CosmoswapRouter02.Contract.Factory(&_CosmoswapRouter02.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_CosmoswapRouter02 *CosmoswapRouter02Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmoswapRouter02.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _CosmoswapRouter02.Contract.GetAmountIn(&_CosmoswapRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_CosmoswapRouter02 *CosmoswapRouter02CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _CosmoswapRouter02.Contract.GetAmountIn(&_CosmoswapRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_CosmoswapRouter02 *CosmoswapRouter02Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmoswapRouter02.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _CosmoswapRouter02.Contract.GetAmountOut(&_CosmoswapRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_CosmoswapRouter02 *CosmoswapRouter02CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _CosmoswapRouter02.Contract.GetAmountOut(&_CosmoswapRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _CosmoswapRouter02.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _CosmoswapRouter02.Contract.GetAmountsIn(&_CosmoswapRouter02.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _CosmoswapRouter02.Contract.GetAmountsIn(&_CosmoswapRouter02.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _CosmoswapRouter02.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _CosmoswapRouter02.Contract.GetAmountsOut(&_CosmoswapRouter02.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _CosmoswapRouter02.Contract.GetAmountsOut(&_CosmoswapRouter02.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmoswapRouter02.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _CosmoswapRouter02.Contract.Quote(&_CosmoswapRouter02.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _CosmoswapRouter02.Contract.Quote(&_CosmoswapRouter02.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.AddLiquidity(&_CosmoswapRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.AddLiquidity(&_CosmoswapRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) AddLiquidityPLUG(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "addLiquidityPLUG", token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) AddLiquidityPLUG(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.AddLiquidityPLUG(&_CosmoswapRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) AddLiquidityPLUG(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.AddLiquidityPLUG(&_CosmoswapRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidity(&_CosmoswapRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidity(&_CosmoswapRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) RemoveLiquidityPLUG(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "removeLiquidityPLUG", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) RemoveLiquidityPLUG(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityPLUG(&_CosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) RemoveLiquidityPLUG(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityPLUG(&_CosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc5b21050.
//
// Solidity: function removeLiquidityPLUGSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) RemoveLiquidityPLUGSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "removeLiquidityPLUGSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc5b21050.
//
// Solidity: function removeLiquidityPLUGSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) RemoveLiquidityPLUGSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityPLUGSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc5b21050.
//
// Solidity: function removeLiquidityPLUGSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) RemoveLiquidityPLUGSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityPLUGSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) RemoveLiquidityPLUGWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "removeLiquidityPLUGWithPermit", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) RemoveLiquidityPLUGWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityPLUGWithPermit(&_CosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) RemoveLiquidityPLUGWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityPLUGWithPermit(&_CosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x89fe6c1f.
//
// Solidity: function removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x89fe6c1f.
//
// Solidity: function removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x89fe6c1f.
//
// Solidity: function removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountPLUG)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityWithPermit(&_CosmoswapRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.RemoveLiquidityWithPermit(&_CosmoswapRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapExactPLUGForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapExactPLUGForTokens", amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapExactPLUGForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactPLUGForTokens(&_CosmoswapRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapExactPLUGForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactPLUGForTokens(&_CosmoswapRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x39940723.
//
// Solidity: function swapExactPLUGForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapExactPLUGForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapExactPLUGForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x39940723.
//
// Solidity: function swapExactPLUGForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapExactPLUGForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactPLUGForTokensSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x39940723.
//
// Solidity: function swapExactPLUGForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapExactPLUGForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactPLUGForTokensSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapExactTokensForPLUG(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapExactTokensForPLUG", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapExactTokensForPLUG(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactTokensForPLUG(&_CosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapExactTokensForPLUG(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactTokensForPLUG(&_CosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9d5148f6.
//
// Solidity: function swapExactTokensForPLUGSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapExactTokensForPLUGSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapExactTokensForPLUGSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9d5148f6.
//
// Solidity: function swapExactTokensForPLUGSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapExactTokensForPLUGSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactTokensForPLUGSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9d5148f6.
//
// Solidity: function swapExactTokensForPLUGSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapExactTokensForPLUGSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactTokensForPLUGSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactTokensForTokens(&_CosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactTokensForTokens(&_CosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_CosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapPLUGForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapPLUGForExactTokens", amountOut, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapPLUGForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapPLUGForExactTokens(&_CosmoswapRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapPLUGForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapPLUGForExactTokens(&_CosmoswapRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapTokensForExactPLUG(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapTokensForExactPLUG", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapTokensForExactPLUG(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapTokensForExactPLUG(&_CosmoswapRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapTokensForExactPLUG(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapTokensForExactPLUG(&_CosmoswapRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapTokensForExactTokens(&_CosmoswapRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.SwapTokensForExactTokens(&_CosmoswapRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmoswapRouter02 *CosmoswapRouter02Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmoswapRouter02.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmoswapRouter02 *CosmoswapRouter02Session) Receive() (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.Receive(&_CosmoswapRouter02.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmoswapRouter02 *CosmoswapRouter02TransactorSession) Receive() (*types.Transaction, error) {
	return _CosmoswapRouter02.Contract.Receive(&_CosmoswapRouter02.TransactOpts)
}

// ICosmoswapFactoryMetaData contains all meta data concerning the ICosmoswapFactory contract.
var ICosmoswapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1e3dd18b": "allPairs(uint256)",
		"574f2ba3": "allPairsLength()",
		"c9c65396": "createPair(address,address)",
		"017e7e58": "feeTo()",
		"094b7415": "feeToSetter()",
		"e6a43905": "getPair(address,address)",
		"f46901ed": "setFeeTo(address)",
		"a2e74af6": "setFeeToSetter(address)",
	},
}

// ICosmoswapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ICosmoswapFactoryMetaData.ABI instead.
var ICosmoswapFactoryABI = ICosmoswapFactoryMetaData.ABI

// Deprecated: Use ICosmoswapFactoryMetaData.Sigs instead.
// ICosmoswapFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ICosmoswapFactoryFuncSigs = ICosmoswapFactoryMetaData.Sigs

// ICosmoswapFactory is an auto generated Go binding around an Ethereum contract.
type ICosmoswapFactory struct {
	ICosmoswapFactoryCaller     // Read-only binding to the contract
	ICosmoswapFactoryTransactor // Write-only binding to the contract
	ICosmoswapFactoryFilterer   // Log filterer for contract events
}

// ICosmoswapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICosmoswapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICosmoswapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICosmoswapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICosmoswapFactorySession struct {
	Contract     *ICosmoswapFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ICosmoswapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICosmoswapFactoryCallerSession struct {
	Contract *ICosmoswapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ICosmoswapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICosmoswapFactoryTransactorSession struct {
	Contract     *ICosmoswapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ICosmoswapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICosmoswapFactoryRaw struct {
	Contract *ICosmoswapFactory // Generic contract binding to access the raw methods on
}

// ICosmoswapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICosmoswapFactoryCallerRaw struct {
	Contract *ICosmoswapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ICosmoswapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICosmoswapFactoryTransactorRaw struct {
	Contract *ICosmoswapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICosmoswapFactory creates a new instance of ICosmoswapFactory, bound to a specific deployed contract.
func NewICosmoswapFactory(address common.Address, backend bind.ContractBackend) (*ICosmoswapFactory, error) {
	contract, err := bindICosmoswapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapFactory{ICosmoswapFactoryCaller: ICosmoswapFactoryCaller{contract: contract}, ICosmoswapFactoryTransactor: ICosmoswapFactoryTransactor{contract: contract}, ICosmoswapFactoryFilterer: ICosmoswapFactoryFilterer{contract: contract}}, nil
}

// NewICosmoswapFactoryCaller creates a new read-only instance of ICosmoswapFactory, bound to a specific deployed contract.
func NewICosmoswapFactoryCaller(address common.Address, caller bind.ContractCaller) (*ICosmoswapFactoryCaller, error) {
	contract, err := bindICosmoswapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapFactoryCaller{contract: contract}, nil
}

// NewICosmoswapFactoryTransactor creates a new write-only instance of ICosmoswapFactory, bound to a specific deployed contract.
func NewICosmoswapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ICosmoswapFactoryTransactor, error) {
	contract, err := bindICosmoswapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapFactoryTransactor{contract: contract}, nil
}

// NewICosmoswapFactoryFilterer creates a new log filterer instance of ICosmoswapFactory, bound to a specific deployed contract.
func NewICosmoswapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ICosmoswapFactoryFilterer, error) {
	contract, err := bindICosmoswapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapFactoryFilterer{contract: contract}, nil
}

// bindICosmoswapFactory binds a generic wrapper to an already deployed contract.
func bindICosmoswapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICosmoswapFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmoswapFactory *ICosmoswapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmoswapFactory.Contract.ICosmoswapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmoswapFactory *ICosmoswapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.ICosmoswapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmoswapFactory *ICosmoswapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.ICosmoswapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmoswapFactory *ICosmoswapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmoswapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmoswapFactory *ICosmoswapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmoswapFactory *ICosmoswapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _ICosmoswapFactory.Contract.AllPairs(&_ICosmoswapFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _ICosmoswapFactory.Contract.AllPairs(&_ICosmoswapFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_ICosmoswapFactory *ICosmoswapFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_ICosmoswapFactory *ICosmoswapFactorySession) AllPairsLength() (*big.Int, error) {
	return _ICosmoswapFactory.Contract.AllPairsLength(&_ICosmoswapFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_ICosmoswapFactory *ICosmoswapFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _ICosmoswapFactory.Contract.AllPairsLength(&_ICosmoswapFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ICosmoswapFactory *ICosmoswapFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ICosmoswapFactory *ICosmoswapFactorySession) FeeTo() (common.Address, error) {
	return _ICosmoswapFactory.Contract.FeeTo(&_ICosmoswapFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ICosmoswapFactory *ICosmoswapFactoryCallerSession) FeeTo() (common.Address, error) {
	return _ICosmoswapFactory.Contract.FeeTo(&_ICosmoswapFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_ICosmoswapFactory *ICosmoswapFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_ICosmoswapFactory *ICosmoswapFactorySession) FeeToSetter() (common.Address, error) {
	return _ICosmoswapFactory.Contract.FeeToSetter(&_ICosmoswapFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_ICosmoswapFactory *ICosmoswapFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _ICosmoswapFactory.Contract.FeeToSetter(&_ICosmoswapFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _ICosmoswapFactory.Contract.GetPair(&_ICosmoswapFactory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _ICosmoswapFactory.Contract.GetPair(&_ICosmoswapFactory.CallOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.CreatePair(&_ICosmoswapFactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_ICosmoswapFactory *ICosmoswapFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.CreatePair(&_ICosmoswapFactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_ICosmoswapFactory *ICosmoswapFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_ICosmoswapFactory *ICosmoswapFactorySession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.SetFeeTo(&_ICosmoswapFactory.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_ICosmoswapFactory *ICosmoswapFactoryTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.SetFeeTo(&_ICosmoswapFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_ICosmoswapFactory *ICosmoswapFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_ICosmoswapFactory *ICosmoswapFactorySession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.SetFeeToSetter(&_ICosmoswapFactory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_ICosmoswapFactory *ICosmoswapFactoryTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _ICosmoswapFactory.Contract.SetFeeToSetter(&_ICosmoswapFactory.TransactOpts, arg0)
}

// ICosmoswapFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the ICosmoswapFactory contract.
type ICosmoswapFactoryPairCreatedIterator struct {
	Event *ICosmoswapFactoryPairCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICosmoswapFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICosmoswapFactoryPairCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICosmoswapFactoryPairCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICosmoswapFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICosmoswapFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICosmoswapFactoryPairCreated represents a PairCreated event raised by the ICosmoswapFactory contract.
type ICosmoswapFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_ICosmoswapFactory *ICosmoswapFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*ICosmoswapFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ICosmoswapFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapFactoryPairCreatedIterator{contract: _ICosmoswapFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_ICosmoswapFactory *ICosmoswapFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *ICosmoswapFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ICosmoswapFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICosmoswapFactoryPairCreated)
				if err := _ICosmoswapFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_ICosmoswapFactory *ICosmoswapFactoryFilterer) ParsePairCreated(log types.Log) (*ICosmoswapFactoryPairCreated, error) {
	event := new(ICosmoswapFactoryPairCreated)
	if err := _ICosmoswapFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICosmoswapPairMetaData contains all meta data concerning the ICosmoswapPair contract.
var ICosmoswapPairMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"ba9a7a56": "MINIMUM_LIQUIDITY()",
		"30adf81f": "PERMIT_TYPEHASH()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"89afcb44": "burn(address)",
		"313ce567": "decimals()",
		"c45a0155": "factory()",
		"0902f1ac": "getReserves()",
		"485cc955": "initialize(address,address)",
		"7464fc3d": "kLast()",
		"6a627842": "mint(address)",
		"06fdde03": "name()",
		"7ecebe00": "nonces(address)",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"5909c0d5": "price0CumulativeLast()",
		"5a3d5493": "price1CumulativeLast()",
		"bc25cf77": "skim(address)",
		"022c0d9f": "swap(uint256,uint256,address,bytes)",
		"95d89b41": "symbol()",
		"fff6cae9": "sync()",
		"0dfe1681": "token0()",
		"d21220a7": "token1()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// ICosmoswapPairABI is the input ABI used to generate the binding from.
// Deprecated: Use ICosmoswapPairMetaData.ABI instead.
var ICosmoswapPairABI = ICosmoswapPairMetaData.ABI

// Deprecated: Use ICosmoswapPairMetaData.Sigs instead.
// ICosmoswapPairFuncSigs maps the 4-byte function signature to its string representation.
var ICosmoswapPairFuncSigs = ICosmoswapPairMetaData.Sigs

// ICosmoswapPair is an auto generated Go binding around an Ethereum contract.
type ICosmoswapPair struct {
	ICosmoswapPairCaller     // Read-only binding to the contract
	ICosmoswapPairTransactor // Write-only binding to the contract
	ICosmoswapPairFilterer   // Log filterer for contract events
}

// ICosmoswapPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICosmoswapPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICosmoswapPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICosmoswapPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICosmoswapPairSession struct {
	Contract     *ICosmoswapPair   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICosmoswapPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICosmoswapPairCallerSession struct {
	Contract *ICosmoswapPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ICosmoswapPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICosmoswapPairTransactorSession struct {
	Contract     *ICosmoswapPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ICosmoswapPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICosmoswapPairRaw struct {
	Contract *ICosmoswapPair // Generic contract binding to access the raw methods on
}

// ICosmoswapPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICosmoswapPairCallerRaw struct {
	Contract *ICosmoswapPairCaller // Generic read-only contract binding to access the raw methods on
}

// ICosmoswapPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICosmoswapPairTransactorRaw struct {
	Contract *ICosmoswapPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICosmoswapPair creates a new instance of ICosmoswapPair, bound to a specific deployed contract.
func NewICosmoswapPair(address common.Address, backend bind.ContractBackend) (*ICosmoswapPair, error) {
	contract, err := bindICosmoswapPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPair{ICosmoswapPairCaller: ICosmoswapPairCaller{contract: contract}, ICosmoswapPairTransactor: ICosmoswapPairTransactor{contract: contract}, ICosmoswapPairFilterer: ICosmoswapPairFilterer{contract: contract}}, nil
}

// NewICosmoswapPairCaller creates a new read-only instance of ICosmoswapPair, bound to a specific deployed contract.
func NewICosmoswapPairCaller(address common.Address, caller bind.ContractCaller) (*ICosmoswapPairCaller, error) {
	contract, err := bindICosmoswapPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairCaller{contract: contract}, nil
}

// NewICosmoswapPairTransactor creates a new write-only instance of ICosmoswapPair, bound to a specific deployed contract.
func NewICosmoswapPairTransactor(address common.Address, transactor bind.ContractTransactor) (*ICosmoswapPairTransactor, error) {
	contract, err := bindICosmoswapPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairTransactor{contract: contract}, nil
}

// NewICosmoswapPairFilterer creates a new log filterer instance of ICosmoswapPair, bound to a specific deployed contract.
func NewICosmoswapPairFilterer(address common.Address, filterer bind.ContractFilterer) (*ICosmoswapPairFilterer, error) {
	contract, err := bindICosmoswapPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairFilterer{contract: contract}, nil
}

// bindICosmoswapPair binds a generic wrapper to an already deployed contract.
func bindICosmoswapPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICosmoswapPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmoswapPair *ICosmoswapPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmoswapPair.Contract.ICosmoswapPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmoswapPair *ICosmoswapPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.ICosmoswapPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmoswapPair *ICosmoswapPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.ICosmoswapPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmoswapPair *ICosmoswapPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmoswapPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmoswapPair *ICosmoswapPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmoswapPair *ICosmoswapPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ICosmoswapPair *ICosmoswapPairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ICosmoswapPair *ICosmoswapPairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ICosmoswapPair.Contract.DOMAINSEPARATOR(&_ICosmoswapPair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ICosmoswapPair.Contract.DOMAINSEPARATOR(&_ICosmoswapPair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _ICosmoswapPair.Contract.MINIMUMLIQUIDITY(&_ICosmoswapPair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _ICosmoswapPair.Contract.MINIMUMLIQUIDITY(&_ICosmoswapPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_ICosmoswapPair *ICosmoswapPairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_ICosmoswapPair *ICosmoswapPairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _ICosmoswapPair.Contract.PERMITTYPEHASH(&_ICosmoswapPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _ICosmoswapPair.Contract.PERMITTYPEHASH(&_ICosmoswapPair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ICosmoswapPair.Contract.Allowance(&_ICosmoswapPair.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ICosmoswapPair.Contract.Allowance(&_ICosmoswapPair.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ICosmoswapPair.Contract.BalanceOf(&_ICosmoswapPair.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ICosmoswapPair.Contract.BalanceOf(&_ICosmoswapPair.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ICosmoswapPair *ICosmoswapPairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ICosmoswapPair *ICosmoswapPairSession) Decimals() (uint8, error) {
	return _ICosmoswapPair.Contract.Decimals(&_ICosmoswapPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Decimals() (uint8, error) {
	return _ICosmoswapPair.Contract.Decimals(&_ICosmoswapPair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairSession) Factory() (common.Address, error) {
	return _ICosmoswapPair.Contract.Factory(&_ICosmoswapPair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Factory() (common.Address, error) {
	return _ICosmoswapPair.Contract.Factory(&_ICosmoswapPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ICosmoswapPair *ICosmoswapPairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ICosmoswapPair *ICosmoswapPairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _ICosmoswapPair.Contract.GetReserves(&_ICosmoswapPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _ICosmoswapPair.Contract.GetReserves(&_ICosmoswapPair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairSession) KLast() (*big.Int, error) {
	return _ICosmoswapPair.Contract.KLast(&_ICosmoswapPair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) KLast() (*big.Int, error) {
	return _ICosmoswapPair.Contract.KLast(&_ICosmoswapPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ICosmoswapPair *ICosmoswapPairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ICosmoswapPair *ICosmoswapPairSession) Name() (string, error) {
	return _ICosmoswapPair.Contract.Name(&_ICosmoswapPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Name() (string, error) {
	return _ICosmoswapPair.Contract.Name(&_ICosmoswapPair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ICosmoswapPair.Contract.Nonces(&_ICosmoswapPair.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ICosmoswapPair.Contract.Nonces(&_ICosmoswapPair.CallOpts, owner)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairSession) Price0CumulativeLast() (*big.Int, error) {
	return _ICosmoswapPair.Contract.Price0CumulativeLast(&_ICosmoswapPair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _ICosmoswapPair.Contract.Price0CumulativeLast(&_ICosmoswapPair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairSession) Price1CumulativeLast() (*big.Int, error) {
	return _ICosmoswapPair.Contract.Price1CumulativeLast(&_ICosmoswapPair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _ICosmoswapPair.Contract.Price1CumulativeLast(&_ICosmoswapPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ICosmoswapPair *ICosmoswapPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ICosmoswapPair *ICosmoswapPairSession) Symbol() (string, error) {
	return _ICosmoswapPair.Contract.Symbol(&_ICosmoswapPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Symbol() (string, error) {
	return _ICosmoswapPair.Contract.Symbol(&_ICosmoswapPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairSession) Token0() (common.Address, error) {
	return _ICosmoswapPair.Contract.Token0(&_ICosmoswapPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Token0() (common.Address, error) {
	return _ICosmoswapPair.Contract.Token0(&_ICosmoswapPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairSession) Token1() (common.Address, error) {
	return _ICosmoswapPair.Contract.Token1(&_ICosmoswapPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) Token1() (common.Address, error) {
	return _ICosmoswapPair.Contract.Token1(&_ICosmoswapPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapPair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairSession) TotalSupply() (*big.Int, error) {
	return _ICosmoswapPair.Contract.TotalSupply(&_ICosmoswapPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ICosmoswapPair *ICosmoswapPairCallerSession) TotalSupply() (*big.Int, error) {
	return _ICosmoswapPair.Contract.TotalSupply(&_ICosmoswapPair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Approve(&_ICosmoswapPair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Approve(&_ICosmoswapPair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_ICosmoswapPair *ICosmoswapPairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_ICosmoswapPair *ICosmoswapPairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Burn(&_ICosmoswapPair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Burn(&_ICosmoswapPair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_ICosmoswapPair *ICosmoswapPairTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "initialize", arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_ICosmoswapPair *ICosmoswapPairSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Initialize(&_ICosmoswapPair.TransactOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Initialize(&_ICosmoswapPair.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ICosmoswapPair *ICosmoswapPairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ICosmoswapPair *ICosmoswapPairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Mint(&_ICosmoswapPair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Mint(&_ICosmoswapPair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ICosmoswapPair *ICosmoswapPairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ICosmoswapPair *ICosmoswapPairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Permit(&_ICosmoswapPair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Permit(&_ICosmoswapPair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_ICosmoswapPair *ICosmoswapPairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_ICosmoswapPair *ICosmoswapPairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Skim(&_ICosmoswapPair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Skim(&_ICosmoswapPair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ICosmoswapPair *ICosmoswapPairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ICosmoswapPair *ICosmoswapPairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Swap(&_ICosmoswapPair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Swap(&_ICosmoswapPair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ICosmoswapPair *ICosmoswapPairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ICosmoswapPair *ICosmoswapPairSession) Sync() (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Sync(&_ICosmoswapPair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Sync() (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Sync(&_ICosmoswapPair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Transfer(&_ICosmoswapPair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.Transfer(&_ICosmoswapPair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.TransferFrom(&_ICosmoswapPair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ICosmoswapPair *ICosmoswapPairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ICosmoswapPair.Contract.TransferFrom(&_ICosmoswapPair.TransactOpts, from, to, value)
}

// ICosmoswapPairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ICosmoswapPair contract.
type ICosmoswapPairApprovalIterator struct {
	Event *ICosmoswapPairApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICosmoswapPairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICosmoswapPairApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICosmoswapPairApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICosmoswapPairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICosmoswapPairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICosmoswapPairApproval represents a Approval event raised by the ICosmoswapPair contract.
type ICosmoswapPairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ICosmoswapPair *ICosmoswapPairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ICosmoswapPairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairApprovalIterator{contract: _ICosmoswapPair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ICosmoswapPair *ICosmoswapPairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ICosmoswapPairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICosmoswapPairApproval)
				if err := _ICosmoswapPair.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ICosmoswapPair *ICosmoswapPairFilterer) ParseApproval(log types.Log) (*ICosmoswapPairApproval, error) {
	event := new(ICosmoswapPairApproval)
	if err := _ICosmoswapPair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICosmoswapPairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ICosmoswapPair contract.
type ICosmoswapPairBurnIterator struct {
	Event *ICosmoswapPairBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICosmoswapPairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICosmoswapPairBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICosmoswapPairBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICosmoswapPairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICosmoswapPairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICosmoswapPairBurn represents a Burn event raised by the ICosmoswapPair contract.
type ICosmoswapPairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_ICosmoswapPair *ICosmoswapPairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ICosmoswapPairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairBurnIterator{contract: _ICosmoswapPair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_ICosmoswapPair *ICosmoswapPairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ICosmoswapPairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICosmoswapPairBurn)
				if err := _ICosmoswapPair.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_ICosmoswapPair *ICosmoswapPairFilterer) ParseBurn(log types.Log) (*ICosmoswapPairBurn, error) {
	event := new(ICosmoswapPairBurn)
	if err := _ICosmoswapPair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICosmoswapPairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ICosmoswapPair contract.
type ICosmoswapPairMintIterator struct {
	Event *ICosmoswapPairMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICosmoswapPairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICosmoswapPairMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICosmoswapPairMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICosmoswapPairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICosmoswapPairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICosmoswapPairMint represents a Mint event raised by the ICosmoswapPair contract.
type ICosmoswapPairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_ICosmoswapPair *ICosmoswapPairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*ICosmoswapPairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairMintIterator{contract: _ICosmoswapPair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_ICosmoswapPair *ICosmoswapPairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ICosmoswapPairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICosmoswapPairMint)
				if err := _ICosmoswapPair.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_ICosmoswapPair *ICosmoswapPairFilterer) ParseMint(log types.Log) (*ICosmoswapPairMint, error) {
	event := new(ICosmoswapPairMint)
	if err := _ICosmoswapPair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICosmoswapPairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the ICosmoswapPair contract.
type ICosmoswapPairSwapIterator struct {
	Event *ICosmoswapPairSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICosmoswapPairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICosmoswapPairSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICosmoswapPairSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICosmoswapPairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICosmoswapPairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICosmoswapPairSwap represents a Swap event raised by the ICosmoswapPair contract.
type ICosmoswapPairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_ICosmoswapPair *ICosmoswapPairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ICosmoswapPairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairSwapIterator{contract: _ICosmoswapPair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_ICosmoswapPair *ICosmoswapPairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ICosmoswapPairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICosmoswapPairSwap)
				if err := _ICosmoswapPair.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_ICosmoswapPair *ICosmoswapPairFilterer) ParseSwap(log types.Log) (*ICosmoswapPairSwap, error) {
	event := new(ICosmoswapPairSwap)
	if err := _ICosmoswapPair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICosmoswapPairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the ICosmoswapPair contract.
type ICosmoswapPairSyncIterator struct {
	Event *ICosmoswapPairSync // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICosmoswapPairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICosmoswapPairSync)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICosmoswapPairSync)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICosmoswapPairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICosmoswapPairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICosmoswapPairSync represents a Sync event raised by the ICosmoswapPair contract.
type ICosmoswapPairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_ICosmoswapPair *ICosmoswapPairFilterer) FilterSync(opts *bind.FilterOpts) (*ICosmoswapPairSyncIterator, error) {

	logs, sub, err := _ICosmoswapPair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairSyncIterator{contract: _ICosmoswapPair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_ICosmoswapPair *ICosmoswapPairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *ICosmoswapPairSync) (event.Subscription, error) {

	logs, sub, err := _ICosmoswapPair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICosmoswapPairSync)
				if err := _ICosmoswapPair.contract.UnpackLog(event, "Sync", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_ICosmoswapPair *ICosmoswapPairFilterer) ParseSync(log types.Log) (*ICosmoswapPairSync, error) {
	event := new(ICosmoswapPairSync)
	if err := _ICosmoswapPair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICosmoswapPairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ICosmoswapPair contract.
type ICosmoswapPairTransferIterator struct {
	Event *ICosmoswapPairTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICosmoswapPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICosmoswapPairTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICosmoswapPairTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICosmoswapPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICosmoswapPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICosmoswapPairTransfer represents a Transfer event raised by the ICosmoswapPair contract.
type ICosmoswapPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ICosmoswapPair *ICosmoswapPairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ICosmoswapPairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapPairTransferIterator{contract: _ICosmoswapPair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ICosmoswapPair *ICosmoswapPairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ICosmoswapPairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ICosmoswapPair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICosmoswapPairTransfer)
				if err := _ICosmoswapPair.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ICosmoswapPair *ICosmoswapPairFilterer) ParseTransfer(log types.Log) (*ICosmoswapPairTransfer, error) {
	event := new(ICosmoswapPairTransfer)
	if err := _ICosmoswapPair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICosmoswapRouter01MetaData contains all meta data concerning the ICosmoswapRouter01 contract.
var ICosmoswapRouter01MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WPLUG\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityPLUG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityPLUG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityPLUGWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactPLUGForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForPLUG\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapPLUGForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactPLUG\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"faa5e5e2": "WPLUG()",
		"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
		"206de5b9": "addLiquidityPLUG(address,uint256,uint256,uint256,address,uint256)",
		"c45a0155": "factory()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
		"c3700e55": "removeLiquidityPLUG(address,uint256,uint256,uint256,address,uint256)",
		"aae97507": "removeLiquidityPLUGWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"e794b31b": "swapExactPLUGForTokens(uint256,address[],address,uint256)",
		"03a618a9": "swapExactTokensForPLUG(uint256,uint256,address[],address,uint256)",
		"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
		"ec301ba2": "swapPLUGForExactTokens(uint256,address[],address,uint256)",
		"9548ea81": "swapTokensForExactPLUG(uint256,uint256,address[],address,uint256)",
		"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
	},
}

// ICosmoswapRouter01ABI is the input ABI used to generate the binding from.
// Deprecated: Use ICosmoswapRouter01MetaData.ABI instead.
var ICosmoswapRouter01ABI = ICosmoswapRouter01MetaData.ABI

// Deprecated: Use ICosmoswapRouter01MetaData.Sigs instead.
// ICosmoswapRouter01FuncSigs maps the 4-byte function signature to its string representation.
var ICosmoswapRouter01FuncSigs = ICosmoswapRouter01MetaData.Sigs

// ICosmoswapRouter01 is an auto generated Go binding around an Ethereum contract.
type ICosmoswapRouter01 struct {
	ICosmoswapRouter01Caller     // Read-only binding to the contract
	ICosmoswapRouter01Transactor // Write-only binding to the contract
	ICosmoswapRouter01Filterer   // Log filterer for contract events
}

// ICosmoswapRouter01Caller is an auto generated read-only Go binding around an Ethereum contract.
type ICosmoswapRouter01Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapRouter01Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ICosmoswapRouter01Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapRouter01Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICosmoswapRouter01Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapRouter01Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICosmoswapRouter01Session struct {
	Contract     *ICosmoswapRouter01 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICosmoswapRouter01CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICosmoswapRouter01CallerSession struct {
	Contract *ICosmoswapRouter01Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ICosmoswapRouter01TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICosmoswapRouter01TransactorSession struct {
	Contract     *ICosmoswapRouter01Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ICosmoswapRouter01Raw is an auto generated low-level Go binding around an Ethereum contract.
type ICosmoswapRouter01Raw struct {
	Contract *ICosmoswapRouter01 // Generic contract binding to access the raw methods on
}

// ICosmoswapRouter01CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICosmoswapRouter01CallerRaw struct {
	Contract *ICosmoswapRouter01Caller // Generic read-only contract binding to access the raw methods on
}

// ICosmoswapRouter01TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICosmoswapRouter01TransactorRaw struct {
	Contract *ICosmoswapRouter01Transactor // Generic write-only contract binding to access the raw methods on
}

// NewICosmoswapRouter01 creates a new instance of ICosmoswapRouter01, bound to a specific deployed contract.
func NewICosmoswapRouter01(address common.Address, backend bind.ContractBackend) (*ICosmoswapRouter01, error) {
	contract, err := bindICosmoswapRouter01(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapRouter01{ICosmoswapRouter01Caller: ICosmoswapRouter01Caller{contract: contract}, ICosmoswapRouter01Transactor: ICosmoswapRouter01Transactor{contract: contract}, ICosmoswapRouter01Filterer: ICosmoswapRouter01Filterer{contract: contract}}, nil
}

// NewICosmoswapRouter01Caller creates a new read-only instance of ICosmoswapRouter01, bound to a specific deployed contract.
func NewICosmoswapRouter01Caller(address common.Address, caller bind.ContractCaller) (*ICosmoswapRouter01Caller, error) {
	contract, err := bindICosmoswapRouter01(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapRouter01Caller{contract: contract}, nil
}

// NewICosmoswapRouter01Transactor creates a new write-only instance of ICosmoswapRouter01, bound to a specific deployed contract.
func NewICosmoswapRouter01Transactor(address common.Address, transactor bind.ContractTransactor) (*ICosmoswapRouter01Transactor, error) {
	contract, err := bindICosmoswapRouter01(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapRouter01Transactor{contract: contract}, nil
}

// NewICosmoswapRouter01Filterer creates a new log filterer instance of ICosmoswapRouter01, bound to a specific deployed contract.
func NewICosmoswapRouter01Filterer(address common.Address, filterer bind.ContractFilterer) (*ICosmoswapRouter01Filterer, error) {
	contract, err := bindICosmoswapRouter01(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapRouter01Filterer{contract: contract}, nil
}

// bindICosmoswapRouter01 binds a generic wrapper to an already deployed contract.
func bindICosmoswapRouter01(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICosmoswapRouter01ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmoswapRouter01 *ICosmoswapRouter01Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmoswapRouter01.Contract.ICosmoswapRouter01Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmoswapRouter01 *ICosmoswapRouter01Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.ICosmoswapRouter01Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmoswapRouter01 *ICosmoswapRouter01Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.ICosmoswapRouter01Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmoswapRouter01 *ICosmoswapRouter01CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmoswapRouter01.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.contract.Transact(opts, method, params...)
}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() pure returns(address)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Caller) WPLUG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapRouter01.contract.Call(opts, &out, "WPLUG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() pure returns(address)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) WPLUG() (common.Address, error) {
	return _ICosmoswapRouter01.Contract.WPLUG(&_ICosmoswapRouter01.CallOpts)
}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() pure returns(address)
func (_ICosmoswapRouter01 *ICosmoswapRouter01CallerSession) WPLUG() (common.Address, error) {
	return _ICosmoswapRouter01.Contract.WPLUG(&_ICosmoswapRouter01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapRouter01.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) Factory() (common.Address, error) {
	return _ICosmoswapRouter01.Contract.Factory(&_ICosmoswapRouter01.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_ICosmoswapRouter01 *ICosmoswapRouter01CallerSession) Factory() (common.Address, error) {
	return _ICosmoswapRouter01.Contract.Factory(&_ICosmoswapRouter01.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter01.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter01.Contract.GetAmountIn(&_ICosmoswapRouter01.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ICosmoswapRouter01 *ICosmoswapRouter01CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter01.Contract.GetAmountIn(&_ICosmoswapRouter01.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter01.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter01.Contract.GetAmountOut(&_ICosmoswapRouter01.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_ICosmoswapRouter01 *ICosmoswapRouter01CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter01.Contract.GetAmountOut(&_ICosmoswapRouter01.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter01.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _ICosmoswapRouter01.Contract.GetAmountsIn(&_ICosmoswapRouter01.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _ICosmoswapRouter01.Contract.GetAmountsIn(&_ICosmoswapRouter01.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter01.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _ICosmoswapRouter01.Contract.GetAmountsOut(&_ICosmoswapRouter01.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _ICosmoswapRouter01.Contract.GetAmountsOut(&_ICosmoswapRouter01.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter01.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter01.Contract.Quote(&_ICosmoswapRouter01.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter01.Contract.Quote(&_ICosmoswapRouter01.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.AddLiquidity(&_ICosmoswapRouter01.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.AddLiquidity(&_ICosmoswapRouter01.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) AddLiquidityPLUG(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "addLiquidityPLUG", token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) AddLiquidityPLUG(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.AddLiquidityPLUG(&_ICosmoswapRouter01.TransactOpts, token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) AddLiquidityPLUG(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.AddLiquidityPLUG(&_ICosmoswapRouter01.TransactOpts, token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.RemoveLiquidity(&_ICosmoswapRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.RemoveLiquidity(&_ICosmoswapRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) RemoveLiquidityPLUG(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "removeLiquidityPLUG", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) RemoveLiquidityPLUG(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.RemoveLiquidityPLUG(&_ICosmoswapRouter01.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) RemoveLiquidityPLUG(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.RemoveLiquidityPLUG(&_ICosmoswapRouter01.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) RemoveLiquidityPLUGWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "removeLiquidityPLUGWithPermit", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) RemoveLiquidityPLUGWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.RemoveLiquidityPLUGWithPermit(&_ICosmoswapRouter01.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) RemoveLiquidityPLUGWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.RemoveLiquidityPLUGWithPermit(&_ICosmoswapRouter01.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.RemoveLiquidityWithPermit(&_ICosmoswapRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.RemoveLiquidityWithPermit(&_ICosmoswapRouter01.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) SwapExactPLUGForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "swapExactPLUGForTokens", amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) SwapExactPLUGForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapExactPLUGForTokens(&_ICosmoswapRouter01.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) SwapExactPLUGForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapExactPLUGForTokens(&_ICosmoswapRouter01.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) SwapExactTokensForPLUG(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "swapExactTokensForPLUG", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) SwapExactTokensForPLUG(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapExactTokensForPLUG(&_ICosmoswapRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) SwapExactTokensForPLUG(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapExactTokensForPLUG(&_ICosmoswapRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapExactTokensForTokens(&_ICosmoswapRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapExactTokensForTokens(&_ICosmoswapRouter01.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) SwapPLUGForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "swapPLUGForExactTokens", amountOut, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) SwapPLUGForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapPLUGForExactTokens(&_ICosmoswapRouter01.TransactOpts, amountOut, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) SwapPLUGForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapPLUGForExactTokens(&_ICosmoswapRouter01.TransactOpts, amountOut, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) SwapTokensForExactPLUG(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "swapTokensForExactPLUG", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) SwapTokensForExactPLUG(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapTokensForExactPLUG(&_ICosmoswapRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) SwapTokensForExactPLUG(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapTokensForExactPLUG(&_ICosmoswapRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapTokensForExactTokens(&_ICosmoswapRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter01 *ICosmoswapRouter01TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter01.Contract.SwapTokensForExactTokens(&_ICosmoswapRouter01.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// ICosmoswapRouter02MetaData contains all meta data concerning the ICosmoswapRouter02 contract.
var ICosmoswapRouter02MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WPLUG\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityPLUG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityPLUG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityPLUGSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityPLUGWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPLUGMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountPLUG\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactPLUGForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactPLUGForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForPLUG\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForPLUGSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapPLUGForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactPLUG\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"faa5e5e2": "WPLUG()",
		"e8e33700": "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)",
		"206de5b9": "addLiquidityPLUG(address,uint256,uint256,uint256,address,uint256)",
		"c45a0155": "factory()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"baa2abde": "removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)",
		"c3700e55": "removeLiquidityPLUG(address,uint256,uint256,uint256,address,uint256)",
		"c5b21050": "removeLiquidityPLUGSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)",
		"aae97507": "removeLiquidityPLUGWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"89fe6c1f": "removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"2195995c": "removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)",
		"e794b31b": "swapExactPLUGForTokens(uint256,address[],address,uint256)",
		"39940723": "swapExactPLUGForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256)",
		"03a618a9": "swapExactTokensForPLUG(uint256,uint256,address[],address,uint256)",
		"9d5148f6": "swapExactTokensForPLUGSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"38ed1739": "swapExactTokensForTokens(uint256,uint256,address[],address,uint256)",
		"5c11d795": "swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)",
		"ec301ba2": "swapPLUGForExactTokens(uint256,address[],address,uint256)",
		"9548ea81": "swapTokensForExactPLUG(uint256,uint256,address[],address,uint256)",
		"8803dbee": "swapTokensForExactTokens(uint256,uint256,address[],address,uint256)",
	},
}

// ICosmoswapRouter02ABI is the input ABI used to generate the binding from.
// Deprecated: Use ICosmoswapRouter02MetaData.ABI instead.
var ICosmoswapRouter02ABI = ICosmoswapRouter02MetaData.ABI

// Deprecated: Use ICosmoswapRouter02MetaData.Sigs instead.
// ICosmoswapRouter02FuncSigs maps the 4-byte function signature to its string representation.
var ICosmoswapRouter02FuncSigs = ICosmoswapRouter02MetaData.Sigs

// ICosmoswapRouter02 is an auto generated Go binding around an Ethereum contract.
type ICosmoswapRouter02 struct {
	ICosmoswapRouter02Caller     // Read-only binding to the contract
	ICosmoswapRouter02Transactor // Write-only binding to the contract
	ICosmoswapRouter02Filterer   // Log filterer for contract events
}

// ICosmoswapRouter02Caller is an auto generated read-only Go binding around an Ethereum contract.
type ICosmoswapRouter02Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapRouter02Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ICosmoswapRouter02Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapRouter02Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICosmoswapRouter02Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICosmoswapRouter02Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICosmoswapRouter02Session struct {
	Contract     *ICosmoswapRouter02 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICosmoswapRouter02CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICosmoswapRouter02CallerSession struct {
	Contract *ICosmoswapRouter02Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ICosmoswapRouter02TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICosmoswapRouter02TransactorSession struct {
	Contract     *ICosmoswapRouter02Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ICosmoswapRouter02Raw is an auto generated low-level Go binding around an Ethereum contract.
type ICosmoswapRouter02Raw struct {
	Contract *ICosmoswapRouter02 // Generic contract binding to access the raw methods on
}

// ICosmoswapRouter02CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICosmoswapRouter02CallerRaw struct {
	Contract *ICosmoswapRouter02Caller // Generic read-only contract binding to access the raw methods on
}

// ICosmoswapRouter02TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICosmoswapRouter02TransactorRaw struct {
	Contract *ICosmoswapRouter02Transactor // Generic write-only contract binding to access the raw methods on
}

// NewICosmoswapRouter02 creates a new instance of ICosmoswapRouter02, bound to a specific deployed contract.
func NewICosmoswapRouter02(address common.Address, backend bind.ContractBackend) (*ICosmoswapRouter02, error) {
	contract, err := bindICosmoswapRouter02(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapRouter02{ICosmoswapRouter02Caller: ICosmoswapRouter02Caller{contract: contract}, ICosmoswapRouter02Transactor: ICosmoswapRouter02Transactor{contract: contract}, ICosmoswapRouter02Filterer: ICosmoswapRouter02Filterer{contract: contract}}, nil
}

// NewICosmoswapRouter02Caller creates a new read-only instance of ICosmoswapRouter02, bound to a specific deployed contract.
func NewICosmoswapRouter02Caller(address common.Address, caller bind.ContractCaller) (*ICosmoswapRouter02Caller, error) {
	contract, err := bindICosmoswapRouter02(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapRouter02Caller{contract: contract}, nil
}

// NewICosmoswapRouter02Transactor creates a new write-only instance of ICosmoswapRouter02, bound to a specific deployed contract.
func NewICosmoswapRouter02Transactor(address common.Address, transactor bind.ContractTransactor) (*ICosmoswapRouter02Transactor, error) {
	contract, err := bindICosmoswapRouter02(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapRouter02Transactor{contract: contract}, nil
}

// NewICosmoswapRouter02Filterer creates a new log filterer instance of ICosmoswapRouter02, bound to a specific deployed contract.
func NewICosmoswapRouter02Filterer(address common.Address, filterer bind.ContractFilterer) (*ICosmoswapRouter02Filterer, error) {
	contract, err := bindICosmoswapRouter02(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICosmoswapRouter02Filterer{contract: contract}, nil
}

// bindICosmoswapRouter02 binds a generic wrapper to an already deployed contract.
func bindICosmoswapRouter02(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICosmoswapRouter02ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmoswapRouter02 *ICosmoswapRouter02Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmoswapRouter02.Contract.ICosmoswapRouter02Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmoswapRouter02 *ICosmoswapRouter02Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.ICosmoswapRouter02Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmoswapRouter02 *ICosmoswapRouter02Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.ICosmoswapRouter02Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICosmoswapRouter02 *ICosmoswapRouter02CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICosmoswapRouter02.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.contract.Transact(opts, method, params...)
}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() pure returns(address)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Caller) WPLUG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapRouter02.contract.Call(opts, &out, "WPLUG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() pure returns(address)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) WPLUG() (common.Address, error) {
	return _ICosmoswapRouter02.Contract.WPLUG(&_ICosmoswapRouter02.CallOpts)
}

// WPLUG is a free data retrieval call binding the contract method 0xfaa5e5e2.
//
// Solidity: function WPLUG() pure returns(address)
func (_ICosmoswapRouter02 *ICosmoswapRouter02CallerSession) WPLUG() (common.Address, error) {
	return _ICosmoswapRouter02.Contract.WPLUG(&_ICosmoswapRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICosmoswapRouter02.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) Factory() (common.Address, error) {
	return _ICosmoswapRouter02.Contract.Factory(&_ICosmoswapRouter02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_ICosmoswapRouter02 *ICosmoswapRouter02CallerSession) Factory() (common.Address, error) {
	return _ICosmoswapRouter02.Contract.Factory(&_ICosmoswapRouter02.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Caller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter02.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter02.Contract.GetAmountIn(&_ICosmoswapRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_ICosmoswapRouter02 *ICosmoswapRouter02CallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter02.Contract.GetAmountIn(&_ICosmoswapRouter02.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Caller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter02.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter02.Contract.GetAmountOut(&_ICosmoswapRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_ICosmoswapRouter02 *ICosmoswapRouter02CallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter02.Contract.GetAmountOut(&_ICosmoswapRouter02.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter02.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _ICosmoswapRouter02.Contract.GetAmountsIn(&_ICosmoswapRouter02.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _ICosmoswapRouter02.Contract.GetAmountsIn(&_ICosmoswapRouter02.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter02.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _ICosmoswapRouter02.Contract.GetAmountsOut(&_ICosmoswapRouter02.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _ICosmoswapRouter02.Contract.GetAmountsOut(&_ICosmoswapRouter02.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Caller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICosmoswapRouter02.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter02.Contract.Quote(&_ICosmoswapRouter02.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02CallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _ICosmoswapRouter02.Contract.Quote(&_ICosmoswapRouter02.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.AddLiquidity(&_ICosmoswapRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.AddLiquidity(&_ICosmoswapRouter02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) AddLiquidityPLUG(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "addLiquidityPLUG", token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) AddLiquidityPLUG(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.AddLiquidityPLUG(&_ICosmoswapRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// AddLiquidityPLUG is a paid mutator transaction binding the contract method 0x206de5b9.
//
// Solidity: function addLiquidityPLUG(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountPLUG, uint256 liquidity)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) AddLiquidityPLUG(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.AddLiquidityPLUG(&_ICosmoswapRouter02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidity(&_ICosmoswapRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidity(&_ICosmoswapRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) RemoveLiquidityPLUG(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "removeLiquidityPLUG", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) RemoveLiquidityPLUG(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityPLUG(&_ICosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUG is a paid mutator transaction binding the contract method 0xc3700e55.
//
// Solidity: function removeLiquidityPLUG(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) RemoveLiquidityPLUG(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityPLUG(&_ICosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc5b21050.
//
// Solidity: function removeLiquidityPLUGSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) RemoveLiquidityPLUGSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "removeLiquidityPLUGSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc5b21050.
//
// Solidity: function removeLiquidityPLUGSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) RemoveLiquidityPLUGSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityPLUGSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc5b21050.
//
// Solidity: function removeLiquidityPLUGSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline) returns(uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) RemoveLiquidityPLUGSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityPLUGSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) RemoveLiquidityPLUGWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "removeLiquidityPLUGWithPermit", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) RemoveLiquidityPLUGWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityPLUGWithPermit(&_ICosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermit is a paid mutator transaction binding the contract method 0xaae97507.
//
// Solidity: function removeLiquidityPLUGWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) RemoveLiquidityPLUGWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityPLUGWithPermit(&_ICosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x89fe6c1f.
//
// Solidity: function removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x89fe6c1f.
//
// Solidity: function removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x89fe6c1f.
//
// Solidity: function removeLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountPLUGMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountPLUG)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountPLUGMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityPLUGWithPermitSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, token, liquidity, amountTokenMin, amountPLUGMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityWithPermit(&_ICosmoswapRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.RemoveLiquidityWithPermit(&_ICosmoswapRouter02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapExactPLUGForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapExactPLUGForTokens", amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapExactPLUGForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactPLUGForTokens(&_ICosmoswapRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokens is a paid mutator transaction binding the contract method 0xe794b31b.
//
// Solidity: function swapExactPLUGForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapExactPLUGForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactPLUGForTokens(&_ICosmoswapRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x39940723.
//
// Solidity: function swapExactPLUGForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapExactPLUGForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapExactPLUGForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x39940723.
//
// Solidity: function swapExactPLUGForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapExactPLUGForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactPLUGForTokensSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactPLUGForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x39940723.
//
// Solidity: function swapExactPLUGForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapExactPLUGForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactPLUGForTokensSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapExactTokensForPLUG(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapExactTokensForPLUG", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapExactTokensForPLUG(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactTokensForPLUG(&_ICosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUG is a paid mutator transaction binding the contract method 0x03a618a9.
//
// Solidity: function swapExactTokensForPLUG(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapExactTokensForPLUG(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactTokensForPLUG(&_ICosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9d5148f6.
//
// Solidity: function swapExactTokensForPLUGSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapExactTokensForPLUGSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapExactTokensForPLUGSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9d5148f6.
//
// Solidity: function swapExactTokensForPLUGSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapExactTokensForPLUGSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactTokensForPLUGSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForPLUGSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9d5148f6.
//
// Solidity: function swapExactTokensForPLUGSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapExactTokensForPLUGSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactTokensForPLUGSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactTokensForTokens(&_ICosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactTokensForTokens(&_ICosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_ICosmoswapRouter02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapPLUGForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapPLUGForExactTokens", amountOut, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapPLUGForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapPLUGForExactTokens(&_ICosmoswapRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapPLUGForExactTokens is a paid mutator transaction binding the contract method 0xec301ba2.
//
// Solidity: function swapPLUGForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapPLUGForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapPLUGForExactTokens(&_ICosmoswapRouter02.TransactOpts, amountOut, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapTokensForExactPLUG(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapTokensForExactPLUG", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapTokensForExactPLUG(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapTokensForExactPLUG(&_ICosmoswapRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactPLUG is a paid mutator transaction binding the contract method 0x9548ea81.
//
// Solidity: function swapTokensForExactPLUG(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapTokensForExactPLUG(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapTokensForExactPLUG(&_ICosmoswapRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapTokensForExactTokens(&_ICosmoswapRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_ICosmoswapRouter02 *ICosmoswapRouter02TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _ICosmoswapRouter02.Contract.SwapTokensForExactTokens(&_ICosmoswapRouter02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWPLUGMetaData contains all meta data concerning the IWPLUG contract.
var IWPLUGMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d0e30db0": "deposit()",
		"a9059cbb": "transfer(address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// IWPLUGABI is the input ABI used to generate the binding from.
// Deprecated: Use IWPLUGMetaData.ABI instead.
var IWPLUGABI = IWPLUGMetaData.ABI

// Deprecated: Use IWPLUGMetaData.Sigs instead.
// IWPLUGFuncSigs maps the 4-byte function signature to its string representation.
var IWPLUGFuncSigs = IWPLUGMetaData.Sigs

// IWPLUG is an auto generated Go binding around an Ethereum contract.
type IWPLUG struct {
	IWPLUGCaller     // Read-only binding to the contract
	IWPLUGTransactor // Write-only binding to the contract
	IWPLUGFilterer   // Log filterer for contract events
}

// IWPLUGCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWPLUGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWPLUGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWPLUGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWPLUGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWPLUGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWPLUGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWPLUGSession struct {
	Contract     *IWPLUG           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWPLUGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWPLUGCallerSession struct {
	Contract *IWPLUGCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWPLUGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWPLUGTransactorSession struct {
	Contract     *IWPLUGTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWPLUGRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWPLUGRaw struct {
	Contract *IWPLUG // Generic contract binding to access the raw methods on
}

// IWPLUGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWPLUGCallerRaw struct {
	Contract *IWPLUGCaller // Generic read-only contract binding to access the raw methods on
}

// IWPLUGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWPLUGTransactorRaw struct {
	Contract *IWPLUGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWPLUG creates a new instance of IWPLUG, bound to a specific deployed contract.
func NewIWPLUG(address common.Address, backend bind.ContractBackend) (*IWPLUG, error) {
	contract, err := bindIWPLUG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWPLUG{IWPLUGCaller: IWPLUGCaller{contract: contract}, IWPLUGTransactor: IWPLUGTransactor{contract: contract}, IWPLUGFilterer: IWPLUGFilterer{contract: contract}}, nil
}

// NewIWPLUGCaller creates a new read-only instance of IWPLUG, bound to a specific deployed contract.
func NewIWPLUGCaller(address common.Address, caller bind.ContractCaller) (*IWPLUGCaller, error) {
	contract, err := bindIWPLUG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWPLUGCaller{contract: contract}, nil
}

// NewIWPLUGTransactor creates a new write-only instance of IWPLUG, bound to a specific deployed contract.
func NewIWPLUGTransactor(address common.Address, transactor bind.ContractTransactor) (*IWPLUGTransactor, error) {
	contract, err := bindIWPLUG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWPLUGTransactor{contract: contract}, nil
}

// NewIWPLUGFilterer creates a new log filterer instance of IWPLUG, bound to a specific deployed contract.
func NewIWPLUGFilterer(address common.Address, filterer bind.ContractFilterer) (*IWPLUGFilterer, error) {
	contract, err := bindIWPLUG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWPLUGFilterer{contract: contract}, nil
}

// bindIWPLUG binds a generic wrapper to an already deployed contract.
func bindIWPLUG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWPLUGABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWPLUG *IWPLUGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWPLUG.Contract.IWPLUGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWPLUG *IWPLUGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWPLUG.Contract.IWPLUGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWPLUG *IWPLUGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWPLUG.Contract.IWPLUGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWPLUG *IWPLUGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWPLUG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWPLUG *IWPLUGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWPLUG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWPLUG *IWPLUGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWPLUG.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWPLUG *IWPLUGTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWPLUG.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWPLUG *IWPLUGSession) Deposit() (*types.Transaction, error) {
	return _IWPLUG.Contract.Deposit(&_IWPLUG.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWPLUG *IWPLUGTransactorSession) Deposit() (*types.Transaction, error) {
	return _IWPLUG.Contract.Deposit(&_IWPLUG.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWPLUG *IWPLUGTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWPLUG.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWPLUG *IWPLUGSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWPLUG.Contract.Transfer(&_IWPLUG.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWPLUG *IWPLUGTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWPLUG.Contract.Transfer(&_IWPLUG.TransactOpts, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWPLUG *IWPLUGTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWPLUG.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWPLUG *IWPLUGSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWPLUG.Contract.Withdraw(&_IWPLUG.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWPLUG *IWPLUGTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWPLUG.Contract.Withdraw(&_IWPLUG.TransactOpts, arg0)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206a1acfbfb0eb40ead7724f20b928a1a865a458bf625f72f66f165bbf6c2067e364736f6c63430006060033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TransferHelperMetaData contains all meta data concerning the TransferHelper contract.
var TransferHelperMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e27dcf2c006334decdf2628122b29523163206f0237e456328e6d4f4fa32040964736f6c63430006060033",
}

// TransferHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferHelperMetaData.ABI instead.
var TransferHelperABI = TransferHelperMetaData.ABI

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferHelperMetaData.Bin instead.
var TransferHelperBin = TransferHelperMetaData.Bin

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := TransferHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}
