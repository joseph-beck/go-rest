package feeder

import (
	"rest/internal/feeder"
	"testing"
)

const conf = "../../conf/service-acc.json"

func TestAdd(t *testing.T) {
	f := feeder.NewRepo(conf)
	f.Add(feeder.Item{})
	f.Update()

	if len(f.Items) <= 0 {
		t.Errorf("Failure adding item")
	}
}

func TestGetAll(t *testing.T) {
	f := feeder.NewRepo(conf)
	f.Add(feeder.Item{})
	results := f.GetAll()

	if len(results) <= 0 {
		t.Errorf("Failure adding item")
	}
}

func TestDelete(t *testing.T) {
	f := feeder.NewRepo(conf)
	f.Delete("3bZTw7sbJCAvZI9yySfE")
}
