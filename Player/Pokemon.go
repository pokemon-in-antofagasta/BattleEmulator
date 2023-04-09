package player

type Pokemon struct {
	Name  string
	Phase int
	Stats
	Moveset []Movement
	Item    Item
}

type Stats struct {
	IntialHP,
	CurrentHP,
	IntialSpeed,
	CurrentSpeed,
	IntialAttack,
	CurrentAttack,
	IntialSpAttack,
	CurrentSpAttack,
	IntialDefense,
	CurrentDefense,
	IntialSpDefense,
	CurrentSpDefense int
}
