package main

import "log"

type item struct {
	cost   int
	damage int
	armor  int
}

var (
	weapons = []item{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}
	armor = []item{
		{0, 0, 0},
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}
	rings = []item{
		{0, 0, 0},
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}
)

func main() {
	bossHealth := 109
	bossDamage := 8
	bossArmor := 2

	maxCost := 0
	for _, w := range weapons {
		for _, a := range armor {
			for i, r1 := range rings {
				for _, r2 := range rings[i+1:] {
					myDamage := w.damage + r1.damage + r2.damage
					myArmor := a.armor + r1.armor + r2.armor
					cost := w.cost + a.cost + r1.cost + r2.cost

					myDamage = myDamage - bossArmor
					if myDamage <= 0 {
						myDamage = 1
					}
					bDamage := bossDamage - myArmor
					if bDamage <= 0 {
						bDamage = 1
					}

					if (bossHealth+myDamage-1)/myDamage > (100+bDamage-1)/bDamage && cost > maxCost {
						maxCost = cost
					}
				}
			}
		}
	}

	log.Printf("Max cost is %d", maxCost)
}
