package store

type Store interface {
	Save(key string, d interface{}) error

	Load(key string, d interface{}) error

	Delete(key string) error
}
