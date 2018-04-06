package main

import (
	"fmt"
	"os"
	"time"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi"
	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/kml"

	"bitbucket.org/friasdesign/pfetcher/internal/clcitybusapi/client"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s in %v\n", what, time.Since(start))
	}
}

func main() {
	defer elapsed("Finished")()
	dumpPath := "testdata"

	if _, err := os.Stat(dumpPath); os.IsNotExist(err) {
		panic("testdata/ folder doesn't exist")
	}
	cli := client.NewClient(nil, "testdata")

	// Create empresa
	fmt.Println("Creating empresa ...")
	emp := clcitybusapi.NewEmpresa(355)
	fmt.Println("Creating empresa ... DONE!")

	// Fetch lineas
	fmt.Println("Fetching lineas ...")
	lin, err := cli.LineaService().LineasPorEmpresa(emp)
	if err != nil {
		panic(err)
	}
	emp.Lineas = lin
	fmt.Println("Fetching lineas ... DONE!")

	// Fetch paradas
	fmt.Println("Fetching paradas ...")
	par, err := cli.ParadaService().ParadasPorEmpresa(emp)
	if err != nil {
		panic(err)
	}
	emp.Paradas = par
	fmt.Println("Fetching paradas ... DONE!")

	fmt.Println("Generating KML file ...")
	kml.Generate(emp, "city_bus.kml")
	if err != nil {
		panic(err)
	}
	fmt.Println("Generating KML file ... DONE!")
}
