// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"database-backup-restore/database"
	"sync"
)

type FakeInteractor struct {
	ActionStub        func(string) error
	actionMutex       sync.RWMutex
	actionArgsForCall []struct {
		arg1 string
	}
	actionReturns struct {
		result1 error
	}
	actionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInteractor) Action(arg1 string) error {
	fake.actionMutex.Lock()
	ret, specificReturn := fake.actionReturnsOnCall[len(fake.actionArgsForCall)]
	fake.actionArgsForCall = append(fake.actionArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ActionStub
	fakeReturns := fake.actionReturns
	fake.recordInvocation("Action", []interface{}{arg1})
	fake.actionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeInteractor) ActionCallCount() int {
	fake.actionMutex.RLock()
	defer fake.actionMutex.RUnlock()
	return len(fake.actionArgsForCall)
}

func (fake *FakeInteractor) ActionCalls(stub func(string) error) {
	fake.actionMutex.Lock()
	defer fake.actionMutex.Unlock()
	fake.ActionStub = stub
}

func (fake *FakeInteractor) ActionArgsForCall(i int) string {
	fake.actionMutex.RLock()
	defer fake.actionMutex.RUnlock()
	argsForCall := fake.actionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInteractor) ActionReturns(result1 error) {
	fake.actionMutex.Lock()
	defer fake.actionMutex.Unlock()
	fake.ActionStub = nil
	fake.actionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInteractor) ActionReturnsOnCall(i int, result1 error) {
	fake.actionMutex.Lock()
	defer fake.actionMutex.Unlock()
	fake.ActionStub = nil
	if fake.actionReturnsOnCall == nil {
		fake.actionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.actionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInteractor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.actionMutex.RLock()
	defer fake.actionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInteractor) recordInvocation(key string, args []interface{}) {
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

var _ database.Interactor = new(FakeInteractor)
