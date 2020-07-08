package wasptest

import (
	waspapi "github.com/iotaledger/wasp/packages/apilib"
	"github.com/iotaledger/wasp/packages/vm/examples/wasmpoc"
	"testing"
	"time"
)

// sending 5 NOP requests with 1 sec sleep between each
func TestWasmVMSend5Requests1Sec(t *testing.T) {
	// setup
	wasps := setup(t, "test_cluster", "TestWasmVMSend5Requests1Sec")

	err := wasps.ListenToMessages(map[string]int{
		"bootuprec":           6,
		"active_committee":    1,
		"dismissed_committee": 0,
		"request_in":          6,
		"request_out":         7,
		"state":               -1, // must be 6 or 7
		"vmmsg":               -1,
	})
	check(err, t)

	err = PutBootupRecords(wasps)
	check(err, t)

	// number 5 is "Wasm VM PoC program" in cluster.json
	sc := &wasps.SmartContractConfig[5]

	err = Activate1SC(wasps, sc)
	check(err, t)

	err = CreateOrigin1SC(wasps, sc)
	check(err, t)

	reqs := []*waspapi.RequestBlockJson{
		{Address: sc.Address,
			RequestCode: wasmpoc.RequestNOP,
		},
	}
	err = SendRequestsNTimes(wasps, sc.OwnerIndexUtxodb, 5, reqs, 1*time.Second)
	check(err, t)

	wasps.CollectMessages(15 * time.Second)

	if !wasps.Report() {
		t.Fail()
	}
}