package service

//go:generate minimock -i Repository

type Repository interface {
	Store(values []int64) ([]string, error)
	Get(keys []string) ([]int64, error)
}
