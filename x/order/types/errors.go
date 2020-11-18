package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// const CodeType
const (
	CodeProductIsEmpty sdk.CodeType = 63001
	CodeSizeIsInvalid  sdk.CodeType = 63002
)

// invalid size
func ErrInvalidSizeParam(size uint) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeSizeIsInvalid, fmt.Sprintf("invalid param: size= %d", size))
}
