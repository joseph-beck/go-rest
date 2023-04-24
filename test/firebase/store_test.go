package firebase

import (
	"context"
	"log"
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

	err, docs := s.ReadInto(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	assert.NotNil(t, docs)
}
