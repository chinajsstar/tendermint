package main

import (
	"fmt"
	"os"

	amino "github.com/tendermint/go-amino"
	crypto "github.com/chinajsstar/tendermint/crypto"
)

func main() {
	cdc := amino.NewCodec()
	crypto.RegisterAmino(cdc)
	cdc.PrintTypes(os.Stdout)
	fmt.Println("")
}
