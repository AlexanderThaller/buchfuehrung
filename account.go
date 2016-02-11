package buchfuehrung

type Accounts []Account

type Account struct {
	ID           int `sql:"AUTO_INCREMENT"`
	Name         string
	Comment      string
	Transactions Transactions
	Type         AccountType
}

//go:generate stringer -type=AccountType
type AccountType byte

const (
	Unkown AccountType = iota
	GiroKonto
)
