package feeder

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Repo struct {
	Items []Item
}

func NewRepo() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
