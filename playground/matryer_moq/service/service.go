package service

import "errors"

type Number interface {
	int64 | float64
}

type Service[T Number] struct {
	r Repository
	v Validator[T]
}

func NewService[T Number](r Repository, v Validator[T]) *Service[T] {
	return &Service[T]{
		r: r,
		v: v,
	}
}

// GetByMap method gets the map, extracts keys (in random order), and transmits them to the repository.
// Method useful to check the way of the mock library to react on in in random order.
func (s *Service[T]) GetByMap(mKeys map[string]struct{}) ([]int64, error) {
	if len(mKeys) == 0 {
		return nil, nil
	}

	keys := make([]string, 0, len(mKeys))
	for key := range mKeys {
		keys = append(keys, key)
	}

	return s.r.Get(keys)
}

// Store method tries to store in in the repository.
// It contains multiple calls of the Check method, so it can be used to check how the mock library reacts to it.
func (s *Service[T]) Store(values []int64) ([]string, error) {
	for _, value := range values {
		ok, err := s.v.Check(value)
		switch {
		case err != nil:
			return nil, err
		case !ok:
			return nil, errors.New("value isn't valid")
		}
	}

	keys, err := s.r.Store(values)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (s *Service[T]) ValidateAny(value T) (bool, error) {
	return s.v.CheckGeneric(value)
}

func (s *Service[T]) ValidateAnyBatch(values []T) (bool, error) {
	return s.v.CheckGenerics(values)
}
