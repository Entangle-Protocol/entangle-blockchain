package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EnsureBech32Address(address string) (string, error) {
	if _, err := sdk.AccAddressFromBech32(address); err == nil {
		// Address is already in bech32 format
		return address, nil
	}

	// If address is not in bech32 format, convert it
	// Remove 0x prefix
	address = strings.TrimPrefix(address, "0x")

	bytes, err := sdk.AccAddressFromHexUnsafe(address)
	if err != nil {
		return "", err
	}

	bech32Address, err := sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), bytes)
	if err != nil {
		return "", err
	}

	return bech32Address, nil
}

func CheckDistrubutorEndDateExpire(ctx sdk.Context, endDate uint64) bool {
	blockTime := uint64(ctx.BlockTime().Unix())
	return (endDate > 0 && endDate < blockTime)
}
