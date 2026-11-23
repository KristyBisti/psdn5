package main

import (
	"fmt"
	"naloga5/vreme"
)

func main() {

	vreme.DodajPostajo("LJ01", "Ljubljana Sever", "Ljubljana")
	vreme.DodajPostajo("LJ02", "Ljubljana Jug", "Ljubljana")
	vreme.DodajPostajo("LJ03", "Ljubljana Zahod", "Ljubljana")
	vreme.DodajPostajo("LJ04", "Ljubljana Vzhod", "Ljubljana")
	vreme.DodajPostajo("MB01", "Maribor", "Maribor")
	vreme.DodajPostajo("KP01", "Koper", "Koper")
	vreme.DodajPostajo("KR01", "Kranj Vzhod", "Kranj")
	vreme.DodajPostajo("KR02", "Kranj Zahod", "Kranj")
	fmt.Println()
	fmt.Println()

	vreme.DodajMeritev("LJ01", 44)   // Primer pravilnega dodajanja meritve 1
	vreme.DodajMeritev("LJ02", 46)   // Primer pravilnega dodajanja meritve 2
	vreme.DodajMeritev("LJ01", 20)   // Primer pravilnega dodajanja meritve 3
	vreme.DodajMeritev("LJ04", 55)   // Primer dodajanja previsoke temperature
	vreme.DodajMeritev("MB01", -120) // Primer dodajanja prenizke temperature
	vreme.DodajMeritev("MB02", 10)   // Primer dodajanja postaji ki ne obstaja

	fmt.Println()
	fmt.Println()
	vreme.IzpisMeritev()

	fmt.Println()
	fmt.Println()
	vreme.IzpisPorocila()

}
