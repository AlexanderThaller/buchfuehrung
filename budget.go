package buchfuehrung

import (
	"math/big"
	"time"
)

type Budgets []Budget

type Budget struct {
	ID        int `sql:"AUTO_INCREMENT"`
	Name      string
	Comment   Comment
	StartTime time.Time
	EndTime   time.Time
	Groups    []BudgetGroup
}

type BudgetGroup struct {
	ID         int `sql:"AUTO_INCREMENT"`
	Name       string
	Categories []BudgetCategory
}

type BudgetCategory struct {
	ID           int `sql:"AUTO_INCREMENT"`
	Category     Category
	Budgeted     *big.Float
	Transactions Transactions
}
