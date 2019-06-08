package main

import (
	"fmt"
	"go-csv-animal/src/api"
	"time"
)

func main() {
	fmt.Print(`test`)
	time.Second
}

func startApp(csv api.CsvReader) {
	gatos := csv.ReadFile("gatos.csv")
	cachorros := csv.ReadFile("cachorros.csv")

}
