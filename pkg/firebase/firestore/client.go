package firestore

import (
	"time"

	"cloud.google.com/go/firestore"
	"gopkg.in/oauth2.v3"
)

const (
	keyCode    = "Code"
	keyAccess  = "Access"
	keyRefresh = "Refresh"

	timeout = 15 * time.Second
)

type client struct {
	store *store
}

func NewClient(cl *firestore.Client, co string) oauth2.TokenStore {
	return newWithTimeout(cl, co, timeout)
}

func newWithTimeout(cl *firestore.Client, co string, t time.Duration) oauth2.TokenStore {
	fs := &store{
		name:    co,
		timeout: t,
		client:  cl,
	}
	
	return &client{
		store: fs,
	}
}

func (c *client) Create(info oauth2.TokenInfo) error {
	t, err := token(info)
	if err != nil {
		return err
	}
	return c.store.Put(t)
}

func (c *client) GetByCode(code string) (oauth2.TokenInfo, error) {
	return c.store.Get(keyCode, code)
}

func (c *client) GetByAccess(access string) (oauth2.TokenInfo, error) {
	return c.store.Get(keyAccess, access)
}

func (c *client) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	return c.store.Get(keyRefresh, refresh)
}

func (c *client) RemoveByCode(code string) error {
	return c.store.Del(keyCode, code)
}

func (c *client) RemoveByAccess(access string) error {
	return c.store.Del(keyCode, access)
}

func (c *client) RemoveByRefresh(refresh string) error {
	return c.store.Del(keyCode, refresh)
}
