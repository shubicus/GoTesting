package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"log"
	"os"
	"time"
)

func main() {
	rate := vegeta.Rate{Freq: 3, Per: time.Second}
	duration := 4 * time.Second
	url := "https://google.com/"
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    url,
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Vegeta-Test!") {
		metrics.Add(res)
	}
	metrics.Close()

	name := "vegeta_histogram.txt"
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("Can not create a file %s: %v ", name, err)
		return
	}
	defer file.Close()

	reporter := vegeta.NewHDRHistogramPlotReporter(&metrics)
	err = reporter(file)
	if err != nil {
		log.Fatalf("Cannot create report: %v", err)
	}

	if _, err := os.Stat(name); os.IsNotExist(err) {
		log.Fatalf("Report file %s does not exist", name)
	}

	fmt.Printf("Report is saved to file [%s]\nYou can observe it at https://hdrhistogram.github.io/HdrHistogram/plotFiles.html", name)
}
