package keeper

import (
	"context"
	"blogs/x/blogs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowPost(goCtx context.Context, req *types.QueryShowPostRequest) (*types.QueryShowPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var write []string
	write = append(write, "Empty comment")
	// TODO: Process the query
		post, found := k.GetPost(ctx, req.Id)
		value, _ := k.GetComment(ctx, req.Id)
	
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	if post.Id == value.Id {
		
		return &types.QueryShowPostResponse{
			post,
			value,
			write,
			}, nil
	}
	value.Id = post.Id
	return &types.QueryShowPostResponse{
		post,
		value,
		write,
		}, nil
}
