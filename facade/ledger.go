package facade

import "fmt"

type Ledger struct {
}

func (l *Ledger) makeEntry(accountId, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountId, txnType, amount)
	return
}
