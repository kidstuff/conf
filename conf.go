package conf

type ConfigObj struct {
	Key *string `bson:"_id" json:",omitempty"`
	Val *string `bson:"Val" json:",omitempty"`
}

type Configurator interface {
	Set(key string, val string) error
	UnSet(key string) error
	Get(key string) (string, error)
	SetMulti(map[string]string) error
	UnSetMulti(keys []string) error
	GetMulti(keys []string) (map[string]string, error)
}
