package feeder

import (
	"testing"
	"rest/internal/feeder"
)

const conf = "../../conf/service-acc.json"

func TestAdd(t *testing.T) {
	feed := feeder.NewRepo(conf)
	feed.Add(feeder.Item{})
	feed.Update()

	if len(feed.Items) <= 0 {
		t.Errorf("Failure adding item")
	}
}

func TestGetAll(t *testing.T) {
	feed := feeder.NewRepo(conf)
	feed.Add(feeder.Item{})
	results := feed.GetAll()

	if len(results) <= 0 {
		t.Errorf("Failure adding item")
	}
}
