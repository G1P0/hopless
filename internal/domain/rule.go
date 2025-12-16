package domain

type Rule struct {
	Src   Node
	Dst   Node
	Allow bool
}
