package main

import (
	"math/rand"
)

type Jelo struct {
	naziv string
	rod   Rod
}

var jela = [...]Jelo{
	{"ćevap", M},
	{"sarma", Z},
	{"pita", Z},
	{"proja", Z},
	{"čorba", Z},
	{"pečenje", S},
	{"ajvar", M},
	{"paprikaš", M},
	{"pasulj", M},
	{"musaka", Z},
	{"ćufta", Z},
	{"jagnjetina", Z},
	{"salata", Z},
	{"kulen", M},
	{"pršuta", Z},
	{"pljeskavica", Z},
	{"ražnjić", M},
	{"kompot", M},
	{"teletina", Z},
	{"karfiol", M},
	{"ravioli", M},
	{"baklava", Z},
	{"roštilj", M},
	{"riba", Z},
	{"piletina", Z},
	{"kesten", M},
	{"knedla", Z},
	{"supa", Z},
	{"pečurka", Z},
	{"jaje", S},
	{"šnicla", Z},
	{"soja", Z},
	{"pasta", Z},
	{"krastavac", M},
	{"krastavčić", M},
	{"tartuf", M},
	{"omlet", M},
	{"sir", M},
	{"susam", M},
	{"sendvič", M},
	{"rolnica", Z},
	{"tuna", Z},
	{"začin", M},
	{"mleko", S},
	{"brašno", S},
	{"gljiva", Z},
	{"šunka", Z},
	{"rakija", Z},
	{"burek", M},
	{"patlidžan", M},
	{"jabuka", Z},
	{"paradajz", M},
	{"krompir", M},
	{"pire", M},
	{"tikvica", Z},
	{"urnebes", M},
	{"šiš", M},
	{"indeks", M},
	{"banana", Z},
	{"lepinja", Z},
	{"palačinka", Z},
}

func randomJelo() Jelo {
	return jela[rand.Intn(len(jela))]
}
