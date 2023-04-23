package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func Add(c *firestore.Client, co string, ctx context.Context) error {
	_, _, err := c.Collection(co).Add(ctx, map[string]interface{}{
		"name": "dave",
		"data": "about david",
	})
	return err
}

func Delete(c *firestore.Client, co string, id string, ctx context.Context) error {
	return nil
}

func Read(c *firestore.Client, co string, ctx context.Context) error {
	iter := c.Collection(co).Documents(ctx)
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

func ReadInto(c *firestore.Client, co string, ctx context.Context) (error, []firestore.DocumentSnapshot) {
	iter := c.Collection(co).Documents(ctx)
	docs := []firestore.DocumentSnapshot{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err, nil
		}
		docs = append(docs, *doc)
	}
	return nil, docs
}
