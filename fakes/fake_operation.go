// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/jamesjoshuahill/go-retry"
)

type FakeOperation struct {
	TryStub          func()
	tryMutex         sync.RWMutex
	tryArgsForCall   []struct{}
	RetryStub        func() bool
	retryMutex       sync.RWMutex
	retryArgsForCall []struct{}
	retryReturns     struct {
		result1 bool
	}
	retryReturnsOnCall map[int]struct {
		result1 bool
	}
	ErrorStub        func() error
	errorMutex       sync.RWMutex
	errorArgsForCall []struct{}
	errorReturns     struct {
		result1 error
	}
	errorReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOperation) Try() {
	fake.tryMutex.Lock()
	fake.tryArgsForCall = append(fake.tryArgsForCall, struct{}{})
	fake.recordInvocation("Try", []interface{}{})
	fake.tryMutex.Unlock()
	if fake.TryStub != nil {
		fake.TryStub()
	}
}

func (fake *FakeOperation) TryCallCount() int {
	fake.tryMutex.RLock()
	defer fake.tryMutex.RUnlock()
	return len(fake.tryArgsForCall)
}

func (fake *FakeOperation) Retry() bool {
	fake.retryMutex.Lock()
	ret, specificReturn := fake.retryReturnsOnCall[len(fake.retryArgsForCall)]
	fake.retryArgsForCall = append(fake.retryArgsForCall, struct{}{})
	fake.recordInvocation("Retry", []interface{}{})
	fake.retryMutex.Unlock()
	if fake.RetryStub != nil {
		return fake.RetryStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.retryReturns.result1
}

func (fake *FakeOperation) RetryCallCount() int {
	fake.retryMutex.RLock()
	defer fake.retryMutex.RUnlock()
	return len(fake.retryArgsForCall)
}

func (fake *FakeOperation) RetryReturns(result1 bool) {
	fake.RetryStub = nil
	fake.retryReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOperation) RetryReturnsOnCall(i int, result1 bool) {
	fake.RetryStub = nil
	if fake.retryReturnsOnCall == nil {
		fake.retryReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.retryReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOperation) Error() error {
	fake.errorMutex.Lock()
	ret, specificReturn := fake.errorReturnsOnCall[len(fake.errorArgsForCall)]
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct{}{})
	fake.recordInvocation("Error", []interface{}{})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		return fake.ErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.errorReturns.result1
}

func (fake *FakeOperation) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeOperation) ErrorReturns(result1 error) {
	fake.ErrorStub = nil
	fake.errorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOperation) ErrorReturnsOnCall(i int, result1 error) {
	fake.ErrorStub = nil
	if fake.errorReturnsOnCall == nil {
		fake.errorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.errorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOperation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tryMutex.RLock()
	defer fake.tryMutex.RUnlock()
	fake.retryMutex.RLock()
	defer fake.retryMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeOperation) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ retry.Operation = new(FakeOperation)
