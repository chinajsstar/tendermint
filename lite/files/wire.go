package files

import (
	"github.com/tendermint/go-amino"
	"github.com/chinajsstar/tendermint/crypto"
)

var cdc = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
}
