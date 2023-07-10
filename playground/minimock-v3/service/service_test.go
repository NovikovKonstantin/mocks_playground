package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mocks[T Number] struct {
	r *RepositoryMock
	v *ValidatorMock[T]
}

func newMocks(t *testing.T) *mocks[int64] {
	return &mocks[int64]{
		r: NewRepositoryMock(t),
		v: NewValidatorMock[int64](t),
	}
}

func (m *mocks[_]) Finish() {
	m.v.MinimockFinish()
	m.r.MinimockFinish()
}

func TestService_GetByMap(t *testing.T) {
	type out struct {
		values []int64
		err    error
	}

	tests := []struct {
		name   string
		in     map[string]struct{}
		setup  func(*mocks[int64])
		assert func(out)
	}{
		{
			name:  "empty map",
			in:    nil,
			setup: func(m *mocks[int64]) {},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.Empty(t, o.values)
			},
		},
		{
			name: "success",
			in: map[string]struct{}{
				"one":   {},
				"two":   {},
				"three": {},
				"four":  {},
				"five":  {},
			},
			setup: func(m *mocks[int64]) {
				m.r.GetMock.Inspect(
					func(keys []string) {
						assert.ElementsMatch(
							t,
							[]string{"one", "two", "three", "four", "five"},
							keys,
						)
					},
				).Return([]int64{1, 2, 3, 4, 5}, nil)
			},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.Equal(t, []int64{1, 2, 3, 4, 5}, o.values)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			defer m.Finish()

			tt.setup(m)

			s := NewService[int64](m.r, m.v)

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
		in     []int64
		setup  func(*mocks[int64])
		assert func(out)
	}{
		{
			name:  "no values given",
			in:    nil,
			setup: func(m *mocks[int64]) {},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.Empty(t, o.keys)
			},
		},
		{
			name: "success",
			in:   []int64{1, 2, 3, 4, 5},
			setup: func(m *mocks[int64]) {
				m.v.CheckMock.Inspect(
					// Build an inspect function, which will check 5 values on 5 mock calls.
					func() func(value int64) {
						expected := []int64{1, 2, 3, 4, 5}

						return func(value int64) {
							// Use built-in counter to check finished call's number.
							i := m.v.CheckAfterCounter()

							assert.Equal(t, expected[i], value)
						}
					}(),
				).Return(true, nil)
				m.r.StoreMock.Expect([]int64{1, 2, 3, 4, 5}).Return([]string{"one", "two", "three", "four", "five"}, nil)
			},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.Equal(t, []string{"one", "two", "three", "four", "five"}, o.keys)
			},
		},
		{
			name: "valid and invalid values are given",
			in:   []int64{1, 2, 3, 4, 5},
			setup: func(m *mocks[int64]) {
				type out struct {
					ok  bool
					err error
				}

				expectations := []struct {
					in  int64
					out out
				}{
					{
						in: 1,
						out: out{
							ok:  true,
							err: nil,
						},
					},
					{
						in: 2,
						out: out{
							ok:  true,
							err: nil,
						},
					},
					{
						in: 3,
						out: out{
							ok:  false,
							err: nil,
						},
					},
				}

				// Mock few calls with different return values.
				// Order isn't specified, so mocks can be called various times.
				for _, expectation := range expectations {
					m.v.CheckMock.When(expectation.in).Then(expectation.out.ok, expectation.out.err)
				}
			},
			assert: func(o out) {
				assert.Error(t, o.err)
				assert.EqualError(t, o.err, "value isn't valid")
				assert.Empty(t, o.keys)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)
			defer m.Finish()

			tt.setup(m)

			s := NewService[int64](m.r, m.v)

			keys, err := s.Store(tt.in)

			tt.assert(out{
				keys: keys,
				err:  err,
			})
		})
	}
}
