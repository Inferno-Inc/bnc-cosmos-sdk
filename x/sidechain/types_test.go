package sidechain

import (
	"encoding/hex"
	"fmt"
	"github.com/magiconair/properties/assert"
	"math/big"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
)

func Test_EncodePackageHeader(t *testing.T){
	bz:=EncodePackageHeader(types.SynCrossChainPackageType, *big.NewInt(10000000000000000))
	fmt.Println(hex.EncodeToString(bz))
	assert.Equal(t, hex.EncodeToString(bz), "hex.EncodeToString(bz)")
}