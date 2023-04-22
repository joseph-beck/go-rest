package feeder

import "testing"

func TestAdd(t *testing.T) {
	feed := NewRepo()
	feed.Add(Item{})

	if len(feed.Items) != 1 {
		t.Errorf("Failure adding item")
	} 
}

func TestGetAll(t *testing.T) {
	feed := NewRepo()
	feed.Add(Item{})
	results := feed.GetAll()

	if len(results) != 1 {
		t.Errorf("Failure adding item")
	}
}