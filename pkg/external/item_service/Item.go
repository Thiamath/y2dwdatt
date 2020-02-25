package item_service

type ItemLabel int

const (
	Book ItemLabel = iota + 1
	Food
	Meds
)

type Item struct {
	Name  string
	Label ItemLabel
}
