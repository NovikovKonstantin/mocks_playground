package service

//go:generate moq -out repository_mock.go . Repository

type Repository interface {
	Store(values []int64) ([]string, error)
	Get(keys []string) ([]int64, error)
}
