package incentive

import (
	"github.com/coinsky/sky-sdk/modules/incentive/internal/keepers"
	"github.com/coinsky/sky-sdk/modules/incentive/internal/types"
)

type (
	GenesisState = types.GenesisState
	State        = types.State
	Params       = types.Params
	Plan         = types.Plan
	Keeper       = keepers.Keeper
)

const (
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
)

var (
	ModuleCdc           = types.ModuleCdc
	DefaultGenesisState = types.DefaultGenesisState
	DefaultParams       = types.DefaultParams
	NewKeeper           = keepers.NewKeeper
)
