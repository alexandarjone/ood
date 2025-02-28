type FinancialInstitution interface {
	GetID() int64
	Deposit(pin string, account Account, amount float64) error
	Withdraw()
	Transfer()
}

type ATM interface {
	Deposit()
	Withdraw()
	Transfer()
}

type atm struct {
	financialInstitution map[int64]FinancialInstitution
}

func (a *atm) Deposit(pin string, account Account, amount float64) {
	f := financialInstitution[account.GetFinancialInstitutionID()]
	f.Deposit(pin, account, amount)
}

type Account interface {
	GetFinancialInstitutionID() int64
}