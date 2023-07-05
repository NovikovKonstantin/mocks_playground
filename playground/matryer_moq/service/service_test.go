package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mocks struct {
	r *RepositoryMock
	v *ValidatorMock[int64]
}

func TestService_GetByMap(t *testing.T) {
	type out struct {
		values []int64
		err    error
	}

	tt := []struct {
		name   string
		in     map[string]struct{}
		mocks  mocks
		assert func(out, mocks)
	}{
		{
			name: "empty map",
			in:   nil,
			mocks: mocks{
				r: &RepositoryMock{},
			},
			assert: func(out out, m mocks) {
				assert.Len(t, m.r.GetCalls(), 0)
				assert.Nil(t, out.values)
				assert.NoError(t, out.err)
			},
		},
		{
			name: "filled map",
			in: map[string]struct{}{
				"key01": {},
				"key02": {},
				"key03": {},
				"key04": {},
			},
			mocks: mocks{
				r: &RepositoryMock{
					GetFunc: func(keys []string) ([]int64, error) {
						expected := []string{
							"key01",
							"key02",
							"key03",
							"key04",
						}
						assert.ElementsMatch(t, keys, expected)

						return []int64{1, 2, 3, 4}, nil
					},
				},
			},
			assert: func(out out, m mocks) {
				assert.NoError(t, out.err)
				assert.Equal(t, out.values, []int64{1, 2, 3, 4})
			},
		},
		{
			name: "error in repo",
			in: map[string]struct{}{
				"key01": {},
				"key02": {},
				"key03": {},
				"key04": {},
			},
			mocks: mocks{
				r: &RepositoryMock{
					GetFunc: func(keys []string) ([]int64, error) {
						expected := []string{
							"key01",
							"key02",
							"key03",
							"key04",
						}
						assert.ElementsMatch(t, keys, expected)

						return nil, errors.New("repo error")
					},
				},
			},
			assert: func(out out, m mocks) {
				if assert.Error(t, out.err) {
					assert.Equal(t, errors.New("repo error"), out.err)
				}
				assert.Empty(t, out.values)
			},
		},
	}

	for _, tt := range tt {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			s := NewService[int64](tt.mocks.r, tt.mocks.v)

			values, err := s.GetByMap(tt.in)

			tt.assert(
				out{
					values: values,
					err:    err,
				},
				tt.mocks,
			)
		})
	}
}

func TestService_Store(t *testing.T) {
	type out struct {
		keys []string
		err  error
	}

	tt := []struct {
		name   string
		in     []int64
		mocks  mocks
		assert func(out, mocks)
	}{
		{
			name: "empty slice",
			in:   nil,
			mocks: mocks{
				r: &RepositoryMock{},
				v: &ValidatorMock[int64]{},
			},
			assert: func(o out, m mocks) {
				assert.NoError(t, o.err)
				assert.Empty(t, o.keys)
			},
		},
		{
			name: "filled with correct values",
			in:   []int64{1, 2, 3, 4, 5},
			mocks: mocks{
				r: &RepositoryMock{
					StoreFunc: func(values []int64) ([]string, error) {
						expected := []int64{1, 2, 3, 4, 5}
						assert.Equal(t, expected, values)
						return []string{"1", "2", "3", "4", "5"}, nil
					},
				},
				v: &ValidatorMock[int64]{
					CheckFunc: func(value int64) (bool, error) {
						return true, nil
					},
				},
			},
			assert: func(out out, m mocks) {
				assert.NoError(t, out.err)
				assert.Len(t, m.r.StoreCalls(), 1)
				assert.Len(t, m.v.CheckCalls(), 5)
				assert.Equal(t, []string{"1", "2", "3", "4", "5"}, out.keys)
			},
		},
		{
			name: "filled with correct and invalid values",
			in:   []int64{1, 2, 3, 4, 5},
			mocks: mocks{
				r: &RepositoryMock{},
				v: &ValidatorMock[int64]{
					CheckFunc: func() func(value int64) (bool, error) {
						values := []int64{1, 2, 3}
						i := -1 // Call's index. It's useful to start from -1 and increase it below.

						return func(value int64) (bool, error) {
							i++

							assert.Equal(t, values[i], value)

							switch i {
							case 2:
								return false, ErrInvalidValue
							default:
								return true, nil
							}
						}
					}(),
				},
			},
			assert: func(out out, m mocks) {
				assert.EqualError(t, out.err, ErrInvalidValue.Error())
				assert.Len(t, m.r.StoreCalls(), 0)
				assert.Len(t, m.v.CheckCalls(), 3)
				assert.Empty(t, out.keys)
			},
		},
	}

	for _, tt := range tt {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			s := NewService[int64](tt.mocks.r, tt.mocks.v)

			keys, err := s.Store(tt.in)

			tt.assert(
				out{
					keys: keys,
					err:  err,
				},
				tt.mocks,
			)
		})
	}
}

func TestService_CheckGeneric(t *testing.T) {
	type out struct {
		result bool
		err    error
	}

	tt := []struct {
		name   string
		in     int64
		mocks  mocks
		assert func(out, mocks)
	}{
		{
			name: "success",
			in:   5,
			mocks: mocks{
				v: &ValidatorMock[int64]{
					CheckGenericFunc: func(value int64) (bool, error) {
						assert.Equal(t, int64(5), value)

						return true, nil
					},
				},
			},
			assert: func(o out, m mocks) {
				assert.NoError(t, o.err)
				assert.True(t, o.result)

				assert.Len(t, m.v.CheckGenericCalls(), 1)
			},
		},
	}

	for _, tt := range tt {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {})
		s := NewService[int64](tt.mocks.r, tt.mocks.v)

		res, err := s.ValidateAny(tt.in)

		tt.assert(
			out{
				result: res,
				err:    err,
			},
			tt.mocks,
		)
	}
}

func TestService_CheckGenerics(t *testing.T) {
	type out struct {
		result bool
		err    error
	}

	tt := []struct {
		name   string
		in     []int64
		mocks  mocks
		assert func(out, mocks)
	}{
		{
			name: "success",
			in:   []int64{1, 2, 3, 4},
			mocks: mocks{
				v: &ValidatorMock[int64]{
					CheckGenericsFunc: func(values []int64) (bool, error) {
						assert.Equal(t, []int64{1, 2, 3, 4}, values)

						return true, nil
					},
				},
			},
			assert: func(o out, m mocks) {
				assert.NoError(t, o.err)
				assert.True(t, o.result)

				assert.Len(t, m.v.CheckGenericsCalls(), 1)
			},
		},
	}

	for _, tt := range tt {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {})
		s := NewService[int64](tt.mocks.r, tt.mocks.v)

		res, err := s.ValidateAnyBatch(tt.in)

		tt.assert(
			out{
				result: res,
				err:    err,
			},
			tt.mocks,
		)
	}
}
