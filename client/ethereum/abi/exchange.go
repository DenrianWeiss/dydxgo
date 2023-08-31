// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// ExchangeMetaData contains all meta data concerning the Exchange contract.
var ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configurationDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"configurationHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalConfigurationHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptedGovernor\",\"type\":\"address\"}],\"name\":\"LogNewGovernorAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominatedGovernor\",\"type\":\"address\"}],\"name\":\"LogNominatedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogNominationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRemovalIntent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"entry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entryId\",\"type\":\"string\"}],\"name\":\"LogRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"removedGovernor\",\"type\":\"address\"}],\"name\":\"LogRemovedGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogUnFrozen\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_CANCEL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAIN_GOVERNANCE_INFO_TAG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FORCED_ACTIONS_REQS_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VERIFIER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNFREEZE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_REMOVAL_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"announceAvailabilityVerifierRemovalIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"announceVerifierRemovalIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredAvailabilityVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"isAvailabilityVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainAcceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainCancelNomination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testGovernor\",\"type\":\"address\"}],\"name\":\"mainIsGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"mainNominateNewGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governorForRemoval\",\"type\":\"address\"}],\"name\":\"mainRemoveGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"registerAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"name\":\"registerVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeAvailabilityVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogDepositCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogDepositNftCancelReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogMintWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogMintableWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositorEthKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"LogNftWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogNftWithdrawalPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"LogSystemAssetType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAdmin\",\"type\":\"address\"}],\"name\":\"LogTokenAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"name\":\"LogTokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAdmin\",\"type\":\"address\"}],\"name\":\"LogUserAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAdmin\",\"type\":\"address\"}],\"name\":\"LogUserAdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogUserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawalAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonQuantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"LogWithdrawalPerformed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"depositNftReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"depositReclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getCancellationRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"}],\"name\":\"getEthKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getQuantizedDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"presumedAssetType\",\"type\":\"uint256\"}],\"name\":\"getQuantum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemAssetType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"isAssetRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedAdmin\",\"type\":\"address\"}],\"name\":\"isTokenAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedAdmin\",\"type\":\"address\"}],\"name\":\"isUserAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"registerAndDepositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"registerAndDepositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"assetInfo\",\"type\":\"bytes\"}],\"name\":\"registerSystemAssetType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"registerTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"registerUserAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"}],\"name\":\"unregisterTokenAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"}],\"name\":\"unregisterUserAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mintingBlob\",\"type\":\"bytes\"}],\"name\":\"withdrawAndMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawNftTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"LogAssetConfigurationApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"LogAssetConfigurationRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"LogAssetConfigurationRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"LogGlobalConfigurationApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"LogGlobalConfigurationRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"LogGlobalConfigurationRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"LogOperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateTransitionFact\",\"type\":\"bytes32\"}],\"name\":\"LogStateTransitionFact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"LogUpdateState\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"applyAssetConfigurationChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"applyGlobalConfigurationChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"escape\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKeyA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"starkKeyB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syntheticAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSynthetic\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"aIsBuyingSynthetic\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getForcedTradeRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"getForcedWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seq\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultTreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"testedOperator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"registerAssetConfigurationChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"registerGlobalConfigurationChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"removeAssetConfigurationChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"configHash\",\"type\":\"bytes32\"}],\"name\":\"removeGlobalConfigurationChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"}],\"name\":\"unregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"programOutput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"applicationData\",\"type\":\"uint256[]\"}],\"name\":\"updateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKeyA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKeyB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultIdA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultIdB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"syntheticAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSynthetic\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"aIsBuyingSynthetic\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"LogForcedTradeRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"LogForcedWithdrawalRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKeyA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"starkKeyB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syntheticAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSynthetic\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"aIsBuyingSynthetic\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submissionExpirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"premiumCost\",\"type\":\"bool\"}],\"name\":\"forcedTradeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"premiumCost\",\"type\":\"bool\"}],\"name\":\"forcedWithdrawalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKeyA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"starkKeyB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultIdB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syntheticAssetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSynthetic\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"aIsBuyingSynthetic\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"freezeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"starkKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantizedAmount\",\"type\":\"uint256\"}],\"name\":\"freezeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeMetaData.ABI instead.
var ExchangeABI = ExchangeMetaData.ABI

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Exchange *ExchangeCaller) DEPOSITCANCELDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "DEPOSIT_CANCEL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Exchange *ExchangeSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _Exchange.Contract.DEPOSITCANCELDELAY(&_Exchange.CallOpts)
}

// DEPOSITCANCELDELAY is a free data retrieval call binding the contract method 0x77e84e0d.
//
// Solidity: function DEPOSIT_CANCEL_DELAY() view returns(uint256)
func (_Exchange *ExchangeCallerSession) DEPOSITCANCELDELAY() (*big.Int, error) {
	return _Exchange.Contract.DEPOSITCANCELDELAY(&_Exchange.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Exchange *ExchangeCaller) FREEZEGRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "FREEZE_GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Exchange *ExchangeSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _Exchange.Contract.FREEZEGRACEPERIOD(&_Exchange.CallOpts)
}

// FREEZEGRACEPERIOD is a free data retrieval call binding the contract method 0x00717542.
//
// Solidity: function FREEZE_GRACE_PERIOD() view returns(uint256)
func (_Exchange *ExchangeCallerSession) FREEZEGRACEPERIOD() (*big.Int, error) {
	return _Exchange.Contract.FREEZEGRACEPERIOD(&_Exchange.CallOpts)
}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_Exchange *ExchangeCaller) MAINGOVERNANCEINFOTAG(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "MAIN_GOVERNANCE_INFO_TAG")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_Exchange *ExchangeSession) MAINGOVERNANCEINFOTAG() (string, error) {
	return _Exchange.Contract.MAINGOVERNANCEINFOTAG(&_Exchange.CallOpts)
}

// MAINGOVERNANCEINFOTAG is a free data retrieval call binding the contract method 0xc23b60ef.
//
// Solidity: function MAIN_GOVERNANCE_INFO_TAG() view returns(string)
func (_Exchange *ExchangeCallerSession) MAINGOVERNANCEINFOTAG() (string, error) {
	return _Exchange.Contract.MAINGOVERNANCEINFOTAG(&_Exchange.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Exchange *ExchangeCaller) MAXFORCEDACTIONSREQSPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "MAX_FORCED_ACTIONS_REQS_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Exchange *ExchangeSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _Exchange.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_Exchange.CallOpts)
}

// MAXFORCEDACTIONSREQSPERBLOCK is a free data retrieval call binding the contract method 0xe30a5cff.
//
// Solidity: function MAX_FORCED_ACTIONS_REQS_PER_BLOCK() view returns(uint256)
func (_Exchange *ExchangeCallerSession) MAXFORCEDACTIONSREQSPERBLOCK() (*big.Int, error) {
	return _Exchange.Contract.MAXFORCEDACTIONSREQSPERBLOCK(&_Exchange.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Exchange *ExchangeCaller) MAXVERIFIERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "MAX_VERIFIER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Exchange *ExchangeSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _Exchange.Contract.MAXVERIFIERCOUNT(&_Exchange.CallOpts)
}

// MAXVERIFIERCOUNT is a free data retrieval call binding the contract method 0xe6de6282.
//
// Solidity: function MAX_VERIFIER_COUNT() view returns(uint256)
func (_Exchange *ExchangeCallerSession) MAXVERIFIERCOUNT() (*big.Int, error) {
	return _Exchange.Contract.MAXVERIFIERCOUNT(&_Exchange.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Exchange *ExchangeCaller) UNFREEZEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "UNFREEZE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Exchange *ExchangeSession) UNFREEZEDELAY() (*big.Int, error) {
	return _Exchange.Contract.UNFREEZEDELAY(&_Exchange.CallOpts)
}

// UNFREEZEDELAY is a free data retrieval call binding the contract method 0x993f3639.
//
// Solidity: function UNFREEZE_DELAY() view returns(uint256)
func (_Exchange *ExchangeCallerSession) UNFREEZEDELAY() (*big.Int, error) {
	return _Exchange.Contract.UNFREEZEDELAY(&_Exchange.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Exchange *ExchangeCaller) VERIFIERREMOVALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "VERIFIER_REMOVAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Exchange *ExchangeSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _Exchange.Contract.VERIFIERREMOVALDELAY(&_Exchange.CallOpts)
}

// VERIFIERREMOVALDELAY is a free data retrieval call binding the contract method 0xb7663112.
//
// Solidity: function VERIFIER_REMOVAL_DELAY() view returns(uint256)
func (_Exchange *ExchangeCallerSession) VERIFIERREMOVALDELAY() (*big.Int, error) {
	return _Exchange.Contract.VERIFIERREMOVALDELAY(&_Exchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Exchange *ExchangeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Exchange *ExchangeSession) VERSION() (string, error) {
	return _Exchange.Contract.VERSION(&_Exchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Exchange *ExchangeCallerSession) VERSION() (string, error) {
	return _Exchange.Contract.VERSION(&_Exchange.CallOpts)
}

// ConfigurationDelay is a free data retrieval call binding the contract method 0xc1a85130.
//
// Solidity: function configurationDelay() view returns(uint256)
func (_Exchange *ExchangeCaller) ConfigurationDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "configurationDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfigurationDelay is a free data retrieval call binding the contract method 0xc1a85130.
//
// Solidity: function configurationDelay() view returns(uint256)
func (_Exchange *ExchangeSession) ConfigurationDelay() (*big.Int, error) {
	return _Exchange.Contract.ConfigurationDelay(&_Exchange.CallOpts)
}

// ConfigurationDelay is a free data retrieval call binding the contract method 0xc1a85130.
//
// Solidity: function configurationDelay() view returns(uint256)
func (_Exchange *ExchangeCallerSession) ConfigurationDelay() (*big.Int, error) {
	return _Exchange.Contract.ConfigurationDelay(&_Exchange.CallOpts)
}

// ConfigurationHash is a free data retrieval call binding the contract method 0xf2011f66.
//
// Solidity: function configurationHash(uint256 ) view returns(bytes32)
func (_Exchange *ExchangeCaller) ConfigurationHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "configurationHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConfigurationHash is a free data retrieval call binding the contract method 0xf2011f66.
//
// Solidity: function configurationHash(uint256 ) view returns(bytes32)
func (_Exchange *ExchangeSession) ConfigurationHash(arg0 *big.Int) ([32]byte, error) {
	return _Exchange.Contract.ConfigurationHash(&_Exchange.CallOpts, arg0)
}

// ConfigurationHash is a free data retrieval call binding the contract method 0xf2011f66.
//
// Solidity: function configurationHash(uint256 ) view returns(bytes32)
func (_Exchange *ExchangeCallerSession) ConfigurationHash(arg0 *big.Int) ([32]byte, error) {
	return _Exchange.Contract.ConfigurationHash(&_Exchange.CallOpts, arg0)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Exchange *ExchangeCaller) GetAssetInfo(opts *bind.CallOpts, assetType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getAssetInfo", assetType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Exchange *ExchangeSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _Exchange.Contract.GetAssetInfo(&_Exchange.CallOpts, assetType)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xf637d950.
//
// Solidity: function getAssetInfo(uint256 assetType) view returns(bytes assetInfo)
func (_Exchange *ExchangeCallerSession) GetAssetInfo(assetType *big.Int) ([]byte, error) {
	return _Exchange.Contract.GetAssetInfo(&_Exchange.CallOpts, assetType)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Exchange *ExchangeCaller) GetCancellationRequest(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getCancellationRequest", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Exchange *ExchangeSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetCancellationRequest(&_Exchange.CallOpts, starkKey, assetId, vaultId)
}

// GetCancellationRequest is a free data retrieval call binding the contract method 0x333ac20b.
//
// Solidity: function getCancellationRequest(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 request)
func (_Exchange *ExchangeCallerSession) GetCancellationRequest(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetCancellationRequest(&_Exchange.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Exchange *ExchangeCaller) GetDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Exchange *ExchangeSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetDepositBalance(&_Exchange.CallOpts, starkKey, assetId, vaultId)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0xabf98fe1.
//
// Solidity: function getDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Exchange *ExchangeCallerSession) GetDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetDepositBalance(&_Exchange.CallOpts, starkKey, assetId, vaultId)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_Exchange *ExchangeCaller) GetEthKey(opts *bind.CallOpts, starkKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getEthKey", starkKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_Exchange *ExchangeSession) GetEthKey(starkKey *big.Int) (common.Address, error) {
	return _Exchange.Contract.GetEthKey(&_Exchange.CallOpts, starkKey)
}

// GetEthKey is a free data retrieval call binding the contract method 0x1dbd1da7.
//
// Solidity: function getEthKey(uint256 starkKey) view returns(address ethKey)
func (_Exchange *ExchangeCallerSession) GetEthKey(starkKey *big.Int) (common.Address, error) {
	return _Exchange.Contract.GetEthKey(&_Exchange.CallOpts, starkKey)
}

// GetForcedTradeRequest is a free data retrieval call binding the contract method 0xf00cec4b.
//
// Solidity: function getForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) view returns(uint256)
func (_Exchange *ExchangeCaller) GetForcedTradeRequest(opts *bind.CallOpts, starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getForcedTradeRequest", starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedTradeRequest is a free data retrieval call binding the contract method 0xf00cec4b.
//
// Solidity: function getForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) view returns(uint256)
func (_Exchange *ExchangeSession) GetForcedTradeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetForcedTradeRequest(&_Exchange.CallOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// GetForcedTradeRequest is a free data retrieval call binding the contract method 0xf00cec4b.
//
// Solidity: function getForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) view returns(uint256)
func (_Exchange *ExchangeCallerSession) GetForcedTradeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetForcedTradeRequest(&_Exchange.CallOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// GetForcedWithdrawalRequest is a free data retrieval call binding the contract method 0x260e96dd.
//
// Solidity: function getForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) view returns(uint256)
func (_Exchange *ExchangeCaller) GetForcedWithdrawalRequest(opts *bind.CallOpts, starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getForcedWithdrawalRequest", starkKey, vaultId, quantizedAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedWithdrawalRequest is a free data retrieval call binding the contract method 0x260e96dd.
//
// Solidity: function getForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) view returns(uint256)
func (_Exchange *ExchangeSession) GetForcedWithdrawalRequest(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetForcedWithdrawalRequest(&_Exchange.CallOpts, starkKey, vaultId, quantizedAmount)
}

// GetForcedWithdrawalRequest is a free data retrieval call binding the contract method 0x260e96dd.
//
// Solidity: function getForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) view returns(uint256)
func (_Exchange *ExchangeCallerSession) GetForcedWithdrawalRequest(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetForcedWithdrawalRequest(&_Exchange.CallOpts, starkKey, vaultId, quantizedAmount)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_Exchange *ExchangeCaller) GetLastBatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getLastBatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_Exchange *ExchangeSession) GetLastBatchId() (*big.Int, error) {
	return _Exchange.Contract.GetLastBatchId(&_Exchange.CallOpts)
}

// GetLastBatchId is a free data retrieval call binding the contract method 0x515535e8.
//
// Solidity: function getLastBatchId() view returns(uint256 batchId)
func (_Exchange *ExchangeCallerSession) GetLastBatchId() (*big.Int, error) {
	return _Exchange.Contract.GetLastBatchId(&_Exchange.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_Exchange *ExchangeCaller) GetOrderRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getOrderRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_Exchange *ExchangeSession) GetOrderRoot() (*big.Int, error) {
	return _Exchange.Contract.GetOrderRoot(&_Exchange.CallOpts)
}

// GetOrderRoot is a free data retrieval call binding the contract method 0x0dded952.
//
// Solidity: function getOrderRoot() view returns(uint256 root)
func (_Exchange *ExchangeCallerSession) GetOrderRoot() (*big.Int, error) {
	return _Exchange.Contract.GetOrderRoot(&_Exchange.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_Exchange *ExchangeCaller) GetOrderTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getOrderTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_Exchange *ExchangeSession) GetOrderTreeHeight() (*big.Int, error) {
	return _Exchange.Contract.GetOrderTreeHeight(&_Exchange.CallOpts)
}

// GetOrderTreeHeight is a free data retrieval call binding the contract method 0x7e9da4c5.
//
// Solidity: function getOrderTreeHeight() view returns(uint256 height)
func (_Exchange *ExchangeCallerSession) GetOrderTreeHeight() (*big.Int, error) {
	return _Exchange.Contract.GetOrderTreeHeight(&_Exchange.CallOpts)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Exchange *ExchangeCaller) GetQuantizedDepositBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getQuantizedDepositBalance", starkKey, assetId, vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Exchange *ExchangeSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetQuantizedDepositBalance(&_Exchange.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantizedDepositBalance is a free data retrieval call binding the contract method 0x4e8912da.
//
// Solidity: function getQuantizedDepositBalance(uint256 starkKey, uint256 assetId, uint256 vaultId) view returns(uint256 balance)
func (_Exchange *ExchangeCallerSession) GetQuantizedDepositBalance(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetQuantizedDepositBalance(&_Exchange.CallOpts, starkKey, assetId, vaultId)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Exchange *ExchangeCaller) GetQuantum(opts *bind.CallOpts, presumedAssetType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getQuantum", presumedAssetType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Exchange *ExchangeSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetQuantum(&_Exchange.CallOpts, presumedAssetType)
}

// GetQuantum is a free data retrieval call binding the contract method 0xdd7202d8.
//
// Solidity: function getQuantum(uint256 presumedAssetType) view returns(uint256 quantum)
func (_Exchange *ExchangeCallerSession) GetQuantum(presumedAssetType *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetQuantum(&_Exchange.CallOpts, presumedAssetType)
}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_Exchange *ExchangeCaller) GetRegisteredAvailabilityVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getRegisteredAvailabilityVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_Exchange *ExchangeSession) GetRegisteredAvailabilityVerifiers() ([]common.Address, error) {
	return _Exchange.Contract.GetRegisteredAvailabilityVerifiers(&_Exchange.CallOpts)
}

// GetRegisteredAvailabilityVerifiers is a free data retrieval call binding the contract method 0x1ac347f2.
//
// Solidity: function getRegisteredAvailabilityVerifiers() view returns(address[] _verifers)
func (_Exchange *ExchangeCallerSession) GetRegisteredAvailabilityVerifiers() ([]common.Address, error) {
	return _Exchange.Contract.GetRegisteredAvailabilityVerifiers(&_Exchange.CallOpts)
}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_Exchange *ExchangeCaller) GetRegisteredVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getRegisteredVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_Exchange *ExchangeSession) GetRegisteredVerifiers() ([]common.Address, error) {
	return _Exchange.Contract.GetRegisteredVerifiers(&_Exchange.CallOpts)
}

// GetRegisteredVerifiers is a free data retrieval call binding the contract method 0x4eab9ed3.
//
// Solidity: function getRegisteredVerifiers() view returns(address[] _verifers)
func (_Exchange *ExchangeCallerSession) GetRegisteredVerifiers() ([]common.Address, error) {
	return _Exchange.Contract.GetRegisteredVerifiers(&_Exchange.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_Exchange *ExchangeCaller) GetSequenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getSequenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_Exchange *ExchangeSession) GetSequenceNumber() (*big.Int, error) {
	return _Exchange.Contract.GetSequenceNumber(&_Exchange.CallOpts)
}

// GetSequenceNumber is a free data retrieval call binding the contract method 0x42af35fd.
//
// Solidity: function getSequenceNumber() view returns(uint256 seq)
func (_Exchange *ExchangeCallerSession) GetSequenceNumber() (*big.Int, error) {
	return _Exchange.Contract.GetSequenceNumber(&_Exchange.CallOpts)
}

// GetSystemAssetType is a free data retrieval call binding the contract method 0x27b66a4d.
//
// Solidity: function getSystemAssetType() view returns(uint256)
func (_Exchange *ExchangeCaller) GetSystemAssetType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getSystemAssetType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSystemAssetType is a free data retrieval call binding the contract method 0x27b66a4d.
//
// Solidity: function getSystemAssetType() view returns(uint256)
func (_Exchange *ExchangeSession) GetSystemAssetType() (*big.Int, error) {
	return _Exchange.Contract.GetSystemAssetType(&_Exchange.CallOpts)
}

// GetSystemAssetType is a free data retrieval call binding the contract method 0x27b66a4d.
//
// Solidity: function getSystemAssetType() view returns(uint256)
func (_Exchange *ExchangeCallerSession) GetSystemAssetType() (*big.Int, error) {
	return _Exchange.Contract.GetSystemAssetType(&_Exchange.CallOpts)
}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_Exchange *ExchangeCaller) GetVaultRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getVaultRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_Exchange *ExchangeSession) GetVaultRoot() (*big.Int, error) {
	return _Exchange.Contract.GetVaultRoot(&_Exchange.CallOpts)
}

// GetVaultRoot is a free data retrieval call binding the contract method 0x64da5dfe.
//
// Solidity: function getVaultRoot() view returns(uint256 root)
func (_Exchange *ExchangeCallerSession) GetVaultRoot() (*big.Int, error) {
	return _Exchange.Contract.GetVaultRoot(&_Exchange.CallOpts)
}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_Exchange *ExchangeCaller) GetVaultTreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getVaultTreeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_Exchange *ExchangeSession) GetVaultTreeHeight() (*big.Int, error) {
	return _Exchange.Contract.GetVaultTreeHeight(&_Exchange.CallOpts)
}

// GetVaultTreeHeight is a free data retrieval call binding the contract method 0xf288a3ff.
//
// Solidity: function getVaultTreeHeight() view returns(uint256 height)
func (_Exchange *ExchangeCallerSession) GetVaultTreeHeight() (*big.Int, error) {
	return _Exchange.Contract.GetVaultTreeHeight(&_Exchange.CallOpts)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 starkKey, uint256 assetId) view returns(uint256 balance)
func (_Exchange *ExchangeCaller) GetWithdrawalBalance(opts *bind.CallOpts, starkKey *big.Int, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getWithdrawalBalance", starkKey, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 starkKey, uint256 assetId) view returns(uint256 balance)
func (_Exchange *ExchangeSession) GetWithdrawalBalance(starkKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetWithdrawalBalance(&_Exchange.CallOpts, starkKey, assetId)
}

// GetWithdrawalBalance is a free data retrieval call binding the contract method 0xec3161b0.
//
// Solidity: function getWithdrawalBalance(uint256 starkKey, uint256 assetId) view returns(uint256 balance)
func (_Exchange *ExchangeCallerSession) GetWithdrawalBalance(starkKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetWithdrawalBalance(&_Exchange.CallOpts, starkKey, assetId)
}

// GlobalConfigurationHash is a free data retrieval call binding the contract method 0xadac3e15.
//
// Solidity: function globalConfigurationHash() view returns(bytes32)
func (_Exchange *ExchangeCaller) GlobalConfigurationHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "globalConfigurationHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GlobalConfigurationHash is a free data retrieval call binding the contract method 0xadac3e15.
//
// Solidity: function globalConfigurationHash() view returns(bytes32)
func (_Exchange *ExchangeSession) GlobalConfigurationHash() ([32]byte, error) {
	return _Exchange.Contract.GlobalConfigurationHash(&_Exchange.CallOpts)
}

// GlobalConfigurationHash is a free data retrieval call binding the contract method 0xadac3e15.
//
// Solidity: function globalConfigurationHash() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) GlobalConfigurationHash() ([32]byte, error) {
	return _Exchange.Contract.GlobalConfigurationHash(&_Exchange.CallOpts)
}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_Exchange *ExchangeCaller) IsAssetRegistered(opts *bind.CallOpts, assetType *big.Int) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isAssetRegistered", assetType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_Exchange *ExchangeSession) IsAssetRegistered(assetType *big.Int) (bool, error) {
	return _Exchange.Contract.IsAssetRegistered(&_Exchange.CallOpts, assetType)
}

// IsAssetRegistered is a free data retrieval call binding the contract method 0x049f5ade.
//
// Solidity: function isAssetRegistered(uint256 assetType) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsAssetRegistered(assetType *big.Int) (bool, error) {
	return _Exchange.Contract.IsAssetRegistered(&_Exchange.CallOpts, assetType)
}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_Exchange *ExchangeCaller) IsAvailabilityVerifier(opts *bind.CallOpts, verifierAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isAvailabilityVerifier", verifierAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_Exchange *ExchangeSession) IsAvailabilityVerifier(verifierAddress common.Address) (bool, error) {
	return _Exchange.Contract.IsAvailabilityVerifier(&_Exchange.CallOpts, verifierAddress)
}

// IsAvailabilityVerifier is a free data retrieval call binding the contract method 0xbd1279ae.
//
// Solidity: function isAvailabilityVerifier(address verifierAddress) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsAvailabilityVerifier(verifierAddress common.Address) (bool, error) {
	return _Exchange.Contract.IsAvailabilityVerifier(&_Exchange.CallOpts, verifierAddress)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Exchange *ExchangeCaller) IsFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Exchange *ExchangeSession) IsFrozen() (bool, error) {
	return _Exchange.Contract.IsFrozen(&_Exchange.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0x33eeb147.
//
// Solidity: function isFrozen() view returns(bool)
func (_Exchange *ExchangeCallerSession) IsFrozen() (bool, error) {
	return _Exchange.Contract.IsFrozen(&_Exchange.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_Exchange *ExchangeCaller) IsOperator(opts *bind.CallOpts, testedOperator common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isOperator", testedOperator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_Exchange *ExchangeSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _Exchange.Contract.IsOperator(&_Exchange.CallOpts, testedOperator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address testedOperator) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsOperator(testedOperator common.Address) (bool, error) {
	return _Exchange.Contract.IsOperator(&_Exchange.CallOpts, testedOperator)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_Exchange *ExchangeCaller) IsTokenAdmin(opts *bind.CallOpts, testedAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isTokenAdmin", testedAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_Exchange *ExchangeSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _Exchange.Contract.IsTokenAdmin(&_Exchange.CallOpts, testedAdmin)
}

// IsTokenAdmin is a free data retrieval call binding the contract method 0xa2bdde3d.
//
// Solidity: function isTokenAdmin(address testedAdmin) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsTokenAdmin(testedAdmin common.Address) (bool, error) {
	return _Exchange.Contract.IsTokenAdmin(&_Exchange.CallOpts, testedAdmin)
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address testedAdmin) view returns(bool)
func (_Exchange *ExchangeCaller) IsUserAdmin(opts *bind.CallOpts, testedAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isUserAdmin", testedAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address testedAdmin) view returns(bool)
func (_Exchange *ExchangeSession) IsUserAdmin(testedAdmin common.Address) (bool, error) {
	return _Exchange.Contract.IsUserAdmin(&_Exchange.CallOpts, testedAdmin)
}

// IsUserAdmin is a free data retrieval call binding the contract method 0x74d523a8.
//
// Solidity: function isUserAdmin(address testedAdmin) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsUserAdmin(testedAdmin common.Address) (bool, error) {
	return _Exchange.Contract.IsUserAdmin(&_Exchange.CallOpts, testedAdmin)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_Exchange *ExchangeCaller) IsVerifier(opts *bind.CallOpts, verifierAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isVerifier", verifierAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_Exchange *ExchangeSession) IsVerifier(verifierAddress common.Address) (bool, error) {
	return _Exchange.Contract.IsVerifier(&_Exchange.CallOpts, verifierAddress)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address verifierAddress) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsVerifier(verifierAddress common.Address) (bool, error) {
	return _Exchange.Contract.IsVerifier(&_Exchange.CallOpts, verifierAddress)
}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_Exchange *ExchangeCaller) MainIsGovernor(opts *bind.CallOpts, testGovernor common.Address) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "mainIsGovernor", testGovernor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_Exchange *ExchangeSession) MainIsGovernor(testGovernor common.Address) (bool, error) {
	return _Exchange.Contract.MainIsGovernor(&_Exchange.CallOpts, testGovernor)
}

// MainIsGovernor is a free data retrieval call binding the contract method 0x45f5cd97.
//
// Solidity: function mainIsGovernor(address testGovernor) view returns(bool)
func (_Exchange *ExchangeCallerSession) MainIsGovernor(testGovernor common.Address) (bool, error) {
	return _Exchange.Contract.MainIsGovernor(&_Exchange.CallOpts, testGovernor)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_Exchange *ExchangeTransactor) AnnounceAvailabilityVerifierRemovalIntent(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "announceAvailabilityVerifierRemovalIntent", verifier)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_Exchange *ExchangeSession) AnnounceAvailabilityVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_Exchange.TransactOpts, verifier)
}

// AnnounceAvailabilityVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x1d078bbb.
//
// Solidity: function announceAvailabilityVerifierRemovalIntent(address verifier) returns()
func (_Exchange *ExchangeTransactorSession) AnnounceAvailabilityVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AnnounceAvailabilityVerifierRemovalIntent(&_Exchange.TransactOpts, verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_Exchange *ExchangeTransactor) AnnounceVerifierRemovalIntent(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "announceVerifierRemovalIntent", verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_Exchange *ExchangeSession) AnnounceVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AnnounceVerifierRemovalIntent(&_Exchange.TransactOpts, verifier)
}

// AnnounceVerifierRemovalIntent is a paid mutator transaction binding the contract method 0x418573b7.
//
// Solidity: function announceVerifierRemovalIntent(address verifier) returns()
func (_Exchange *ExchangeTransactorSession) AnnounceVerifierRemovalIntent(verifier common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.AnnounceVerifierRemovalIntent(&_Exchange.TransactOpts, verifier)
}

// ApplyAssetConfigurationChange is a paid mutator transaction binding the contract method 0x9bd4f229.
//
// Solidity: function applyAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeTransactor) ApplyAssetConfigurationChange(opts *bind.TransactOpts, assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "applyAssetConfigurationChange", assetId, configHash)
}

// ApplyAssetConfigurationChange is a paid mutator transaction binding the contract method 0x9bd4f229.
//
// Solidity: function applyAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeSession) ApplyAssetConfigurationChange(assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.ApplyAssetConfigurationChange(&_Exchange.TransactOpts, assetId, configHash)
}

// ApplyAssetConfigurationChange is a paid mutator transaction binding the contract method 0x9bd4f229.
//
// Solidity: function applyAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeTransactorSession) ApplyAssetConfigurationChange(assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.ApplyAssetConfigurationChange(&_Exchange.TransactOpts, assetId, configHash)
}

// ApplyGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x80f235b0.
//
// Solidity: function applyGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeTransactor) ApplyGlobalConfigurationChange(opts *bind.TransactOpts, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "applyGlobalConfigurationChange", configHash)
}

// ApplyGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x80f235b0.
//
// Solidity: function applyGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeSession) ApplyGlobalConfigurationChange(configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.ApplyGlobalConfigurationChange(&_Exchange.TransactOpts, configHash)
}

// ApplyGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x80f235b0.
//
// Solidity: function applyGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeTransactorSession) ApplyGlobalConfigurationChange(configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.ApplyGlobalConfigurationChange(&_Exchange.TransactOpts, configHash)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeTransactor) Deposit(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "deposit", starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Deposit(&_Exchange.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0x00aeef8a.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeTransactorSession) Deposit(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Deposit(&_Exchange.TransactOpts, starkKey, assetType, vaultId)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactor) Deposit0(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "deposit0", starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Deposit0(&_Exchange.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x2505c3d9.
//
// Solidity: function deposit(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactorSession) Deposit0(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Deposit0(&_Exchange.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Exchange *ExchangeTransactor) DepositCancel(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositCancel", starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Exchange *ExchangeSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositCancel(&_Exchange.TransactOpts, starkKey, assetId, vaultId)
}

// DepositCancel is a paid mutator transaction binding the contract method 0x7df7dc04.
//
// Solidity: function depositCancel(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Exchange *ExchangeTransactorSession) DepositCancel(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositCancel(&_Exchange.TransactOpts, starkKey, assetId, vaultId)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactor) DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositERC20", starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositERC20(&_Exchange.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x9ed17084.
//
// Solidity: function depositERC20(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactorSession) DepositERC20(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositERC20(&_Exchange.TransactOpts, starkKey, assetType, vaultId, quantizedAmount)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeTransactor) DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositEth", starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositEth(&_Exchange.TransactOpts, starkKey, assetType, vaultId)
}

// DepositEth is a paid mutator transaction binding the contract method 0x6ce5d957.
//
// Solidity: function depositEth(uint256 starkKey, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeTransactorSession) DepositEth(starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositEth(&_Exchange.TransactOpts, starkKey, assetType, vaultId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Exchange *ExchangeTransactor) DepositNft(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositNft", starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Exchange *ExchangeSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositNft(&_Exchange.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNft is a paid mutator transaction binding the contract method 0xae1cdde6.
//
// Solidity: function depositNft(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Exchange *ExchangeTransactorSession) DepositNft(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositNft(&_Exchange.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Exchange *ExchangeTransactor) DepositNftReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositNftReclaim", starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Exchange *ExchangeSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositNftReclaim(&_Exchange.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositNftReclaim is a paid mutator transaction binding the contract method 0xfcb05822.
//
// Solidity: function depositNftReclaim(uint256 starkKey, uint256 assetType, uint256 vaultId, uint256 tokenId) returns()
func (_Exchange *ExchangeTransactorSession) DepositNftReclaim(starkKey *big.Int, assetType *big.Int, vaultId *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositNftReclaim(&_Exchange.TransactOpts, starkKey, assetType, vaultId, tokenId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Exchange *ExchangeTransactor) DepositReclaim(opts *bind.TransactOpts, starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositReclaim", starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Exchange *ExchangeSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositReclaim(&_Exchange.TransactOpts, starkKey, assetId, vaultId)
}

// DepositReclaim is a paid mutator transaction binding the contract method 0xae873816.
//
// Solidity: function depositReclaim(uint256 starkKey, uint256 assetId, uint256 vaultId) returns()
func (_Exchange *ExchangeTransactorSession) DepositReclaim(starkKey *big.Int, assetId *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositReclaim(&_Exchange.TransactOpts, starkKey, assetId, vaultId)
}

// Escape is a paid mutator transaction binding the contract method 0x366a2670.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactor) Escape(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "escape", starkKey, vaultId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x366a2670.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeSession) Escape(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Escape(&_Exchange.TransactOpts, starkKey, vaultId, quantizedAmount)
}

// Escape is a paid mutator transaction binding the contract method 0x366a2670.
//
// Solidity: function escape(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactorSession) Escape(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Escape(&_Exchange.TransactOpts, starkKey, vaultId, quantizedAmount)
}

// ForcedTradeRequest is a paid mutator transaction binding the contract method 0x2ecb8162.
//
// Solidity: function forcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 submissionExpirationTime, uint256 nonce, bytes signature, bool premiumCost) returns()
func (_Exchange *ExchangeTransactor) ForcedTradeRequest(opts *bind.TransactOpts, starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, submissionExpirationTime *big.Int, nonce *big.Int, signature []byte, premiumCost bool) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "forcedTradeRequest", starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, submissionExpirationTime, nonce, signature, premiumCost)
}

// ForcedTradeRequest is a paid mutator transaction binding the contract method 0x2ecb8162.
//
// Solidity: function forcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 submissionExpirationTime, uint256 nonce, bytes signature, bool premiumCost) returns()
func (_Exchange *ExchangeSession) ForcedTradeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, submissionExpirationTime *big.Int, nonce *big.Int, signature []byte, premiumCost bool) (*types.Transaction, error) {
	return _Exchange.Contract.ForcedTradeRequest(&_Exchange.TransactOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, submissionExpirationTime, nonce, signature, premiumCost)
}

// ForcedTradeRequest is a paid mutator transaction binding the contract method 0x2ecb8162.
//
// Solidity: function forcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 submissionExpirationTime, uint256 nonce, bytes signature, bool premiumCost) returns()
func (_Exchange *ExchangeTransactorSession) ForcedTradeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, submissionExpirationTime *big.Int, nonce *big.Int, signature []byte, premiumCost bool) (*types.Transaction, error) {
	return _Exchange.Contract.ForcedTradeRequest(&_Exchange.TransactOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, submissionExpirationTime, nonce, signature, premiumCost)
}

// ForcedWithdrawalRequest is a paid mutator transaction binding the contract method 0xaf1437a3.
//
// Solidity: function forcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount, bool premiumCost) returns()
func (_Exchange *ExchangeTransactor) ForcedWithdrawalRequest(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int, premiumCost bool) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "forcedWithdrawalRequest", starkKey, vaultId, quantizedAmount, premiumCost)
}

// ForcedWithdrawalRequest is a paid mutator transaction binding the contract method 0xaf1437a3.
//
// Solidity: function forcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount, bool premiumCost) returns()
func (_Exchange *ExchangeSession) ForcedWithdrawalRequest(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int, premiumCost bool) (*types.Transaction, error) {
	return _Exchange.Contract.ForcedWithdrawalRequest(&_Exchange.TransactOpts, starkKey, vaultId, quantizedAmount, premiumCost)
}

// ForcedWithdrawalRequest is a paid mutator transaction binding the contract method 0xaf1437a3.
//
// Solidity: function forcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount, bool premiumCost) returns()
func (_Exchange *ExchangeTransactorSession) ForcedWithdrawalRequest(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int, premiumCost bool) (*types.Transaction, error) {
	return _Exchange.Contract.ForcedWithdrawalRequest(&_Exchange.TransactOpts, starkKey, vaultId, quantizedAmount, premiumCost)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x3153a386.
//
// Solidity: function freezeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) returns()
func (_Exchange *ExchangeTransactor) FreezeRequest(opts *bind.TransactOpts, starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "freezeRequest", starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x3153a386.
//
// Solidity: function freezeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) returns()
func (_Exchange *ExchangeSession) FreezeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.FreezeRequest(&_Exchange.TransactOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// FreezeRequest is a paid mutator transaction binding the contract method 0x3153a386.
//
// Solidity: function freezeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce) returns()
func (_Exchange *ExchangeTransactorSession) FreezeRequest(starkKeyA *big.Int, starkKeyB *big.Int, vaultIdA *big.Int, vaultIdB *big.Int, collateralAssetId *big.Int, syntheticAssetId *big.Int, amountCollateral *big.Int, amountSynthetic *big.Int, aIsBuyingSynthetic bool, nonce *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.FreezeRequest(&_Exchange.TransactOpts, starkKeyA, starkKeyB, vaultIdA, vaultIdB, collateralAssetId, syntheticAssetId, amountCollateral, amountSynthetic, aIsBuyingSynthetic, nonce)
}

// FreezeRequest0 is a paid mutator transaction binding the contract method 0x75d4eefd.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactor) FreezeRequest0(opts *bind.TransactOpts, starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "freezeRequest0", starkKey, vaultId, quantizedAmount)
}

// FreezeRequest0 is a paid mutator transaction binding the contract method 0x75d4eefd.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeSession) FreezeRequest0(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.FreezeRequest0(&_Exchange.TransactOpts, starkKey, vaultId, quantizedAmount)
}

// FreezeRequest0 is a paid mutator transaction binding the contract method 0x75d4eefd.
//
// Solidity: function freezeRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactorSession) FreezeRequest0(starkKey *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.FreezeRequest0(&_Exchange.TransactOpts, starkKey, vaultId, quantizedAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_Exchange *ExchangeTransactor) Initialize(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "initialize", data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_Exchange *ExchangeSession) Initialize(data []byte) (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_Exchange *ExchangeTransactorSession) Initialize(data []byte) (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts, data)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_Exchange *ExchangeTransactor) MainAcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "mainAcceptGovernance")
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_Exchange *ExchangeSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _Exchange.Contract.MainAcceptGovernance(&_Exchange.TransactOpts)
}

// MainAcceptGovernance is a paid mutator transaction binding the contract method 0x28700a15.
//
// Solidity: function mainAcceptGovernance() returns()
func (_Exchange *ExchangeTransactorSession) MainAcceptGovernance() (*types.Transaction, error) {
	return _Exchange.Contract.MainAcceptGovernance(&_Exchange.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_Exchange *ExchangeTransactor) MainCancelNomination(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "mainCancelNomination")
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_Exchange *ExchangeSession) MainCancelNomination() (*types.Transaction, error) {
	return _Exchange.Contract.MainCancelNomination(&_Exchange.TransactOpts)
}

// MainCancelNomination is a paid mutator transaction binding the contract method 0x72eb3688.
//
// Solidity: function mainCancelNomination() returns()
func (_Exchange *ExchangeTransactorSession) MainCancelNomination() (*types.Transaction, error) {
	return _Exchange.Contract.MainCancelNomination(&_Exchange.TransactOpts)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_Exchange *ExchangeTransactor) MainNominateNewGovernor(opts *bind.TransactOpts, newGovernor common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "mainNominateNewGovernor", newGovernor)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_Exchange *ExchangeSession) MainNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.MainNominateNewGovernor(&_Exchange.TransactOpts, newGovernor)
}

// MainNominateNewGovernor is a paid mutator transaction binding the contract method 0x8c4bce1c.
//
// Solidity: function mainNominateNewGovernor(address newGovernor) returns()
func (_Exchange *ExchangeTransactorSession) MainNominateNewGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.MainNominateNewGovernor(&_Exchange.TransactOpts, newGovernor)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_Exchange *ExchangeTransactor) MainRemoveGovernor(opts *bind.TransactOpts, governorForRemoval common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "mainRemoveGovernor", governorForRemoval)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_Exchange *ExchangeSession) MainRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.MainRemoveGovernor(&_Exchange.TransactOpts, governorForRemoval)
}

// MainRemoveGovernor is a paid mutator transaction binding the contract method 0xa1cc921e.
//
// Solidity: function mainRemoveGovernor(address governorForRemoval) returns()
func (_Exchange *ExchangeTransactorSession) MainRemoveGovernor(governorForRemoval common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.MainRemoveGovernor(&_Exchange.TransactOpts, governorForRemoval)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Exchange *ExchangeTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Exchange *ExchangeSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Exchange.Contract.OnERC721Received(&_Exchange.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Exchange *ExchangeTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Exchange.Contract.OnERC721Received(&_Exchange.TransactOpts, arg0, arg1, arg2, arg3)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactor) RegisterAndDepositERC20(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerAndDepositERC20", ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeSession) RegisterAndDepositERC20(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterAndDepositERC20(&_Exchange.TransactOpts, ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositERC20 is a paid mutator transaction binding the contract method 0xd5280589.
//
// Solidity: function registerAndDepositERC20(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId, uint256 quantizedAmount) returns()
func (_Exchange *ExchangeTransactorSession) RegisterAndDepositERC20(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterAndDepositERC20(&_Exchange.TransactOpts, ethKey, starkKey, signature, assetType, vaultId, quantizedAmount)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeTransactor) RegisterAndDepositEth(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerAndDepositEth", ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeSession) RegisterAndDepositEth(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterAndDepositEth(&_Exchange.TransactOpts, ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAndDepositEth is a paid mutator transaction binding the contract method 0x3ccfc8ed.
//
// Solidity: function registerAndDepositEth(address ethKey, uint256 starkKey, bytes signature, uint256 assetType, uint256 vaultId) payable returns()
func (_Exchange *ExchangeTransactorSession) RegisterAndDepositEth(ethKey common.Address, starkKey *big.Int, signature []byte, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterAndDepositEth(&_Exchange.TransactOpts, ethKey, starkKey, signature, assetType, vaultId)
}

// RegisterAssetConfigurationChange is a paid mutator transaction binding the contract method 0xad548467.
//
// Solidity: function registerAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeTransactor) RegisterAssetConfigurationChange(opts *bind.TransactOpts, assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerAssetConfigurationChange", assetId, configHash)
}

// RegisterAssetConfigurationChange is a paid mutator transaction binding the contract method 0xad548467.
//
// Solidity: function registerAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeSession) RegisterAssetConfigurationChange(assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterAssetConfigurationChange(&_Exchange.TransactOpts, assetId, configHash)
}

// RegisterAssetConfigurationChange is a paid mutator transaction binding the contract method 0xad548467.
//
// Solidity: function registerAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeTransactorSession) RegisterAssetConfigurationChange(assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterAssetConfigurationChange(&_Exchange.TransactOpts, assetId, configHash)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_Exchange *ExchangeTransactor) RegisterAvailabilityVerifier(opts *bind.TransactOpts, verifier common.Address, identifier string) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerAvailabilityVerifier", verifier, identifier)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_Exchange *ExchangeSession) RegisterAvailabilityVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterAvailabilityVerifier(&_Exchange.TransactOpts, verifier, identifier)
}

// RegisterAvailabilityVerifier is a paid mutator transaction binding the contract method 0xbdb75785.
//
// Solidity: function registerAvailabilityVerifier(address verifier, string identifier) returns()
func (_Exchange *ExchangeTransactorSession) RegisterAvailabilityVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterAvailabilityVerifier(&_Exchange.TransactOpts, verifier, identifier)
}

// RegisterGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x9f8a5044.
//
// Solidity: function registerGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeTransactor) RegisterGlobalConfigurationChange(opts *bind.TransactOpts, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerGlobalConfigurationChange", configHash)
}

// RegisterGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x9f8a5044.
//
// Solidity: function registerGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeSession) RegisterGlobalConfigurationChange(configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterGlobalConfigurationChange(&_Exchange.TransactOpts, configHash)
}

// RegisterGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x9f8a5044.
//
// Solidity: function registerGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeTransactorSession) RegisterGlobalConfigurationChange(configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterGlobalConfigurationChange(&_Exchange.TransactOpts, configHash)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_Exchange *ExchangeTransactor) RegisterOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerOperator", newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_Exchange *ExchangeSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterOperator(&_Exchange.TransactOpts, newOperator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address newOperator) returns()
func (_Exchange *ExchangeTransactorSession) RegisterOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterOperator(&_Exchange.TransactOpts, newOperator)
}

// RegisterSystemAssetType is a paid mutator transaction binding the contract method 0x7fbf9ba9.
//
// Solidity: function registerSystemAssetType(uint256 assetType, bytes assetInfo) returns()
func (_Exchange *ExchangeTransactor) RegisterSystemAssetType(opts *bind.TransactOpts, assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerSystemAssetType", assetType, assetInfo)
}

// RegisterSystemAssetType is a paid mutator transaction binding the contract method 0x7fbf9ba9.
//
// Solidity: function registerSystemAssetType(uint256 assetType, bytes assetInfo) returns()
func (_Exchange *ExchangeSession) RegisterSystemAssetType(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterSystemAssetType(&_Exchange.TransactOpts, assetType, assetInfo)
}

// RegisterSystemAssetType is a paid mutator transaction binding the contract method 0x7fbf9ba9.
//
// Solidity: function registerSystemAssetType(uint256 assetType, bytes assetInfo) returns()
func (_Exchange *ExchangeTransactorSession) RegisterSystemAssetType(assetType *big.Int, assetInfo []byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterSystemAssetType(&_Exchange.TransactOpts, assetType, assetInfo)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_Exchange *ExchangeTransactor) RegisterToken(opts *bind.TransactOpts, arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerToken", arg0, arg1)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_Exchange *ExchangeSession) RegisterToken(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterToken(&_Exchange.TransactOpts, arg0, arg1)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xc8b1031a.
//
// Solidity: function registerToken(uint256 , bytes ) returns()
func (_Exchange *ExchangeTransactorSession) RegisterToken(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterToken(&_Exchange.TransactOpts, arg0, arg1)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 , bytes , uint256 ) returns()
func (_Exchange *ExchangeTransactor) RegisterToken0(opts *bind.TransactOpts, arg0 *big.Int, arg1 []byte, arg2 *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerToken0", arg0, arg1, arg2)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 , bytes , uint256 ) returns()
func (_Exchange *ExchangeSession) RegisterToken0(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterToken0(&_Exchange.TransactOpts, arg0, arg1, arg2)
}

// RegisterToken0 is a paid mutator transaction binding the contract method 0xd88d8b38.
//
// Solidity: function registerToken(uint256 , bytes , uint256 ) returns()
func (_Exchange *ExchangeTransactorSession) RegisterToken0(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterToken0(&_Exchange.TransactOpts, arg0, arg1, arg2)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_Exchange *ExchangeTransactor) RegisterTokenAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerTokenAdmin", newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_Exchange *ExchangeSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterTokenAdmin(&_Exchange.TransactOpts, newAdmin)
}

// RegisterTokenAdmin is a paid mutator transaction binding the contract method 0x0b3a2d21.
//
// Solidity: function registerTokenAdmin(address newAdmin) returns()
func (_Exchange *ExchangeTransactorSession) RegisterTokenAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterTokenAdmin(&_Exchange.TransactOpts, newAdmin)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address ethKey, uint256 starkKey, bytes signature) returns()
func (_Exchange *ExchangeTransactor) RegisterUser(opts *bind.TransactOpts, ethKey common.Address, starkKey *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerUser", ethKey, starkKey, signature)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address ethKey, uint256 starkKey, bytes signature) returns()
func (_Exchange *ExchangeSession) RegisterUser(ethKey common.Address, starkKey *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterUser(&_Exchange.TransactOpts, ethKey, starkKey, signature)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xdd2414d4.
//
// Solidity: function registerUser(address ethKey, uint256 starkKey, bytes signature) returns()
func (_Exchange *ExchangeTransactorSession) RegisterUser(ethKey common.Address, starkKey *big.Int, signature []byte) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterUser(&_Exchange.TransactOpts, ethKey, starkKey, signature)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address newAdmin) returns()
func (_Exchange *ExchangeTransactor) RegisterUserAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerUserAdmin", newAdmin)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address newAdmin) returns()
func (_Exchange *ExchangeSession) RegisterUserAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterUserAdmin(&_Exchange.TransactOpts, newAdmin)
}

// RegisterUserAdmin is a paid mutator transaction binding the contract method 0xf3e0c3fb.
//
// Solidity: function registerUserAdmin(address newAdmin) returns()
func (_Exchange *ExchangeTransactorSession) RegisterUserAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterUserAdmin(&_Exchange.TransactOpts, newAdmin)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_Exchange *ExchangeTransactor) RegisterVerifier(opts *bind.TransactOpts, verifier common.Address, identifier string) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "registerVerifier", verifier, identifier)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_Exchange *ExchangeSession) RegisterVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterVerifier(&_Exchange.TransactOpts, verifier, identifier)
}

// RegisterVerifier is a paid mutator transaction binding the contract method 0x3776fe2a.
//
// Solidity: function registerVerifier(address verifier, string identifier) returns()
func (_Exchange *ExchangeTransactorSession) RegisterVerifier(verifier common.Address, identifier string) (*types.Transaction, error) {
	return _Exchange.Contract.RegisterVerifier(&_Exchange.TransactOpts, verifier, identifier)
}

// RemoveAssetConfigurationChange is a paid mutator transaction binding the contract method 0xfc86484a.
//
// Solidity: function removeAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeTransactor) RemoveAssetConfigurationChange(opts *bind.TransactOpts, assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeAssetConfigurationChange", assetId, configHash)
}

// RemoveAssetConfigurationChange is a paid mutator transaction binding the contract method 0xfc86484a.
//
// Solidity: function removeAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeSession) RemoveAssetConfigurationChange(assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAssetConfigurationChange(&_Exchange.TransactOpts, assetId, configHash)
}

// RemoveAssetConfigurationChange is a paid mutator transaction binding the contract method 0xfc86484a.
//
// Solidity: function removeAssetConfigurationChange(uint256 assetId, bytes32 configHash) returns()
func (_Exchange *ExchangeTransactorSession) RemoveAssetConfigurationChange(assetId *big.Int, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAssetConfigurationChange(&_Exchange.TransactOpts, assetId, configHash)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_Exchange *ExchangeTransactor) RemoveAvailabilityVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeAvailabilityVerifier", verifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_Exchange *ExchangeSession) RemoveAvailabilityVerifier(verifier common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAvailabilityVerifier(&_Exchange.TransactOpts, verifier)
}

// RemoveAvailabilityVerifier is a paid mutator transaction binding the contract method 0xb1e640bf.
//
// Solidity: function removeAvailabilityVerifier(address verifier) returns()
func (_Exchange *ExchangeTransactorSession) RemoveAvailabilityVerifier(verifier common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveAvailabilityVerifier(&_Exchange.TransactOpts, verifier)
}

// RemoveGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x18290577.
//
// Solidity: function removeGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeTransactor) RemoveGlobalConfigurationChange(opts *bind.TransactOpts, configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeGlobalConfigurationChange", configHash)
}

// RemoveGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x18290577.
//
// Solidity: function removeGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeSession) RemoveGlobalConfigurationChange(configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveGlobalConfigurationChange(&_Exchange.TransactOpts, configHash)
}

// RemoveGlobalConfigurationChange is a paid mutator transaction binding the contract method 0x18290577.
//
// Solidity: function removeGlobalConfigurationChange(bytes32 configHash) returns()
func (_Exchange *ExchangeTransactorSession) RemoveGlobalConfigurationChange(configHash [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveGlobalConfigurationChange(&_Exchange.TransactOpts, configHash)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_Exchange *ExchangeTransactor) RemoveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeVerifier", verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_Exchange *ExchangeSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveVerifier(&_Exchange.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_Exchange *ExchangeTransactorSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveVerifier(&_Exchange.TransactOpts, verifier)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_Exchange *ExchangeTransactor) UnFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "unFreeze")
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_Exchange *ExchangeSession) UnFreeze() (*types.Transaction, error) {
	return _Exchange.Contract.UnFreeze(&_Exchange.TransactOpts)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_Exchange *ExchangeTransactorSession) UnFreeze() (*types.Transaction, error) {
	return _Exchange.Contract.UnFreeze(&_Exchange.TransactOpts)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_Exchange *ExchangeTransactor) UnregisterOperator(opts *bind.TransactOpts, removedOperator common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "unregisterOperator", removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_Exchange *ExchangeSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UnregisterOperator(&_Exchange.TransactOpts, removedOperator)
}

// UnregisterOperator is a paid mutator transaction binding the contract method 0x96115bc2.
//
// Solidity: function unregisterOperator(address removedOperator) returns()
func (_Exchange *ExchangeTransactorSession) UnregisterOperator(removedOperator common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UnregisterOperator(&_Exchange.TransactOpts, removedOperator)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_Exchange *ExchangeTransactor) UnregisterTokenAdmin(opts *bind.TransactOpts, oldAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "unregisterTokenAdmin", oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_Exchange *ExchangeSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UnregisterTokenAdmin(&_Exchange.TransactOpts, oldAdmin)
}

// UnregisterTokenAdmin is a paid mutator transaction binding the contract method 0xa6fa6e90.
//
// Solidity: function unregisterTokenAdmin(address oldAdmin) returns()
func (_Exchange *ExchangeTransactorSession) UnregisterTokenAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UnregisterTokenAdmin(&_Exchange.TransactOpts, oldAdmin)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address oldAdmin) returns()
func (_Exchange *ExchangeTransactor) UnregisterUserAdmin(opts *bind.TransactOpts, oldAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "unregisterUserAdmin", oldAdmin)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address oldAdmin) returns()
func (_Exchange *ExchangeSession) UnregisterUserAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UnregisterUserAdmin(&_Exchange.TransactOpts, oldAdmin)
}

// UnregisterUserAdmin is a paid mutator transaction binding the contract method 0xb04b6179.
//
// Solidity: function unregisterUserAdmin(address oldAdmin) returns()
func (_Exchange *ExchangeTransactorSession) UnregisterUserAdmin(oldAdmin common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UnregisterUserAdmin(&_Exchange.TransactOpts, oldAdmin)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] programOutput, uint256[] applicationData) returns()
func (_Exchange *ExchangeTransactor) UpdateState(opts *bind.TransactOpts, programOutput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateState", programOutput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] programOutput, uint256[] applicationData) returns()
func (_Exchange *ExchangeSession) UpdateState(programOutput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateState(&_Exchange.TransactOpts, programOutput, applicationData)
}

// UpdateState is a paid mutator transaction binding the contract method 0x538f9406.
//
// Solidity: function updateState(uint256[] programOutput, uint256[] applicationData) returns()
func (_Exchange *ExchangeTransactorSession) UpdateState(programOutput []*big.Int, applicationData []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateState(&_Exchange.TransactOpts, programOutput, applicationData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 starkKey, uint256 assetType) returns()
func (_Exchange *ExchangeTransactor) Withdraw(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdraw", starkKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 starkKey, uint256 assetType) returns()
func (_Exchange *ExchangeSession) Withdraw(starkKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Withdraw(&_Exchange.TransactOpts, starkKey, assetType)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 starkKey, uint256 assetType) returns()
func (_Exchange *ExchangeTransactorSession) Withdraw(starkKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Withdraw(&_Exchange.TransactOpts, starkKey, assetType)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 starkKey, uint256 assetType, bytes mintingBlob) returns()
func (_Exchange *ExchangeTransactor) WithdrawAndMint(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdrawAndMint", starkKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 starkKey, uint256 assetType, bytes mintingBlob) returns()
func (_Exchange *ExchangeSession) WithdrawAndMint(starkKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawAndMint(&_Exchange.TransactOpts, starkKey, assetType, mintingBlob)
}

// WithdrawAndMint is a paid mutator transaction binding the contract method 0xd91443b7.
//
// Solidity: function withdrawAndMint(uint256 starkKey, uint256 assetType, bytes mintingBlob) returns()
func (_Exchange *ExchangeTransactorSession) WithdrawAndMint(starkKey *big.Int, assetType *big.Int, mintingBlob []byte) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawAndMint(&_Exchange.TransactOpts, starkKey, assetType, mintingBlob)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 starkKey, uint256 assetType, uint256 tokenId) returns()
func (_Exchange *ExchangeTransactor) WithdrawNft(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdrawNft", starkKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 starkKey, uint256 assetType, uint256 tokenId) returns()
func (_Exchange *ExchangeSession) WithdrawNft(starkKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawNft(&_Exchange.TransactOpts, starkKey, assetType, tokenId)
}

// WithdrawNft is a paid mutator transaction binding the contract method 0x019b417a.
//
// Solidity: function withdrawNft(uint256 starkKey, uint256 assetType, uint256 tokenId) returns()
func (_Exchange *ExchangeTransactorSession) WithdrawNft(starkKey *big.Int, assetType *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawNft(&_Exchange.TransactOpts, starkKey, assetType, tokenId)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 starkKey, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_Exchange *ExchangeTransactor) WithdrawNftTo(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdrawNftTo", starkKey, assetType, tokenId, recipient)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 starkKey, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_Exchange *ExchangeSession) WithdrawNftTo(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawNftTo(&_Exchange.TransactOpts, starkKey, assetType, tokenId, recipient)
}

// WithdrawNftTo is a paid mutator transaction binding the contract method 0xebef0fd0.
//
// Solidity: function withdrawNftTo(uint256 starkKey, uint256 assetType, uint256 tokenId, address recipient) returns()
func (_Exchange *ExchangeTransactorSession) WithdrawNftTo(starkKey *big.Int, assetType *big.Int, tokenId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawNftTo(&_Exchange.TransactOpts, starkKey, assetType, tokenId, recipient)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 starkKey, uint256 assetType, address recipient) returns()
func (_Exchange *ExchangeTransactor) WithdrawTo(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdrawTo", starkKey, assetType, recipient)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 starkKey, uint256 assetType, address recipient) returns()
func (_Exchange *ExchangeSession) WithdrawTo(starkKey *big.Int, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawTo(&_Exchange.TransactOpts, starkKey, assetType, recipient)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 starkKey, uint256 assetType, address recipient) returns()
func (_Exchange *ExchangeTransactorSession) WithdrawTo(starkKey *big.Int, assetType *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawTo(&_Exchange.TransactOpts, starkKey, assetType, recipient)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Exchange *ExchangeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Exchange.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Exchange *ExchangeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Exchange.Contract.Fallback(&_Exchange.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Exchange *ExchangeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Exchange.Contract.Fallback(&_Exchange.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeSession) Receive() (*types.Transaction, error) {
	return _Exchange.Contract.Receive(&_Exchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Exchange *ExchangeTransactorSession) Receive() (*types.Transaction, error) {
	return _Exchange.Contract.Receive(&_Exchange.TransactOpts)
}

// ExchangeLogAssetConfigurationAppliedIterator is returned from FilterLogAssetConfigurationApplied and is used to iterate over the raw logs and unpacked data for LogAssetConfigurationApplied events raised by the Exchange contract.
type ExchangeLogAssetConfigurationAppliedIterator struct {
	Event *ExchangeLogAssetConfigurationApplied // Event containing the contract specifics and raw log

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
func (it *ExchangeLogAssetConfigurationAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogAssetConfigurationApplied)
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
		it.Event = new(ExchangeLogAssetConfigurationApplied)
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
func (it *ExchangeLogAssetConfigurationAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogAssetConfigurationAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogAssetConfigurationApplied represents a LogAssetConfigurationApplied event raised by the Exchange contract.
type ExchangeLogAssetConfigurationApplied struct {
	AssetId    *big.Int
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogAssetConfigurationApplied is a free log retrieval operation binding the contract event 0x617a24590fcf7d5a2650aa9c10572915cb4bf853b1ef135cc1918460946a7a2f.
//
// Solidity: event LogAssetConfigurationApplied(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) FilterLogAssetConfigurationApplied(opts *bind.FilterOpts) (*ExchangeLogAssetConfigurationAppliedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogAssetConfigurationApplied")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogAssetConfigurationAppliedIterator{contract: _Exchange.contract, event: "LogAssetConfigurationApplied", logs: logs, sub: sub}, nil
}

// WatchLogAssetConfigurationApplied is a free log subscription operation binding the contract event 0x617a24590fcf7d5a2650aa9c10572915cb4bf853b1ef135cc1918460946a7a2f.
//
// Solidity: event LogAssetConfigurationApplied(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) WatchLogAssetConfigurationApplied(opts *bind.WatchOpts, sink chan<- *ExchangeLogAssetConfigurationApplied) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogAssetConfigurationApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogAssetConfigurationApplied)
				if err := _Exchange.contract.UnpackLog(event, "LogAssetConfigurationApplied", log); err != nil {
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

// ParseLogAssetConfigurationApplied is a log parse operation binding the contract event 0x617a24590fcf7d5a2650aa9c10572915cb4bf853b1ef135cc1918460946a7a2f.
//
// Solidity: event LogAssetConfigurationApplied(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) ParseLogAssetConfigurationApplied(log types.Log) (*ExchangeLogAssetConfigurationApplied, error) {
	event := new(ExchangeLogAssetConfigurationApplied)
	if err := _Exchange.contract.UnpackLog(event, "LogAssetConfigurationApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogAssetConfigurationRegisteredIterator is returned from FilterLogAssetConfigurationRegistered and is used to iterate over the raw logs and unpacked data for LogAssetConfigurationRegistered events raised by the Exchange contract.
type ExchangeLogAssetConfigurationRegisteredIterator struct {
	Event *ExchangeLogAssetConfigurationRegistered // Event containing the contract specifics and raw log

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
func (it *ExchangeLogAssetConfigurationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogAssetConfigurationRegistered)
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
		it.Event = new(ExchangeLogAssetConfigurationRegistered)
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
func (it *ExchangeLogAssetConfigurationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogAssetConfigurationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogAssetConfigurationRegistered represents a LogAssetConfigurationRegistered event raised by the Exchange contract.
type ExchangeLogAssetConfigurationRegistered struct {
	AssetId    *big.Int
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogAssetConfigurationRegistered is a free log retrieval operation binding the contract event 0xfbe49fc520438bdf956c8964ba0394d2da4fa3f34686e826a2ded762fa687c2d.
//
// Solidity: event LogAssetConfigurationRegistered(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) FilterLogAssetConfigurationRegistered(opts *bind.FilterOpts) (*ExchangeLogAssetConfigurationRegisteredIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogAssetConfigurationRegistered")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogAssetConfigurationRegisteredIterator{contract: _Exchange.contract, event: "LogAssetConfigurationRegistered", logs: logs, sub: sub}, nil
}

// WatchLogAssetConfigurationRegistered is a free log subscription operation binding the contract event 0xfbe49fc520438bdf956c8964ba0394d2da4fa3f34686e826a2ded762fa687c2d.
//
// Solidity: event LogAssetConfigurationRegistered(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) WatchLogAssetConfigurationRegistered(opts *bind.WatchOpts, sink chan<- *ExchangeLogAssetConfigurationRegistered) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogAssetConfigurationRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogAssetConfigurationRegistered)
				if err := _Exchange.contract.UnpackLog(event, "LogAssetConfigurationRegistered", log); err != nil {
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

// ParseLogAssetConfigurationRegistered is a log parse operation binding the contract event 0xfbe49fc520438bdf956c8964ba0394d2da4fa3f34686e826a2ded762fa687c2d.
//
// Solidity: event LogAssetConfigurationRegistered(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) ParseLogAssetConfigurationRegistered(log types.Log) (*ExchangeLogAssetConfigurationRegistered, error) {
	event := new(ExchangeLogAssetConfigurationRegistered)
	if err := _Exchange.contract.UnpackLog(event, "LogAssetConfigurationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogAssetConfigurationRemovedIterator is returned from FilterLogAssetConfigurationRemoved and is used to iterate over the raw logs and unpacked data for LogAssetConfigurationRemoved events raised by the Exchange contract.
type ExchangeLogAssetConfigurationRemovedIterator struct {
	Event *ExchangeLogAssetConfigurationRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeLogAssetConfigurationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogAssetConfigurationRemoved)
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
		it.Event = new(ExchangeLogAssetConfigurationRemoved)
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
func (it *ExchangeLogAssetConfigurationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogAssetConfigurationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogAssetConfigurationRemoved represents a LogAssetConfigurationRemoved event raised by the Exchange contract.
type ExchangeLogAssetConfigurationRemoved struct {
	AssetId    *big.Int
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogAssetConfigurationRemoved is a free log retrieval operation binding the contract event 0xbedb78be82c75282d73923c8287b6b3e3ef57379cdc69f7fbd3d227a9969b48e.
//
// Solidity: event LogAssetConfigurationRemoved(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) FilterLogAssetConfigurationRemoved(opts *bind.FilterOpts) (*ExchangeLogAssetConfigurationRemovedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogAssetConfigurationRemoved")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogAssetConfigurationRemovedIterator{contract: _Exchange.contract, event: "LogAssetConfigurationRemoved", logs: logs, sub: sub}, nil
}

// WatchLogAssetConfigurationRemoved is a free log subscription operation binding the contract event 0xbedb78be82c75282d73923c8287b6b3e3ef57379cdc69f7fbd3d227a9969b48e.
//
// Solidity: event LogAssetConfigurationRemoved(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) WatchLogAssetConfigurationRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeLogAssetConfigurationRemoved) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogAssetConfigurationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogAssetConfigurationRemoved)
				if err := _Exchange.contract.UnpackLog(event, "LogAssetConfigurationRemoved", log); err != nil {
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

// ParseLogAssetConfigurationRemoved is a log parse operation binding the contract event 0xbedb78be82c75282d73923c8287b6b3e3ef57379cdc69f7fbd3d227a9969b48e.
//
// Solidity: event LogAssetConfigurationRemoved(uint256 assetId, bytes32 configHash)
func (_Exchange *ExchangeFilterer) ParseLogAssetConfigurationRemoved(log types.Log) (*ExchangeLogAssetConfigurationRemoved, error) {
	event := new(ExchangeLogAssetConfigurationRemoved)
	if err := _Exchange.contract.UnpackLog(event, "LogAssetConfigurationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the Exchange contract.
type ExchangeLogDepositIterator struct {
	Event *ExchangeLogDeposit // Event containing the contract specifics and raw log

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
func (it *ExchangeLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogDeposit)
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
		it.Event = new(ExchangeLogDeposit)
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
func (it *ExchangeLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogDeposit represents a LogDeposit event raised by the Exchange contract.
type ExchangeLogDeposit struct {
	DepositorEthKey    common.Address
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDeposit is a free log retrieval operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) FilterLogDeposit(opts *bind.FilterOpts) (*ExchangeLogDepositIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogDepositIterator{contract: _Exchange.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *ExchangeLogDeposit) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogDeposit)
				if err := _Exchange.contract.UnpackLog(event, "LogDeposit", log); err != nil {
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

// ParseLogDeposit is a log parse operation binding the contract event 0x06724742ccc8c330a39a641ef02a0b419bd09248360680bb38159b0a8c2635d6.
//
// Solidity: event LogDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) ParseLogDeposit(log types.Log) (*ExchangeLogDeposit, error) {
	event := new(ExchangeLogDeposit)
	if err := _Exchange.contract.UnpackLog(event, "LogDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogDepositCancelIterator is returned from FilterLogDepositCancel and is used to iterate over the raw logs and unpacked data for LogDepositCancel events raised by the Exchange contract.
type ExchangeLogDepositCancelIterator struct {
	Event *ExchangeLogDepositCancel // Event containing the contract specifics and raw log

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
func (it *ExchangeLogDepositCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogDepositCancel)
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
		it.Event = new(ExchangeLogDepositCancel)
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
func (it *ExchangeLogDepositCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogDepositCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogDepositCancel represents a LogDepositCancel event raised by the Exchange contract.
type ExchangeLogDepositCancel struct {
	StarkKey *big.Int
	VaultId  *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancel is a free log retrieval operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Exchange *ExchangeFilterer) FilterLogDepositCancel(opts *bind.FilterOpts) (*ExchangeLogDepositCancelIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogDepositCancelIterator{contract: _Exchange.contract, event: "LogDepositCancel", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancel is a free log subscription operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Exchange *ExchangeFilterer) WatchLogDepositCancel(opts *bind.WatchOpts, sink chan<- *ExchangeLogDepositCancel) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogDepositCancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogDepositCancel)
				if err := _Exchange.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
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

// ParseLogDepositCancel is a log parse operation binding the contract event 0x0bc1df35228095c37da66a6ffcc755ea79dfc437345685f618e05fafad6b445e.
//
// Solidity: event LogDepositCancel(uint256 starkKey, uint256 vaultId, uint256 assetId)
func (_Exchange *ExchangeFilterer) ParseLogDepositCancel(log types.Log) (*ExchangeLogDepositCancel, error) {
	event := new(ExchangeLogDepositCancel)
	if err := _Exchange.contract.UnpackLog(event, "LogDepositCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogDepositCancelReclaimedIterator is returned from FilterLogDepositCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositCancelReclaimed events raised by the Exchange contract.
type ExchangeLogDepositCancelReclaimedIterator struct {
	Event *ExchangeLogDepositCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *ExchangeLogDepositCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogDepositCancelReclaimed)
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
		it.Event = new(ExchangeLogDepositCancelReclaimed)
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
func (it *ExchangeLogDepositCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogDepositCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogDepositCancelReclaimed represents a LogDepositCancelReclaimed event raised by the Exchange contract.
type ExchangeLogDepositCancelReclaimed struct {
	StarkKey           *big.Int
	VaultId            *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogDepositCancelReclaimed is a free log retrieval operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) FilterLogDepositCancelReclaimed(opts *bind.FilterOpts) (*ExchangeLogDepositCancelReclaimedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogDepositCancelReclaimedIterator{contract: _Exchange.contract, event: "LogDepositCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositCancelReclaimed is a free log subscription operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) WatchLogDepositCancelReclaimed(opts *bind.WatchOpts, sink chan<- *ExchangeLogDepositCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogDepositCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogDepositCancelReclaimed)
				if err := _Exchange.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
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

// ParseLogDepositCancelReclaimed is a log parse operation binding the contract event 0xe3e46ecf1138180bf93cba62a0b7e661d976a8ab3d40243f7b082667d8f500af.
//
// Solidity: event LogDepositCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) ParseLogDepositCancelReclaimed(log types.Log) (*ExchangeLogDepositCancelReclaimed, error) {
	event := new(ExchangeLogDepositCancelReclaimed)
	if err := _Exchange.contract.UnpackLog(event, "LogDepositCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogDepositNftCancelReclaimedIterator is returned from FilterLogDepositNftCancelReclaimed and is used to iterate over the raw logs and unpacked data for LogDepositNftCancelReclaimed events raised by the Exchange contract.
type ExchangeLogDepositNftCancelReclaimedIterator struct {
	Event *ExchangeLogDepositNftCancelReclaimed // Event containing the contract specifics and raw log

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
func (it *ExchangeLogDepositNftCancelReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogDepositNftCancelReclaimed)
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
		it.Event = new(ExchangeLogDepositNftCancelReclaimed)
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
func (it *ExchangeLogDepositNftCancelReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogDepositNftCancelReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogDepositNftCancelReclaimed represents a LogDepositNftCancelReclaimed event raised by the Exchange contract.
type ExchangeLogDepositNftCancelReclaimed struct {
	StarkKey  *big.Int
	VaultId   *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogDepositNftCancelReclaimed is a free log retrieval operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Exchange *ExchangeFilterer) FilterLogDepositNftCancelReclaimed(opts *bind.FilterOpts) (*ExchangeLogDepositNftCancelReclaimedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogDepositNftCancelReclaimedIterator{contract: _Exchange.contract, event: "LogDepositNftCancelReclaimed", logs: logs, sub: sub}, nil
}

// WatchLogDepositNftCancelReclaimed is a free log subscription operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Exchange *ExchangeFilterer) WatchLogDepositNftCancelReclaimed(opts *bind.WatchOpts, sink chan<- *ExchangeLogDepositNftCancelReclaimed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogDepositNftCancelReclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogDepositNftCancelReclaimed)
				if err := _Exchange.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
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

// ParseLogDepositNftCancelReclaimed is a log parse operation binding the contract event 0xf00c0c1a754f6df7545d96a7e12aad552728b94ca6aa94f81e297bdbcf1dab9c.
//
// Solidity: event LogDepositNftCancelReclaimed(uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Exchange *ExchangeFilterer) ParseLogDepositNftCancelReclaimed(log types.Log) (*ExchangeLogDepositNftCancelReclaimed, error) {
	event := new(ExchangeLogDepositNftCancelReclaimed)
	if err := _Exchange.contract.UnpackLog(event, "LogDepositNftCancelReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogForcedTradeRequestIterator is returned from FilterLogForcedTradeRequest and is used to iterate over the raw logs and unpacked data for LogForcedTradeRequest events raised by the Exchange contract.
type ExchangeLogForcedTradeRequestIterator struct {
	Event *ExchangeLogForcedTradeRequest // Event containing the contract specifics and raw log

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
func (it *ExchangeLogForcedTradeRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogForcedTradeRequest)
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
		it.Event = new(ExchangeLogForcedTradeRequest)
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
func (it *ExchangeLogForcedTradeRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogForcedTradeRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogForcedTradeRequest represents a LogForcedTradeRequest event raised by the Exchange contract.
type ExchangeLogForcedTradeRequest struct {
	StarkKeyA          *big.Int
	StarkKeyB          *big.Int
	VaultIdA           *big.Int
	VaultIdB           *big.Int
	CollateralAssetId  *big.Int
	SyntheticAssetId   *big.Int
	AmountCollateral   *big.Int
	AmountSynthetic    *big.Int
	AIsBuyingSynthetic bool
	Nonce              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogForcedTradeRequest is a free log retrieval operation binding the contract event 0x79acb9227c98b68d7628d2c99a7719926eff77e8c2275f6ffe7cf255b32772be.
//
// Solidity: event LogForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce)
func (_Exchange *ExchangeFilterer) FilterLogForcedTradeRequest(opts *bind.FilterOpts) (*ExchangeLogForcedTradeRequestIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogForcedTradeRequest")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogForcedTradeRequestIterator{contract: _Exchange.contract, event: "LogForcedTradeRequest", logs: logs, sub: sub}, nil
}

// WatchLogForcedTradeRequest is a free log subscription operation binding the contract event 0x79acb9227c98b68d7628d2c99a7719926eff77e8c2275f6ffe7cf255b32772be.
//
// Solidity: event LogForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce)
func (_Exchange *ExchangeFilterer) WatchLogForcedTradeRequest(opts *bind.WatchOpts, sink chan<- *ExchangeLogForcedTradeRequest) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogForcedTradeRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogForcedTradeRequest)
				if err := _Exchange.contract.UnpackLog(event, "LogForcedTradeRequest", log); err != nil {
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

// ParseLogForcedTradeRequest is a log parse operation binding the contract event 0x79acb9227c98b68d7628d2c99a7719926eff77e8c2275f6ffe7cf255b32772be.
//
// Solidity: event LogForcedTradeRequest(uint256 starkKeyA, uint256 starkKeyB, uint256 vaultIdA, uint256 vaultIdB, uint256 collateralAssetId, uint256 syntheticAssetId, uint256 amountCollateral, uint256 amountSynthetic, bool aIsBuyingSynthetic, uint256 nonce)
func (_Exchange *ExchangeFilterer) ParseLogForcedTradeRequest(log types.Log) (*ExchangeLogForcedTradeRequest, error) {
	event := new(ExchangeLogForcedTradeRequest)
	if err := _Exchange.contract.UnpackLog(event, "LogForcedTradeRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogForcedWithdrawalRequestIterator is returned from FilterLogForcedWithdrawalRequest and is used to iterate over the raw logs and unpacked data for LogForcedWithdrawalRequest events raised by the Exchange contract.
type ExchangeLogForcedWithdrawalRequestIterator struct {
	Event *ExchangeLogForcedWithdrawalRequest // Event containing the contract specifics and raw log

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
func (it *ExchangeLogForcedWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogForcedWithdrawalRequest)
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
		it.Event = new(ExchangeLogForcedWithdrawalRequest)
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
func (it *ExchangeLogForcedWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogForcedWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogForcedWithdrawalRequest represents a LogForcedWithdrawalRequest event raised by the Exchange contract.
type ExchangeLogForcedWithdrawalRequest struct {
	StarkKey        *big.Int
	VaultId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogForcedWithdrawalRequest is a free log retrieval operation binding the contract event 0x5f3a3d7d7c2b8cc20fa315d7504580e81d5f74ea46e1b3dc2985281df08e204d.
//
// Solidity: event LogForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) FilterLogForcedWithdrawalRequest(opts *bind.FilterOpts) (*ExchangeLogForcedWithdrawalRequestIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogForcedWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogForcedWithdrawalRequestIterator{contract: _Exchange.contract, event: "LogForcedWithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchLogForcedWithdrawalRequest is a free log subscription operation binding the contract event 0x5f3a3d7d7c2b8cc20fa315d7504580e81d5f74ea46e1b3dc2985281df08e204d.
//
// Solidity: event LogForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) WatchLogForcedWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *ExchangeLogForcedWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogForcedWithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogForcedWithdrawalRequest)
				if err := _Exchange.contract.UnpackLog(event, "LogForcedWithdrawalRequest", log); err != nil {
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

// ParseLogForcedWithdrawalRequest is a log parse operation binding the contract event 0x5f3a3d7d7c2b8cc20fa315d7504580e81d5f74ea46e1b3dc2985281df08e204d.
//
// Solidity: event LogForcedWithdrawalRequest(uint256 starkKey, uint256 vaultId, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) ParseLogForcedWithdrawalRequest(log types.Log) (*ExchangeLogForcedWithdrawalRequest, error) {
	event := new(ExchangeLogForcedWithdrawalRequest)
	if err := _Exchange.contract.UnpackLog(event, "LogForcedWithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogFrozenIterator is returned from FilterLogFrozen and is used to iterate over the raw logs and unpacked data for LogFrozen events raised by the Exchange contract.
type ExchangeLogFrozenIterator struct {
	Event *ExchangeLogFrozen // Event containing the contract specifics and raw log

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
func (it *ExchangeLogFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogFrozen)
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
		it.Event = new(ExchangeLogFrozen)
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
func (it *ExchangeLogFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogFrozen represents a LogFrozen event raised by the Exchange contract.
type ExchangeLogFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogFrozen is a free log retrieval operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_Exchange *ExchangeFilterer) FilterLogFrozen(opts *bind.FilterOpts) (*ExchangeLogFrozenIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogFrozen")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogFrozenIterator{contract: _Exchange.contract, event: "LogFrozen", logs: logs, sub: sub}, nil
}

// WatchLogFrozen is a free log subscription operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_Exchange *ExchangeFilterer) WatchLogFrozen(opts *bind.WatchOpts, sink chan<- *ExchangeLogFrozen) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogFrozen)
				if err := _Exchange.contract.UnpackLog(event, "LogFrozen", log); err != nil {
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

// ParseLogFrozen is a log parse operation binding the contract event 0xf5b8e6419478ab140eb98026ab5bd607038cb0ac4d4dad5b1fc0848dfd203d1f.
//
// Solidity: event LogFrozen()
func (_Exchange *ExchangeFilterer) ParseLogFrozen(log types.Log) (*ExchangeLogFrozen, error) {
	event := new(ExchangeLogFrozen)
	if err := _Exchange.contract.UnpackLog(event, "LogFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogGlobalConfigurationAppliedIterator is returned from FilterLogGlobalConfigurationApplied and is used to iterate over the raw logs and unpacked data for LogGlobalConfigurationApplied events raised by the Exchange contract.
type ExchangeLogGlobalConfigurationAppliedIterator struct {
	Event *ExchangeLogGlobalConfigurationApplied // Event containing the contract specifics and raw log

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
func (it *ExchangeLogGlobalConfigurationAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogGlobalConfigurationApplied)
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
		it.Event = new(ExchangeLogGlobalConfigurationApplied)
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
func (it *ExchangeLogGlobalConfigurationAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogGlobalConfigurationAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogGlobalConfigurationApplied represents a LogGlobalConfigurationApplied event raised by the Exchange contract.
type ExchangeLogGlobalConfigurationApplied struct {
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogGlobalConfigurationApplied is a free log retrieval operation binding the contract event 0x215e2cbb35cc7dd3f085fdb4ede08d32f4619001c73f5b10419389e7949c28ba.
//
// Solidity: event LogGlobalConfigurationApplied(bytes32 configHash)
func (_Exchange *ExchangeFilterer) FilterLogGlobalConfigurationApplied(opts *bind.FilterOpts) (*ExchangeLogGlobalConfigurationAppliedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogGlobalConfigurationApplied")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogGlobalConfigurationAppliedIterator{contract: _Exchange.contract, event: "LogGlobalConfigurationApplied", logs: logs, sub: sub}, nil
}

// WatchLogGlobalConfigurationApplied is a free log subscription operation binding the contract event 0x215e2cbb35cc7dd3f085fdb4ede08d32f4619001c73f5b10419389e7949c28ba.
//
// Solidity: event LogGlobalConfigurationApplied(bytes32 configHash)
func (_Exchange *ExchangeFilterer) WatchLogGlobalConfigurationApplied(opts *bind.WatchOpts, sink chan<- *ExchangeLogGlobalConfigurationApplied) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogGlobalConfigurationApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogGlobalConfigurationApplied)
				if err := _Exchange.contract.UnpackLog(event, "LogGlobalConfigurationApplied", log); err != nil {
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

// ParseLogGlobalConfigurationApplied is a log parse operation binding the contract event 0x215e2cbb35cc7dd3f085fdb4ede08d32f4619001c73f5b10419389e7949c28ba.
//
// Solidity: event LogGlobalConfigurationApplied(bytes32 configHash)
func (_Exchange *ExchangeFilterer) ParseLogGlobalConfigurationApplied(log types.Log) (*ExchangeLogGlobalConfigurationApplied, error) {
	event := new(ExchangeLogGlobalConfigurationApplied)
	if err := _Exchange.contract.UnpackLog(event, "LogGlobalConfigurationApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogGlobalConfigurationRegisteredIterator is returned from FilterLogGlobalConfigurationRegistered and is used to iterate over the raw logs and unpacked data for LogGlobalConfigurationRegistered events raised by the Exchange contract.
type ExchangeLogGlobalConfigurationRegisteredIterator struct {
	Event *ExchangeLogGlobalConfigurationRegistered // Event containing the contract specifics and raw log

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
func (it *ExchangeLogGlobalConfigurationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogGlobalConfigurationRegistered)
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
		it.Event = new(ExchangeLogGlobalConfigurationRegistered)
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
func (it *ExchangeLogGlobalConfigurationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogGlobalConfigurationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogGlobalConfigurationRegistered represents a LogGlobalConfigurationRegistered event raised by the Exchange contract.
type ExchangeLogGlobalConfigurationRegistered struct {
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogGlobalConfigurationRegistered is a free log retrieval operation binding the contract event 0xfae91148d75f3789f216345abe08b5415e0e727d51cab9f4e41efc23f71193f7.
//
// Solidity: event LogGlobalConfigurationRegistered(bytes32 configHash)
func (_Exchange *ExchangeFilterer) FilterLogGlobalConfigurationRegistered(opts *bind.FilterOpts) (*ExchangeLogGlobalConfigurationRegisteredIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogGlobalConfigurationRegistered")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogGlobalConfigurationRegisteredIterator{contract: _Exchange.contract, event: "LogGlobalConfigurationRegistered", logs: logs, sub: sub}, nil
}

// WatchLogGlobalConfigurationRegistered is a free log subscription operation binding the contract event 0xfae91148d75f3789f216345abe08b5415e0e727d51cab9f4e41efc23f71193f7.
//
// Solidity: event LogGlobalConfigurationRegistered(bytes32 configHash)
func (_Exchange *ExchangeFilterer) WatchLogGlobalConfigurationRegistered(opts *bind.WatchOpts, sink chan<- *ExchangeLogGlobalConfigurationRegistered) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogGlobalConfigurationRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogGlobalConfigurationRegistered)
				if err := _Exchange.contract.UnpackLog(event, "LogGlobalConfigurationRegistered", log); err != nil {
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

// ParseLogGlobalConfigurationRegistered is a log parse operation binding the contract event 0xfae91148d75f3789f216345abe08b5415e0e727d51cab9f4e41efc23f71193f7.
//
// Solidity: event LogGlobalConfigurationRegistered(bytes32 configHash)
func (_Exchange *ExchangeFilterer) ParseLogGlobalConfigurationRegistered(log types.Log) (*ExchangeLogGlobalConfigurationRegistered, error) {
	event := new(ExchangeLogGlobalConfigurationRegistered)
	if err := _Exchange.contract.UnpackLog(event, "LogGlobalConfigurationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogGlobalConfigurationRemovedIterator is returned from FilterLogGlobalConfigurationRemoved and is used to iterate over the raw logs and unpacked data for LogGlobalConfigurationRemoved events raised by the Exchange contract.
type ExchangeLogGlobalConfigurationRemovedIterator struct {
	Event *ExchangeLogGlobalConfigurationRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeLogGlobalConfigurationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogGlobalConfigurationRemoved)
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
		it.Event = new(ExchangeLogGlobalConfigurationRemoved)
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
func (it *ExchangeLogGlobalConfigurationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogGlobalConfigurationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogGlobalConfigurationRemoved represents a LogGlobalConfigurationRemoved event raised by the Exchange contract.
type ExchangeLogGlobalConfigurationRemoved struct {
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogGlobalConfigurationRemoved is a free log retrieval operation binding the contract event 0x21e13101b7a52038b3abfab6f17d3a3d091c620b978b98c9de7f3aa91e7ac063.
//
// Solidity: event LogGlobalConfigurationRemoved(bytes32 configHash)
func (_Exchange *ExchangeFilterer) FilterLogGlobalConfigurationRemoved(opts *bind.FilterOpts) (*ExchangeLogGlobalConfigurationRemovedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogGlobalConfigurationRemoved")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogGlobalConfigurationRemovedIterator{contract: _Exchange.contract, event: "LogGlobalConfigurationRemoved", logs: logs, sub: sub}, nil
}

// WatchLogGlobalConfigurationRemoved is a free log subscription operation binding the contract event 0x21e13101b7a52038b3abfab6f17d3a3d091c620b978b98c9de7f3aa91e7ac063.
//
// Solidity: event LogGlobalConfigurationRemoved(bytes32 configHash)
func (_Exchange *ExchangeFilterer) WatchLogGlobalConfigurationRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeLogGlobalConfigurationRemoved) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogGlobalConfigurationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogGlobalConfigurationRemoved)
				if err := _Exchange.contract.UnpackLog(event, "LogGlobalConfigurationRemoved", log); err != nil {
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

// ParseLogGlobalConfigurationRemoved is a log parse operation binding the contract event 0x21e13101b7a52038b3abfab6f17d3a3d091c620b978b98c9de7f3aa91e7ac063.
//
// Solidity: event LogGlobalConfigurationRemoved(bytes32 configHash)
func (_Exchange *ExchangeFilterer) ParseLogGlobalConfigurationRemoved(log types.Log) (*ExchangeLogGlobalConfigurationRemoved, error) {
	event := new(ExchangeLogGlobalConfigurationRemoved)
	if err := _Exchange.contract.UnpackLog(event, "LogGlobalConfigurationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogMintWithdrawalPerformedIterator is returned from FilterLogMintWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogMintWithdrawalPerformed events raised by the Exchange contract.
type ExchangeLogMintWithdrawalPerformedIterator struct {
	Event *ExchangeLogMintWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *ExchangeLogMintWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogMintWithdrawalPerformed)
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
		it.Event = new(ExchangeLogMintWithdrawalPerformed)
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
func (it *ExchangeLogMintWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogMintWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogMintWithdrawalPerformed represents a LogMintWithdrawalPerformed event raised by the Exchange contract.
type ExchangeLogMintWithdrawalPerformed struct {
	StarkKey           *big.Int
	TokenId            *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	AssetId            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogMintWithdrawalPerformed is a free log retrieval operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 starkKey, uint256 tokenId, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Exchange *ExchangeFilterer) FilterLogMintWithdrawalPerformed(opts *bind.FilterOpts) (*ExchangeLogMintWithdrawalPerformedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogMintWithdrawalPerformedIterator{contract: _Exchange.contract, event: "LogMintWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogMintWithdrawalPerformed is a free log subscription operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 starkKey, uint256 tokenId, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Exchange *ExchangeFilterer) WatchLogMintWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *ExchangeLogMintWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogMintWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogMintWithdrawalPerformed)
				if err := _Exchange.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
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

// ParseLogMintWithdrawalPerformed is a log parse operation binding the contract event 0x7e6e15df814c1a309a57686de672b2bedd128eacde35c5370c36d6840d4e9a92.
//
// Solidity: event LogMintWithdrawalPerformed(uint256 starkKey, uint256 tokenId, uint256 nonQuantizedAmount, uint256 quantizedAmount, uint256 assetId)
func (_Exchange *ExchangeFilterer) ParseLogMintWithdrawalPerformed(log types.Log) (*ExchangeLogMintWithdrawalPerformed, error) {
	event := new(ExchangeLogMintWithdrawalPerformed)
	if err := _Exchange.contract.UnpackLog(event, "LogMintWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogMintableWithdrawalAllowedIterator is returned from FilterLogMintableWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogMintableWithdrawalAllowed events raised by the Exchange contract.
type ExchangeLogMintableWithdrawalAllowedIterator struct {
	Event *ExchangeLogMintableWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *ExchangeLogMintableWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogMintableWithdrawalAllowed)
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
		it.Event = new(ExchangeLogMintableWithdrawalAllowed)
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
func (it *ExchangeLogMintableWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogMintableWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogMintableWithdrawalAllowed represents a LogMintableWithdrawalAllowed event raised by the Exchange contract.
type ExchangeLogMintableWithdrawalAllowed struct {
	StarkKey        *big.Int
	AssetId         *big.Int
	QuantizedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogMintableWithdrawalAllowed is a free log retrieval operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 starkKey, uint256 assetId, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) FilterLogMintableWithdrawalAllowed(opts *bind.FilterOpts) (*ExchangeLogMintableWithdrawalAllowedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogMintableWithdrawalAllowedIterator{contract: _Exchange.contract, event: "LogMintableWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogMintableWithdrawalAllowed is a free log subscription operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 starkKey, uint256 assetId, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) WatchLogMintableWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *ExchangeLogMintableWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogMintableWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogMintableWithdrawalAllowed)
				if err := _Exchange.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
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

// ParseLogMintableWithdrawalAllowed is a log parse operation binding the contract event 0x2acce0cedb29dc4927e6c891b57ef5bc8858eea4bf52787ea94873aebd4aeca0.
//
// Solidity: event LogMintableWithdrawalAllowed(uint256 starkKey, uint256 assetId, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) ParseLogMintableWithdrawalAllowed(log types.Log) (*ExchangeLogMintableWithdrawalAllowed, error) {
	event := new(ExchangeLogMintableWithdrawalAllowed)
	if err := _Exchange.contract.UnpackLog(event, "LogMintableWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogNewGovernorAcceptedIterator is returned from FilterLogNewGovernorAccepted and is used to iterate over the raw logs and unpacked data for LogNewGovernorAccepted events raised by the Exchange contract.
type ExchangeLogNewGovernorAcceptedIterator struct {
	Event *ExchangeLogNewGovernorAccepted // Event containing the contract specifics and raw log

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
func (it *ExchangeLogNewGovernorAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogNewGovernorAccepted)
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
		it.Event = new(ExchangeLogNewGovernorAccepted)
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
func (it *ExchangeLogNewGovernorAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogNewGovernorAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogNewGovernorAccepted represents a LogNewGovernorAccepted event raised by the Exchange contract.
type ExchangeLogNewGovernorAccepted struct {
	AcceptedGovernor common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogNewGovernorAccepted is a free log retrieval operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_Exchange *ExchangeFilterer) FilterLogNewGovernorAccepted(opts *bind.FilterOpts) (*ExchangeLogNewGovernorAcceptedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogNewGovernorAcceptedIterator{contract: _Exchange.contract, event: "LogNewGovernorAccepted", logs: logs, sub: sub}, nil
}

// WatchLogNewGovernorAccepted is a free log subscription operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_Exchange *ExchangeFilterer) WatchLogNewGovernorAccepted(opts *bind.WatchOpts, sink chan<- *ExchangeLogNewGovernorAccepted) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogNewGovernorAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogNewGovernorAccepted)
				if err := _Exchange.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
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

// ParseLogNewGovernorAccepted is a log parse operation binding the contract event 0xcfb473e6c03f9a29ddaf990e736fa3de5188a0bd85d684f5b6e164ebfbfff5d2.
//
// Solidity: event LogNewGovernorAccepted(address acceptedGovernor)
func (_Exchange *ExchangeFilterer) ParseLogNewGovernorAccepted(log types.Log) (*ExchangeLogNewGovernorAccepted, error) {
	event := new(ExchangeLogNewGovernorAccepted)
	if err := _Exchange.contract.UnpackLog(event, "LogNewGovernorAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogNftDepositIterator is returned from FilterLogNftDeposit and is used to iterate over the raw logs and unpacked data for LogNftDeposit events raised by the Exchange contract.
type ExchangeLogNftDepositIterator struct {
	Event *ExchangeLogNftDeposit // Event containing the contract specifics and raw log

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
func (it *ExchangeLogNftDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogNftDeposit)
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
		it.Event = new(ExchangeLogNftDeposit)
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
func (it *ExchangeLogNftDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogNftDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogNftDeposit represents a LogNftDeposit event raised by the Exchange contract.
type ExchangeLogNftDeposit struct {
	DepositorEthKey common.Address
	StarkKey        *big.Int
	VaultId         *big.Int
	AssetType       *big.Int
	TokenId         *big.Int
	AssetId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogNftDeposit is a free log retrieval operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Exchange *ExchangeFilterer) FilterLogNftDeposit(opts *bind.FilterOpts) (*ExchangeLogNftDepositIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogNftDepositIterator{contract: _Exchange.contract, event: "LogNftDeposit", logs: logs, sub: sub}, nil
}

// WatchLogNftDeposit is a free log subscription operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Exchange *ExchangeFilterer) WatchLogNftDeposit(opts *bind.WatchOpts, sink chan<- *ExchangeLogNftDeposit) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogNftDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogNftDeposit)
				if err := _Exchange.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
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

// ParseLogNftDeposit is a log parse operation binding the contract event 0x0fcf2162832b2d6033d4d34d2f45a28d9cfee523f1899945bbdd32529cfda67b.
//
// Solidity: event LogNftDeposit(address depositorEthKey, uint256 starkKey, uint256 vaultId, uint256 assetType, uint256 tokenId, uint256 assetId)
func (_Exchange *ExchangeFilterer) ParseLogNftDeposit(log types.Log) (*ExchangeLogNftDeposit, error) {
	event := new(ExchangeLogNftDeposit)
	if err := _Exchange.contract.UnpackLog(event, "LogNftDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogNftWithdrawalAllowedIterator is returned from FilterLogNftWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalAllowed events raised by the Exchange contract.
type ExchangeLogNftWithdrawalAllowedIterator struct {
	Event *ExchangeLogNftWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *ExchangeLogNftWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogNftWithdrawalAllowed)
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
		it.Event = new(ExchangeLogNftWithdrawalAllowed)
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
func (it *ExchangeLogNftWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogNftWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogNftWithdrawalAllowed represents a LogNftWithdrawalAllowed event raised by the Exchange contract.
type ExchangeLogNftWithdrawalAllowed struct {
	StarkKey *big.Int
	AssetId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalAllowed is a free log retrieval operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 starkKey, uint256 assetId)
func (_Exchange *ExchangeFilterer) FilterLogNftWithdrawalAllowed(opts *bind.FilterOpts) (*ExchangeLogNftWithdrawalAllowedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogNftWithdrawalAllowedIterator{contract: _Exchange.contract, event: "LogNftWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalAllowed is a free log subscription operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 starkKey, uint256 assetId)
func (_Exchange *ExchangeFilterer) WatchLogNftWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *ExchangeLogNftWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogNftWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogNftWithdrawalAllowed)
				if err := _Exchange.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
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

// ParseLogNftWithdrawalAllowed is a log parse operation binding the contract event 0xf07608f26256bce78d87220cfc0e7b1cc69b48e55104bfa591b2818161386627.
//
// Solidity: event LogNftWithdrawalAllowed(uint256 starkKey, uint256 assetId)
func (_Exchange *ExchangeFilterer) ParseLogNftWithdrawalAllowed(log types.Log) (*ExchangeLogNftWithdrawalAllowed, error) {
	event := new(ExchangeLogNftWithdrawalAllowed)
	if err := _Exchange.contract.UnpackLog(event, "LogNftWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogNftWithdrawalPerformedIterator is returned from FilterLogNftWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogNftWithdrawalPerformed events raised by the Exchange contract.
type ExchangeLogNftWithdrawalPerformedIterator struct {
	Event *ExchangeLogNftWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *ExchangeLogNftWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogNftWithdrawalPerformed)
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
		it.Event = new(ExchangeLogNftWithdrawalPerformed)
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
func (it *ExchangeLogNftWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogNftWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogNftWithdrawalPerformed represents a LogNftWithdrawalPerformed event raised by the Exchange contract.
type ExchangeLogNftWithdrawalPerformed struct {
	StarkKey  *big.Int
	AssetType *big.Int
	TokenId   *big.Int
	AssetId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogNftWithdrawalPerformed is a free log retrieval operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Exchange *ExchangeFilterer) FilterLogNftWithdrawalPerformed(opts *bind.FilterOpts) (*ExchangeLogNftWithdrawalPerformedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogNftWithdrawalPerformedIterator{contract: _Exchange.contract, event: "LogNftWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogNftWithdrawalPerformed is a free log subscription operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Exchange *ExchangeFilterer) WatchLogNftWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *ExchangeLogNftWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogNftWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogNftWithdrawalPerformed)
				if err := _Exchange.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
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

// ParseLogNftWithdrawalPerformed is a log parse operation binding the contract event 0xa5cfa8e2199ec5b8ca319288bcab72734207d30569756ee594a74b4df7abbf41.
//
// Solidity: event LogNftWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 tokenId, uint256 assetId, address recipient)
func (_Exchange *ExchangeFilterer) ParseLogNftWithdrawalPerformed(log types.Log) (*ExchangeLogNftWithdrawalPerformed, error) {
	event := new(ExchangeLogNftWithdrawalPerformed)
	if err := _Exchange.contract.UnpackLog(event, "LogNftWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogNominatedGovernorIterator is returned from FilterLogNominatedGovernor and is used to iterate over the raw logs and unpacked data for LogNominatedGovernor events raised by the Exchange contract.
type ExchangeLogNominatedGovernorIterator struct {
	Event *ExchangeLogNominatedGovernor // Event containing the contract specifics and raw log

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
func (it *ExchangeLogNominatedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogNominatedGovernor)
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
		it.Event = new(ExchangeLogNominatedGovernor)
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
func (it *ExchangeLogNominatedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogNominatedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogNominatedGovernor represents a LogNominatedGovernor event raised by the Exchange contract.
type ExchangeLogNominatedGovernor struct {
	NominatedGovernor common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogNominatedGovernor is a free log retrieval operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_Exchange *ExchangeFilterer) FilterLogNominatedGovernor(opts *bind.FilterOpts) (*ExchangeLogNominatedGovernorIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogNominatedGovernorIterator{contract: _Exchange.contract, event: "LogNominatedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogNominatedGovernor is a free log subscription operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_Exchange *ExchangeFilterer) WatchLogNominatedGovernor(opts *bind.WatchOpts, sink chan<- *ExchangeLogNominatedGovernor) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogNominatedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogNominatedGovernor)
				if err := _Exchange.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
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

// ParseLogNominatedGovernor is a log parse operation binding the contract event 0x6166272c8d3f5f579082f2827532732f97195007983bb5b83ac12c56700b01a6.
//
// Solidity: event LogNominatedGovernor(address nominatedGovernor)
func (_Exchange *ExchangeFilterer) ParseLogNominatedGovernor(log types.Log) (*ExchangeLogNominatedGovernor, error) {
	event := new(ExchangeLogNominatedGovernor)
	if err := _Exchange.contract.UnpackLog(event, "LogNominatedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogNominationCancelledIterator is returned from FilterLogNominationCancelled and is used to iterate over the raw logs and unpacked data for LogNominationCancelled events raised by the Exchange contract.
type ExchangeLogNominationCancelledIterator struct {
	Event *ExchangeLogNominationCancelled // Event containing the contract specifics and raw log

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
func (it *ExchangeLogNominationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogNominationCancelled)
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
		it.Event = new(ExchangeLogNominationCancelled)
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
func (it *ExchangeLogNominationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogNominationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogNominationCancelled represents a LogNominationCancelled event raised by the Exchange contract.
type ExchangeLogNominationCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNominationCancelled is a free log retrieval operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_Exchange *ExchangeFilterer) FilterLogNominationCancelled(opts *bind.FilterOpts) (*ExchangeLogNominationCancelledIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogNominationCancelledIterator{contract: _Exchange.contract, event: "LogNominationCancelled", logs: logs, sub: sub}, nil
}

// WatchLogNominationCancelled is a free log subscription operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_Exchange *ExchangeFilterer) WatchLogNominationCancelled(opts *bind.WatchOpts, sink chan<- *ExchangeLogNominationCancelled) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogNominationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogNominationCancelled)
				if err := _Exchange.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
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

// ParseLogNominationCancelled is a log parse operation binding the contract event 0x7a8dc7dd7fffb43c4807438fa62729225156941e641fd877938f4edade3429f5.
//
// Solidity: event LogNominationCancelled()
func (_Exchange *ExchangeFilterer) ParseLogNominationCancelled(log types.Log) (*ExchangeLogNominationCancelled, error) {
	event := new(ExchangeLogNominationCancelled)
	if err := _Exchange.contract.UnpackLog(event, "LogNominationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogOperatorAddedIterator is returned from FilterLogOperatorAdded and is used to iterate over the raw logs and unpacked data for LogOperatorAdded events raised by the Exchange contract.
type ExchangeLogOperatorAddedIterator struct {
	Event *ExchangeLogOperatorAdded // Event containing the contract specifics and raw log

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
func (it *ExchangeLogOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogOperatorAdded)
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
		it.Event = new(ExchangeLogOperatorAdded)
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
func (it *ExchangeLogOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogOperatorAdded represents a LogOperatorAdded event raised by the Exchange contract.
type ExchangeLogOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorAdded is a free log retrieval operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_Exchange *ExchangeFilterer) FilterLogOperatorAdded(opts *bind.FilterOpts) (*ExchangeLogOperatorAddedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogOperatorAddedIterator{contract: _Exchange.contract, event: "LogOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchLogOperatorAdded is a free log subscription operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_Exchange *ExchangeFilterer) WatchLogOperatorAdded(opts *bind.WatchOpts, sink chan<- *ExchangeLogOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogOperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogOperatorAdded)
				if err := _Exchange.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
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

// ParseLogOperatorAdded is a log parse operation binding the contract event 0x50a18c352ee1c02ffe058e15c2eb6e58be387c81e73cc1e17035286e54c19a57.
//
// Solidity: event LogOperatorAdded(address operator)
func (_Exchange *ExchangeFilterer) ParseLogOperatorAdded(log types.Log) (*ExchangeLogOperatorAdded, error) {
	event := new(ExchangeLogOperatorAdded)
	if err := _Exchange.contract.UnpackLog(event, "LogOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogOperatorRemovedIterator is returned from FilterLogOperatorRemoved and is used to iterate over the raw logs and unpacked data for LogOperatorRemoved events raised by the Exchange contract.
type ExchangeLogOperatorRemovedIterator struct {
	Event *ExchangeLogOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeLogOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogOperatorRemoved)
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
		it.Event = new(ExchangeLogOperatorRemoved)
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
func (it *ExchangeLogOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogOperatorRemoved represents a LogOperatorRemoved event raised by the Exchange contract.
type ExchangeLogOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogOperatorRemoved is a free log retrieval operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_Exchange *ExchangeFilterer) FilterLogOperatorRemoved(opts *bind.FilterOpts) (*ExchangeLogOperatorRemovedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogOperatorRemovedIterator{contract: _Exchange.contract, event: "LogOperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchLogOperatorRemoved is a free log subscription operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_Exchange *ExchangeFilterer) WatchLogOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeLogOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogOperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogOperatorRemoved)
				if err := _Exchange.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
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

// ParseLogOperatorRemoved is a log parse operation binding the contract event 0xec5f6c3a91a1efb1f9a308bb33c6e9e66bf9090fad0732f127dfdbf516d0625d.
//
// Solidity: event LogOperatorRemoved(address operator)
func (_Exchange *ExchangeFilterer) ParseLogOperatorRemoved(log types.Log) (*ExchangeLogOperatorRemoved, error) {
	event := new(ExchangeLogOperatorRemoved)
	if err := _Exchange.contract.UnpackLog(event, "LogOperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogRegisteredIterator is returned from FilterLogRegistered and is used to iterate over the raw logs and unpacked data for LogRegistered events raised by the Exchange contract.
type ExchangeLogRegisteredIterator struct {
	Event *ExchangeLogRegistered // Event containing the contract specifics and raw log

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
func (it *ExchangeLogRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogRegistered)
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
		it.Event = new(ExchangeLogRegistered)
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
func (it *ExchangeLogRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogRegistered represents a LogRegistered event raised by the Exchange contract.
type ExchangeLogRegistered struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRegistered is a free log retrieval operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_Exchange *ExchangeFilterer) FilterLogRegistered(opts *bind.FilterOpts) (*ExchangeLogRegisteredIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogRegistered")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogRegisteredIterator{contract: _Exchange.contract, event: "LogRegistered", logs: logs, sub: sub}, nil
}

// WatchLogRegistered is a free log subscription operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_Exchange *ExchangeFilterer) WatchLogRegistered(opts *bind.WatchOpts, sink chan<- *ExchangeLogRegistered) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogRegistered)
				if err := _Exchange.contract.UnpackLog(event, "LogRegistered", log); err != nil {
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

// ParseLogRegistered is a log parse operation binding the contract event 0xaa7f29c97c6763919ef56006f07fbf5c1ac734b0414665df43cecdbce9010c9b.
//
// Solidity: event LogRegistered(address entry, string entryId)
func (_Exchange *ExchangeFilterer) ParseLogRegistered(log types.Log) (*ExchangeLogRegistered, error) {
	event := new(ExchangeLogRegistered)
	if err := _Exchange.contract.UnpackLog(event, "LogRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogRemovalIntentIterator is returned from FilterLogRemovalIntent and is used to iterate over the raw logs and unpacked data for LogRemovalIntent events raised by the Exchange contract.
type ExchangeLogRemovalIntentIterator struct {
	Event *ExchangeLogRemovalIntent // Event containing the contract specifics and raw log

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
func (it *ExchangeLogRemovalIntentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogRemovalIntent)
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
		it.Event = new(ExchangeLogRemovalIntent)
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
func (it *ExchangeLogRemovalIntentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogRemovalIntentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogRemovalIntent represents a LogRemovalIntent event raised by the Exchange contract.
type ExchangeLogRemovalIntent struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRemovalIntent is a free log retrieval operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_Exchange *ExchangeFilterer) FilterLogRemovalIntent(opts *bind.FilterOpts) (*ExchangeLogRemovalIntentIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogRemovalIntent")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogRemovalIntentIterator{contract: _Exchange.contract, event: "LogRemovalIntent", logs: logs, sub: sub}, nil
}

// WatchLogRemovalIntent is a free log subscription operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_Exchange *ExchangeFilterer) WatchLogRemovalIntent(opts *bind.WatchOpts, sink chan<- *ExchangeLogRemovalIntent) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogRemovalIntent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogRemovalIntent)
				if err := _Exchange.contract.UnpackLog(event, "LogRemovalIntent", log); err != nil {
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

// ParseLogRemovalIntent is a log parse operation binding the contract event 0xa98ac1f696177f16ca482653307aa4a02b68cf03b14409a546de5adf946484fc.
//
// Solidity: event LogRemovalIntent(address entry, string entryId)
func (_Exchange *ExchangeFilterer) ParseLogRemovalIntent(log types.Log) (*ExchangeLogRemovalIntent, error) {
	event := new(ExchangeLogRemovalIntent)
	if err := _Exchange.contract.UnpackLog(event, "LogRemovalIntent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogRemovedIterator is returned from FilterLogRemoved and is used to iterate over the raw logs and unpacked data for LogRemoved events raised by the Exchange contract.
type ExchangeLogRemovedIterator struct {
	Event *ExchangeLogRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeLogRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogRemoved)
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
		it.Event = new(ExchangeLogRemoved)
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
func (it *ExchangeLogRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogRemoved represents a LogRemoved event raised by the Exchange contract.
type ExchangeLogRemoved struct {
	Entry   common.Address
	EntryId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogRemoved is a free log retrieval operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_Exchange *ExchangeFilterer) FilterLogRemoved(opts *bind.FilterOpts) (*ExchangeLogRemovedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogRemoved")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogRemovedIterator{contract: _Exchange.contract, event: "LogRemoved", logs: logs, sub: sub}, nil
}

// WatchLogRemoved is a free log subscription operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_Exchange *ExchangeFilterer) WatchLogRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeLogRemoved) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogRemoved)
				if err := _Exchange.contract.UnpackLog(event, "LogRemoved", log); err != nil {
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

// ParseLogRemoved is a log parse operation binding the contract event 0x35b176cf4f09df896aa55e10eec90bb4c4689ea1d076135a8de3138d829d0142.
//
// Solidity: event LogRemoved(address entry, string entryId)
func (_Exchange *ExchangeFilterer) ParseLogRemoved(log types.Log) (*ExchangeLogRemoved, error) {
	event := new(ExchangeLogRemoved)
	if err := _Exchange.contract.UnpackLog(event, "LogRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogRemovedGovernorIterator is returned from FilterLogRemovedGovernor and is used to iterate over the raw logs and unpacked data for LogRemovedGovernor events raised by the Exchange contract.
type ExchangeLogRemovedGovernorIterator struct {
	Event *ExchangeLogRemovedGovernor // Event containing the contract specifics and raw log

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
func (it *ExchangeLogRemovedGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogRemovedGovernor)
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
		it.Event = new(ExchangeLogRemovedGovernor)
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
func (it *ExchangeLogRemovedGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogRemovedGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogRemovedGovernor represents a LogRemovedGovernor event raised by the Exchange contract.
type ExchangeLogRemovedGovernor struct {
	RemovedGovernor common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogRemovedGovernor is a free log retrieval operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_Exchange *ExchangeFilterer) FilterLogRemovedGovernor(opts *bind.FilterOpts) (*ExchangeLogRemovedGovernorIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogRemovedGovernorIterator{contract: _Exchange.contract, event: "LogRemovedGovernor", logs: logs, sub: sub}, nil
}

// WatchLogRemovedGovernor is a free log subscription operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_Exchange *ExchangeFilterer) WatchLogRemovedGovernor(opts *bind.WatchOpts, sink chan<- *ExchangeLogRemovedGovernor) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogRemovedGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogRemovedGovernor)
				if err := _Exchange.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
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

// ParseLogRemovedGovernor is a log parse operation binding the contract event 0xd75f94825e770b8b512be8e74759e252ad00e102e38f50cce2f7c6f868a29599.
//
// Solidity: event LogRemovedGovernor(address removedGovernor)
func (_Exchange *ExchangeFilterer) ParseLogRemovedGovernor(log types.Log) (*ExchangeLogRemovedGovernor, error) {
	event := new(ExchangeLogRemovedGovernor)
	if err := _Exchange.contract.UnpackLog(event, "LogRemovedGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogStateTransitionFactIterator is returned from FilterLogStateTransitionFact and is used to iterate over the raw logs and unpacked data for LogStateTransitionFact events raised by the Exchange contract.
type ExchangeLogStateTransitionFactIterator struct {
	Event *ExchangeLogStateTransitionFact // Event containing the contract specifics and raw log

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
func (it *ExchangeLogStateTransitionFactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogStateTransitionFact)
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
		it.Event = new(ExchangeLogStateTransitionFact)
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
func (it *ExchangeLogStateTransitionFactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogStateTransitionFactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogStateTransitionFact represents a LogStateTransitionFact event raised by the Exchange contract.
type ExchangeLogStateTransitionFact struct {
	StateTransitionFact [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogStateTransitionFact is a free log retrieval operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_Exchange *ExchangeFilterer) FilterLogStateTransitionFact(opts *bind.FilterOpts) (*ExchangeLogStateTransitionFactIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogStateTransitionFactIterator{contract: _Exchange.contract, event: "LogStateTransitionFact", logs: logs, sub: sub}, nil
}

// WatchLogStateTransitionFact is a free log subscription operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_Exchange *ExchangeFilterer) WatchLogStateTransitionFact(opts *bind.WatchOpts, sink chan<- *ExchangeLogStateTransitionFact) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogStateTransitionFact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogStateTransitionFact)
				if err := _Exchange.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
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

// ParseLogStateTransitionFact is a log parse operation binding the contract event 0x9866f8ddfe70bb512b2f2b28b49d4017c43f7ba775f1a20c61c13eea8cdac111.
//
// Solidity: event LogStateTransitionFact(bytes32 stateTransitionFact)
func (_Exchange *ExchangeFilterer) ParseLogStateTransitionFact(log types.Log) (*ExchangeLogStateTransitionFact, error) {
	event := new(ExchangeLogStateTransitionFact)
	if err := _Exchange.contract.UnpackLog(event, "LogStateTransitionFact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogSystemAssetTypeIterator is returned from FilterLogSystemAssetType and is used to iterate over the raw logs and unpacked data for LogSystemAssetType events raised by the Exchange contract.
type ExchangeLogSystemAssetTypeIterator struct {
	Event *ExchangeLogSystemAssetType // Event containing the contract specifics and raw log

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
func (it *ExchangeLogSystemAssetTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogSystemAssetType)
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
		it.Event = new(ExchangeLogSystemAssetType)
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
func (it *ExchangeLogSystemAssetTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogSystemAssetTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogSystemAssetType represents a LogSystemAssetType event raised by the Exchange contract.
type ExchangeLogSystemAssetType struct {
	AssetType *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSystemAssetType is a free log retrieval operation binding the contract event 0x51f48293a5ef1940e2b4eb2580372cf384aaa5bc739639e4624ce8d18c9644ab.
//
// Solidity: event LogSystemAssetType(uint256 assetType)
func (_Exchange *ExchangeFilterer) FilterLogSystemAssetType(opts *bind.FilterOpts) (*ExchangeLogSystemAssetTypeIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogSystemAssetType")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogSystemAssetTypeIterator{contract: _Exchange.contract, event: "LogSystemAssetType", logs: logs, sub: sub}, nil
}

// WatchLogSystemAssetType is a free log subscription operation binding the contract event 0x51f48293a5ef1940e2b4eb2580372cf384aaa5bc739639e4624ce8d18c9644ab.
//
// Solidity: event LogSystemAssetType(uint256 assetType)
func (_Exchange *ExchangeFilterer) WatchLogSystemAssetType(opts *bind.WatchOpts, sink chan<- *ExchangeLogSystemAssetType) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogSystemAssetType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogSystemAssetType)
				if err := _Exchange.contract.UnpackLog(event, "LogSystemAssetType", log); err != nil {
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

// ParseLogSystemAssetType is a log parse operation binding the contract event 0x51f48293a5ef1940e2b4eb2580372cf384aaa5bc739639e4624ce8d18c9644ab.
//
// Solidity: event LogSystemAssetType(uint256 assetType)
func (_Exchange *ExchangeFilterer) ParseLogSystemAssetType(log types.Log) (*ExchangeLogSystemAssetType, error) {
	event := new(ExchangeLogSystemAssetType)
	if err := _Exchange.contract.UnpackLog(event, "LogSystemAssetType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogTokenAdminAddedIterator is returned from FilterLogTokenAdminAdded and is used to iterate over the raw logs and unpacked data for LogTokenAdminAdded events raised by the Exchange contract.
type ExchangeLogTokenAdminAddedIterator struct {
	Event *ExchangeLogTokenAdminAdded // Event containing the contract specifics and raw log

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
func (it *ExchangeLogTokenAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogTokenAdminAdded)
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
		it.Event = new(ExchangeLogTokenAdminAdded)
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
func (it *ExchangeLogTokenAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogTokenAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogTokenAdminAdded represents a LogTokenAdminAdded event raised by the Exchange contract.
type ExchangeLogTokenAdminAdded struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminAdded is a free log retrieval operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_Exchange *ExchangeFilterer) FilterLogTokenAdminAdded(opts *bind.FilterOpts) (*ExchangeLogTokenAdminAddedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogTokenAdminAddedIterator{contract: _Exchange.contract, event: "LogTokenAdminAdded", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminAdded is a free log subscription operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_Exchange *ExchangeFilterer) WatchLogTokenAdminAdded(opts *bind.WatchOpts, sink chan<- *ExchangeLogTokenAdminAdded) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogTokenAdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogTokenAdminAdded)
				if err := _Exchange.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
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

// ParseLogTokenAdminAdded is a log parse operation binding the contract event 0x9085a9044aeb6daeeb5b4bf84af42b1a1613d4056f503c4e992b6396c16bd52f.
//
// Solidity: event LogTokenAdminAdded(address tokenAdmin)
func (_Exchange *ExchangeFilterer) ParseLogTokenAdminAdded(log types.Log) (*ExchangeLogTokenAdminAdded, error) {
	event := new(ExchangeLogTokenAdminAdded)
	if err := _Exchange.contract.UnpackLog(event, "LogTokenAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogTokenAdminRemovedIterator is returned from FilterLogTokenAdminRemoved and is used to iterate over the raw logs and unpacked data for LogTokenAdminRemoved events raised by the Exchange contract.
type ExchangeLogTokenAdminRemovedIterator struct {
	Event *ExchangeLogTokenAdminRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeLogTokenAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogTokenAdminRemoved)
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
		it.Event = new(ExchangeLogTokenAdminRemoved)
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
func (it *ExchangeLogTokenAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogTokenAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogTokenAdminRemoved represents a LogTokenAdminRemoved event raised by the Exchange contract.
type ExchangeLogTokenAdminRemoved struct {
	TokenAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdminRemoved is a free log retrieval operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_Exchange *ExchangeFilterer) FilterLogTokenAdminRemoved(opts *bind.FilterOpts) (*ExchangeLogTokenAdminRemovedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogTokenAdminRemovedIterator{contract: _Exchange.contract, event: "LogTokenAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdminRemoved is a free log subscription operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_Exchange *ExchangeFilterer) WatchLogTokenAdminRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeLogTokenAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogTokenAdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogTokenAdminRemoved)
				if err := _Exchange.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
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

// ParseLogTokenAdminRemoved is a log parse operation binding the contract event 0xfa49aecb996ea8d99950bb051552dfcc0b5460a0bb209867a1ed8067c32c2177.
//
// Solidity: event LogTokenAdminRemoved(address tokenAdmin)
func (_Exchange *ExchangeFilterer) ParseLogTokenAdminRemoved(log types.Log) (*ExchangeLogTokenAdminRemoved, error) {
	event := new(ExchangeLogTokenAdminRemoved)
	if err := _Exchange.contract.UnpackLog(event, "LogTokenAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogTokenRegisteredIterator is returned from FilterLogTokenRegistered and is used to iterate over the raw logs and unpacked data for LogTokenRegistered events raised by the Exchange contract.
type ExchangeLogTokenRegisteredIterator struct {
	Event *ExchangeLogTokenRegistered // Event containing the contract specifics and raw log

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
func (it *ExchangeLogTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogTokenRegistered)
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
		it.Event = new(ExchangeLogTokenRegistered)
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
func (it *ExchangeLogTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogTokenRegistered represents a LogTokenRegistered event raised by the Exchange contract.
type ExchangeLogTokenRegistered struct {
	AssetType *big.Int
	AssetInfo []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogTokenRegistered is a free log retrieval operation binding the contract event 0x4d2c7bfd8df1ba4f331f1abd2562bf3088e8b378c7dd1308113a82c64e518dbf.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo)
func (_Exchange *ExchangeFilterer) FilterLogTokenRegistered(opts *bind.FilterOpts) (*ExchangeLogTokenRegisteredIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogTokenRegisteredIterator{contract: _Exchange.contract, event: "LogTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchLogTokenRegistered is a free log subscription operation binding the contract event 0x4d2c7bfd8df1ba4f331f1abd2562bf3088e8b378c7dd1308113a82c64e518dbf.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo)
func (_Exchange *ExchangeFilterer) WatchLogTokenRegistered(opts *bind.WatchOpts, sink chan<- *ExchangeLogTokenRegistered) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogTokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogTokenRegistered)
				if err := _Exchange.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
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

// ParseLogTokenRegistered is a log parse operation binding the contract event 0x4d2c7bfd8df1ba4f331f1abd2562bf3088e8b378c7dd1308113a82c64e518dbf.
//
// Solidity: event LogTokenRegistered(uint256 assetType, bytes assetInfo)
func (_Exchange *ExchangeFilterer) ParseLogTokenRegistered(log types.Log) (*ExchangeLogTokenRegistered, error) {
	event := new(ExchangeLogTokenRegistered)
	if err := _Exchange.contract.UnpackLog(event, "LogTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogUnFrozenIterator is returned from FilterLogUnFrozen and is used to iterate over the raw logs and unpacked data for LogUnFrozen events raised by the Exchange contract.
type ExchangeLogUnFrozenIterator struct {
	Event *ExchangeLogUnFrozen // Event containing the contract specifics and raw log

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
func (it *ExchangeLogUnFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogUnFrozen)
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
		it.Event = new(ExchangeLogUnFrozen)
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
func (it *ExchangeLogUnFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogUnFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogUnFrozen represents a LogUnFrozen event raised by the Exchange contract.
type ExchangeLogUnFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogUnFrozen is a free log retrieval operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_Exchange *ExchangeFilterer) FilterLogUnFrozen(opts *bind.FilterOpts) (*ExchangeLogUnFrozenIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogUnFrozen")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogUnFrozenIterator{contract: _Exchange.contract, event: "LogUnFrozen", logs: logs, sub: sub}, nil
}

// WatchLogUnFrozen is a free log subscription operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_Exchange *ExchangeFilterer) WatchLogUnFrozen(opts *bind.WatchOpts, sink chan<- *ExchangeLogUnFrozen) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogUnFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogUnFrozen)
				if err := _Exchange.contract.UnpackLog(event, "LogUnFrozen", log); err != nil {
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

// ParseLogUnFrozen is a log parse operation binding the contract event 0x07017fe9180629cfffba412f65a9affcf9a121de02294179f5c058f881dcc9f8.
//
// Solidity: event LogUnFrozen()
func (_Exchange *ExchangeFilterer) ParseLogUnFrozen(log types.Log) (*ExchangeLogUnFrozen, error) {
	event := new(ExchangeLogUnFrozen)
	if err := _Exchange.contract.UnpackLog(event, "LogUnFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogUpdateStateIterator is returned from FilterLogUpdateState and is used to iterate over the raw logs and unpacked data for LogUpdateState events raised by the Exchange contract.
type ExchangeLogUpdateStateIterator struct {
	Event *ExchangeLogUpdateState // Event containing the contract specifics and raw log

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
func (it *ExchangeLogUpdateStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogUpdateState)
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
		it.Event = new(ExchangeLogUpdateState)
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
func (it *ExchangeLogUpdateStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogUpdateStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogUpdateState represents a LogUpdateState event raised by the Exchange contract.
type ExchangeLogUpdateState struct {
	SequenceNumber *big.Int
	BatchId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateState is a free log retrieval operation binding the contract event 0x2672b53d25204094519f7b0fba8d2b5cd0cc1e426f49554c89461cdb9dcec08f.
//
// Solidity: event LogUpdateState(uint256 sequenceNumber, uint256 batchId)
func (_Exchange *ExchangeFilterer) FilterLogUpdateState(opts *bind.FilterOpts) (*ExchangeLogUpdateStateIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogUpdateState")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogUpdateStateIterator{contract: _Exchange.contract, event: "LogUpdateState", logs: logs, sub: sub}, nil
}

// WatchLogUpdateState is a free log subscription operation binding the contract event 0x2672b53d25204094519f7b0fba8d2b5cd0cc1e426f49554c89461cdb9dcec08f.
//
// Solidity: event LogUpdateState(uint256 sequenceNumber, uint256 batchId)
func (_Exchange *ExchangeFilterer) WatchLogUpdateState(opts *bind.WatchOpts, sink chan<- *ExchangeLogUpdateState) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogUpdateState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogUpdateState)
				if err := _Exchange.contract.UnpackLog(event, "LogUpdateState", log); err != nil {
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

// ParseLogUpdateState is a log parse operation binding the contract event 0x2672b53d25204094519f7b0fba8d2b5cd0cc1e426f49554c89461cdb9dcec08f.
//
// Solidity: event LogUpdateState(uint256 sequenceNumber, uint256 batchId)
func (_Exchange *ExchangeFilterer) ParseLogUpdateState(log types.Log) (*ExchangeLogUpdateState, error) {
	event := new(ExchangeLogUpdateState)
	if err := _Exchange.contract.UnpackLog(event, "LogUpdateState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogUserAdminAddedIterator is returned from FilterLogUserAdminAdded and is used to iterate over the raw logs and unpacked data for LogUserAdminAdded events raised by the Exchange contract.
type ExchangeLogUserAdminAddedIterator struct {
	Event *ExchangeLogUserAdminAdded // Event containing the contract specifics and raw log

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
func (it *ExchangeLogUserAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogUserAdminAdded)
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
		it.Event = new(ExchangeLogUserAdminAdded)
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
func (it *ExchangeLogUserAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogUserAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogUserAdminAdded represents a LogUserAdminAdded event raised by the Exchange contract.
type ExchangeLogUserAdminAdded struct {
	UserAdmin common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogUserAdminAdded is a free log retrieval operation binding the contract event 0x7284e8b42a1333a4f23e858e513b3b28d2667a3531b7c1872cce3f7720a25046.
//
// Solidity: event LogUserAdminAdded(address userAdmin)
func (_Exchange *ExchangeFilterer) FilterLogUserAdminAdded(opts *bind.FilterOpts) (*ExchangeLogUserAdminAddedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogUserAdminAdded")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogUserAdminAddedIterator{contract: _Exchange.contract, event: "LogUserAdminAdded", logs: logs, sub: sub}, nil
}

// WatchLogUserAdminAdded is a free log subscription operation binding the contract event 0x7284e8b42a1333a4f23e858e513b3b28d2667a3531b7c1872cce3f7720a25046.
//
// Solidity: event LogUserAdminAdded(address userAdmin)
func (_Exchange *ExchangeFilterer) WatchLogUserAdminAdded(opts *bind.WatchOpts, sink chan<- *ExchangeLogUserAdminAdded) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogUserAdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogUserAdminAdded)
				if err := _Exchange.contract.UnpackLog(event, "LogUserAdminAdded", log); err != nil {
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

// ParseLogUserAdminAdded is a log parse operation binding the contract event 0x7284e8b42a1333a4f23e858e513b3b28d2667a3531b7c1872cce3f7720a25046.
//
// Solidity: event LogUserAdminAdded(address userAdmin)
func (_Exchange *ExchangeFilterer) ParseLogUserAdminAdded(log types.Log) (*ExchangeLogUserAdminAdded, error) {
	event := new(ExchangeLogUserAdminAdded)
	if err := _Exchange.contract.UnpackLog(event, "LogUserAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogUserAdminRemovedIterator is returned from FilterLogUserAdminRemoved and is used to iterate over the raw logs and unpacked data for LogUserAdminRemoved events raised by the Exchange contract.
type ExchangeLogUserAdminRemovedIterator struct {
	Event *ExchangeLogUserAdminRemoved // Event containing the contract specifics and raw log

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
func (it *ExchangeLogUserAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogUserAdminRemoved)
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
		it.Event = new(ExchangeLogUserAdminRemoved)
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
func (it *ExchangeLogUserAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogUserAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogUserAdminRemoved represents a LogUserAdminRemoved event raised by the Exchange contract.
type ExchangeLogUserAdminRemoved struct {
	UserAdmin common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogUserAdminRemoved is a free log retrieval operation binding the contract event 0xb32f8aed6bedf93605e95bc99e0e229b8bbfcd0fe2e76a6748450d3e9522db46.
//
// Solidity: event LogUserAdminRemoved(address userAdmin)
func (_Exchange *ExchangeFilterer) FilterLogUserAdminRemoved(opts *bind.FilterOpts) (*ExchangeLogUserAdminRemovedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogUserAdminRemoved")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogUserAdminRemovedIterator{contract: _Exchange.contract, event: "LogUserAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchLogUserAdminRemoved is a free log subscription operation binding the contract event 0xb32f8aed6bedf93605e95bc99e0e229b8bbfcd0fe2e76a6748450d3e9522db46.
//
// Solidity: event LogUserAdminRemoved(address userAdmin)
func (_Exchange *ExchangeFilterer) WatchLogUserAdminRemoved(opts *bind.WatchOpts, sink chan<- *ExchangeLogUserAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogUserAdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogUserAdminRemoved)
				if err := _Exchange.contract.UnpackLog(event, "LogUserAdminRemoved", log); err != nil {
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

// ParseLogUserAdminRemoved is a log parse operation binding the contract event 0xb32f8aed6bedf93605e95bc99e0e229b8bbfcd0fe2e76a6748450d3e9522db46.
//
// Solidity: event LogUserAdminRemoved(address userAdmin)
func (_Exchange *ExchangeFilterer) ParseLogUserAdminRemoved(log types.Log) (*ExchangeLogUserAdminRemoved, error) {
	event := new(ExchangeLogUserAdminRemoved)
	if err := _Exchange.contract.UnpackLog(event, "LogUserAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogUserRegisteredIterator is returned from FilterLogUserRegistered and is used to iterate over the raw logs and unpacked data for LogUserRegistered events raised by the Exchange contract.
type ExchangeLogUserRegisteredIterator struct {
	Event *ExchangeLogUserRegistered // Event containing the contract specifics and raw log

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
func (it *ExchangeLogUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogUserRegistered)
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
		it.Event = new(ExchangeLogUserRegistered)
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
func (it *ExchangeLogUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogUserRegistered represents a LogUserRegistered event raised by the Exchange contract.
type ExchangeLogUserRegistered struct {
	EthKey   common.Address
	StarkKey *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogUserRegistered is a free log retrieval operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_Exchange *ExchangeFilterer) FilterLogUserRegistered(opts *bind.FilterOpts) (*ExchangeLogUserRegisteredIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogUserRegisteredIterator{contract: _Exchange.contract, event: "LogUserRegistered", logs: logs, sub: sub}, nil
}

// WatchLogUserRegistered is a free log subscription operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_Exchange *ExchangeFilterer) WatchLogUserRegistered(opts *bind.WatchOpts, sink chan<- *ExchangeLogUserRegistered) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogUserRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogUserRegistered)
				if err := _Exchange.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
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

// ParseLogUserRegistered is a log parse operation binding the contract event 0xcab1cf17c190e4e2195a7b8f7b362023246fa774390432b4704ab6b29d56b07b.
//
// Solidity: event LogUserRegistered(address ethKey, uint256 starkKey, address sender)
func (_Exchange *ExchangeFilterer) ParseLogUserRegistered(log types.Log) (*ExchangeLogUserRegistered, error) {
	event := new(ExchangeLogUserRegistered)
	if err := _Exchange.contract.UnpackLog(event, "LogUserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogWithdrawalAllowedIterator is returned from FilterLogWithdrawalAllowed and is used to iterate over the raw logs and unpacked data for LogWithdrawalAllowed events raised by the Exchange contract.
type ExchangeLogWithdrawalAllowedIterator struct {
	Event *ExchangeLogWithdrawalAllowed // Event containing the contract specifics and raw log

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
func (it *ExchangeLogWithdrawalAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogWithdrawalAllowed)
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
		it.Event = new(ExchangeLogWithdrawalAllowed)
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
func (it *ExchangeLogWithdrawalAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogWithdrawalAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogWithdrawalAllowed represents a LogWithdrawalAllowed event raised by the Exchange contract.
type ExchangeLogWithdrawalAllowed struct {
	StarkKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalAllowed is a free log retrieval operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) FilterLogWithdrawalAllowed(opts *bind.FilterOpts) (*ExchangeLogWithdrawalAllowedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogWithdrawalAllowedIterator{contract: _Exchange.contract, event: "LogWithdrawalAllowed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalAllowed is a free log subscription operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) WatchLogWithdrawalAllowed(opts *bind.WatchOpts, sink chan<- *ExchangeLogWithdrawalAllowed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogWithdrawalAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogWithdrawalAllowed)
				if err := _Exchange.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
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

// ParseLogWithdrawalAllowed is a log parse operation binding the contract event 0x03c10a82c955f7bcd0c934147fb39cafca947a4294425b1751d884c8ac954287.
//
// Solidity: event LogWithdrawalAllowed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount)
func (_Exchange *ExchangeFilterer) ParseLogWithdrawalAllowed(log types.Log) (*ExchangeLogWithdrawalAllowed, error) {
	event := new(ExchangeLogWithdrawalAllowed)
	if err := _Exchange.contract.UnpackLog(event, "LogWithdrawalAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeLogWithdrawalPerformedIterator is returned from FilterLogWithdrawalPerformed and is used to iterate over the raw logs and unpacked data for LogWithdrawalPerformed events raised by the Exchange contract.
type ExchangeLogWithdrawalPerformedIterator struct {
	Event *ExchangeLogWithdrawalPerformed // Event containing the contract specifics and raw log

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
func (it *ExchangeLogWithdrawalPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogWithdrawalPerformed)
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
		it.Event = new(ExchangeLogWithdrawalPerformed)
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
func (it *ExchangeLogWithdrawalPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogWithdrawalPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogWithdrawalPerformed represents a LogWithdrawalPerformed event raised by the Exchange contract.
type ExchangeLogWithdrawalPerformed struct {
	StarkKey           *big.Int
	AssetType          *big.Int
	NonQuantizedAmount *big.Int
	QuantizedAmount    *big.Int
	Recipient          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawalPerformed is a free log retrieval operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Exchange *ExchangeFilterer) FilterLogWithdrawalPerformed(opts *bind.FilterOpts) (*ExchangeLogWithdrawalPerformedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogWithdrawalPerformedIterator{contract: _Exchange.contract, event: "LogWithdrawalPerformed", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawalPerformed is a free log subscription operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Exchange *ExchangeFilterer) WatchLogWithdrawalPerformed(opts *bind.WatchOpts, sink chan<- *ExchangeLogWithdrawalPerformed) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogWithdrawalPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogWithdrawalPerformed)
				if err := _Exchange.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
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

// ParseLogWithdrawalPerformed is a log parse operation binding the contract event 0xb7477a7b93b2addc5272bbd7ad0986ef1c0d0bd265f26c3dc4bbe42727c2ac0c.
//
// Solidity: event LogWithdrawalPerformed(uint256 starkKey, uint256 assetType, uint256 nonQuantizedAmount, uint256 quantizedAmount, address recipient)
func (_Exchange *ExchangeFilterer) ParseLogWithdrawalPerformed(log types.Log) (*ExchangeLogWithdrawalPerformed, error) {
	event := new(ExchangeLogWithdrawalPerformed)
	if err := _Exchange.contract.UnpackLog(event, "LogWithdrawalPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
