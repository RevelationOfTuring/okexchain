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
	CodeUnmarshalJSONFailed								CodeType		  = 67805
	CodeInvalideBasic									CodeType		  = 67806
	CodeWithdrawValidatorRewardsAndCommissionFailed		CodeType		  = 67806

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
