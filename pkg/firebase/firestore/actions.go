package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (s *Store) Add(ctx context.Context) error {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	_, _, err := s.Client.Collection(s.Collection).Add(ctx, map[string]interface{}{
		"name": "dave",
		"data": "about david",
	})
	return err
}

func (s *Store) AddStruct(st interface{}, ctx context.Context) error {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	m, err := MapStruct(st)
	if err != nil {
		return err
	}

	_, _, err = s.Client.Collection(s.Collection).Add(ctx, m)
	return err
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

func (s *Store) Delete(id string, ctx context.Context) error {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	iter := s.Client.Collection(s.Collection).DocumentRefs(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		if doc.ID == id {
			_, err := doc.Delete(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Store) DeleteAll(ctx context.Context) error {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	iter := s.Client.Collection(s.Collection).DocumentRefs(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		_, err = doc.Delete(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
