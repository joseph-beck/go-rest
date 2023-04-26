package firestore

import (
	"context"
	"sync"

	"cloud.google.com/go/firestore"
)

type StoreCreator interface {
	NewStore(path string, collection string) *Store
}

type StoreAdder interface {
	Add(ctx context.Context) error
	AddStruct(st interface{}, ctx context.Context) error
	AddStructs(st []interface{}, ctx context.Context) error
}

type StoreReader interface {
	Read(ctx context.Context) error
	ReadInto(ctx context.Context) ([]firestore.DocumentSnapshot, error)
	ReadStruct(id string, ctx context.Context) (interface{}, error)
}

type StoreUpdater interface {
	Update(ctx context.Context) error
	UpdateStruct(id string, u interface{}, ctx context.Context) error
}

type StoreDeleter interface {
	Delete(id string, ctx context.Context) error
	DeleteAll(ctx context.Context) error
}

type StoreCloser interface {
	Close() error
}

type Storer interface {
	StoreCreator
	StoreAdder
	StoreReader
	StoreUpdater
	StoreDeleter
	StoreCloser
}

type Store struct {
	ClientMu   sync.Mutex
	Client     firestore.Client
	Collection string
}

func NewStore(path string, collection string) *Store {
	return &Store{
		Client:     *newConn(path),
		Collection: collection,
	}
}

func (s *Store) Close() error {
	s.ClientMu.Lock()
	defer s.ClientMu.Unlock()

	err := s.Client.Close()
	return err
}
