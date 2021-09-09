package db

type (
	UniqueKey string
	Label     string
	Record    map[Label][]string
	Data      map[UniqueKey]Record
)

type Query struct {
}

type DB interface {
	Name() string

	CloseDB() error
	Open(string) error
	// Lock() error
	// Unlock() error

	Keys() ([]UniqueKey, error)
	Data() (Data, error)
	Get(UniqueKey) (Record, error)
	IncrementalInsert(UniqueKey, Record) error
	DeleteKey(UniqueKey) error
	DeleteLabel(UniqueKey, Label) error
	Raw(Query) (Data, error)
}
