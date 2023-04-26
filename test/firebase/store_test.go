package firebase

import (
	"context"
	"log"
	"rest/internal/feeder"
	fs "rest/pkg/firebase/firestore"

	"testing"

	"github.com/stretchr/testify/assert"
)

const conf = "../../conf/service-acc.json"

var s *fs.Store

func TestConn(t *testing.T) {
	s = fs.NewStore(conf, "test")
	defer s.Close()
}

func TestAdd(t *testing.T) {
	s = fs.NewStore(conf, "test")
	defer s.Close()

	ctx := context.Background()
	err := s.Add(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestAddStruct(t *testing.T) {
	s = fs.NewStore(conf, "test")
	defer s.Close()

	u := feeder.Item{
		Name: "davide",
		Data: "something about davide",
	}

	ctx := context.Background()
	err := s.AddStruct(u, ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRead(t *testing.T) {
	s = fs.NewStore(conf, "test")
	defer s.Close()

	ctx := context.Background()
	err := s.Read(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestReadInto(t *testing.T) {
	s = fs.NewStore(conf, "test")
	defer s.Close()
	ctx := context.Background()

	docs, err := s.ReadInto(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	assert.NotNil(t, docs)
}

func TestDelete(t *testing.T) {
	s = fs.NewStore(conf, "test")
	defer s.Close()
	ctx := context.Background()

	err := s.Delete("d59qe7H7lfwy6HFuqIFN", ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

// func TestDeleteAll(t *testing.T) {
// 	s = fs.NewStore(conf, "test")
// 	defer s.Close()
// 	ctx := context.Background()

// 	err := s.DeleteAll(ctx)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }