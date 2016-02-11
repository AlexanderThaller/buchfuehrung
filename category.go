package buchfuehrung

type Categories []Category

type Category struct {
	ID   int `sql:"AUTO_INCREMENT"`
	Name string
}
