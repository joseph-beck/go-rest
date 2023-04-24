package firebase

import (
	"errors"
	"fmt"
	"log"
	"rest/pkg/firebase/firestore"
	"testing"
)

type user struct {
	Name   string
	Age    int
	Height float32
}

func TestParse(t *testing.T) {
	u := user{
		Name:   "dave",
		Age:    27,
		Height: 1.7,
	}

	m, err := firestore.MapStruct(u)
	if err != nil {
		log.Fatalln(err)
	}
	if len(m) != 3 {
		mp := "map\n"
		for field, val := range m {
			mp += fmt.Sprint("[", field, ": ", val, "]\n")
		}
		log.Fatalln(errors.New("empty map created"), ":", mp)
	}
}
