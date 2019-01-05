package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	bolt "go.etcd.io/bbolt"
)

func main() {

	var addr = flag.String("addr", ":3000", "address to host the server on")
	flag.Parse()

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("CryptPad"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	s := &server{
		db:     db,
		router: chi.NewRouter(),
	}
	s.routes()

	log.Printf("Starting server on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, s.router))
}
