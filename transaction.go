package buchfuehrung

import "time"

type Transactions []Transaction

type Transaction struct {
	ID        int `sql:"AUTO_INCREMENT"`
	Category  Category
	Comment   string
	Inflow    float64
	Outflow   float64
	Payee     Payee
	TimeStamp time.Time
}
