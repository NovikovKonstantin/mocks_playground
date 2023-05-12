package service

// Validator checks value with business rules
type Validator interface {
	Check(value int64) (bool, error)
}
