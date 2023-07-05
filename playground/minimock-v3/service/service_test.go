package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mocks[T Number] struct{
	r *RepositoryMock
	v *ValidatorMock[T]
}

func newMocks(t *testing.T) *mocks[int64] {
	return &mocks[int64]{
		r: NewRepositoryMock(t),
		v: NewValidatorMock[int64](t),
	}
}

func TestService_GetByMap(t *testing.T) {
	type out struct {
		values []int64
		err error
	}

	tests := []struct{
		name string
		in map[string]struct{}
		setup func(*mocks[int64])
		assert func(out)
	}{
		{
			name: "empty map",
			in: nil,
			setup: func(m *mocks[int64]) {},
			assert: func(o out) {
				assert.NoError(t, o.err)
				assert.Empty(t, o.values)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			m := newMocks(t)

			tt.setup(m)

			s := NewService[int64](m.r, m.v)

			values, err := s.GetByMap(tt.in)

			tt.assert(out{
				values: values,
				err: err,
			})
		})
	}
}
