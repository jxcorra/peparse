package common

type Key struct {
	Element *string `json:"element"`
	Class   *string `json:"class"`
}

type Item struct {
	Attribute *string `json:"attr"`
	Text      *string `json:"text"`
}

type Search struct {
	Key      Key    `json:"key"`
	WithText bool   `json:"withText"`
	Parse    []Item `json:"parse"`
}

type ResourceConfig struct {
	Url    string   `json:"url"`
	Search []Search `json:"search"`
}

type Resources struct {
	Resources []ResourceConfig `json:"resources"`
}

type Data map[string]string

type DataItem struct {
	Data     Data
	WithText bool
}

func (di *DataItem) GetData() map[string]string {
	return di.Data
}

type Parsed []DataItem

type Parameters struct {
	Period       int
	NumOfWorkers int
	Resources    *Resources
	Tasks        chan ResourceConfig
	Output       chan Parsed
}
