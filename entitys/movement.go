package entitys

type Movement struct {
	Name    string
	Dmg     int
	Dmgtype string
}

func (f *Fighter) AttackWith(movement Movement) Movement {

	return movement

}

func (attackingMovement Movement) To(enemyFighter *Fighter) *Fighter {

	switch attackingMovement.Dmgtype {

	case "Normal":

		if attackingMovement.Dmg < enemyFighter.Normaldfs {

			enemyFighter.Life--
			return enemyFighter

		} else {

			enemyFighter.Life = enemyFighter.Life - (attackingMovement.Dmg - enemyFighter.Normaldfs)
			return enemyFighter
		}

	case "Especial":

		if attackingMovement.Dmg < enemyFighter.Specialdfs {

			enemyFighter.Life--
			return enemyFighter

		} else {

			enemyFighter.Life = enemyFighter.Life - (attackingMovement.Dmg - enemyFighter.Specialdfs)
			return enemyFighter
		}

	}

	return enemyFighter

}
