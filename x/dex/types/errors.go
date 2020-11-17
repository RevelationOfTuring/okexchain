package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// const CodeType
const (
	CodeAddrAndProductAllRequired sdk.CodeType = 64000
	codeInvalidProduct            sdk.CodeType = 64001
	codeTokenPairNotFound         sdk.CodeType = 64002
	codeDelistOwnerNotMatch       sdk.CodeType = 64003
	codeInvalidBalanceNotEnough   sdk.CodeType = 64004
	codeInvalidAsset              sdk.CodeType = 64005
	codeUnknownOperator           sdk.CodeType = 64006
	codeExistOperator             sdk.CodeType = 64007
	codeInvalidWebsiteLength      sdk.CodeType = 64008
	codeInvalidWebsiteURL         sdk.CodeType = 64009
	CodeTokenPairIsInvalid        sdk.CodeType = 64010
)

// CodeType to Message
func codeToDefaultMsg(code sdk.CodeType) string {
	switch code {
	case codeInvalidProduct:
		return "invalid product"
	case codeTokenPairNotFound:
		return "tokenpair not found"
	case codeDelistOwnerNotMatch:
		return "tokenpair delistor should be it's owner "
	default:
		return fmt.Sprintf("unknown code %d", code)
	}
}

// Addr and Product All Required
func ErrAddrAndProductAllRequired() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAddrAndProductAllRequired, "bad request: address、base_asset and quote_asset could not be empty at the same time")
}

// invalid tokenpair
func ErrTokenPairIsInvalid() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeTokenPairIsInvalid, "the nil pointer is not expected")
}

// ErrInvalidProduct returns invalid product error
func ErrInvalidProduct(msg string) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeInvalidProduct, codeToDefaultMsg(codeInvalidProduct)+": %s", msg)
}

// ErrTokenPairNotFound returns token pair not found error
func ErrTokenPairNotFound(msg string) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeTokenPairNotFound, codeToDefaultMsg(codeTokenPairNotFound)+": %s", msg)
}

// ErrDelistOwnerNotMatch returns delist owner not match error
func ErrDelistOwnerNotMatch(msg string) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeDelistOwnerNotMatch, codeToDefaultMsg(codeDelistOwnerNotMatch)+": %s", msg)
}

// ErrInvalidBalanceNotEnough returns invalid balance not enough error
func ErrInvalidBalanceNotEnough(message string) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeInvalidBalanceNotEnough, message)
}

// ErrInvalidAsset returns invalid asset error
func ErrInvalidAsset(message string) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeInvalidAsset, message)
}

func ErrUnknownOperator(addr sdk.AccAddress) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeUnknownOperator, fmt.Sprintf("unknown dex operator with address %s", addr.String()))
}

func ErrExistOperator(addr sdk.AccAddress) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeExistOperator, fmt.Sprintf("dex operator already exists with address %s", addr.String()))
}

func ErrInvalidWebsiteLength(got, max int) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeInvalidWebsiteLength, fmt.Sprintf("invalid website length, got length %v, max is %v", got, max))
}

func ErrInvalidWebsiteURL(msg string) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeInvalidWebsiteURL, fmt.Sprintf("invalid website URL: %s", msg))
}

// ErrTokenPairExisted returns an error when the token pair is existed during the process of listing
// ErrTokenPairExisted returns an error when the token pair is existing during the process of listing
func ErrTokenPairExisted(baseAsset, quoteAsset string) sdk.Error {
	return sdk.NewError(DefaultCodespace, codeInvalidAsset,
		fmt.Sprintf("failed. the token pair exists with %s and %s", baseAsset, quoteAsset))
}
