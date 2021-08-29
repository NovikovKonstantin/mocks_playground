package service

type Repository interface {
	Store(value int64) (string, error)
	GetByKey(key string) (int64, error)
	GetList() ([]int64, error)
}
