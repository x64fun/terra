package util

type OrderDirection string

const (
	OrderDirectionASC  OrderDirection = "ASC"
	OrderDirectionDESC OrderDirection = "DESC"
)

type Order struct {
	Direction OrderDirection
	Field     string
}
