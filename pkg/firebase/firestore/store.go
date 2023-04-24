package firestore

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Store struct {
	ClientMu   sync.Mutex
	Client     firestore.Client
	Collection string
}

func NewStore(path string, collection string) *Store {
	return &Store{
		Client: *newConn(path),
		Collection: collection,
	}
}

func (s *Store) Close() {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	defer s.Client.Close()
}

func (s *Store) Add(ctx context.Context) error {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	_, _, err := s.Client.Collection(s.Collection).Add(ctx, map[string]interface{}{
		"name": "dave",
		"data": "about david",
	})
	return err
}

func (s *Store) Delete(c *firestore.Client, co string, id string, ctx context.Context) error {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	return nil
}

func (s *Store) Read(ctx context.Context) error {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	iter := s.Client.Collection(s.Collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(doc.Data())
	}
	return nil
}

func (s *Store) ReadInto(ctx context.Context) ([]firestore.DocumentSnapshot, error) {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	iter := s.Client.Collection(s.Collection).Documents(ctx)
	docs := []firestore.DocumentSnapshot{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		docs = append(docs, *doc)
	}
	return docs, nil
}
