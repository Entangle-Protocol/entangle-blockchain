package ante

import (
	// "github.com/Entangle-Protocol/entangle-blockchain/x/evm/types"
	// "github.com/Entangle-Protocol/entangle-blockchain/x/evm/types"

	"github.com/Entangle-Protocol/entangle-blockchain/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DistributorsAuthDecorator keeps track if transaction have correct distributor signer
// NOTE: This decorator does not perform any validation
type DistributorsAuthDecorator struct {
	distrKeeper DistributorsAuthKeeper
}

// NewGasWantedDecorator creates a new NewGasWantedDecorator
func NewDistributorsAuthDecorator(
	distrKeeper DistributorsAuthKeeper,
) DistributorsAuthDecorator {
	return DistributorsAuthDecorator{
		distrKeeper,
	}
}

func (dg DistributorsAuthDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if !simulate && !dg.distrKeeper.IsTestingMode() {
		for i, msg := range tx.GetMsgs() {
			ethTx, isEthTx := msg.(*types.MsgEthereumTx)
			if isEthTx {
				txData, err := types.UnpackTxData(ethTx.Data)
				if err != nil {
					return ctx, sdkerrors.Wrapf(err, "failed to unpack tx data any for tx %d", i)
				}

				if txData.GetTo() == nil {
					if dg.distrKeeper.ValidateTransaction(ctx, ethTx.From) != nil {
						return ctx, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Transaction have no valid Distributor signers")
					}
				}
			}
		}
	}
	return next(ctx, tx, simulate)
}
