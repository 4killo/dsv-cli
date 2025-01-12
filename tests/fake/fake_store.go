// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/DelineaXPM/dsv-cli/errors"
	"github.com/DelineaXPM/dsv-cli/store"
)

type FakeStore struct {
	DeleteStub        func(string) *errors.ApiError
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 *errors.ApiError
	}
	deleteReturnsOnCall map[int]struct {
		result1 *errors.ApiError
	}
	GetStub        func(string, interface{}) *errors.ApiError
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	getReturns struct {
		result1 *errors.ApiError
	}
	getReturnsOnCall map[int]struct {
		result1 *errors.ApiError
	}
	ListStub        func(string) ([]string, *errors.ApiError)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 string
	}
	listReturns struct {
		result1 []string
		result2 *errors.ApiError
	}
	listReturnsOnCall map[int]struct {
		result1 []string
		result2 *errors.ApiError
	}
	StoreStub        func(string, interface{}) *errors.ApiError
	storeMutex       sync.RWMutex
	storeArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	storeReturns struct {
		result1 *errors.ApiError
	}
	storeReturnsOnCall map[int]struct {
		result1 *errors.ApiError
	}
	StoreStringStub        func(string, string) *errors.ApiError
	storeStringMutex       sync.RWMutex
	storeStringArgsForCall []struct {
		arg1 string
		arg2 string
	}
	storeStringReturns struct {
		result1 *errors.ApiError
	}
	storeStringReturnsOnCall map[int]struct {
		result1 *errors.ApiError
	}
	WipeStub        func(string) *errors.ApiError
	wipeMutex       sync.RWMutex
	wipeArgsForCall []struct {
		arg1 string
	}
	wipeReturns struct {
		result1 *errors.ApiError
	}
	wipeReturnsOnCall map[int]struct {
		result1 *errors.ApiError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStore) Delete(arg1 string) *errors.ApiError {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeStore) DeleteCalls(stub func(string) *errors.ApiError) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeStore) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) DeleteReturns(result1 *errors.ApiError) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) DeleteReturnsOnCall(i int, result1 *errors.ApiError) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *errors.ApiError
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) Get(arg1 string, arg2 interface{}) *errors.ApiError {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeStore) GetCalls(stub func(string, interface{}) *errors.ApiError) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeStore) GetArgsForCall(i int) (string, interface{}) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) GetReturns(result1 *errors.ApiError) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) GetReturnsOnCall(i int, result1 *errors.ApiError) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *errors.ApiError
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) List(arg1 string) ([]string, *errors.ApiError) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeStore) ListCalls(stub func(string) ([]string, *errors.ApiError)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeStore) ListArgsForCall(i int) string {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) ListReturns(result1 []string, result2 *errors.ApiError) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []string
		result2 *errors.ApiError
	}{result1, result2}
}

func (fake *FakeStore) ListReturnsOnCall(i int, result1 []string, result2 *errors.ApiError) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 *errors.ApiError
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []string
		result2 *errors.ApiError
	}{result1, result2}
}

func (fake *FakeStore) Store(arg1 string, arg2 interface{}) *errors.ApiError {
	fake.storeMutex.Lock()
	ret, specificReturn := fake.storeReturnsOnCall[len(fake.storeArgsForCall)]
	fake.storeArgsForCall = append(fake.storeArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.StoreStub
	fakeReturns := fake.storeReturns
	fake.recordInvocation("Store", []interface{}{arg1, arg2})
	fake.storeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) StoreCallCount() int {
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	return len(fake.storeArgsForCall)
}

func (fake *FakeStore) StoreCalls(stub func(string, interface{}) *errors.ApiError) {
	fake.storeMutex.Lock()
	defer fake.storeMutex.Unlock()
	fake.StoreStub = stub
}

func (fake *FakeStore) StoreArgsForCall(i int) (string, interface{}) {
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	argsForCall := fake.storeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) StoreReturns(result1 *errors.ApiError) {
	fake.storeMutex.Lock()
	defer fake.storeMutex.Unlock()
	fake.StoreStub = nil
	fake.storeReturns = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) StoreReturnsOnCall(i int, result1 *errors.ApiError) {
	fake.storeMutex.Lock()
	defer fake.storeMutex.Unlock()
	fake.StoreStub = nil
	if fake.storeReturnsOnCall == nil {
		fake.storeReturnsOnCall = make(map[int]struct {
			result1 *errors.ApiError
		})
	}
	fake.storeReturnsOnCall[i] = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) StoreString(arg1 string, arg2 string) *errors.ApiError {
	fake.storeStringMutex.Lock()
	ret, specificReturn := fake.storeStringReturnsOnCall[len(fake.storeStringArgsForCall)]
	fake.storeStringArgsForCall = append(fake.storeStringArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.StoreStringStub
	fakeReturns := fake.storeStringReturns
	fake.recordInvocation("StoreString", []interface{}{arg1, arg2})
	fake.storeStringMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) StoreStringCallCount() int {
	fake.storeStringMutex.RLock()
	defer fake.storeStringMutex.RUnlock()
	return len(fake.storeStringArgsForCall)
}

func (fake *FakeStore) StoreStringCalls(stub func(string, string) *errors.ApiError) {
	fake.storeStringMutex.Lock()
	defer fake.storeStringMutex.Unlock()
	fake.StoreStringStub = stub
}

func (fake *FakeStore) StoreStringArgsForCall(i int) (string, string) {
	fake.storeStringMutex.RLock()
	defer fake.storeStringMutex.RUnlock()
	argsForCall := fake.storeStringArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) StoreStringReturns(result1 *errors.ApiError) {
	fake.storeStringMutex.Lock()
	defer fake.storeStringMutex.Unlock()
	fake.StoreStringStub = nil
	fake.storeStringReturns = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) StoreStringReturnsOnCall(i int, result1 *errors.ApiError) {
	fake.storeStringMutex.Lock()
	defer fake.storeStringMutex.Unlock()
	fake.StoreStringStub = nil
	if fake.storeStringReturnsOnCall == nil {
		fake.storeStringReturnsOnCall = make(map[int]struct {
			result1 *errors.ApiError
		})
	}
	fake.storeStringReturnsOnCall[i] = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) Wipe(arg1 string) *errors.ApiError {
	fake.wipeMutex.Lock()
	ret, specificReturn := fake.wipeReturnsOnCall[len(fake.wipeArgsForCall)]
	fake.wipeArgsForCall = append(fake.wipeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.WipeStub
	fakeReturns := fake.wipeReturns
	fake.recordInvocation("Wipe", []interface{}{arg1})
	fake.wipeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStore) WipeCallCount() int {
	fake.wipeMutex.RLock()
	defer fake.wipeMutex.RUnlock()
	return len(fake.wipeArgsForCall)
}

func (fake *FakeStore) WipeCalls(stub func(string) *errors.ApiError) {
	fake.wipeMutex.Lock()
	defer fake.wipeMutex.Unlock()
	fake.WipeStub = stub
}

func (fake *FakeStore) WipeArgsForCall(i int) string {
	fake.wipeMutex.RLock()
	defer fake.wipeMutex.RUnlock()
	argsForCall := fake.wipeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) WipeReturns(result1 *errors.ApiError) {
	fake.wipeMutex.Lock()
	defer fake.wipeMutex.Unlock()
	fake.WipeStub = nil
	fake.wipeReturns = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) WipeReturnsOnCall(i int, result1 *errors.ApiError) {
	fake.wipeMutex.Lock()
	defer fake.wipeMutex.Unlock()
	fake.WipeStub = nil
	if fake.wipeReturnsOnCall == nil {
		fake.wipeReturnsOnCall = make(map[int]struct {
			result1 *errors.ApiError
		})
	}
	fake.wipeReturnsOnCall[i] = struct {
		result1 *errors.ApiError
	}{result1}
}

func (fake *FakeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.storeMutex.RLock()
	defer fake.storeMutex.RUnlock()
	fake.storeStringMutex.RLock()
	defer fake.storeStringMutex.RUnlock()
	fake.wipeMutex.RLock()
	defer fake.wipeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStore) recordInvocation(key string, args []interface{}) {
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

var _ store.Store = new(FakeStore)
