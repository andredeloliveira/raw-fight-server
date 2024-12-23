package fighter

type Stats struct {
	Attack  int
	Defense int
	Dodge   int
}

type Fighter struct {
	Name      string
	HP        int
	Stamina   int
	BaseStats Stats
}

type ActionType string

const (
	Kick      ActionType = "kick"
	Punch     ActionType = "punch"
	Block     ActionType = "block"
	DodgeBack ActionType = "dodge_back"
	Jump      ActionType = "jump"
	Duck      ActionType = "duck"
)
