package player

type Movement struct {
	Name    string
	Dmg     int
	Dmgtype string
}

func (f *Pokemon) AttackWith(movement Movement) Movement {

	return movement

}

func (attackingMovement Movement) To(enemyPokemon *Pokemon) *Pokemon {

	switch attackingMovement.Dmgtype {

	case "Normal":

		if attackingMovement.Dmg < enemyPokemon.CurrentDefense {

			enemyPokemon.CurrentHP--
			return enemyPokemon

		} else {

			enemyPokemon.CurrentHP = enemyPokemon.CurrentHP - (attackingMovement.Dmg - enemyPokemon.CurrentDefense)
			return enemyPokemon
		}

	case "Especial":

		if attackingMovement.Dmg < enemyPokemon.CurrentSpDefense {

			enemyPokemon.CurrentHP--
			return enemyPokemon

		} else {

			enemyPokemon.CurrentHP = enemyPokemon.CurrentHP - (attackingMovement.Dmg - enemyPokemon.CurrentSpDefense)
			return enemyPokemon
		}

	}

	return enemyPokemon

}
