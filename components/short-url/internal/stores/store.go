package stores

// Store represents
type Store interface {
	GetRecord(key string) (string, error)
	PutRecord(key string, value string) error
}
