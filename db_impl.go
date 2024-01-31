package in_memory_database

type inMemDB struct{}

func InMemoryDatabase() DB {
	return new(inMemDB)
}

func (db *inMemDB) Get(key string) (string, bool) {
	panic("implement me")
}

func (db *inMemDB) Set(key string, value string) {
	panic("implement me")
}

func (db *inMemDB) Delete(key string) {
	panic("implement me")
}

func (db *inMemDB) StartTransaction() {
	panic("implement me")
}

func (db *inMemDB) Commit() {
	panic("implement me")
}

func (db *inMemDB) Rollback() {
	panic("implement me")
}
