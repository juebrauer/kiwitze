package kiwitze

import (
	"fmt"
	"math/rand"
)

var witze = []string{
	"Warum gehen Mathematikbücher immer zum Therapeuten? Weil sie zu viele Probleme haben!",
	"Warum nehmen Programmierer immer eine zweite Brille mit? Für den Fall, dass sie einen Bug finden!",
	"Warum können sich Computer nicht gegenseitig beleidigen? Weil sie keine Bytes verschwenden wollen!",
	"Wann ist ein Schritt nach hinten ein Schritt in die richtige Richtung? Beim Rückwärtssalto!",
	"Warum gehen Funktionen nie alleine zur Party? Weil sie immer einen Parameter mitbringen müssen!",
}

func BringMichZumLachen() {
	// Warnung! KI-Humor ist etwas seltsam ...
	fmt.Println(rand.Intn(len(witze)))
}
