package buchfuehrung

type Accounts []Account

type Account struct {
	ID           int    `sql:"AUTO_INCREMENT"`
	Name         string `sql:"not null;unique"`
	Comment      string
	Transactions Transactions
	Type         AccountType
}

func NewAccount(name, comment, accounttype string) *Account {
	account := new(Account)
	account.Name = name
	account.Comment = comment
	account.Type = ParseAccountType(accounttype)

	return account
}

//go:generate stringer -type=AccountType
type AccountType byte

const (
	Unkown AccountType = iota
	GiroKonto
)

func ParseAccountType(accounttype string) AccountType {
	switch accounttype {
	case "GiroKonto":
		return GiroKonto

	default:
		return Unkown
	}
}
