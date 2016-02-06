package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s FILE\n", os.Args[0])
		return
	}

	rep, err := GeoTIFFRepresentation(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rep)
}
