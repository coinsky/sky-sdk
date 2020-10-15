package comment

import (
	"github.com/coinsky/sky-sdk/modules/comment/internal/keepers"
	"github.com/coinsky/sky-sdk/modules/comment/internal/types"
)

const (
	StoreKey   = types.StoreKey
	ModuleName = types.ModuleName
)

var (
	NewBaseKeeper = keepers.NewKeeper
)

type (
	Keeper          = keepers.Keeper
	TokenComment    = types.TokenComment
	CommentRef      = types.CommentRef
	MsgCommentToken = types.MsgCommentToken
)
