package in_memory_database

import "github.com/ttdsuen/golang-stack"

type inMemDB struct {
	values  map[string]string
	history *stack.Stack[map[string]*string]
}

func InMemoryDatabase() DB {
	db := new(inMemDB)
	db.values = make(map[string]string)
	db.history = stack.NewStack[map[string]*string]()
	return db
}

func (db *inMemDB) Get(key string) (string, bool) {
	value, ok := db.values[key]
	return value, ok
}

func (db *inMemDB) Set(key string, value string) {
	db.storeChange(key)
	db.values[key] = value
}

func (db *inMemDB) Delete(key string) {
	db.storeChange(key)
	delete(db.values, key)
}

func (db *inMemDB) StartTransaction() {
	changes := make(map[string]*string)
	db.history.Push(changes)
}

func (db *inMemDB) Commit() {
	if db.history.IsEmpty() {
		// TODO: we can return error/panic in this case - that means that we try to commit transaction that not started
		return
	}
	_, _ = db.history.Pop()
}

func (db *inMemDB) Rollback() {
	changes, ok := db.history.Pop()
	if !ok {
		// TODO: we can return error/panic in this case - that means that we try to rollback transaction that not started
		return
	}
	for key, value := range changes {
		if value == nil {
			delete(db.values, key)
		} else {
			db.values[key] = *value
		}
	}
}

func (db *inMemDB) storeChange(key string) {
	changes, ok := db.history.Top()
	if !ok {
		return // skip storing if no transaction started yet
	}
	_, ok = changes[key]
	if ok {
		return // store original value before 1st change only, otherwise skip storing
	}
	original, originalOk := db.values[key]
	if !originalOk {
		changes[key] = nil
	} else {
		changes[key] = &original
	}
}
