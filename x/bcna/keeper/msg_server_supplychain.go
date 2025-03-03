package keeper

import (
	"context"

	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSupplychain(goCtx context.Context, msg *types.MsgCreateSupplychain) (*types.MsgCreateSupplychainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var supplychain = types.Supplychain{
		Creator:     msg.Creator,
		Product:     msg.Product,
		Info:        msg.Info,
		Supplyinfo:  msg.Supplyinfo,
		Supplyextra: msg.Supplyextra,
	}

	id := k.AppendSupplychain(
		ctx,
		supplychain,
	)

	return &types.MsgCreateSupplychainResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSupplychain(goCtx context.Context, msg *types.MsgUpdateSupplychain) (*types.MsgUpdateSupplychainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var supplychain = types.Supplychain{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Product:     msg.Product,
		Info:        msg.Info,
		Supplyinfo:  msg.Supplyinfo,
		Supplyextra: msg.Supplyextra,
	}

	// Checks that the element exists
	val, found := k.GetSupplychain(ctx, msg.Id)
	if !found {
		return nil, types.ErrKeyNotFound.Wrapf("key doesn't exist: %d", msg.Id)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, types.ErrUnauthorized.Wrapf("Unauthorized: %s,", msg.Creator)
	}

	k.SetSupplychain(ctx, supplychain)

	return &types.MsgUpdateSupplychainResponse{}, nil
}

func (k msgServer) DeleteSupplychain(goCtx context.Context, msg *types.MsgDeleteSupplychain) (*types.MsgDeleteSupplychainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetSupplychain(ctx, msg.Id)
	if !found {
		return nil, types.ErrKeyNotFound.Wrapf("key doesn't exist: %d", msg.Id)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, types.ErrUnauthorized.Wrapf("Unauthorized: %s,", msg.Creator)
	}

	k.RemoveSupplychain(ctx, msg.Id)

	return &types.MsgDeleteSupplychainResponse{}, nil
}
