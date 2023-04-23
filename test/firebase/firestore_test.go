package firebase

import (
	"context"
	"log"
	fs "rest/pkg/firebase/firestore"
	"testing"

	"cloud.google.com/go/firestore"
)

const conf = "../../conf/service-acc.json"

var c *firestore.Client

func TestConn(t *testing.T) {
	c := fs.Conn(conf)
	defer fs.Close(c)
}

func TestAdd(t *testing.T) {
	c = fs.Conn(conf)
	defer fs.Close(c)

	ctx := context.Background()
	err := fs.Add(c, "test", ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRead(t *testing.T) {
	c = fs.Conn(conf)
	defer fs.Close(c)

	ctx := context.Background()
	err := fs.Read(c, "test", ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestReadInto(t *testing.T) {
	c = fs.Conn(conf)
	defer fs.Close(c)
	ctx := context.Background()

	err, docs := fs.ReadInto(c, "test", ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for doc := range docs {
		log.Println(doc)
	}
}
