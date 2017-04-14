// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/silk/controller"
)

type LeaseRenewer struct {
	RenewSubnetLeaseStub        func(lease controller.Lease) error
	renewSubnetLeaseMutex       sync.RWMutex
	renewSubnetLeaseArgsForCall []struct {
		lease controller.Lease
	}
	renewSubnetLeaseReturns struct {
		result1 error
	}
	renewSubnetLeaseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LeaseRenewer) RenewSubnetLease(lease controller.Lease) error {
	fake.renewSubnetLeaseMutex.Lock()
	ret, specificReturn := fake.renewSubnetLeaseReturnsOnCall[len(fake.renewSubnetLeaseArgsForCall)]
	fake.renewSubnetLeaseArgsForCall = append(fake.renewSubnetLeaseArgsForCall, struct {
		lease controller.Lease
	}{lease})
	fake.recordInvocation("RenewSubnetLease", []interface{}{lease})
	fake.renewSubnetLeaseMutex.Unlock()
	if fake.RenewSubnetLeaseStub != nil {
		return fake.RenewSubnetLeaseStub(lease)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.renewSubnetLeaseReturns.result1
}

func (fake *LeaseRenewer) RenewSubnetLeaseCallCount() int {
	fake.renewSubnetLeaseMutex.RLock()
	defer fake.renewSubnetLeaseMutex.RUnlock()
	return len(fake.renewSubnetLeaseArgsForCall)
}

func (fake *LeaseRenewer) RenewSubnetLeaseArgsForCall(i int) controller.Lease {
	fake.renewSubnetLeaseMutex.RLock()
	defer fake.renewSubnetLeaseMutex.RUnlock()
	return fake.renewSubnetLeaseArgsForCall[i].lease
}

func (fake *LeaseRenewer) RenewSubnetLeaseReturns(result1 error) {
	fake.RenewSubnetLeaseStub = nil
	fake.renewSubnetLeaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *LeaseRenewer) RenewSubnetLeaseReturnsOnCall(i int, result1 error) {
	fake.RenewSubnetLeaseStub = nil
	if fake.renewSubnetLeaseReturnsOnCall == nil {
		fake.renewSubnetLeaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.renewSubnetLeaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LeaseRenewer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.renewSubnetLeaseMutex.RLock()
	defer fake.renewSubnetLeaseMutex.RUnlock()
	return fake.invocations
}

func (fake *LeaseRenewer) recordInvocation(key string, args []interface{}) {
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