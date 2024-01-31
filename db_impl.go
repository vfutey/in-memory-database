package in_memory_database

import "github.com/ttdsuen/golang-stack"

type inMemDB struct {
	values  map[string]string
	history *stack.Stack[map[string]string]
}

func InMemoryDatabase() DB {
	db := new(inMemDB)
	db.values = make(map[string]string)
	db.history = stack.NewStack[map[string]string]()
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
	clonedMap := cloneMap(db.values)
	db.history.Push(clonedMap)
}

func (db *inMemDB) Commit() {
	if db.history.IsEmpty() {
		// TODO: we can return error/panic in this case - that means that we try to commit transaction that not started
		return
	}
	_, _ = db.history.Pop()
}

func (db *inMemDB) Rollback() {
	values, ok := db.history.Pop()
	if !ok {
		// TODO: we can return error/panic in this case - that means that we try to rollback transaction that not started
		return
	}
	db.values = values // revert values to previous snapshot
}

func cloneMap(values map[string]string) map[string]string {
	res := make(map[string]string, len(values))
	for key, value := range values {
		res[key] = value
	}
	return res
}
