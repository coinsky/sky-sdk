package alias

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/coinsky/sky-sdk/modules/alias/client/cli"
	"github.com/coinsky/sky-sdk/modules/alias/client/rest"
	"github.com/coinsky/sky-sdk/modules/alias/internal/keepers"
	"github.com/coinsky/sky-sdk/modules/alias/internal/types"
)

// app module basics object
type AppModuleBasic struct {
}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

// genesis
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return types.ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(data json.RawMessage) error {
	var state GenesisState
	if err := types.ModuleCdc.UnmarshalJSON(data, &state); err != nil {
		return err
	}
	return state.Validate()
}

// client functionality
func (amb AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr, types.ModuleCdc)
}

func (amb AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(cdc)
}

func (amb AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(cdc)
}

// ___________________________
// app module object
type AppModule struct {
	AppModuleBasic
	blKeeper keepers.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(blKeeper keepers.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		blKeeper:       blKeeper,
	}
}

// registers
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// routes
func (am AppModule) Route() string {
	return types.RouterKey
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.blKeeper)
}

func (am AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return keepers.NewQuerier(am.blKeeper)
}

func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
}

func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	types.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.blKeeper, genesisState)
	return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.blKeeper)
	return types.ModuleCdc.MustMarshalJSON(gs)
}
