package buchfuehrung

import (
	"math/big"
	"time"
)

type Transactions []Transaction

type Transaction struct {
	ID        int `sql:"AUTO_INCREMENT"`
	Category  Category
	Comment   Comment
	Inflow    *big.Float
	Outflow   *big.Float
	Payee     Payee
	TimeStamp time.Time
}
