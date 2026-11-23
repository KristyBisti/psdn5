// Package vreme allows us to create and manage weather stations
package vreme

import "fmt"

// Postaja is a structure with the name, location and meassurements of a weather station
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

var postaje map[string]Postaja = make(map[string]Postaja)
var minTemp float64 = -50
var maxTemp float64 = 50
var stMeritev int = 3

// DodajMeritev checks if the measurement is within minTemp and maxTemp.
// If it is, it is added to the list of measurements, otherwise the error text is displayed.
func DodajMeritev(oznaka string, temperatura float64) {
	if temperatura < minTemp || temperatura > maxTemp {
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

// IzpisMeritev displays all of the saved weather stations, their tag, location, name and meassurements
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
	if !obstaja || len(postaja.meritve) < stMeritev {
		return 0, 0, 0, false
	}
	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	var vsota float64 = 0
	for i := range postaja.meritve {
		meritev := postaja.meritve[i]
		vsota += meritev
		if meritev < najnizja {
			najnizja = meritev
		}
		if meritev > najvisja {
			najvisja = meritev
		}
	}

	return najnizja, najvisja, vsota / (float64(len(postaja.meritve))), true
}

// IzpisPorocila displays the most important information for each of the weather stations.
// It displays the minimum, maxiumum and average temperature as well as the description of the average weather temperature.
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

// DodajPostajo adds a new station with a tag, name and location
func DodajPostajo(oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{ime, kraj, []float64{}}
}

// NastaviNastavitve sets the lowest and highest allowed temperature as well as the minimum meassurements to display statistics
func NastaviNastavitve(newMinTemp float64, newMaxTemp float64, minStMeritev int) {
	minTemp = newMinTemp
	maxTemp = newMaxTemp
	stMeritev = minStMeritev
}
