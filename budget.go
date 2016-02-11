package buchfuehrung

import (
	"math/big"
	"time"
)

type Budgets []Budget

type Budget struct {
	Name      string
	Comment   Comment
	StartTime time.Time
	EndTime   time.Time
	Groups    []BudgetGroups
}

type BudgetGroups struct {
	Name       string
	Categories []BudgetCategory
}

type BudgetCategory struct {
	Category     Category
	Budgeted     *big.Float
	Transactions Transactions
}
