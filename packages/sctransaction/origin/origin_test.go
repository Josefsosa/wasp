package origin

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	valuetransaction "github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/transaction"
	"github.com/iotaledger/goshimmer/dapps/waspconn/packages/utxodb"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/stretchr/testify/assert"
)

const (
	scAddrStr = "behcN8An9CQDU3rcHSUucamZreVWfJepmiuBoSpomfi1"
	dscr      = "test test sc"
)

var nodeLocations = []string{"127.0.0.1:4000", "127.0.0.1:4001", "127.0.0.1:4002", "127.0.0.1:4003"}

func TestReadWrite(t *testing.T) {
	scAddr, err := address.FromBase58(scAddrStr)
	assert.NoError(t, err)

	ownerAddress := utxodb.GetAddress(1)
	tx, err := NewOriginTransaction(NewOriginTransactionParams{
		Address:              scAddr,
		OwnerSignatureScheme: utxodb.GetSigScheme(ownerAddress),
		AllInputs:            utxodb.GetAddressOutputs(ownerAddress),
		InputColor:           balance.ColorIOTA,
		ProgramHash:          *hashing.HashStrings(dscr),
	})
	assert.NoError(t, err)
	t.Logf("created transaction txid = %s", tx.ID().String())

	data := tx.Bytes()
	vtx, _, err := valuetransaction.FromBytes(data)
	assert.NoError(t, err)

	txback, err := sctransaction.ParseValueTransaction(vtx)
	assert.NoError(t, err)

	assert.EqualValues(t, tx.ID(), txback.ID())
}