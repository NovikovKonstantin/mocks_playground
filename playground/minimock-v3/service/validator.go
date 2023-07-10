package service

//go:generate minimock -i Validator

// Validator checks value with business rules.
// It contains generics, so you can check the mock library's feature of mocking generics.
type Validator[T Number] interface {
	Check(value int64) (bool, error)
	CheckGeneric(value T) (bool, error)
	CheckGenerics(values []T) (bool, error)
}
