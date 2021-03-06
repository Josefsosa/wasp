package frcmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/tools/wwallet/sc/fr"
	"github.com/iotaledger/wasp/tools/wwallet/util"
	"github.com/iotaledger/wasp/tools/wwallet/wallet"
)

func adminCmd(args []string) {
	if len(args) < 1 {
		adminUsage()
	}

	switch args[0] {
	case "deploy":
		check(fr.Config.Deploy(wallet.Load().SignatureScheme()))

	case "set-period":
		if len(args) != 2 {
			fr.Config.PrintUsage("admin set-period <seconds>")
			os.Exit(1)
		}
		s, err := strconv.Atoi(args[1])
		check(err)

		util.WithSCRequest(fr.Config, func() (*sctransaction.Transaction, error) {
			return fr.Client().SetPeriod(s)
		})

	default:
		adminUsage()
	}
}

func adminUsage() {
	fmt.Printf("Usage: %s fr admin [deploy|set-period <seconds>]\n", os.Args[0])
	os.Exit(1)
}
