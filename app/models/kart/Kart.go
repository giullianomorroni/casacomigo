package kart

type (
	Kart struct {
		Casal	string
		Lista	map[int]int
	}
)

func (k *Kart) AddItem(code, quantity int) {
	k.Lista[code] = quantity;
}
