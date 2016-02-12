package buchfuehrung

import "time"

type Budgets []Budget

type Budget struct {
	ID        int `sql:"AUTO_INCREMENT"`
	Name      string
	Comment   string
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
	Budgeted     float64
	Transactions Transactions
}
