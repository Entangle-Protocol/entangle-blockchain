package distributorsauth

import (
	"math/rand"

	"github.com/Entangle-Protocol/entangle-blockchain/testutil/sample"
	distributorsauthsimulation "github.com/Entangle-Protocol/entangle-blockchain/x/distributorsauth/simulation"
	"github.com/Entangle-Protocol/entangle-blockchain/x/distributorsauth/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = distributorsauthsimulation.FindAccount
	//_ = simappparams.StakePerAccount
	_ = rand.Rand{}
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	admins := make([]types.Admin, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		admins[i] = types.Admin{
			Address:    acc.Address.String(),
			EditOption: false,
		}
	}
	distributorsauthGenesis := types.GenesisState{
		//Params:   types.DefaultParams(),
		Admins: admins,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&distributorsauthGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
