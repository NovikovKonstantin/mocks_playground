package service

type Repository interface {
	Store(values []int64) ([]string, error)
	Get(keys []string) ([]int64, error)
}
