package keeper

import (
	"github.com/Entangle-Protocol/entangle-blockchain/x/distributorsauth/types"
	"github.com/Entangle-Protocol/entangle-blockchain/x/evm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker
func BeginBlocker(_ sdk.Context, _ abci.RequestBeginBlock, _ keeper.Keeper) {}

func (k *Keeper) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	distributors, err := k.GetDistributors(ctx)
	if err != nil {
		panic(err)
	}

	for _, distributor := range distributors {
		if types.CheckDistrubutorEndDateExpire(ctx, distributor.GetEndDate()) {
			k.RemoveDistributor(ctx, distributor.Address)
		}
	}

	return []abci.ValidatorUpdate{}
}
