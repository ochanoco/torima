package tee

import (
	"fmt"
	"reflect"

	"github.com/ochanoco/proxy/core"
)

func verifyDB(proxy *core.OchanocoProxy) error {
	fmt.Println("logger is verifing DB...")

	serviceLogs, err := proxy.Database.SelectAllServiceLogs()

	if err != nil {
		return err
	}

	if len(serviceLogs) == 0 {
		fmt.Printf("no record!!!!\n")
		return nil
	}

	var previous core.HashChain

	for _, log := range serviceLogs[1:] {
		previous = core.HashChain{
			Hash: core.CalcHashChain(log, &previous),
		}
	}

	last, err := proxy.Database.FindLastHashChain()

	if err != nil {
		return err
	}

	if !reflect.DeepEqual(last.Hash, previous.Hash) {
		return fmt.Errorf("verifyDB: hashchain is invalid (%v, %v)", last.Hash, previous.Hash)
	}

	fmt.Printf("DB is valid (%v, %v)\n", last.Hash, previous.Hash)

	return nil
}
