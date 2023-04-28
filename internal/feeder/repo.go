package feeder

import (
	"context"
	"encoding/json"
	"log"
	"rest/pkg/firestore"
)

const (
	collection = "feed"
)

type RepoGetter interface {
	GetAll() []Item
}

type RepoAdder interface {
	Add(item Item)
}

type RepoUpdater interface {
	Update()
}

type RepoDeleter interface {
	Delete(id string)
}

type RepoManager interface {
	RepoGetter
	RepoAdder
	RepoUpdater
}

type Repo struct {
	Items []Item
	Store *firestore.Store
}

func NewRepo(path string) *Repo {
	return &Repo{
		Items: []Item{},
		Store: firestore.NewStore(path, collection),
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	r.Update()
	return r.Items
}

func (r *Repo) Update() {
	ctx := context.Background()
	docs, err := r.Store.ReadInto(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var i []Item
	for _, doc := range docs {
		d := doc.Data()
		b, err := json.Marshal(d)
		if err != nil {
			log.Fatalln(err)
		}

		var item *Item
		err = json.Unmarshal(b, &item)
		if err != nil {
			log.Fatalln(err)
		}
		i = append(i, *item)
	}
	r.Items = i
}

func (r *Repo) Delete(id string) {
	ctx := context.Background()
	err := r.Store.Delete(id, ctx)
	if err != nil {
		log.Fatalln(err)
	}
}