package proxy

import (
	"bytes"

	"github.com/pkg/errors"

	"github.com/chinajsstar/tendermint/lite"
	certerr "github.com/chinajsstar/tendermint/lite/errors"
	"github.com/chinajsstar/tendermint/types"
)

func ValidateBlockMeta(meta *types.BlockMeta, check lite.Commit) error {
	if meta == nil {
		return errors.New("expecting a non-nil BlockMeta")
	}
	// TODO: check the BlockID??
	return ValidateHeader(meta.Header, check)
}

func ValidateBlock(meta *types.Block, check lite.Commit) error {
	if meta == nil {
		return errors.New("expecting a non-nil Block")
	}
	err := ValidateHeader(meta.Header, check)
	if err != nil {
		return err
	}
	if !bytes.Equal(meta.Data.Hash(), meta.Header.DataHash) {
		return errors.New("Data hash doesn't match header")
	}
	return nil
}

func ValidateHeader(head *types.Header, check lite.Commit) error {
	if head == nil {
		return errors.New("expecting a non-nil Header")
	}
	// make sure they are for the same height (obvious fail)
	if head.Height != check.Height() {
		return certerr.ErrHeightMismatch(head.Height, check.Height())
	}
	// check if they are equal by using hashes
	chead := check.Header
	if !bytes.Equal(head.Hash(), chead.Hash()) {
		return errors.New("Headers don't match")
	}
	return nil
}
