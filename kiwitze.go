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
	"Wie nennt man einen Bumerang, der nicht zurückkommt? Einen Stock!",
	"Warum sitzen Ameisen nie in der Kirche? Weil sie Insekten sind!",
	"Was ist der Unterschied zwischen einem Schneeball und einem klugen Gedanken? Beide sind kalt, aber einer schmilzt sofort!",
	"Was haben Wolken und Programmierer gemeinsam? Beide nutzen gerne die Cloud!",
	"Warum bekommen Computer immer so schnell Fieber? Weil sie keine Antiviren-Software haben!",
}

func BringMichZumLachen() {
	// Warnung! KI-Humor ist etwas seltsam ...
	fmt.Println(witze[rand.Intn(len(witze))])
}
