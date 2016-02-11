package buchfuehrung

type Payee struct {
	ID   int `sql:"AUTO_INCREMENT"`
	Name string
}
