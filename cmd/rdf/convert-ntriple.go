package main

import (
	"io"
	"log"
	"os"

	"github.com/knakk/rdf"
)

const data = "data/rdf/"

func main() {
	in, err := os.Open(data + "PERSEE_hom_collection_2019-04-30.rdf")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(data + "test.nt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	dec := rdf.NewTripleDecoder(in, rdf.RDFXML)

	enc := rdf.NewTripleEncoder(out, rdf.NTriples)
	defer enc.Close()

	for triple, err := dec.Decode(); err != io.EOF; triple, err = dec.Decode() {
		if err != nil {
			log.Fatal(err)
		}
		err = enc.Encode(triple)
		if err != nil {
			log.Fatal(err)
		}
	}
}
