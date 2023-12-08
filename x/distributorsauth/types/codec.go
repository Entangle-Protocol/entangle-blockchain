package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgAddAdmin{}, "entangle/distributorsauth/MsgAddAdmin", nil)
	cdc.RegisterConcrete(&MsgRemoveAdmin{}, "entangle/distributorsauth/MsgRemoveAdmin", nil)
	cdc.RegisterConcrete(&MsgAddDistributor{}, "entangle/distributorsauth/MsgAddDistributor", nil)
	cdc.RegisterConcrete(&MsgRemoveDistributor{}, "entangle/distributorsauth/MsgRemoveDistributor", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	registry.RegisterImplementations((*govtypes.Content)(nil), &AddDistributorProposal{})

}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
