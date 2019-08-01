package starter

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	abci "github.com/tendermint/tendermint/abci/types"
)

// app module Basics object
type BlankModuleBasic struct {
	ModuleName string
}

// TODO comment
type BlankModule struct {
	BlankModuleBasic
	keeper interface{}
}

// TODO comment
type BlankModuleGenesisState []string

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = BlankModule{}
	_ module.AppModuleBasic = BlankModuleBasic{}
)

// TODO comment
func NewBlankModule(name string, keeper interface{}) BlankModule {

	return BlankModule{BlankModuleBasic{name}, keeper}
}

func (bm BlankModuleBasic) Name() string {
	return bm.ModuleName
}

// TODO comment
func (BlankModuleBasic) RegisterCodec(cdc *codec.Codec) {
	panic("RegisterCodec not implemented")
}

// Validation check of the Genesis
func (bm BlankModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return nil
	//panic("ValidateGenesis not implemented")
}

// TODO comment
func (bm BlankModuleBasic) DefaultGenesis() json.RawMessage {
	data := BlankModuleGenesisState{bm.ModuleName}
	cdc := codec.New()
	return cdc.MustMarshalJSON(data)
}

// TODO comment
func (bm BlankModule) Name() string {
	return bm.ModuleName
}

// TODO comment
func (bm BlankModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// TODO comment
func (bm BlankModule) Route() string {
	return bm.ModuleName
}

// TODO comment
func (bm BlankModule) NewQuerierHandler() sdk.Querier {
	panic("NewQuerierHandler not implemented")
}

// TODO comment
func (bm BlankModuleBasic) GetQueryCmd(*codec.Codec) *cobra.Command {
	panic("GetQueryCmd not implemented")
}

// TODO comment
func (bm BlankModuleBasic) GetTxCmd(*codec.Codec) *cobra.Command {
	panic("GetTxCmd not implemented")
}

// Register rest routes
func (BlankModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	//rest.RegisterRoutes(ctx, rtr, cdc, StoreKey)
	panic("RegisterRESTRoutes not implemented")
}

// TODO comment
func (bm BlankModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// TODO comment
func (bm BlankModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// TODO comment
func (bm BlankModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// TODO comment
func (bm BlankModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return nil
}

// TODO comment
func (bm BlankModule) NewHandler() sdk.Handler {
	panic("NewHandler not implemented")
}

// TODO comment
func (bm BlankModule) QuerierRoute() string {
	return bm.ModuleName
}
