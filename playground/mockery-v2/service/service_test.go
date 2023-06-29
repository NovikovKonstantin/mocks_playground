package service_test

import (
	"errors"
	"mockery-v2/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type m struct {
	r *service.MockRepository
	v *service.MockValidator[int64]
}

func newMocks(t *testing.T) m {
	return m{
		r: service.NewMockRepository(t),
		v: service.NewMockValidator[int64](t),
	}
}

func TestService_GetByMap(t *testing.T) {
	type out struct {
		values []int64
		err    error
	}

	tt := []struct {
		name   string
		m      m
		in     map[string]struct{}
		setup  func(m)
		assert func(out out)
	}{
		{
			name:  "empty map",
			in:    nil,
			m:     newMocks(t),
			setup: func(m m) {},
			assert: func(out out) {
				assert.NoError(t, out.err)
			},
		},
		{
			name: "filled map",
			in: map[string]struct{}{
				"one":   {},
				"two":   {},
				"three": {},
				"four":  {},
				"five":  {},
			},
			m: newMocks(t),
			setup: func(m m) {
				// It's so bad.
				m.r.On("Get", mock.Anything).Return(func(keys []string) ([]int64, error) {
					expected := []string{
						"one",
						"two",
						"three",
						"four",
						"five",
					}
					assert.ElementsMatch(t, expected, keys)

					return []int64{1, 2, 3, 4, 5}, nil
				})
			},
			assert: func(out out) {
				assert.NoError(t, out.err)
				assert.Equal(t, out.values, []int64{1, 2, 3, 4, 5})
			},
		},
		{
			name: "error",
			in: map[string]struct{}{
				"one":   {},
				"two":   {},
				"three": {},
				"four":  {},
				"five":  {},
			},
			m: newMocks(t),
			setup: func(m m) {
				m.r.EXPECT().Get(mock.AnythingOfType("[]string")).Return(nil, errors.New("my error"))
			},
			assert: func(out out) {
				assert.Error(t, out.err, "my error")
				assert.Empty(t, out.values)
			},
		},
	}

	for _, tt := range tt {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			tt.setup(tt.m)

			s := service.NewService[int64](tt.m.r, tt.m.v)

			values, err := s.GetByMap(tt.in)

			tt.assert(out{
				values: values,
				err:    err,
			})
		})
	}
}
