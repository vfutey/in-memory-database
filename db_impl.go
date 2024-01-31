package in_memory_database

type inMemDB struct {
	values map[string]string
}

func InMemoryDatabase() DB {
	db := new(inMemDB)
	db.values = make(map[string]string)
	return db
}

func (db *inMemDB) Get(key string) (string, bool) {
	value, ok := db.values[key]
	return value, ok
}

func (db *inMemDB) Set(key string, value string) {
	db.values[key] = value
}

func (db *inMemDB) Delete(key string) {
	delete(db.values, key)
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
