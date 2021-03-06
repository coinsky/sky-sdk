package incentive

import (
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/coinsky/sky-sdk/modules/incentive/internal/keepers"
	dex "github.com/coinsky/sky-sdk/types"
)

var (
	PoolAddr = sdk.AccAddress(crypto.AddressHash([]byte("incentive_pool")))
)

func BeginBlocker(ctx sdk.Context, k keepers.Keeper) {
	blockRewards := calcRewards(ctx, k)
	if k.HasCoins(ctx, PoolAddr, blockRewards) {
		if err := collectRewardsFromPool(k, ctx, blockRewards); err != nil {
			panic(err)
		}
	}
}

func collectRewardsFromPool(k keepers.Keeper, ctx sdk.Context, blockRewards sdk.Coins) sdk.Error {
	//transfer rewards into collected_fees for further distribution
	if err := k.SendCoinsFromAccountToModule(ctx, PoolAddr, auth.FeeCollectorName, blockRewards); err != nil {
		return err
	}
	return nil
}

func calcRewards(ctx sdk.Context, k keepers.Keeper) sdk.Coins {
	height := ctx.BlockHeader().Height
	adjustmentHeight := k.GetState(ctx).HeightAdjustment
	height += adjustmentHeight

	rewardAmount := int64(0)
	inPlan := false
	plans := k.GetParams(ctx).Plans

	for _, plan := range plans {
		if height > plan.StartHeight && height <= plan.EndHeight {
			// the height may be in different plans, do not break
			rewardAmount += plan.RewardPerBlock
			inPlan = true
		}
	}

	if !inPlan {
		rewardAmount = k.GetParams(ctx).DefaultRewardPerBlock
	}

	blockRewardsCoins := sdk.NewCoins(sdk.NewInt64Coin(dex.DefaultBondDenom, rewardAmount))
	return blockRewardsCoins
}
