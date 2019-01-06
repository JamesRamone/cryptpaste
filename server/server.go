package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/uuid"

	"github.com/go-chi/chi"

	bolt "go.etcd.io/bbolt"
)

type server struct {
	db     *bolt.DB
	router *chi.Mux
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("web/dist")).ServeHTTP(w, r)
	}
}

func (s *server) handleAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello API")
	}
}

func (s *server) handleGetPad() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			respondHTTPErr(w, r, 404)
			return
		}
		s.db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("CryptPad"))
			v := b.Get(id.Bytes())
			if v == nil {
				respondHTTPErr(w, r, 404)
				return errors.New("Not found")
			}
			p := &Pad{}
			err := json.Unmarshal(v, p)
			if err != nil {
				respondHTTPErr(w, r, http.StatusInternalServerError)
				return errors.New("Error")
			}
			respond(w, r, 200, p)
			return nil
		})
	}
}

func (s *server) handleCreatePad() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := NewPad()
		err := json.NewDecoder(r.Body).Decode(p)
		if err != nil {
			respondHTTPErr(w, r, http.StatusBadRequest)
			log.Println(err)
			return
		}
		log.Printf("%s ::: %s \n", p.ID.String(), p.Data)
		enc, err := json.Marshal(p)
		if err != nil {
			respondHTTPErr(w, r, http.StatusInternalServerError)
			return
		}
		s.db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("CryptPad"))
			err := b.Put(p.ID.Bytes(), enc)
			if err != nil {
				respondHTTPErr(w, r, http.StatusInternalServerError)
				return errors.New("Error storing in database")
			}
			log.Printf("created pad %s \n", p.ID.String())
			respond(w, r, 200, p.ID)
			return nil
		})
	}
}
