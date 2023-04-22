package feeder

import (
	"testing"
	"rest/internal/feeder"
)

func TestAdd(t *testing.T) {
	feed := feeder.NewRepo()
	feed.Add(feeder.Item{})

	if len(feed.Items) != 1 {
		t.Errorf("Failure adding item")
	}
}

func TestGetAll(t *testing.T) {
	feed := feeder.NewRepo()
	feed.Add(feeder.Item{})
	results := feed.GetAll()

	if len(results) != 1 {
		t.Errorf("Failure adding item")
	}
}
