package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
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
