package firestore

import (
	"sync"

	"cloud.google.com/go/firestore"
)

type StoreCreator interface {
	NewStore(path string, collection string) *Store
}

type StoreAdder interface {
	Add() error
	AddStruct() error
	AddStructs() error
}

type StoreReader interface {
	Read() error
	ReadInto() error
	ReadStruct() error
}

type StoreUpdater interface {
	Update() error
	UpdateStruct() error
}

type StoreDeleter interface {
	Delete() error
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
