package store

var client Factory

type Factory interface {
	Users() UserStore
	CommonDB() CommonStore
	Close() error
	InitTables() error
}

func Client() Factory {
	return client
}
func SetClient(factory Factory) {
	client = factory
}