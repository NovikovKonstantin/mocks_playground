package service

import (
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/tj/assert"
)

type mocks struct {
	repo      *RepositoryMock
	validator *ValidatorMock
}

func newMocks(t *testing.T) *mocks {
	mc := minimock.NewController(t)

	mockRepo := NewRepositoryMock(mc)
	mockValidator := NewValidatorMock(mc)

	return &mocks{
		repo:      mockRepo,
		validator: mockValidator,
	}
}

func (m *mocks) meetExpectations() {
	m.repo.MinimockFinish()
	m.validator.MinimockFinish()
}

func TestNewService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		s := NewService(mocks.repo, mocks.validator)
		assert.NotNil(t, s)
	})
}

func TestService_GetList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		storedValues := []int64{0, 1, 2, 3}

		mocks.repo.GetListMock.Return(storedValues, nil)

		s := NewService(mocks.repo, mocks.validator)
		retValues, err := s.GetList()
		assert.NoError(t, err)
		assert.Equal(t, storedValues, retValues)
	})

	t.Run("repository error", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		expectedErr := errors.New("some error")

		mocks.repo.GetListMock.Return(nil, expectedErr)

		s := NewService(mocks.repo, mocks.validator)
		retValues, err := s.GetList()
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, retValues)
	})
}

func TestService_GetByKey(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		storedValue := int64(1)
		storedKey := "some key"

		mocks.repo.GetByKeyMock.Set(func(key string) (ret int64, err error) {
			assert.Equal(t, storedKey, key)
			return storedValue, nil
		})

		s := NewService(mocks.repo, mocks.validator)
		retValue, err := s.GetByKey(storedKey)
		assert.NoError(t, err)
		assert.Equal(t, storedValue, retValue)
	})

	t.Run("multiple calls test", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		storedValues := map[string]int64{
			"one":   1,
			"two":   2,
			"three": 3,
		}

		for k, v := range storedValues {
			mocks.repo.GetByKeyMock.When(k).Then(v, nil)
		}

		s := NewService(mocks.repo, mocks.validator)

		callsCountByKey := 5
		expectedCallsCount := callsCountByKey * len(storedValues)
		for i := 0; i < callsCountByKey; i++ {
			for k, expectedValue := range storedValues {
				retValue, err := s.GetByKey(k)
				assert.NoError(t, err)
				assert.Equal(t, expectedValue, retValue)
			}
		}
		assert.Equal(t, uint64(expectedCallsCount), mocks.repo.GetByKeyAfterCounter())
	})

	t.Run("repository error", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		storedValue := int64(1)
		storedKey := "some key"

		mocks.repo.GetByKeyMock.Expect(storedKey).Return(storedValue, nil)

		s := NewService(mocks.repo, mocks.validator)
		retValue, err := s.GetByKey(storedKey)
		assert.NoError(t, err)
		assert.Equal(t, storedValue, retValue)
	})

}

func TestService_Store(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		value := int64(1)
		key := "some key"

		mocks.validator.CheckMock.Expect(value).Return(true, nil)
		mocks.repo.StoreMock.Expect(value).Return(key, nil)

		s := NewService(mocks.repo, mocks.validator)
		retKey, err := s.Store(value)
		assert.NoError(t, err)
		assert.Equal(t, key, retKey)
	})

	t.Run("validator forbid", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		value := int64(1)

		mocks.validator.CheckMock.Expect(value).Return(false, nil)

		s := NewService(mocks.repo, mocks.validator)
		retKey, err := s.Store(value)
		assert.NotEmpty(t, err)
		assert.Equal(t, "", retKey)
	})

	t.Run("repo fail", func(t *testing.T) {
		mocks := newMocks(t)
		defer mocks.meetExpectations()

		value := int64(1)
		expectedError := errors.New("some error")

		mocks.validator.CheckMock.Expect(value).Return(true, nil)
		mocks.repo.StoreMock.Expect(value).Return("", expectedError)

		s := NewService(mocks.repo, mocks.validator)
		retKey, err := s.Store(value)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, "", retKey)
	})

}
