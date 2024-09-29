package service

type MineralIface interface {
	IsMatchAssignValue(text string) bool
	AssignValue(text string) error
	GetValue(currencies, mineral string) (Mineral, error)
}
