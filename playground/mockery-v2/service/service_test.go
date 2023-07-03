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
			name: "matched by return func",
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
			name: "use matched by func",
			in: map[string]struct{}{
				"one":   {},
				"two":   {},
				"three": {},
				"four":  {},
				"five":  {},
			},
			m: newMocks(t),
			setup: func(m m) {
				m.r.EXPECT().
					Get(mock.MatchedBy(func(keys []string) bool {
						// Simple matcher function, which compares input slice of string with expected slice without order.
						// Also can be implemented by .On function and assert.ElementsMatch, but it won't be a matcher function.
						expected := map[string]struct{}{
							"one":   {},
							"two":   {},
							"three": {},
							"four":  {},
							"five":  {},
						}

						if len(keys) != len(expected) {
							return false
						}

						for _, key := range keys {
							delete(expected, key)
						}

						return len(expected) == 0
					})).
					Return([]int64{1, 2, 3, 4, 5}, nil)
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
				// Send a type in a string. No consts for primitives, make your own consts or type string everytime.
				// No methods as mock.String().
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

func TestService_Store(t *testing.T) {
	type out struct {
		keys []string
		err  error
	}

	tests := []struct {
		name   string
		m      m
		in     []int64
		setup  func(m)
		assert func(out)
	}{
		{
			name:  "zero values",
			in:    nil,
			m:     newMocks(t),
			setup: func(m m) {},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.Empty(t, o.keys)
			},
		},
		{
			name: "check type only, not value",
			in:   []int64{1, 2, 3, 4, 5},
			m:    newMocks(t),
			setup: func(m m) {
				m.v.EXPECT().
					Check(mock.AnythingOfType("int64")). // Check type only, not the value. Bad unit test, we'll pass random values as expected ones.
					Return(true, nil).
					Times(5) // Specify count of calls. Very useful.
				m.r.EXPECT().Store([]int64{1, 2, 3, 4, 5}).Return([]string{"one", "two", "three", "four", "five"}, nil)
			},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.Equal(t, o.keys, []string{"one", "two", "three", "four", "five"})
			},
		},
		{
			name: "check values properly",
			in:   []int64{1, 2, 3, 4, 5},
			m:    newMocks(t),
			setup: func(m m) {
				// Register EXPECT for every entry in the input.
				for _, value := range []int64{1, 2, 3, 4, 5} {
					m.v.EXPECT().Check(value).Return(true, nil)
				}

				m.r.EXPECT().Store([]int64{1, 2, 3, 4, 5}).Return([]string{"one", "two", "three", "four", "five"}, nil)
			},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.Equal(t, o.keys, []string{"one", "two", "three", "four", "five"})
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			tt.setup(tt.m)

			s := service.NewService[int64](tt.m.r, tt.m.v)

			keys, err := s.Store(tt.in)

			tt.assert(out{
				keys: keys,
				err:  err,
			})
		})
	}
}

func TestService_ValidateAny(t *testing.T) {
	type out struct {
		ok  bool
		err error
	}

	tests := []struct {
		name   string
		m      m
		in     int64
		setup  func(m)
		assert func(out)
	}{
		{
			name: "success",
			m:    newMocks(t),
			in:   1,
			setup: func(m m) {
				m.v.EXPECT().CheckGeneric(int64(1)).Return(true, nil)
			},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.True(t, o.ok)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			tt.setup(tt.m)

			s := service.NewService[int64](tt.m.r, tt.m.v)

			ok, err := s.ValidateAny(tt.in)

			tt.assert(out{
				ok:  ok,
				err: err,
			})
		})
	}
}

func TestService_ValidateAnyBatch(t *testing.T) {
	type out struct {
		ok  bool
		err error
	}

	tests := []struct {
		name   string
		m      m
		in     []int64
		setup  func(m)
		assert func(out)
	}{
		{
			name: "success",
			m:    newMocks(t),
			in:   []int64{1, 2, 3, 4, 5},
			setup: func(m m) {
				m.v.EXPECT().CheckGenerics([]int64{1, 2, 3, 4, 5}).Return(true, nil)
			},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.True(t, o.ok)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			tt.setup(tt.m)

			s := service.NewService[int64](tt.m.r, tt.m.v)

			ok, err := s.ValidateAnyBatch(tt.in)

			tt.assert(out{
				ok:  ok,
				err: err,
			})
		})
	}
}
