package newfeeds

type Item struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

type Repo struct {
	Items []Item
}

//实例化Repo
func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}