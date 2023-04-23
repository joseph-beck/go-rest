package firebase

import (
	"os"
	"reflect"
	fs "rest/pkg/firebase/firestore"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iterator"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
)

const conf = "../../conf/service-acc.json"

var c *firestore.Client

func TestMain(m *testing.M) {
	c = fs.Conn(conf)
	os.Exit(func() int {
		c.Close()
		return m.Run()
	}())
}

func TestStoreClient(t *testing.T) {
	client := fs.NewClient(c, "tests")
	type holder struct {
		key string
		get func(string) (oauth2.TokenInfo, error)
		del func(string) error
	}
	tokens := map[*models.Token]holder{
		{Access: "access"}:   {key: "access", get: client.GetByAccess, del: client.RemoveByAccess},
		{Code: "code"}:       {key: "code", get: client.GetByCode, del: client.RemoveByCode},
		{Refresh: "refresh"}: {key: "refresh", get: client.GetByRefresh, del: client.RemoveByRefresh},
	}
	for i, h := range tokens {
		err := client.Create(i)
		assert.Nil(t, err)

		tok, err := h.get(h.key)
		assert.Nil(t, err)
		assert.Equal(t, i, tok)

		err = h.del(h.key)
		assert.Nil(t, err)

		_, err = h.get(h.key)
		assert.NotNil(t, err)

		err = h.del(h.key)
		assert.Nil(t, err)
	}
}

func TestNoDocument(t *testing.T) {
	client := fs.NewClient(c, "tests")
	info, err := client.GetByRefresh("whoops")
	assert.Nil(t, info)
	assert.Equal(t, iterator.Done, err)
}

func TestIsNilOrZero(t *testing.T) {
	tokens := map[oauth2.TokenInfo]bool{
		nil:                               true,
		&models.Token{}:                   true,
		&models.Token{Access: "access"}:   false,
		&models.Token{Code: "code"}:       false,
		&models.Token{Refresh: "refresh"}: false,
	}
	for tok, expected := range tokens {
		result := func(info oauth2.TokenInfo) bool {
			if info == nil {
				return true
			}
			if v := reflect.ValueOf(info); v.IsNil() {
				return true
			}
			return reflect.DeepEqual(info, info.New())
		}(tok)
		assert.Equal(t, expected, result)
	}
}
