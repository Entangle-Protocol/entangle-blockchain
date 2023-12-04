package keeper

import (
	"fmt"
	"strconv"

	"github.com/Entangle-Protocol/entangle-blockchain/x/distributorsauth/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"
)

// var (
// 	pollPrefix    = "poll"
// 	votesPrefix   = "votes"
// 	countKey      = "count"
// 	pollQueueName = "pending_poll_queue"
// )

type (
	Keeper struct {
		cdc         codec.BinaryCodec
		storeKey    storetypes.StoreKey
		memKey      storetypes.StoreKey
		AuthKeeper  types.AccountKeeper
		testingMode bool
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	at types.AccountKeeper,
	testingMode bool,
) *Keeper {
	return &Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		memKey:      memKey,
		AuthKeeper:  at,
		testingMode: testingMode,
	}
}

func (k Keeper) IsTestingMode() bool {
	return k.testingMode
}

func (k *Keeper) SetTestingMode(testingMode bool) {
	k.testingMode = testingMode
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) AddDistributor(ctx sdk.Context, info types.DistributorInfo) error {
	bech32Address, err := types.EnsureBech32Address(info.Address)
	if err != nil {
		return err
	}

	info.Address = bech32Address
	k.addDistributor(ctx, info)
	return nil
}

func (k *Keeper) addDistributor(ctx sdk.Context, info types.DistributorInfo) {
	store := ctx.KVStore(k.storeKey)
	stringKey := types.DistributorPrefix + info.Address
	key := []byte(stringKey)
	value := k.cdc.MustMarshalLengthPrefixed(&info)
	store.Set(key, value)

	ctx.EventManager().EmitEvent(
		types.AddDistributorEvent(info.Address, strconv.FormatUint(info.EndDate, 10)),
	)
}

func (k *Keeper) GetDistributor(ctx sdk.Context, address string) (types.DistributorInfo, error) {
	bech32Address, err := types.EnsureBech32Address(address)
	if err != nil {
		return types.DistributorInfo{}, sdkerrors.Wrap(types.ErrWrongDistributorAddress, address)
	}

	return k.getDistributor(ctx, bech32Address)
}

func (k *Keeper) getDistributor(ctx sdk.Context, address string) (types.DistributorInfo, error) {
	store := ctx.KVStore(k.storeKey)
	stringKey := types.DistributorPrefix + address
	key := []byte(stringKey)
	value := store.Get(key)
	if value == nil {
		return types.DistributorInfo{
			Address: "",
			EndDate: 0,
		}, sdkerrors.Wrap(types.ErrNoDistributorInStoreForAddress, address)
	}

	var info types.DistributorInfo
	k.cdc.MustUnmarshalLengthPrefixed(value, &info)

	return info, nil
}

func (k *Keeper) GetDistributors(ctx sdk.Context) ([]types.DistributorInfo, error) {
	var distrList []types.DistributorInfo
	store := ctx.KVStore(k.storeKey)
	bytes_prefix := []byte(types.DistributorPrefix)
	iterator := sdk.KVStorePrefixIterator(store, bytes_prefix)
	for ; iterator.Valid(); iterator.Next() {
		var info types.DistributorInfo
		k.cdc.MustUnmarshalLengthPrefixed(store.Get(iterator.Key()), &info)
		distrList = append(distrList, info)
	}
	// k.cdc.Unmarshal(iter.Value(), &authorization)
	// res := codec.MustMarshalJSONIndent(&k.cdc, distrList)
	return distrList, nil
}

func (k *Keeper) RemoveDistributor(ctx sdk.Context, address string) {
	bech32Address, err := types.EnsureBech32Address(address)
	if err != nil {
		return
	}

	k.removeDistributor(ctx, bech32Address)
}

func (k *Keeper) removeDistributor(ctx sdk.Context, address string) {
	key := []byte(types.DistributorPrefix + address)
	store := ctx.KVStore(k.storeKey)
	store.Delete(key)
	ctx.EventManager().EmitEvent(
		types.RemoveDistributorEvent(address),
	)

}

func (k Keeper) IterateDistributors(ctx sdk.Context,
	handler func(address string, end_date uint64) bool,
) {
	store := ctx.KVStore(k.storeKey)
	bytesPrefix := []byte(types.DistributorPrefix)
	iter := sdk.KVStorePrefixIterator(store, bytesPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var distributor types.DistributorInfo
		k.cdc.MustUnmarshalLengthPrefixed(store.Get(iter.Key()), &distributor)
		if handler(distributor.Address, distributor.EndDate) {
			break
		}
	}
}

func (k *Keeper) AddAdmin(ctx sdk.Context, admin types.Admin) error {
	bech32Address, err := types.EnsureBech32Address(admin.Address)
	if err != nil {
		return err
	}

	admin.Address = bech32Address
	return k.addAdmin(ctx, admin)
}

func (k *Keeper) addAdmin(ctx sdk.Context, admin types.Admin) error {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.DistributorAdminPrefix + admin.Address)
	value := k.cdc.MustMarshalLengthPrefixed(&admin)
	store.Set(key, value)
	ctx.EventManager().EmitEvent(
		types.AddAdminEvent(admin.Address, strconv.FormatBool(admin.EditOption)),
	)
	return nil
}

func (k *Keeper) GetAdmin(ctx sdk.Context, address string) (types.Admin, error) {
	bech32Address, err := types.EnsureBech32Address(address)
	if err != nil {
		return types.Admin{}, sdkerrors.Wrap(types.ErrWrongAdminAddress, address)
	}

	return k.getAdmin(ctx, bech32Address)
}

func (k *Keeper) getAdmin(ctx sdk.Context, address string) (types.Admin, error) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.DistributorAdminPrefix + address)
	value := store.Get(key)
	if value == nil {
		return types.Admin{}, sdkerrors.Wrap(types.ErrNoAdminInStoreForAddress, address)
	}

	var admin types.Admin
	k.cdc.MustUnmarshalLengthPrefixed(value, &admin)

	return admin, nil
}

func (k Keeper) GetAdmins(ctx sdk.Context) ([]types.Admin, error) {
	var AdminList []types.Admin
	store := ctx.KVStore(k.storeKey)
	bytes_prefix := []byte(types.DistributorAdminPrefix)
	iterator := sdk.KVStorePrefixIterator(store, bytes_prefix)
	for ; iterator.Valid(); iterator.Next() {
		var admin types.Admin
		k.cdc.MustUnmarshalLengthPrefixed(store.Get(iterator.Key()), &admin)
		AdminList = append(AdminList, admin)
	}
	// k.cdc.Unmarshal(iter.Value(), &authorization)
	// res := codec.MustMarshalJSONIndent(&k.cdc, distrList)
	return AdminList, nil
}

func (k *Keeper) RemoveAdmin(ctx sdk.Context, address string) {
	bech32Address, err := types.EnsureBech32Address(address)
	if err != nil {
		return
	}

	k.removeAdmin(ctx, bech32Address)
}

func (k *Keeper) removeAdmin(ctx sdk.Context, address string) {
	key := []byte(types.DistributorAdminPrefix + address)
	store := ctx.KVStore(k.storeKey)
	store.Delete(key)
	ctx.EventManager().EmitEvent(
		types.RemoveAdminEvent(address),
	)
}

func (k Keeper) IterateAdmins(ctx sdk.Context,
	handler func(address string, editOption bool) bool,
) {
	store := ctx.KVStore(k.storeKey)
	bytes_prefix := []byte(types.DistributorAdminPrefix)
	iter := sdk.KVStorePrefixIterator(store, bytes_prefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var admin types.Admin
		k.cdc.MustUnmarshalLengthPrefixed(store.Get(iter.Key()), &admin)
		if handler(admin.Address, admin.EditOption) {
			break
		}
	}
}

// ValidateDistributor returns error if the given adress is not allowed distributor.
func (k Keeper) ValidateDistributor(ctx sdk.Context, address string) error {
	distributor, err := k.GetDistributor(ctx, address)
	if err != nil {
		return err
	}

	if distributor.Address == "" {
		return sdkerrors.Wrap(types.ErrNoDistributorInStoreForAddress, address)
	}

	return nil
}

// ValidateAdmin returns error if the given adress is not allowed admin.
func (k Keeper) ValidateAdmin(ctx sdk.Context, address string) error {
	if address == "" {
		return sdkerrors.Wrap(types.ErrWrongAdminAddress, address)
	}

	admin, err := k.GetAdmin(ctx, address)
	if err != nil {
		return err
	}

	if admin.Address == "" {
		return sdkerrors.Wrap(types.ErrNoAdminInStoreForAddress, address)
	}

	return nil
}

func (k Keeper) ValidateTransaction(ctx sdk.Context, address string) error {
	if address == "" {
		return sdkerrors.Wrap(types.ErrWrongDistributorAddress, address)
	}

	// // Remove 0x prefix
	// address = strings.TrimPrefix(address, "0x")

	// // Convert Ethereum hex address to bytes
	// decodedBytes, err := hex.DecodeString(address)
	// if err != nil {
	// 	panic(err)
	// }

	// // Convert bytes to bech32
	// bech32Address := sdk.AccAddress(decodedBytes).String()
	// bech32Address, err := types.EnsureBech32Address(address)
	// if err != nil {
	// 	return err
	// }

	err := k.ValidateAdmin(ctx, address)
	if err != nil {
		err := k.ValidateDistributor(ctx, address)
		if err != nil {
			return err
		}
	}

	return nil

}

func (k *Keeper) HandleAddDistributorProposal(ctx sdk.Context, p *types.AddDistributorProposal) error {
	// check address is valid
	_, err := types.EnsureBech32Address(p.Address)
	if err != nil {
		return sdkerrors.Wrap(types.ErrWrongDistributorAddress, p.Address)
	}

	// Create DistributorInfo
	distributorInfo := types.DistributorInfo{
		Address: p.Address,
		EndDate: p.EndDate,
	}

	// Add distributor to store
	err = k.AddDistributor(ctx, distributorInfo)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		types.AddDistributoProposalEvent(p.Address, strconv.FormatUint(p.EndDate, 10)),
	)
	return nil
}
