package deck

type Player struct {
	Turn  bool  `json: "turn"`
	Hand  *Hand `json: "hand"`
	Pen   *Hand `json: "pen"`
	Final *Hand `json: "final"`
}

func NewPlayer() *Player {
	return &Player{Turn: false, Hand: &Hand{}, Pen: &Hand{}, Final: &Hand{}}
}
func IsWinner(p *Player) bool {
	return len(p.Hand) == 0 && len(p.Pen) == 0 && len(p.Final) == 0
}
