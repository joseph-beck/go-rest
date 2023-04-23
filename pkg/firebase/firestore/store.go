package firestore

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"gopkg.in/oauth2.v3/models"
)

var ErrDocDoesNotExist = errors.New("document doesn't exist")

type store struct {
	clientMut sync.Mutex
	client    *firestore.Client
	name      string
	timeout   time.Duration
}

func (s *store) Put(tok *models.Token) error {
	s.clientMut.Lock()
	defer s.clientMut.Unlock()

	c, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	_, _, err := s.client.Collection(s.name).Add(c, tok)
	return err
}

func (s *store) Get(key string, val interface{}) (*models.Token, error) {
	s.clientMut.Lock()
	defer s.clientMut.Unlock()

	c, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	iter := s.client.Collection(s.name).Where(key, "==", val).Limit(1).Documents(c)
	doc, err := func(iter *firestore.DocumentIterator) (*firestore.DocumentSnapshot, error) {
		doc, err := iter.Next()
		if err != nil {
			return nil, err
		}
		if !doc.Exists() {
			return nil, errors.New("document doesn't exist")
		}
		return doc, err
	}(iter)
	if err != nil {
		return nil, err
	}

	info := &models.Token{}
	err = doc.DataTo(info)
	return info, err
}

func (s *store) Del(key string, val interface{}) error {
	s.clientMut.Lock()
	defer s.clientMut.Unlock()

	c, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	return s.client.RunTransaction(c, func(c context.Context, t *firestore.Transaction) error {
		query := s.client.Collection(s.name).Where(key, "==", val).Limit(1)
		iter := t.Documents(query)
		doc, err := func(iter *firestore.DocumentIterator) (*firestore.DocumentSnapshot, error) {
			doc, err := iter.Next()
			if err != nil {
				return nil, err
			}
			if !doc.Exists() {
				return nil, ErrDocDoesNotExist
			}
			return doc, err
		}(iter)

		if err != nil {
			if err == iterator.Done || err == ErrDocDoesNotExist {
				return nil
			}
			return err
		}
		return t.Delete(doc.Ref)
	})
}
