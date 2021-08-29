package service

import "errors"

type Service struct {
	r Repository
	v Validator
}

func NewService(r Repository, v Validator) *Service {
	return &Service{
		r: r,
		v: v,
	}
}

func (s *Service) GetList() ([]int64, error) {
	values, err := s.r.GetList()
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (s *Service) GetByKey(key string) (int64, error) {
	value, err := s.r.GetByKey(key)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) Store(value int64) (string, error) {
	ok, err := s.v.Check(value)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("value is not valid")
	}

	key, err := s.r.Store(value)
	if err != nil {
		return "", err
	}

	return key, nil
}
