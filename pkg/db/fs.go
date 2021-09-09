package db

var _ DB = &File{}

type File struct {
	name    string
	indexes []UniqueKey
	data    Data

	// sync.Mutex
}

func (f *File) Name() string {
	return f.name
}

func (f *File) CloseDB() error {
	return nil
}

func (f *File) Open(filepath string) error {
	return nil
}

// func (f *File) Lock() error {
// 	return nil
// }
//
// func (f *File) Unlock() error {
// 	return nil
// }

func (f *File) Keys() ([]UniqueKey, error) {
	return f.indexes, nil
}

func (f *File) Data() (Data, error) {
	return f.data, nil
}

func (f *File) Get(UniqueKey) (Record, error) {
	return nil, nil
}

func (f *File) IncrementalInsert(key UniqueKey, record Record) error {
	return nil
}

func (f *File) save() error {
	return nil
}

func (f *File) DeleteKey(key UniqueKey) error {
	return nil
}

func (f *File) DeleteLabel(key UniqueKey, label Label) error {
	return nil
}

func (f *File) Raw(query Query) (Data, error) {
	return nil, nil
}
