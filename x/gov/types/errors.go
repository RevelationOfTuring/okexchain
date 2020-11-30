//nolint
package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = "gov"

	CodeUnknownProposal         sdk.CodeType = 68000
	CodeInvalidProposalStatus   sdk.CodeType = 68001
	CodeInitialDepositNotEnough sdk.CodeType = 68002
	CodeInvalidProposer         sdk.CodeType = 68003
	CodeInvalidHeight           sdk.CodeType = 68004
	CodeInsufficientCoins		sdk.CodeType = 68005
	CodeUnknownRequest			sdk.CodeType = 68006
	CodeInvalidCoins			sdk.CodeType = 68007
)

func ErrUnknownProposal(codespace sdk.CodespaceType, proposalID uint64) sdk.Error {
	return sdk.NewError(codespace, CodeUnknownProposal, fmt.Sprintf("unknown proposal with id %d", proposalID))
}

func ErrInvalidateProposalStatus(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidProposalStatus, msg)
}

func ErrInitialDepositNotEnough(codespace sdk.CodespaceType, initDeposit string) sdk.Error {
	return sdk.NewError(codespace, CodeInitialDepositNotEnough,
		fmt.Sprintf("InitialDeposit must be greater than or equal to %s", initDeposit))
}

func ErrInvalidProposer(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidProposer, message)
}

func ErrInvalidHeight(codespace sdk.CodespaceType, h, ch, max uint64) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidHeight,
		fmt.Sprintf("Height %d must be greater than current block height %d and less than %d + %d.",
			h, ch, ch, max))
}

func ErrInsufficientCoins(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeInsufficientCoins, message)
}

func ErrUnknownRequest(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeUnknownRequest, message)
}

func ErrInvalidCoins(message string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidCoins, message)
}