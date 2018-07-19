package blockchain

import (
	"github.com/tendermint/go-amino"
	"github.com/chinajsstar/tendermint/crypto"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	crypto.RegisterAmino(cdc)
}
