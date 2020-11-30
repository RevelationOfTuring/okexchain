package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace 						sdk.CodespaceType = "ammswap"

	CodeUnexistwapTokenPair       			sdk.CodeType = 65000
	CodeUnexistPoolToken					sdk.CodeType = 65001
	CodeMintCoinsFailed						sdk.CodeType = 65002
	CodeSendCoinsFromAccountToModule 		sdk.CodeType = 65003
	CodeBaseAmountNameBigerQuoteAmountName	sdk.CodeType = 65004
	CodeValidateSwapAmountName				sdk.CodeType = 65005
	CodeBaseAmountNameEqualQuoteAmountName	sdk.CodeType = 65006
	CodeValidateDenom						sdk.CodeType = 65007
	CodeNotAllowedOriginSymbol				sdk.CodeType = 65008
	CodeInsufficientPoolToken				sdk.CodeType = 65009
	CodeUnknownRequest						sdk.CodeType = 65010
	CodeInsufficientCoins					sdk.CodeType = 65011
	CodeTokenNotExist						sdk.CodeType = 65012
	CodeInvalidCoins						sdk.CodeType = 65013
	CodeInternalError						sdk.CodeType = 65014
)

func ErrUnexistswapTokenPair(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeUnexistwapTokenPair, message)
}

func ErrUnexistPoolToken(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeUnexistPoolToken, message)
}

func ErrCodeMinCoinsFailed(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeMintCoinsFailed, message)
}

func ErrSendCoinsFromAccountToModule(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeSendCoinsFromAccountToModule, message)
}

func ErrBaseAmountNameBigerQuoteAmountName(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeBaseAmountNameBigerQuoteAmountName, message)
}

func ErrValidateSwapAmountName(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeValidateSwapAmountName, message)
}

func ErrBaseAmountNameEqualQuoteAmountName(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeBaseAmountNameEqualQuoteAmountName, message)
}

func ErrValidateDenom(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeValidateDenom, message)
}

func ErrNotAllowedOriginSymbol(codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeNotAllowedOriginSymbol, message)
}

func ErrInsufficientPoolToken (codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeInsufficientPoolToken, message)
}

func ErrUnknownRequest (codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeUnknownRequest, message)
}

func ErrTokenNotExist (codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotExist, message)
}

func ErrInvalidCoins (codespace sdk.CodespaceType, message string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidCoins, message)
}