package main

import (
	"fmt"
	"math/rand"
)

func randomPridevNazivBroj() string {
	jelo := randomJelo()
	pridev := randomPridev()
	broj := rand.Intn(100)

	switch jelo.rod {
	case M:
		return fmt.Sprintf("%s%s%d", pridev.mrod, jelo.naziv, broj)
	case Z:
		return fmt.Sprintf("%s%s%d", pridev.zrod, jelo.naziv, broj)
	case S:
		return fmt.Sprintf("%s%s%d", pridev.srod, jelo.naziv, broj)
	}

	panic("Switch statement is incorrect!")
}
