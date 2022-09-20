package keeper

import (
	"github.com/gadost/gaia/v7/x/icaauth/types"
)

var _ types.QueryServer = Keeper{}
