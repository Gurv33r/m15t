package deck

type Hand struct {
	Cards []*Card `json:"cards"`
}

func (h *Hand) Size() int {
	return len(h.Cards)
}
func (h *Hand) Remove(c *Card) {
	index := h.IndexOf(c)
	if index > -1 {
		h.Cards = append(h.Cards[:index], h.Cards[index+1:]...)
	}
}
func (h *Hand) IndexOf(target *Card) int {
	i := -1
	for index, value := range h.Cards {
		if value == target {
			i = index
		}
	}
	return i
}

func (h *Hand) Collect(incoming []*Card) {
	h.Cards = append(h.Cards, incoming...)
}
func (h *Hand) LowestCard() *Card {
	min := h.Cards[0]
	for _, value := range h.Cards {
		if min.Num > value.Num {
			min = value
		}
	}
	return min
}
