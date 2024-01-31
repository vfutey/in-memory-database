package in_memory_database

type DB interface {
	Get(key string) (string, bool)
	Set(key, value string)
	Delete(key string)

	StartTransaction()
	Commit()
	Rollback()
}
