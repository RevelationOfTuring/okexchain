package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// const CodeType
const (
	DefaultCodespace = "backend"

	CodeNewErrorsMergedFailed         sdk.CodeType = 62001
	CodeProductIsRequired             sdk.CodeType = 62002
	CodeAddressIsEmpty                sdk.CodeType = 62003
	CodeOrderStatusMustBeOpenOrClosed sdk.CodeType = 62004
	CodeAddressAndProductRequired     sdk.CodeType = 62005
	CodeGetChainHeightFailed          sdk.CodeType = 62006
	CodeGetBlockTxHashesFailed        sdk.CodeType = 62007
	CodeOrderSideMustBuyOrSell        sdk.CodeType = 62008
	CodeProductDoesNotExist           sdk.CodeType = 62009
)

// invalid param side, must be buy or sell
func ErrOrderSideParamMustBuyOrSell(side string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeOrderSideMustBuyOrSell, fmt.Sprintf("Side should not be %s", side))
}

// product is required
func ErrProductIsRequired() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeProductIsRequired, "invalid params: product is required")
}

// product does not exist
func ErrProductDoesNotExist(product string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeProductDoesNotExist, fmt.Sprintf("product %s does not exist", product))
}
