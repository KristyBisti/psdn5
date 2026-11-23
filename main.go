package main

import (
	"fmt"
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func dodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	if temperatura < -50 || temperatura > 50 {
		fmt.Printf("temperatura %.1f izven dovoljenega intervala\n", temperatura)
		return
	}
	var postaja Postaja
	var obstaja bool
	postaja, obstaja = postaje[oznaka]
	if !obstaja {
		fmt.Printf("Postaja %s ne obstaja\n", oznaka)
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

func izpisMeritev(postaje map[string]Postaja) {
	fmt.Printf("MERITVE:\n")
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): [", oznaka, postaja.ime, postaja.kraj)
		for i, meritev := range postaja.meritve {
			if i > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%.1f", meritev)
		}
		fmt.Printf("]\n")
	}

}

func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	var postaja Postaja
	var obstaja bool
	postaja, obstaja = postaje[oznaka]
	if !obstaja || len(postaja.meritve) < 3 {
		return 0, 0, 0, false
	}
	var minTemp float64 = postaja.meritve[0]
	var maxTemp float64 = postaja.meritve[0]
	var averageTemp float64 = 0
	for i := range postaja.meritve {
		meritev := postaja.meritve[i]
		averageTemp += meritev
		if meritev < minTemp {
			minTemp = meritev
		}
		if meritev > maxTemp {
			maxTemp = meritev
		}
	}

	return minTemp, maxTemp, averageTemp / (float64(len(postaja.meritve))), true
}

func izpisPorocila(postaje map[string]Postaja) {
	for oznaka, postaja := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, oznaka)
		fmt.Printf("%s: ", postaja.ime)
		if !ok {
			fmt.Printf("premalo meritev\n")
			continue
		}
		fmt.Printf("min %.1f °C, max %.1f °C, povprečje %.1f °C -> ", najnizja, najvisja, povprecna)
		switch {
		case povprecna >= 25:
			fmt.Printf("Vroče")
		case povprecna >= 15:
			fmt.Printf("Toplo")
		case povprecna >= 5:
			fmt.Printf("Hladno")
		default:
			fmt.Printf("Mrzlo")
		}
		fmt.Printf("\n")

	}
}

func main() {

	postaje := map[string]Postaja{
		"LJ01": {"Ljubljana Sever", "Ljubljana", []float64{10, 15, 20}},
		"LJ02": {"Ljubljana Jug", "Ljubljana", []float64{10, 15, 10}},
		"LJ03": {"Ljubljana Zahod", "Ljubljana", []float64{10, 15, 30}},
		"LJ04": {"Ljubljana Vzhod", "Ljubljana", []float64{10, 15, 0}},
		"MB01": {"Maribor", "Maribor", []float64{20, 30, 40}},
		"KP01": {"Koper", "Koper", []float64{46, 48, 50}},
		"KR01": {"Kranj Vzhod", "Kranj", []float64{-12, -18, 3}},
		"KR02": {"Kranj Zahod", "Kranj", []float64{30, 10}},
	}
	fmt.Println()
	fmt.Println()
	dodajMeritev(postaje, "LJ01", 44)   // Primer pravilnega dodajanja meritve 1
	dodajMeritev(postaje, "LJ02", 46)   // Primer pravilnega dodajanja meritve 2
	dodajMeritev(postaje, "LJ01", 20)   // Primer pravilnega dodajanja meritve 3
	dodajMeritev(postaje, "LJ04", 55)   // Primer dodajanja previsoke temperature
	dodajMeritev(postaje, "MB01", -120) // Primer dodajanja prenizke temperature
	dodajMeritev(postaje, "MB02", 10)   // Primer dodajanja postaji ki ne obstaja

	fmt.Println()
	fmt.Println()
	izpisMeritev(postaje)

	fmt.Println()
	fmt.Println()
	izpisPorocila(postaje)

	fmt.Println()
	fmt.Print(statistika(postaje, "KR01")) // Statistika obstoječega kraja z dovolj meritvami
	fmt.Println()
	fmt.Print(statistika(postaje, "KR02")) // Statistika obstoječega kraja z pemalo ( < 3) meritvami
	fmt.Println()
	fmt.Print(statistika(postaje, "KR03")) // Statistika ne-obstoječega kraja
}
