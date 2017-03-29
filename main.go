package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var nowstr string
	flag.StringVar(&nowstr, "now", time.Now().UTC().Format(time.RFC3339), "now (formatted as RFC3339)")
	flag.Parse()

	now, err := time.Parse(time.RFC3339, nowstr)
	if err != nil {
		log.Fatal(err)
	}
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		txt := scn.Text()
		then, err := time.Parse(time.RFC3339, txt)
		if err != nil {
			log.Printf("could not parse time from %s\n", txt)
			continue
		}
		if then.Equal(now) {
			fmt.Println("equal")
		} else if then.Before(now) {
			fmt.Println("before")
		} else {
			fmt.Println("after")
		}
	}
	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}
}
