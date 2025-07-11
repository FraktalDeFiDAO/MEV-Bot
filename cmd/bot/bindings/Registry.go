// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// ExchangeLibExchangeInfo is an auto generated low-level Go binding around an user-defined struct.
type ExchangeLibExchangeInfo struct {
	Name    string
	Router  common.Address
	Enabled bool
}

// PoolLibPoolInfo is an auto generated low-level Go binding around an user-defined struct.
type PoolLibPoolInfo struct {
	Token0     common.Address
	Token1     common.Address
	ExchangeId *big.Int
	Enabled    bool
}

// TokenLibTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenLibTokenInfo struct {
	Decimals uint8
	Enabled  bool
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addExchange\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPool\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"exchangeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addTokens\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"decimals\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getExchange\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structExchangeLib.ExchangeInfo\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExchangeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExchanges\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structExchangeLib.ExchangeInfo[]\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPool\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoolLib.PoolInfo\",\"components\":[{\"name\":\"token0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"exchangeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTokenLib.TokenInfo\",\"components\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPoolEnabled\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTokenEnabled\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setExchangeEnabled\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolEnabled\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenEnabled\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608080604052346100165761142b908161001b8239f35b5f80fdfe604060808152600480361015610013575f80fd5b5f90813560e01c80630b9d584714610e7c57806313d59a8314610cd05780631e2e3a6b14610aff5780632798795a14610ac95780634c4523a614610a8b57806359770438146109e9578063673a2a1f1461094e57806369581898146106d7578063748538d9146106a157806378a89567146106635780638eec5d7014610625578063a74ea63f146105f2578063aa6ca8081461054a578063b27925ff14610495578063bbe4f6db146103d6578063fb93e5da14610248578063feaef9bd1461018e5763ff612cd0146100e3575f80fd5b3461018a578260031936011261018a578035906100fe611066565b918084525f805160206113d68339815191529182602052610121868620546111ad565b1561015457508352602052918120600101805460ff60a01b191692151560a01b60ff60a01b169290921790915580f35b80f35b606490602087519162461bcd60e51b8352820152601060248201526f65786368616e6765206d697373696e6760801b6044820152fd5b5080fd5b50913461018a578060031936011261018a576101a8610ff3565b906101b1611066565b9160018060a01b0316938484527f9cff66c0348144d548e1d853e2a1e81687b8d40edb52adef6b217b14ec4feb01908160205260ff83862054161561021557508394610151945260205283209061ff00825491151560081b169061ff001916179055565b606490602084519162461bcd60e51b8352820152600d60248201526c746f6b656e206d697373696e6760981b6044820152fd5b503461018a578260031936011261018a5767ffffffffffffffff9080358281116103d2576102799036908301611075565b90926024359081116103ce576102929036908401611075565b959061029d83611130565b946102aa8351968761110e565b8386526020938487019060051b8201913683116103a357905b8282106103ab575050506102d687611130565b966102e38351988961110e565b8088528388019060051b8201913683116103a757905b82821061038a57505050835186510361035757505050815b81518110156103535761034e906103496001600160a01b0361033383866111e5565b511660ff61034184896111e5565b5116906112af565b61118b565b610311565b8280f35b5162461bcd60e51b815291820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b604482015260649150fd5b813560ff811681036103a35781529084019084016102f9565b8880fd5b8780fd5b81356001600160a01b03811681036103ca5781529085019085016102c3565b8980fd5b8480fd5b8380fd5b82843461018a57602036600319011261018a5790816080926103f6610ff3565b926103ff611167565b50610408611167565b506001600160a01b0393841681527f3d3ff1109b55fec8aa0fb0c1a802263897546ac48f762cb91dcb472b76f407d56020522081519190610448836110d6565b83815416938484528060018301541660208501908152606060ff60036002860154958789019687520154169501941515855283519586525116602085015251908301525115156060820152f35b503461018a578260031936011261018a576104ae610ff3565b6104b6611066565b9160018060a01b03809216908185527f3d3ff1109b55fec8aa0fb0c1a802263897546ac48f762cb91dcb472b76f407d59283602052868620541615610518575093600391849561015195526020528420019060ff801983541691151516179055565b606490602087519162461bcd60e51b8352820152600c60248201526b706f6f6c206d697373696e6760a01b6044820152fd5b5050346105ef57806003193601126105ef577f9cff66c0348144d548e1d853e2a1e81687b8d40edb52adef6b217b14ec4feb029182549261058a8461127d565b92805b8581106105a5578351806105a18782611023565b0390f35b8282527f6e57dbff476de205dddc9d57d8a327789e699b4afc629b0e11092fa826aa6a158101546105ea91906001600160a01b03166105e482886111e5565b5261118b565b61058d565b80fd5b82843461018a573660031901126105ef5761060b610ff3565b60243560ff8116810361062157610151916112af565b8280fd5b82843461018a578160031936011261018a576020907f3d3ff1109b55fec8aa0fb0c1a802263897546ac48f762cb91dcb472b76f407d6549051908152f35b82843461018a578160031936011261018a576020907f9cff66c0348144d548e1d853e2a1e81687b8d40edb52adef6b217b14ec4feb02549051908152f35b82843461018a57602036600319011261018a5760209060ff6106c96106c4610ff3565b611245565b5460081c1690519015158152f35b503461018a578260031936011261018a5780359267ffffffffffffffff918285116103d257366023860112156103d25784810135948386116103ce5736602487830101116103ce5761072761100d565b83519286602089819a6024601f1997610747858a601f860116018b61110e565b828a520183890137860101526001600160a01b0391821693841561091d577f9e53c3621b95d93ac0d3ea4bb0fd8278bb4efe432ca3f32eebd986397c6cc1b7968754976107938961118b565b90558651916107a1836110a6565b825289820195865286820194600193848752898b525f805160206113d68339815191528c52888b20935190815193841161090a57508b6107e185546111ad565b601f81116108c2575b50508b91601f84116001146108615750908291610850999a9b92610856575b50505f19600383901b1c191690831b1781555b019251166001600160601b0360a01b83541617825551151581549060ff60a01b90151560a01b169060ff60a01b1916179055565b51908152f35b015190505f80610809565b839b93168484528c8420935b8d8282106108ae575050918593918c6108509b9c9d9410610896575b505050811b01815561081c565b01515f1960f88460031b161c191690555f8080610889565b83850151865594870194938401930161086d565b858d52818d2090601f860160051c8201928610610900575b601f0160051c019086905b8281106108f557508d91506107ea565b8d81550186906108e5565b90915081906108da565b634e487b7160e01b8c526041905260248bfd5b855162461bcd60e51b81528083018a9052600b60248201526a3d32b937903937baba32b960a91b6044820152606490fd5b5050346105ef57806003193601126105ef577f3d3ff1109b55fec8aa0fb0c1a802263897546ac48f762cb91dcb472b76f407d69182549261098e8461127d565b92805b8581106109a5578351806105a18782611023565b8282527f4292a41d0bbcaa2ddf074acbd09c985b25675965ea5c85b4fab56861d1b7d0b58101546109e491906001600160a01b03166105e482886111e5565b610991565b5050346105ef57602090816003193601126105ef578290610a08610ff3565b81848451610a15816110f2565b828152015281848451610a27816110f2565b82815201526001600160a01b031681527f9cff66c0348144d548e1d853e2a1e81687b8d40edb52adef6b217b14ec4feb0183522082519190610a68836110f2565b5460ff8281831694858152019160081c1615158152835192835251151590820152f35b82843461018a578160031936011261018a576020907f9e53c3621b95d93ac0d3ea4bb0fd8278bb4efe432ca3f32eebd986397c6cc1b7549051908152f35b82843461018a57602036600319011261018a5760209060ff6003610af3610aee610ff3565b61120d565b01541690519015158152f35b5050346105ef57806003193601126105ef577f9e53c3621b95d93ac0d3ea4bb0fd8278bb4efe432ca3f32eebd986397c6cc1b754610b3c81611130565b90610b498451928361110e565b808252601f19610b5882611130565b01835b818110610cb9575050825b818110610bca575050825191602080840190808552835180925280868601968360051b870101940192955b828710610b9e5785850386f35b909192938280610bba600193603f198a82030186528851610f8b565b9601920196019592919092610b91565b80845260205f805160206113d6833981519152815285852090865191610bef836110a6565b87518154919088610bff846111ad565b808352600194808616908115610c9f5750600114610c6f575b509181610c3060ff9593610c6a99989795038261110e565b855201546001600160a01b0381169184019190915260a01c16151587820152610c5982866111e5565b52610c6481856111e5565b5061118b565b610b66565b838b52858b208b92505b818310610c8c5750508101840181610c18565b8054848401880152918601918501610c79565b60ff19168488015250151560051b82018501905081610c18565b602090610cc4611148565b82828701015201610b5b565b50913461018a57608036600319011261018a57610ceb610ff3565b610cf361100d565b916001600160a01b03906044358281169190829003610e78578284168015610e495783610d1f8661120d565b541615610da4575b509161015194939183826003955196610d3f886110d6565b16865260208601918252850191606435835280610d6360608801956001875261120d565b965116906001600160601b0360a01b9182885416178755600187019251169082541617905551600284015551151591019060ff801983541691151516179055565b7f3d3ff1109b55fec8aa0fb0c1a802263897546ac48f762cb91dcb472b76f407d6805490600160401b821015610e365760018201808255821015610e235788527f4292a41d0bbcaa2ddf074acbd09c985b25675965ea5c85b4fab56861d1b7d0b50180546001600160a01b031916909117905594955085946003610d27565b634e487b7160e01b895260328a52602489fd5b634e487b7160e01b895260418a52602489fd5b815162461bcd60e51b81526020818a015260096024820152681e995c9bc81c1bdbdb60ba1b6044820152606490fd5b8580fd5b50823461062157602091826003193601126103d257610e99611148565b50610ea2611148565b503583525f805160206113d6833981519152825280832092815193610ec6856110a6565b825182825493610ed5856111ad565b908184526001958887821691825f14610f69575050600114610f34575b50509181610f096105a196959360ff95038261110e565b875201546001600160a01b0381168587015260a01c1615158185015251828152928392830190610f8b565b908792508482528282205b818310610f5457505082010181610f09610ef2565b80548386018501528893909201918601610f3f565b60ff19168187015292151560051b85019092019250839150610f099050610ef2565b9190825192606082528351908160608401525f5b828110610fdd575060809394506040905f85848601015260018060a01b036020820151166020850152015115156040830152601f8019910116010190565b8060208092880101516080828701015201610f9f565b600435906001600160a01b038216820361100957565b5f80fd5b602435906001600160a01b038216820361100957565b602090816040818301928281528551809452019301915f5b828110611049575050505090565b83516001600160a01b03168552938101939281019260010161103b565b60243590811515820361100957565b9181601f840112156110095782359167ffffffffffffffff8311611009576020808501948460051b01011161100957565b6060810190811067ffffffffffffffff8211176110c257604052565b634e487b7160e01b5f52604160045260245ffd5b6080810190811067ffffffffffffffff8211176110c257604052565b6040810190811067ffffffffffffffff8211176110c257604052565b90601f8019910116810190811067ffffffffffffffff8211176110c257604052565b67ffffffffffffffff81116110c25760051b60200190565b60405190611155826110a6565b5f604083606081528260208201520152565b60405190611174826110d6565b5f6060838281528260208201528260408201520152565b5f1981146111995760010190565b634e487b7160e01b5f52601160045260245ffd5b90600182811c921680156111db575b60208310146111c757565b634e487b7160e01b5f52602260045260245ffd5b91607f16916111bc565b80518210156111f95760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b6001600160a01b03165f9081527f3d3ff1109b55fec8aa0fb0c1a802263897546ac48f762cb91dcb472b76f407d56020526040902090565b6001600160a01b03165f9081527f9cff66c0348144d548e1d853e2a1e81687b8d40edb52adef6b217b14ec4feb016020526040902090565b9061128782611130565b611294604051918261110e565b82815280926112a5601f1991611130565b0190602036910137565b6001600160a01b038116919082156113a35760ff6112cc82611245565b541615611326575b611324925060ff604051926112e8846110f2565b16825260ff6112fe602084019260018452611245565b92511660ff1983541617825551151561ff00825491151560081b169061ff001916179055565b565b7f9cff66c0348144d548e1d853e2a1e81687b8d40edb52adef6b217b14ec4feb02928354600160401b8110156110c257600181018086558110156111f957611324945f527f6e57dbff476de205dddc9d57d8a327789e699b4afc629b0e11092fa826aa6a1501906001600160601b0360a01b8254161790556112d4565b60405162461bcd60e51b815260206004820152600a6024820152693d32b937903a37b5b2b760b11b6044820152606490fdfe9e53c3621b95d93ac0d3ea4bb0fd8278bb4efe432ca3f32eebd986397c6cc1b6a26469706673582212203e687d896e25b147ddc7bcd03111e29efcd2e68a0dd68eca713732815afa224064736f6c63430008140033",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// GetExchange is a free data retrieval call binding the contract method 0x0b9d5847.
//
// Solidity: function getExchange(uint256 id) view returns((string,address,bool))
func (_Registry *RegistryCaller) GetExchange(opts *bind.CallOpts, id *big.Int) (ExchangeLibExchangeInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getExchange", id)

	if err != nil {
		return *new(ExchangeLibExchangeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ExchangeLibExchangeInfo)).(*ExchangeLibExchangeInfo)

	return out0, err

}

// GetExchange is a free data retrieval call binding the contract method 0x0b9d5847.
//
// Solidity: function getExchange(uint256 id) view returns((string,address,bool))
func (_Registry *RegistrySession) GetExchange(id *big.Int) (ExchangeLibExchangeInfo, error) {
	return _Registry.Contract.GetExchange(&_Registry.CallOpts, id)
}

// GetExchange is a free data retrieval call binding the contract method 0x0b9d5847.
//
// Solidity: function getExchange(uint256 id) view returns((string,address,bool))
func (_Registry *RegistryCallerSession) GetExchange(id *big.Int) (ExchangeLibExchangeInfo, error) {
	return _Registry.Contract.GetExchange(&_Registry.CallOpts, id)
}

// GetExchangeCount is a free data retrieval call binding the contract method 0x4c4523a6.
//
// Solidity: function getExchangeCount() view returns(uint256)
func (_Registry *RegistryCaller) GetExchangeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getExchangeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeCount is a free data retrieval call binding the contract method 0x4c4523a6.
//
// Solidity: function getExchangeCount() view returns(uint256)
func (_Registry *RegistrySession) GetExchangeCount() (*big.Int, error) {
	return _Registry.Contract.GetExchangeCount(&_Registry.CallOpts)
}

// GetExchangeCount is a free data retrieval call binding the contract method 0x4c4523a6.
//
// Solidity: function getExchangeCount() view returns(uint256)
func (_Registry *RegistryCallerSession) GetExchangeCount() (*big.Int, error) {
	return _Registry.Contract.GetExchangeCount(&_Registry.CallOpts)
}

// GetExchanges is a free data retrieval call binding the contract method 0x1e2e3a6b.
//
// Solidity: function getExchanges() view returns((string,address,bool)[])
func (_Registry *RegistryCaller) GetExchanges(opts *bind.CallOpts) ([]ExchangeLibExchangeInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getExchanges")

	if err != nil {
		return *new([]ExchangeLibExchangeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ExchangeLibExchangeInfo)).(*[]ExchangeLibExchangeInfo)

	return out0, err

}

// GetExchanges is a free data retrieval call binding the contract method 0x1e2e3a6b.
//
// Solidity: function getExchanges() view returns((string,address,bool)[])
func (_Registry *RegistrySession) GetExchanges() ([]ExchangeLibExchangeInfo, error) {
	return _Registry.Contract.GetExchanges(&_Registry.CallOpts)
}

// GetExchanges is a free data retrieval call binding the contract method 0x1e2e3a6b.
//
// Solidity: function getExchanges() view returns((string,address,bool)[])
func (_Registry *RegistryCallerSession) GetExchanges() ([]ExchangeLibExchangeInfo, error) {
	return _Registry.Contract.GetExchanges(&_Registry.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0xbbe4f6db.
//
// Solidity: function getPool(address pool) view returns((address,address,uint256,bool))
func (_Registry *RegistryCaller) GetPool(opts *bind.CallOpts, pool common.Address) (PoolLibPoolInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getPool", pool)

	if err != nil {
		return *new(PoolLibPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolLibPoolInfo)).(*PoolLibPoolInfo)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0xbbe4f6db.
//
// Solidity: function getPool(address pool) view returns((address,address,uint256,bool))
func (_Registry *RegistrySession) GetPool(pool common.Address) (PoolLibPoolInfo, error) {
	return _Registry.Contract.GetPool(&_Registry.CallOpts, pool)
}

// GetPool is a free data retrieval call binding the contract method 0xbbe4f6db.
//
// Solidity: function getPool(address pool) view returns((address,address,uint256,bool))
func (_Registry *RegistryCallerSession) GetPool(pool common.Address) (PoolLibPoolInfo, error) {
	return _Registry.Contract.GetPool(&_Registry.CallOpts, pool)
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() view returns(uint256)
func (_Registry *RegistryCaller) GetPoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getPoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() view returns(uint256)
func (_Registry *RegistrySession) GetPoolCount() (*big.Int, error) {
	return _Registry.Contract.GetPoolCount(&_Registry.CallOpts)
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() view returns(uint256)
func (_Registry *RegistryCallerSession) GetPoolCount() (*big.Int, error) {
	return _Registry.Contract.GetPoolCount(&_Registry.CallOpts)
}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns(address[])
func (_Registry *RegistryCaller) GetPools(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getPools")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns(address[])
func (_Registry *RegistrySession) GetPools() ([]common.Address, error) {
	return _Registry.Contract.GetPools(&_Registry.CallOpts)
}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns(address[])
func (_Registry *RegistryCallerSession) GetPools() ([]common.Address, error) {
	return _Registry.Contract.GetPools(&_Registry.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address token) view returns((uint8,bool))
func (_Registry *RegistryCaller) GetToken(opts *bind.CallOpts, token common.Address) (TokenLibTokenInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getToken", token)

	if err != nil {
		return *new(TokenLibTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenLibTokenInfo)).(*TokenLibTokenInfo)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address token) view returns((uint8,bool))
func (_Registry *RegistrySession) GetToken(token common.Address) (TokenLibTokenInfo, error) {
	return _Registry.Contract.GetToken(&_Registry.CallOpts, token)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address token) view returns((uint8,bool))
func (_Registry *RegistryCallerSession) GetToken(token common.Address) (TokenLibTokenInfo, error) {
	return _Registry.Contract.GetToken(&_Registry.CallOpts, token)
}

// GetTokenCount is a free data retrieval call binding the contract method 0x78a89567.
//
// Solidity: function getTokenCount() view returns(uint256)
func (_Registry *RegistryCaller) GetTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCount is a free data retrieval call binding the contract method 0x78a89567.
//
// Solidity: function getTokenCount() view returns(uint256)
func (_Registry *RegistrySession) GetTokenCount() (*big.Int, error) {
	return _Registry.Contract.GetTokenCount(&_Registry.CallOpts)
}

// GetTokenCount is a free data retrieval call binding the contract method 0x78a89567.
//
// Solidity: function getTokenCount() view returns(uint256)
func (_Registry *RegistryCallerSession) GetTokenCount() (*big.Int, error) {
	return _Registry.Contract.GetTokenCount(&_Registry.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Registry *RegistryCaller) GetTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Registry *RegistrySession) GetTokens() ([]common.Address, error) {
	return _Registry.Contract.GetTokens(&_Registry.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Registry *RegistryCallerSession) GetTokens() ([]common.Address, error) {
	return _Registry.Contract.GetTokens(&_Registry.CallOpts)
}

// IsPoolEnabled is a free data retrieval call binding the contract method 0x2798795a.
//
// Solidity: function isPoolEnabled(address pool) view returns(bool)
func (_Registry *RegistryCaller) IsPoolEnabled(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isPoolEnabled", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolEnabled is a free data retrieval call binding the contract method 0x2798795a.
//
// Solidity: function isPoolEnabled(address pool) view returns(bool)
func (_Registry *RegistrySession) IsPoolEnabled(pool common.Address) (bool, error) {
	return _Registry.Contract.IsPoolEnabled(&_Registry.CallOpts, pool)
}

// IsPoolEnabled is a free data retrieval call binding the contract method 0x2798795a.
//
// Solidity: function isPoolEnabled(address pool) view returns(bool)
func (_Registry *RegistryCallerSession) IsPoolEnabled(pool common.Address) (bool, error) {
	return _Registry.Contract.IsPoolEnabled(&_Registry.CallOpts, pool)
}

// IsTokenEnabled is a free data retrieval call binding the contract method 0x748538d9.
//
// Solidity: function isTokenEnabled(address token) view returns(bool)
func (_Registry *RegistryCaller) IsTokenEnabled(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isTokenEnabled", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenEnabled is a free data retrieval call binding the contract method 0x748538d9.
//
// Solidity: function isTokenEnabled(address token) view returns(bool)
func (_Registry *RegistrySession) IsTokenEnabled(token common.Address) (bool, error) {
	return _Registry.Contract.IsTokenEnabled(&_Registry.CallOpts, token)
}

// IsTokenEnabled is a free data retrieval call binding the contract method 0x748538d9.
//
// Solidity: function isTokenEnabled(address token) view returns(bool)
func (_Registry *RegistryCallerSession) IsTokenEnabled(token common.Address) (bool, error) {
	return _Registry.Contract.IsTokenEnabled(&_Registry.CallOpts, token)
}

// AddExchange is a paid mutator transaction binding the contract method 0x69581898.
//
// Solidity: function addExchange(string name, address router) returns(uint256 id)
func (_Registry *RegistryTransactor) AddExchange(opts *bind.TransactOpts, name string, router common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addExchange", name, router)
}

// AddExchange is a paid mutator transaction binding the contract method 0x69581898.
//
// Solidity: function addExchange(string name, address router) returns(uint256 id)
func (_Registry *RegistrySession) AddExchange(name string, router common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddExchange(&_Registry.TransactOpts, name, router)
}

// AddExchange is a paid mutator transaction binding the contract method 0x69581898.
//
// Solidity: function addExchange(string name, address router) returns(uint256 id)
func (_Registry *RegistryTransactorSession) AddExchange(name string, router common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddExchange(&_Registry.TransactOpts, name, router)
}

// AddPool is a paid mutator transaction binding the contract method 0x13d59a83.
//
// Solidity: function addPool(address pool, address token0, address token1, uint256 exchangeId) returns()
func (_Registry *RegistryTransactor) AddPool(opts *bind.TransactOpts, pool common.Address, token0 common.Address, token1 common.Address, exchangeId *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addPool", pool, token0, token1, exchangeId)
}

// AddPool is a paid mutator transaction binding the contract method 0x13d59a83.
//
// Solidity: function addPool(address pool, address token0, address token1, uint256 exchangeId) returns()
func (_Registry *RegistrySession) AddPool(pool common.Address, token0 common.Address, token1 common.Address, exchangeId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.AddPool(&_Registry.TransactOpts, pool, token0, token1, exchangeId)
}

// AddPool is a paid mutator transaction binding the contract method 0x13d59a83.
//
// Solidity: function addPool(address pool, address token0, address token1, uint256 exchangeId) returns()
func (_Registry *RegistryTransactorSession) AddPool(pool common.Address, token0 common.Address, token1 common.Address, exchangeId *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.AddPool(&_Registry.TransactOpts, pool, token0, token1, exchangeId)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address token, uint8 decimals) returns()
func (_Registry *RegistryTransactor) AddToken(opts *bind.TransactOpts, token common.Address, decimals uint8) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addToken", token, decimals)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address token, uint8 decimals) returns()
func (_Registry *RegistrySession) AddToken(token common.Address, decimals uint8) (*types.Transaction, error) {
	return _Registry.Contract.AddToken(&_Registry.TransactOpts, token, decimals)
}

// AddToken is a paid mutator transaction binding the contract method 0xa74ea63f.
//
// Solidity: function addToken(address token, uint8 decimals) returns()
func (_Registry *RegistryTransactorSession) AddToken(token common.Address, decimals uint8) (*types.Transaction, error) {
	return _Registry.Contract.AddToken(&_Registry.TransactOpts, token, decimals)
}

// AddTokens is a paid mutator transaction binding the contract method 0xfb93e5da.
//
// Solidity: function addTokens(address[] tokens, uint8[] decimals) returns()
func (_Registry *RegistryTransactor) AddTokens(opts *bind.TransactOpts, tokens []common.Address, decimals []uint8) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addTokens", tokens, decimals)
}

// AddTokens is a paid mutator transaction binding the contract method 0xfb93e5da.
//
// Solidity: function addTokens(address[] tokens, uint8[] decimals) returns()
func (_Registry *RegistrySession) AddTokens(tokens []common.Address, decimals []uint8) (*types.Transaction, error) {
	return _Registry.Contract.AddTokens(&_Registry.TransactOpts, tokens, decimals)
}

// AddTokens is a paid mutator transaction binding the contract method 0xfb93e5da.
//
// Solidity: function addTokens(address[] tokens, uint8[] decimals) returns()
func (_Registry *RegistryTransactorSession) AddTokens(tokens []common.Address, decimals []uint8) (*types.Transaction, error) {
	return _Registry.Contract.AddTokens(&_Registry.TransactOpts, tokens, decimals)
}

// SetExchangeEnabled is a paid mutator transaction binding the contract method 0xff612cd0.
//
// Solidity: function setExchangeEnabled(uint256 id, bool enabled) returns()
func (_Registry *RegistryTransactor) SetExchangeEnabled(opts *bind.TransactOpts, id *big.Int, enabled bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setExchangeEnabled", id, enabled)
}

// SetExchangeEnabled is a paid mutator transaction binding the contract method 0xff612cd0.
//
// Solidity: function setExchangeEnabled(uint256 id, bool enabled) returns()
func (_Registry *RegistrySession) SetExchangeEnabled(id *big.Int, enabled bool) (*types.Transaction, error) {
	return _Registry.Contract.SetExchangeEnabled(&_Registry.TransactOpts, id, enabled)
}

// SetExchangeEnabled is a paid mutator transaction binding the contract method 0xff612cd0.
//
// Solidity: function setExchangeEnabled(uint256 id, bool enabled) returns()
func (_Registry *RegistryTransactorSession) SetExchangeEnabled(id *big.Int, enabled bool) (*types.Transaction, error) {
	return _Registry.Contract.SetExchangeEnabled(&_Registry.TransactOpts, id, enabled)
}

// SetPoolEnabled is a paid mutator transaction binding the contract method 0xb27925ff.
//
// Solidity: function setPoolEnabled(address pool, bool enabled) returns()
func (_Registry *RegistryTransactor) SetPoolEnabled(opts *bind.TransactOpts, pool common.Address, enabled bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setPoolEnabled", pool, enabled)
}

// SetPoolEnabled is a paid mutator transaction binding the contract method 0xb27925ff.
//
// Solidity: function setPoolEnabled(address pool, bool enabled) returns()
func (_Registry *RegistrySession) SetPoolEnabled(pool common.Address, enabled bool) (*types.Transaction, error) {
	return _Registry.Contract.SetPoolEnabled(&_Registry.TransactOpts, pool, enabled)
}

// SetPoolEnabled is a paid mutator transaction binding the contract method 0xb27925ff.
//
// Solidity: function setPoolEnabled(address pool, bool enabled) returns()
func (_Registry *RegistryTransactorSession) SetPoolEnabled(pool common.Address, enabled bool) (*types.Transaction, error) {
	return _Registry.Contract.SetPoolEnabled(&_Registry.TransactOpts, pool, enabled)
}

// SetTokenEnabled is a paid mutator transaction binding the contract method 0xfeaef9bd.
//
// Solidity: function setTokenEnabled(address token, bool enabled) returns()
func (_Registry *RegistryTransactor) SetTokenEnabled(opts *bind.TransactOpts, token common.Address, enabled bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setTokenEnabled", token, enabled)
}

// SetTokenEnabled is a paid mutator transaction binding the contract method 0xfeaef9bd.
//
// Solidity: function setTokenEnabled(address token, bool enabled) returns()
func (_Registry *RegistrySession) SetTokenEnabled(token common.Address, enabled bool) (*types.Transaction, error) {
	return _Registry.Contract.SetTokenEnabled(&_Registry.TransactOpts, token, enabled)
}

// SetTokenEnabled is a paid mutator transaction binding the contract method 0xfeaef9bd.
//
// Solidity: function setTokenEnabled(address token, bool enabled) returns()
func (_Registry *RegistryTransactorSession) SetTokenEnabled(token common.Address, enabled bool) (*types.Transaction, error) {
	return _Registry.Contract.SetTokenEnabled(&_Registry.TransactOpts, token, enabled)
}
