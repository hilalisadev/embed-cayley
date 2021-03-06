package main

import (
	"log"

	"github.com/google/cayley"
	_ "github.com/google/cayley/graph/bolt"
)

import "github.com/google/cayley/graph"

func open() {
	path := "./db"
	// Initialize the database
	graph.InitQuadStore("bolt", path, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", path, nil)
	if err != nil {
		log.Fatalln(err)
	}

	p := cayley.StartPath(store, "").Out("type").Is("Person")

	it := p.BuildIterator()
	for cayley.RawNext(it) {
		log.Println(store.NameOf(it.Result()))
	}

}

func main() {
	open()
}
