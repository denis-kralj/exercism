package dndcharacter

import "math/rand"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	if score % 2 != 0 {
        score --
    }

    return (score/2)-5
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := [4]int{}

    rolls[0] = rand.Intn(6) + 1
    rolls[1] = rand.Intn(6) + 1
    rolls[2] = rand.Intn(6) + 1
    rolls[3] = rand.Intn(6) + 1

    smallest := rolls[0]

    for i := 1; i < len(rolls); i++ {
        if smallest > rolls[i] {
            smallest = rolls[i]
        }
    }

    return rolls[0] + rolls[1] + rolls[2] + rolls[3] - smallest
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	con := Ability()
    hp := 10 + Modifier(con)

    return Character {
        Ability(),
        Ability(),
        con,
        Ability(),
        Ability(),
        Ability(),
        hp,
    }
}
