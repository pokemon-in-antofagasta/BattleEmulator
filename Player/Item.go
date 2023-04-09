package player

type Item struct {
	Name   string
	Effect string
}

func (p *Pokemon) Use(item Item) {

	/*
		Para facilitar esto se implementa convención de guardado del efecto:
		El string será un conjunto de pequeños efectos seguidos de un punto con estas reglas
		Heal -> "Cure X hp." (Donde X es la cantidad de puntos expresado en un int. un int = 0 significa curar full)
		Estados efimeros -> "Clean X." (Donde X es el estado que "limpia")
		Estados persistentes -> "Cure X." (Donde X es el estado que "limpia". Si X="all" entonces cura todo)


	*/

}
