package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/KristyBisti/psdn5/vreme"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "vremenskaPostaja",
		Usage: "Postaja meri temperaturo",
		Flags: []cli.Flag{
			&cli.FloatFlag{Name: "minTemp", Value: -50, Usage: "najnižja dovoljena temperatura meritve"},
			&cli.FloatFlag{Name: "maxTemp", Value: 50, Usage: "najvišja dovoljena temperatura meritve"},
			&cli.IntFlag{Name: "stMeritev", Value: 3, Usage: "najmanjše število meritev za statistiko"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			vreme.NastaviNastavitve(minTemp, maxTemp, stMeritev)

			simulirajPostaje()

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

func simulirajPostaje() {
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

	vreme.DodajMeritev("LJ01", -10)
	vreme.DodajMeritev("LJ01", -20)
	vreme.DodajMeritev("LJ01", -80)
	vreme.DodajMeritev("LJ02", 10)
	vreme.DodajMeritev("LJ02", 10)
	vreme.DodajMeritev("LJ02", 10)
	vreme.DodajMeritev("LJ03", 16)
	vreme.DodajMeritev("LJ03", 17)
	vreme.DodajMeritev("LJ03", 19)
	fmt.Println()
	fmt.Println()
	vreme.IzpisMeritev()

	fmt.Println()
	fmt.Println()
	vreme.IzpisPorocila()
}
