package vreme

import "fmt"

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

var postaje map[string]Postaja

func DodajMeritev(oznaka string, temperatura float64) {
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

func IzpisMeritev() {
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

func statistika(oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
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

func IzpisPorocila() {
	for oznaka, postaja := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(oznaka)
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

func DodajPostajo(oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{ime, kraj, []float64{30, 10}}
}
