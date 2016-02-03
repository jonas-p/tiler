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

	proj4, err := GeoTIFFProj4Def(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Proj4: %s\n", proj4)
}
