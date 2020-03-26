package main

import "log"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	myHealth := 50
	myMana := 500

	bossHealth := 51
	bossDamage := 9

	minMana := 99999

	var f func(bool, int, int, int, int, int, int, int)
	f = func(
		myTurn bool,
		manaSpent int,
		myHealth int,
		myMana int,
		bossHealth int,
		shieldTurns int,
		poisonTurns int,
		rechargeTurns int,
	) {
		if manaSpent > minMana {
			return
		}

		armor := 0
		if myHealth <= 0 {
			return
		}

		if poisonTurns > 0 {
			bossHealth -= 3
			poisonTurns--
		}
		if rechargeTurns > 0 {
			myMana += 101
			rechargeTurns--
		}
		if shieldTurns > 0 {
			armor = 7
			shieldTurns--
		}

		if bossHealth <= 0 {
			if manaSpent < minMana {
				minMana = manaSpent
			}
			return
		}

		if !myTurn {
			f(!myTurn, manaSpent, myHealth-(bossDamage-armor), myMana, bossHealth, shieldTurns, poisonTurns, rechargeTurns)
			return
		}

		if myMana < 53 {
			return
		}

		// Magic Missile
		f(!myTurn, manaSpent+53, myHealth, myMana-53, bossHealth-4, shieldTurns, poisonTurns, rechargeTurns)
		// Drain
		if myMana >= 73 {
			f(!myTurn, manaSpent+73, myHealth+2, myMana-73, bossHealth-2, shieldTurns, poisonTurns, rechargeTurns)
		}
		// Shield
		if myMana >= 113 && shieldTurns == 0 {
			f(!myTurn, manaSpent+113, myHealth, myMana-113, bossHealth, 6, poisonTurns, rechargeTurns)
		}
		// Poison
		if myMana >= 173 && poisonTurns == 0 {
			f(!myTurn, manaSpent+173, myHealth, myMana-173, bossHealth, shieldTurns, 6, rechargeTurns)
		}
		// Recharge
		if myMana >= 229 && rechargeTurns == 0 {
			f(!myTurn, manaSpent+229, myHealth, myMana-229, bossHealth, shieldTurns, poisonTurns, 5)
		}
	}

	f(true, 0, myHealth, myMana, bossHealth, 0, 0, 0)

	log.Printf("Min mana is %d", minMana)
}
