// nolint
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace            sdk.CodespaceType = "distr"
	CodeInvalidInput									CodeType          = 67800
	CodeNoValidatorCommission							CodeType          = 67801
	CodeSetWithdrawAddrDisabled							CodeType          = 67802
	CodeInvalideData									CodeType		  = 67803
	CodeInvalideRoute									CodeType		  = 67804
	CodeInvalideBasic									CodeType		  = 67805
	CodeWithdrawValidatorRewardsAndCommissionFailed		CodeType		  = 67806
	CodeAccAddressFromBech32Failed						CodeType		  = 67807
	CodeValAddressFromBech32							CodeType		  = 67808
	CodeReadRESTReqFailed								CodeType		  = 67809
	CodeSendCoinsFromModuleToAccountFailed				CodeType		  = 67810
	CodeUnknownRequest									CodeType		  = 67811
	CodeUnauthorized									CodeType		  = 67812
	CodeSetWithdrawAddrFailed							CodeType		  = 67813
	CodeWithdrawValidatorCommissionFailed				CodeType		  = 67814
)

func ErrNilDelegatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "delegator address is nil")
}
func ErrNilWithdrawAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "withdraw address is nil")
}
func ErrNilValidatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "validator address is nil")
}
func ErrNoValidatorCommission(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNoValidatorCommission, "no validator commission to withdraw")
}
func ErrSetWithdrawAddrDisabled(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeSetWithdrawAddrDisabled, "set withdraw address disabled")
}
func ErrBadDistribution(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "community pool does not have sufficient coins to distribute")
}
func ErrInvalidProposalAmount(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "invalid community pool spend proposal amount")
}
func ErrEmptyProposalRecipient(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "invalid community pool spend proposal recipient")
}
func ErrSendCoinsFromModuleToAccountFailed(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeSendCoinsFromModuleToAccountFailed, "invalid withdrawAddr or commission")
}
func ErrUnknownRequest(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeUnknownRequest, "incorrectly formatted request data")
}
func ErrUnauthorized(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeUnknownRequest, "blacklisted from receiving external funds")
}
func ErrSetWithdrawAddrFailed(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeSetWithdrawAddrFailed, "delegators addr or withdraw addr is invalid")
}
func ERRWithdrawValidatorCommissionFailed(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeWithdrawValidatorCommissionFailed, "withdraw validator commission failed")
}