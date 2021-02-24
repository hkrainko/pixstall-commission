package model

// true = ascending, false = descending
type CommissionSorter struct {
	Price          *bool
	State          *bool
	CreateTime     *bool
	LastUpdateTime *bool
}
