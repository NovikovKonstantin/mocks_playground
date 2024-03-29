package service

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i minimock-v3/service.Validator -o ./validator_mock_test.go -n ValidatorMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ValidatorMock implements Validator
type ValidatorMock[T Number] struct {
	t minimock.Tester

	funcCheck          func(value int64) (b1 bool, err error)
	inspectFuncCheck   func(value int64)
	afterCheckCounter  uint64
	beforeCheckCounter uint64
	CheckMock          mValidatorMockCheck[T]

	funcCheckGeneric          func(value T) (b1 bool, err error)
	inspectFuncCheckGeneric   func(value T)
	afterCheckGenericCounter  uint64
	beforeCheckGenericCounter uint64
	CheckGenericMock          mValidatorMockCheckGeneric[T]

	funcCheckGenerics          func(values []T) (b1 bool, err error)
	inspectFuncCheckGenerics   func(values []T)
	afterCheckGenericsCounter  uint64
	beforeCheckGenericsCounter uint64
	CheckGenericsMock          mValidatorMockCheckGenerics[T]
}

// NewValidatorMock returns a mock for Validator
func NewValidatorMock[T Number](t minimock.Tester) *ValidatorMock[T] {
	m := &ValidatorMock[T]{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CheckMock = mValidatorMockCheck[T]{mock: m}
	m.CheckMock.callArgs = []*ValidatorMockCheckParams[T]{}

	m.CheckGenericMock = mValidatorMockCheckGeneric[T]{mock: m}
	m.CheckGenericMock.callArgs = []*ValidatorMockCheckGenericParams[T]{}

	m.CheckGenericsMock = mValidatorMockCheckGenerics[T]{mock: m}
	m.CheckGenericsMock.callArgs = []*ValidatorMockCheckGenericsParams[T]{}

	return m
}

type mValidatorMockCheck[T Number] struct {
	mock               *ValidatorMock[T]
	defaultExpectation *ValidatorMockCheckExpectation[T]
	expectations       []*ValidatorMockCheckExpectation[T]

	callArgs []*ValidatorMockCheckParams[T]
	mutex    sync.RWMutex
}

// ValidatorMockCheckExpectation specifies expectation struct of the Validator.Check
type ValidatorMockCheckExpectation[T Number] struct {
	mock    *ValidatorMock[T]
	params  *ValidatorMockCheckParams[T]
	results *ValidatorMockCheckResults[T]
	Counter uint64
}

// ValidatorMockCheckParams contains parameters of the Validator.Check
type ValidatorMockCheckParams[T Number] struct {
	value int64
}

// ValidatorMockCheckResults contains results of the Validator.Check
type ValidatorMockCheckResults[T Number] struct {
	b1  bool
	err error
}

// Expect sets up expected params for Validator.Check
func (mmCheck *mValidatorMockCheck[T]) Expect(value int64) *mValidatorMockCheck[T] {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("ValidatorMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &ValidatorMockCheckExpectation[T]{}
	}

	mmCheck.defaultExpectation.params = &ValidatorMockCheckParams[T]{value}
	for _, e := range mmCheck.expectations {
		if minimock.Equal(e.params, mmCheck.defaultExpectation.params) {
			mmCheck.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCheck.defaultExpectation.params)
		}
	}

	return mmCheck
}

// Inspect accepts an inspector function that has same arguments as the Validator.Check
func (mmCheck *mValidatorMockCheck[T]) Inspect(f func(value int64)) *mValidatorMockCheck[T] {
	if mmCheck.mock.inspectFuncCheck != nil {
		mmCheck.mock.t.Fatalf("Inspect function is already set for ValidatorMock.Check")
	}

	mmCheck.mock.inspectFuncCheck = f

	return mmCheck
}

// Return sets up results that will be returned by Validator.Check
func (mmCheck *mValidatorMockCheck[T]) Return(b1 bool, err error) *ValidatorMock[T] {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("ValidatorMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &ValidatorMockCheckExpectation[T]{mock: mmCheck.mock}
	}
	mmCheck.defaultExpectation.results = &ValidatorMockCheckResults[T]{b1, err}
	return mmCheck.mock
}

// Set uses given function f to mock the Validator.Check method
func (mmCheck *mValidatorMockCheck[T]) Set(f func(value int64) (b1 bool, err error)) *ValidatorMock[T] {
	if mmCheck.defaultExpectation != nil {
		mmCheck.mock.t.Fatalf("Default expectation is already set for the Validator.Check method")
	}

	if len(mmCheck.expectations) > 0 {
		mmCheck.mock.t.Fatalf("Some expectations are already set for the Validator.Check method")
	}

	mmCheck.mock.funcCheck = f
	return mmCheck.mock
}

// When sets expectation for the Validator.Check which will trigger the result defined by the following
// Then helper
func (mmCheck *mValidatorMockCheck[T]) When(value int64) *ValidatorMockCheckExpectation[T] {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("ValidatorMock.Check mock is already set by Set")
	}

	expectation := &ValidatorMockCheckExpectation[T]{
		mock:   mmCheck.mock,
		params: &ValidatorMockCheckParams[T]{value},
	}
	mmCheck.expectations = append(mmCheck.expectations, expectation)
	return expectation
}

// Then sets up Validator.Check return parameters for the expectation previously defined by the When method
func (e *ValidatorMockCheckExpectation[T]) Then(b1 bool, err error) *ValidatorMock[T] {
	e.results = &ValidatorMockCheckResults[T]{b1, err}
	return e.mock
}

// Check implements Validator
func (mmCheck *ValidatorMock[T]) Check(value int64) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmCheck.beforeCheckCounter, 1)
	defer mm_atomic.AddUint64(&mmCheck.afterCheckCounter, 1)

	if mmCheck.inspectFuncCheck != nil {
		mmCheck.inspectFuncCheck(value)
	}

	mm_params := &ValidatorMockCheckParams[T]{value}

	// Record call args
	mmCheck.CheckMock.mutex.Lock()
	mmCheck.CheckMock.callArgs = append(mmCheck.CheckMock.callArgs, mm_params)
	mmCheck.CheckMock.mutex.Unlock()

	for _, e := range mmCheck.CheckMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmCheck.CheckMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCheck.CheckMock.defaultExpectation.Counter, 1)
		mm_want := mmCheck.CheckMock.defaultExpectation.params
		mm_got := ValidatorMockCheckParams[T]{value}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCheck.t.Errorf("ValidatorMock.Check got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCheck.CheckMock.defaultExpectation.results
		if mm_results == nil {
			mmCheck.t.Fatal("No results are set for the ValidatorMock.Check")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmCheck.funcCheck != nil {
		return mmCheck.funcCheck(value)
	}
	mmCheck.t.Fatalf("Unexpected call to ValidatorMock.Check. %v", value)
	return
}

// CheckAfterCounter returns a count of finished ValidatorMock.Check invocations
func (mmCheck *ValidatorMock[T]) CheckAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheck.afterCheckCounter)
}

// CheckBeforeCounter returns a count of ValidatorMock.Check invocations
func (mmCheck *ValidatorMock[T]) CheckBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheck.beforeCheckCounter)
}

// Calls returns a list of arguments used in each call to ValidatorMock.Check.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCheck *mValidatorMockCheck[T]) Calls() []*ValidatorMockCheckParams[T] {
	mmCheck.mutex.RLock()

	argCopy := make([]*ValidatorMockCheckParams[T], len(mmCheck.callArgs))
	copy(argCopy, mmCheck.callArgs)

	mmCheck.mutex.RUnlock()

	return argCopy
}

// MinimockCheckDone returns true if the count of the Check invocations corresponds
// the number of defined expectations
func (m *ValidatorMock[T]) MinimockCheckDone() bool {
	for _, e := range m.CheckMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CheckMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCheckCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheck != nil && mm_atomic.LoadUint64(&m.afterCheckCounter) < 1 {
		return false
	}
	return true
}

// MinimockCheckInspect logs each unmet expectation
func (m *ValidatorMock[T]) MinimockCheckInspect() {
	for _, e := range m.CheckMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ValidatorMock.Check with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CheckMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCheckCounter) < 1 {
		if m.CheckMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ValidatorMock.Check")
		} else {
			m.t.Errorf("Expected call to ValidatorMock.Check with params: %#v", *m.CheckMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheck != nil && mm_atomic.LoadUint64(&m.afterCheckCounter) < 1 {
		m.t.Error("Expected call to ValidatorMock.Check")
	}
}

type mValidatorMockCheckGeneric[T Number] struct {
	mock               *ValidatorMock[T]
	defaultExpectation *ValidatorMockCheckGenericExpectation[T]
	expectations       []*ValidatorMockCheckGenericExpectation[T]

	callArgs []*ValidatorMockCheckGenericParams[T]
	mutex    sync.RWMutex
}

// ValidatorMockCheckGenericExpectation specifies expectation struct of the Validator.CheckGeneric
type ValidatorMockCheckGenericExpectation[T Number] struct {
	mock    *ValidatorMock[T]
	params  *ValidatorMockCheckGenericParams[T]
	results *ValidatorMockCheckGenericResults[T]
	Counter uint64
}

// ValidatorMockCheckGenericParams contains parameters of the Validator.CheckGeneric
type ValidatorMockCheckGenericParams[T Number] struct {
	value T
}

// ValidatorMockCheckGenericResults contains results of the Validator.CheckGeneric
type ValidatorMockCheckGenericResults[T Number] struct {
	b1  bool
	err error
}

// Expect sets up expected params for Validator.CheckGeneric
func (mmCheckGeneric *mValidatorMockCheckGeneric[T]) Expect(value T) *mValidatorMockCheckGeneric[T] {
	if mmCheckGeneric.mock.funcCheckGeneric != nil {
		mmCheckGeneric.mock.t.Fatalf("ValidatorMock.CheckGeneric mock is already set by Set")
	}

	if mmCheckGeneric.defaultExpectation == nil {
		mmCheckGeneric.defaultExpectation = &ValidatorMockCheckGenericExpectation[T]{}
	}

	mmCheckGeneric.defaultExpectation.params = &ValidatorMockCheckGenericParams[T]{value}
	for _, e := range mmCheckGeneric.expectations {
		if minimock.Equal(e.params, mmCheckGeneric.defaultExpectation.params) {
			mmCheckGeneric.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCheckGeneric.defaultExpectation.params)
		}
	}

	return mmCheckGeneric
}

// Inspect accepts an inspector function that has same arguments as the Validator.CheckGeneric
func (mmCheckGeneric *mValidatorMockCheckGeneric[T]) Inspect(f func(value T)) *mValidatorMockCheckGeneric[T] {
	if mmCheckGeneric.mock.inspectFuncCheckGeneric != nil {
		mmCheckGeneric.mock.t.Fatalf("Inspect function is already set for ValidatorMock.CheckGeneric")
	}

	mmCheckGeneric.mock.inspectFuncCheckGeneric = f

	return mmCheckGeneric
}

// Return sets up results that will be returned by Validator.CheckGeneric
func (mmCheckGeneric *mValidatorMockCheckGeneric[T]) Return(b1 bool, err error) *ValidatorMock[T] {
	if mmCheckGeneric.mock.funcCheckGeneric != nil {
		mmCheckGeneric.mock.t.Fatalf("ValidatorMock.CheckGeneric mock is already set by Set")
	}

	if mmCheckGeneric.defaultExpectation == nil {
		mmCheckGeneric.defaultExpectation = &ValidatorMockCheckGenericExpectation[T]{mock: mmCheckGeneric.mock}
	}
	mmCheckGeneric.defaultExpectation.results = &ValidatorMockCheckGenericResults[T]{b1, err}
	return mmCheckGeneric.mock
}

// Set uses given function f to mock the Validator.CheckGeneric method
func (mmCheckGeneric *mValidatorMockCheckGeneric[T]) Set(f func(value T) (b1 bool, err error)) *ValidatorMock[T] {
	if mmCheckGeneric.defaultExpectation != nil {
		mmCheckGeneric.mock.t.Fatalf("Default expectation is already set for the Validator.CheckGeneric method")
	}

	if len(mmCheckGeneric.expectations) > 0 {
		mmCheckGeneric.mock.t.Fatalf("Some expectations are already set for the Validator.CheckGeneric method")
	}

	mmCheckGeneric.mock.funcCheckGeneric = f
	return mmCheckGeneric.mock
}

// When sets expectation for the Validator.CheckGeneric which will trigger the result defined by the following
// Then helper
func (mmCheckGeneric *mValidatorMockCheckGeneric[T]) When(value T) *ValidatorMockCheckGenericExpectation[T] {
	if mmCheckGeneric.mock.funcCheckGeneric != nil {
		mmCheckGeneric.mock.t.Fatalf("ValidatorMock.CheckGeneric mock is already set by Set")
	}

	expectation := &ValidatorMockCheckGenericExpectation[T]{
		mock:   mmCheckGeneric.mock,
		params: &ValidatorMockCheckGenericParams[T]{value},
	}
	mmCheckGeneric.expectations = append(mmCheckGeneric.expectations, expectation)
	return expectation
}

// Then sets up Validator.CheckGeneric return parameters for the expectation previously defined by the When method
func (e *ValidatorMockCheckGenericExpectation[T]) Then(b1 bool, err error) *ValidatorMock[T] {
	e.results = &ValidatorMockCheckGenericResults[T]{b1, err}
	return e.mock
}

// CheckGeneric implements Validator
func (mmCheckGeneric *ValidatorMock[T]) CheckGeneric(value T) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmCheckGeneric.beforeCheckGenericCounter, 1)
	defer mm_atomic.AddUint64(&mmCheckGeneric.afterCheckGenericCounter, 1)

	if mmCheckGeneric.inspectFuncCheckGeneric != nil {
		mmCheckGeneric.inspectFuncCheckGeneric(value)
	}

	mm_params := &ValidatorMockCheckGenericParams[T]{value}

	// Record call args
	mmCheckGeneric.CheckGenericMock.mutex.Lock()
	mmCheckGeneric.CheckGenericMock.callArgs = append(mmCheckGeneric.CheckGenericMock.callArgs, mm_params)
	mmCheckGeneric.CheckGenericMock.mutex.Unlock()

	for _, e := range mmCheckGeneric.CheckGenericMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmCheckGeneric.CheckGenericMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCheckGeneric.CheckGenericMock.defaultExpectation.Counter, 1)
		mm_want := mmCheckGeneric.CheckGenericMock.defaultExpectation.params
		mm_got := ValidatorMockCheckGenericParams[T]{value}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCheckGeneric.t.Errorf("ValidatorMock.CheckGeneric got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCheckGeneric.CheckGenericMock.defaultExpectation.results
		if mm_results == nil {
			mmCheckGeneric.t.Fatal("No results are set for the ValidatorMock.CheckGeneric")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmCheckGeneric.funcCheckGeneric != nil {
		return mmCheckGeneric.funcCheckGeneric(value)
	}
	mmCheckGeneric.t.Fatalf("Unexpected call to ValidatorMock.CheckGeneric. %v", value)
	return
}

// CheckGenericAfterCounter returns a count of finished ValidatorMock.CheckGeneric invocations
func (mmCheckGeneric *ValidatorMock[T]) CheckGenericAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheckGeneric.afterCheckGenericCounter)
}

// CheckGenericBeforeCounter returns a count of ValidatorMock.CheckGeneric invocations
func (mmCheckGeneric *ValidatorMock[T]) CheckGenericBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheckGeneric.beforeCheckGenericCounter)
}

// Calls returns a list of arguments used in each call to ValidatorMock.CheckGeneric.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCheckGeneric *mValidatorMockCheckGeneric[T]) Calls() []*ValidatorMockCheckGenericParams[T] {
	mmCheckGeneric.mutex.RLock()

	argCopy := make([]*ValidatorMockCheckGenericParams[T], len(mmCheckGeneric.callArgs))
	copy(argCopy, mmCheckGeneric.callArgs)

	mmCheckGeneric.mutex.RUnlock()

	return argCopy
}

// MinimockCheckGenericDone returns true if the count of the CheckGeneric invocations corresponds
// the number of defined expectations
func (m *ValidatorMock[T]) MinimockCheckGenericDone() bool {
	for _, e := range m.CheckGenericMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CheckGenericMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCheckGenericCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheckGeneric != nil && mm_atomic.LoadUint64(&m.afterCheckGenericCounter) < 1 {
		return false
	}
	return true
}

// MinimockCheckGenericInspect logs each unmet expectation
func (m *ValidatorMock[T]) MinimockCheckGenericInspect() {
	for _, e := range m.CheckGenericMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ValidatorMock.CheckGeneric with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CheckGenericMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCheckGenericCounter) < 1 {
		if m.CheckGenericMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ValidatorMock.CheckGeneric")
		} else {
			m.t.Errorf("Expected call to ValidatorMock.CheckGeneric with params: %#v", *m.CheckGenericMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheckGeneric != nil && mm_atomic.LoadUint64(&m.afterCheckGenericCounter) < 1 {
		m.t.Error("Expected call to ValidatorMock.CheckGeneric")
	}
}

type mValidatorMockCheckGenerics[T Number] struct {
	mock               *ValidatorMock[T]
	defaultExpectation *ValidatorMockCheckGenericsExpectation[T]
	expectations       []*ValidatorMockCheckGenericsExpectation[T]

	callArgs []*ValidatorMockCheckGenericsParams[T]
	mutex    sync.RWMutex
}

// ValidatorMockCheckGenericsExpectation specifies expectation struct of the Validator.CheckGenerics
type ValidatorMockCheckGenericsExpectation[T Number] struct {
	mock    *ValidatorMock[T]
	params  *ValidatorMockCheckGenericsParams[T]
	results *ValidatorMockCheckGenericsResults[T]
	Counter uint64
}

// ValidatorMockCheckGenericsParams contains parameters of the Validator.CheckGenerics
type ValidatorMockCheckGenericsParams[T Number] struct {
	values []T
}

// ValidatorMockCheckGenericsResults contains results of the Validator.CheckGenerics
type ValidatorMockCheckGenericsResults[T Number] struct {
	b1  bool
	err error
}

// Expect sets up expected params for Validator.CheckGenerics
func (mmCheckGenerics *mValidatorMockCheckGenerics[T]) Expect(values []T) *mValidatorMockCheckGenerics[T] {
	if mmCheckGenerics.mock.funcCheckGenerics != nil {
		mmCheckGenerics.mock.t.Fatalf("ValidatorMock.CheckGenerics mock is already set by Set")
	}

	if mmCheckGenerics.defaultExpectation == nil {
		mmCheckGenerics.defaultExpectation = &ValidatorMockCheckGenericsExpectation[T]{}
	}

	mmCheckGenerics.defaultExpectation.params = &ValidatorMockCheckGenericsParams[T]{values}
	for _, e := range mmCheckGenerics.expectations {
		if minimock.Equal(e.params, mmCheckGenerics.defaultExpectation.params) {
			mmCheckGenerics.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCheckGenerics.defaultExpectation.params)
		}
	}

	return mmCheckGenerics
}

// Inspect accepts an inspector function that has same arguments as the Validator.CheckGenerics
func (mmCheckGenerics *mValidatorMockCheckGenerics[T]) Inspect(f func(values []T)) *mValidatorMockCheckGenerics[T] {
	if mmCheckGenerics.mock.inspectFuncCheckGenerics != nil {
		mmCheckGenerics.mock.t.Fatalf("Inspect function is already set for ValidatorMock.CheckGenerics")
	}

	mmCheckGenerics.mock.inspectFuncCheckGenerics = f

	return mmCheckGenerics
}

// Return sets up results that will be returned by Validator.CheckGenerics
func (mmCheckGenerics *mValidatorMockCheckGenerics[T]) Return(b1 bool, err error) *ValidatorMock[T] {
	if mmCheckGenerics.mock.funcCheckGenerics != nil {
		mmCheckGenerics.mock.t.Fatalf("ValidatorMock.CheckGenerics mock is already set by Set")
	}

	if mmCheckGenerics.defaultExpectation == nil {
		mmCheckGenerics.defaultExpectation = &ValidatorMockCheckGenericsExpectation[T]{mock: mmCheckGenerics.mock}
	}
	mmCheckGenerics.defaultExpectation.results = &ValidatorMockCheckGenericsResults[T]{b1, err}
	return mmCheckGenerics.mock
}

// Set uses given function f to mock the Validator.CheckGenerics method
func (mmCheckGenerics *mValidatorMockCheckGenerics[T]) Set(f func(values []T) (b1 bool, err error)) *ValidatorMock[T] {
	if mmCheckGenerics.defaultExpectation != nil {
		mmCheckGenerics.mock.t.Fatalf("Default expectation is already set for the Validator.CheckGenerics method")
	}

	if len(mmCheckGenerics.expectations) > 0 {
		mmCheckGenerics.mock.t.Fatalf("Some expectations are already set for the Validator.CheckGenerics method")
	}

	mmCheckGenerics.mock.funcCheckGenerics = f
	return mmCheckGenerics.mock
}

// When sets expectation for the Validator.CheckGenerics which will trigger the result defined by the following
// Then helper
func (mmCheckGenerics *mValidatorMockCheckGenerics[T]) When(values []T) *ValidatorMockCheckGenericsExpectation[T] {
	if mmCheckGenerics.mock.funcCheckGenerics != nil {
		mmCheckGenerics.mock.t.Fatalf("ValidatorMock.CheckGenerics mock is already set by Set")
	}

	expectation := &ValidatorMockCheckGenericsExpectation[T]{
		mock:   mmCheckGenerics.mock,
		params: &ValidatorMockCheckGenericsParams[T]{values},
	}
	mmCheckGenerics.expectations = append(mmCheckGenerics.expectations, expectation)
	return expectation
}

// Then sets up Validator.CheckGenerics return parameters for the expectation previously defined by the When method
func (e *ValidatorMockCheckGenericsExpectation[T]) Then(b1 bool, err error) *ValidatorMock[T] {
	e.results = &ValidatorMockCheckGenericsResults[T]{b1, err}
	return e.mock
}

// CheckGenerics implements Validator
func (mmCheckGenerics *ValidatorMock[T]) CheckGenerics(values []T) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmCheckGenerics.beforeCheckGenericsCounter, 1)
	defer mm_atomic.AddUint64(&mmCheckGenerics.afterCheckGenericsCounter, 1)

	if mmCheckGenerics.inspectFuncCheckGenerics != nil {
		mmCheckGenerics.inspectFuncCheckGenerics(values)
	}

	mm_params := &ValidatorMockCheckGenericsParams[T]{values}

	// Record call args
	mmCheckGenerics.CheckGenericsMock.mutex.Lock()
	mmCheckGenerics.CheckGenericsMock.callArgs = append(mmCheckGenerics.CheckGenericsMock.callArgs, mm_params)
	mmCheckGenerics.CheckGenericsMock.mutex.Unlock()

	for _, e := range mmCheckGenerics.CheckGenericsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmCheckGenerics.CheckGenericsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCheckGenerics.CheckGenericsMock.defaultExpectation.Counter, 1)
		mm_want := mmCheckGenerics.CheckGenericsMock.defaultExpectation.params
		mm_got := ValidatorMockCheckGenericsParams[T]{values}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCheckGenerics.t.Errorf("ValidatorMock.CheckGenerics got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCheckGenerics.CheckGenericsMock.defaultExpectation.results
		if mm_results == nil {
			mmCheckGenerics.t.Fatal("No results are set for the ValidatorMock.CheckGenerics")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmCheckGenerics.funcCheckGenerics != nil {
		return mmCheckGenerics.funcCheckGenerics(values)
	}
	mmCheckGenerics.t.Fatalf("Unexpected call to ValidatorMock.CheckGenerics. %v", values)
	return
}

// CheckGenericsAfterCounter returns a count of finished ValidatorMock.CheckGenerics invocations
func (mmCheckGenerics *ValidatorMock[T]) CheckGenericsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheckGenerics.afterCheckGenericsCounter)
}

// CheckGenericsBeforeCounter returns a count of ValidatorMock.CheckGenerics invocations
func (mmCheckGenerics *ValidatorMock[T]) CheckGenericsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheckGenerics.beforeCheckGenericsCounter)
}

// Calls returns a list of arguments used in each call to ValidatorMock.CheckGenerics.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCheckGenerics *mValidatorMockCheckGenerics[T]) Calls() []*ValidatorMockCheckGenericsParams[T] {
	mmCheckGenerics.mutex.RLock()

	argCopy := make([]*ValidatorMockCheckGenericsParams[T], len(mmCheckGenerics.callArgs))
	copy(argCopy, mmCheckGenerics.callArgs)

	mmCheckGenerics.mutex.RUnlock()

	return argCopy
}

// MinimockCheckGenericsDone returns true if the count of the CheckGenerics invocations corresponds
// the number of defined expectations
func (m *ValidatorMock[T]) MinimockCheckGenericsDone() bool {
	for _, e := range m.CheckGenericsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CheckGenericsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCheckGenericsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheckGenerics != nil && mm_atomic.LoadUint64(&m.afterCheckGenericsCounter) < 1 {
		return false
	}
	return true
}

// MinimockCheckGenericsInspect logs each unmet expectation
func (m *ValidatorMock[T]) MinimockCheckGenericsInspect() {
	for _, e := range m.CheckGenericsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ValidatorMock.CheckGenerics with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CheckGenericsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCheckGenericsCounter) < 1 {
		if m.CheckGenericsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ValidatorMock.CheckGenerics")
		} else {
			m.t.Errorf("Expected call to ValidatorMock.CheckGenerics with params: %#v", *m.CheckGenericsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheckGenerics != nil && mm_atomic.LoadUint64(&m.afterCheckGenericsCounter) < 1 {
		m.t.Error("Expected call to ValidatorMock.CheckGenerics")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ValidatorMock[T]) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCheckInspect()

		m.MinimockCheckGenericInspect()

		m.MinimockCheckGenericsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ValidatorMock[T]) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ValidatorMock[T]) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCheckDone() &&
		m.MinimockCheckGenericDone() &&
		m.MinimockCheckGenericsDone()
}
