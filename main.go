package main

import (
	"flag"
	"fmt"
	"regexp"
	"sync"

	"github.com/xssnick/tonutils-go/ton/wallet"
)

type result struct {
	seed   []string
	adress string
}

func worker(r *regexp.Regexp, results chan<- result) {
	var (
		tonapi wallet.TonAPI = nil
	)
	for {
		words := wallet.NewSeed()
		w, err := wallet.FromSeed(tonapi, words, wallet.V3)
		if err != nil {
			fmt.Printf("[ERROR]FromSeed: %s\n", err)
			continue
		}
		adress := w.Address()
		matched := r.MatchString(adress.String())
		if matched {
			res := result{
				seed:   words,
				adress: adress.String(),
			}
			results <- res
		}
	}
}

func PrinterWorker(results <-chan result) {
	for res := range results {
		fmt.Printf("[MATCH] %v -> %s\n", res.seed, res.adress)
	}
}

func main() {
	rr := flag.String("r", `^`, "Regexp pattern")
	icase := flag.Bool("i", false, "Ignore letter case")
	workers := flag.Int("w", 1, "Number of workers")
	flag.Parse()
	if *icase {
		*rr = fmt.Sprintf("%s%s", "(?i)", *rr)
	}

	r := regexp.MustCompile(*rr)
	fmt.Printf("[INFO] Regex compiled successfully: %v\n", *rr)

	results := make(chan result, *workers)
	go PrinterWorker(results)
	var wg sync.WaitGroup
	wg.Add(*workers)
	for w := 1; w <= *workers; w++ {
		go func(r *regexp.Regexp, results chan<- result) {
			worker(r, results)
			wg.Done()
		}(r, results)
		fmt.Printf("[INFO] Worker #%v started\n", w)
	}
	wg.Wait()
}
