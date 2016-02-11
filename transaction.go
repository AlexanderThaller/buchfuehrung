package buchfuehrung

import (
	"math/big"
	"time"
)

type Transactions []Transaction

type Transaction struct {
	Category  Category
	Comment   Comment
	Inflow    *big.Float
	Outflow   *big.Float
	Payee     Payee
	TimeStamp time.Time
}
