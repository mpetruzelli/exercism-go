package account

import "sync"

const (
	testVersion = 1
)

type Account struct {
	balance int64
	opened  bool
	mux     sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, opened: true}
}

func (a *Account) Balance() (int64, bool) {
	a.mux.Lock()
	if a.opened == false {
		balance := a.balance
		a.mux.Unlock()
		return balance, false
	}
	balance := a.balance
	a.mux.Unlock()
	return balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	if a.opened == false {
		balance := a.balance
		a.mux.Unlock()
		return balance, false
	}
	payout := a.balance
	a.balance -= payout
	a.opened = false
	a.mux.Unlock()
	return payout, true
}
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	if a.balance+amount < 0 || a.opened == false {
		balance := a.balance
		a.mux.Unlock()
		return balance, false
	}
	a.balance += amount
	balance := a.balance
	a.mux.Unlock()
	return balance, true
}
