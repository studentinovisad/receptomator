package main

import "math/rand"

type Pridev struct {
	mrod string
	zrod string
	srod string
}

var pridevi = [...]Pridev{
	{"sočan", "sočna", "sočno"},
	{"ukusan", "ukusna", "ukusno"},
	{"svež", "sveža", "sveže"},
	{"slan", "slana", "slano"},
	{"sladak", "slatka", "slatko"},
	{"kisel", "kisela", "kiselo"},
	{"pržen", "pržena", "prženo"},
	{"pečen", "pečena", "pečeno"},
	{"sirov", "sirova", "sirovo"},
	{"kuvan", "kuvana", "kuvano"},
	{"pohovan", "pohovana", "pohovano"},
	{"dimljen", "dimljena", "dimljeno"},
	{"začinjen", "začinjena", "začinjeno"},
	{"pikantan", "pikantna", "pikantno"},
	{"ljut", "ljuta", "ljuto"},
	{"blagi", "blaga", "blago"},
	{"hrskav", "hrskava", "hrskavo"},
	{"neukusan", "neukusna", "neukusno"},
	{"aromatičan", "aromatična", "aromatično"},
	{"vruć", "vruća", "vruće"},
	{"hladan", "hladna", "hladno"},
	{"sitan", "sitna", "sitno"},
	{"kremast", "kremasta", "kremasto"},
	{"prirodan", "prirodna", "prirodno"},
	{"gorak", "gorka", "gorko"},
	{"plav", "plava", "plavo"},
	{"crven", "crvena", "crveno"},
	{"žut", "žuta", "žuto"},
	{"živ", "živa", "živo"},
	{"pokvaren", "pokvarena", "pokvareno"},
	{"otrovni", "otrovna", "otrovno"},
}

func randomPridev() Pridev {
	return pridevi[rand.Intn(len(pridevi))]
}
