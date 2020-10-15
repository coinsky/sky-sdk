package bancorlite

import (
	"github.com/coinsky/sky-sdk/modules/bancorlite/internal/keepers"
	"github.com/coinsky/sky-sdk/modules/bancorlite/internal/types"
)

const (
	StoreKey   = types.StoreKey
	ModuleName = types.ModuleName
)

var (
	NewBaseKeeper       = keepers.NewKeeper
	NewBancorInfoKeeper = keepers.NewBancorInfoKeeper
	DefaultParams       = types.DefaultParams
	ModuleCdc           = types.ModuleCdc
)

type (
	Keeper                     = keepers.Keeper
	BancorInfo                 = keepers.BancorInfo
	MsgBancorTradeInfoForKafka = types.MsgBancorTradeInfoForKafka
	MsgBancorInfoForKafka      = types.MsgBancorInfoForKafka
	MsgBancorInit              = types.MsgBancorInit
	MsgBancorTrade             = types.MsgBancorTrade
	MsgBancorCancel            = types.MsgBancorCancel
)
