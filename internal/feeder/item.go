package feeder

import "fmt"

type Item struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (i *Item) Str() string {
	return fmt.Sprintf(
		"Name: %s\nData: %s\n",
		i.Name,
		i.Data,
	)
}
