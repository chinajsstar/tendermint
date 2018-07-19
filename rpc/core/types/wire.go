package core_types

import (
	"github.com/tendermint/go-amino"
	"github.com/chinajsstar/tendermint/crypto"
	"github.com/chinajsstar/tendermint/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterEvidences(cdc)
	crypto.RegisterAmino(cdc)
}
