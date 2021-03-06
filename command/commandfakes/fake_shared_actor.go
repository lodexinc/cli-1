// This file was generated by counterfeiter
package commandfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/command"
)

type FakeSharedActor struct {
	CheckTargetStub        func(config sharedaction.Config, targetedOrganizationRequired bool, targetedSpaceRequired bool) error
	checkTargetMutex       sync.RWMutex
	checkTargetArgsForCall []struct {
		config                       sharedaction.Config
		targetedOrganizationRequired bool
		targetedSpaceRequired        bool
	}
	checkTargetReturns struct {
		result1 error
	}
	checkTargetReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSharedActor) CheckTarget(config sharedaction.Config, targetedOrganizationRequired bool, targetedSpaceRequired bool) error {
	fake.checkTargetMutex.Lock()
	ret, specificReturn := fake.checkTargetReturnsOnCall[len(fake.checkTargetArgsForCall)]
	fake.checkTargetArgsForCall = append(fake.checkTargetArgsForCall, struct {
		config                       sharedaction.Config
		targetedOrganizationRequired bool
		targetedSpaceRequired        bool
	}{config, targetedOrganizationRequired, targetedSpaceRequired})
	fake.recordInvocation("CheckTarget", []interface{}{config, targetedOrganizationRequired, targetedSpaceRequired})
	fake.checkTargetMutex.Unlock()
	if fake.CheckTargetStub != nil {
		return fake.CheckTargetStub(config, targetedOrganizationRequired, targetedSpaceRequired)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkTargetReturns.result1
}

func (fake *FakeSharedActor) CheckTargetCallCount() int {
	fake.checkTargetMutex.RLock()
	defer fake.checkTargetMutex.RUnlock()
	return len(fake.checkTargetArgsForCall)
}

func (fake *FakeSharedActor) CheckTargetArgsForCall(i int) (sharedaction.Config, bool, bool) {
	fake.checkTargetMutex.RLock()
	defer fake.checkTargetMutex.RUnlock()
	return fake.checkTargetArgsForCall[i].config, fake.checkTargetArgsForCall[i].targetedOrganizationRequired, fake.checkTargetArgsForCall[i].targetedSpaceRequired
}

func (fake *FakeSharedActor) CheckTargetReturns(result1 error) {
	fake.CheckTargetStub = nil
	fake.checkTargetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSharedActor) CheckTargetReturnsOnCall(i int, result1 error) {
	fake.CheckTargetStub = nil
	if fake.checkTargetReturnsOnCall == nil {
		fake.checkTargetReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkTargetReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSharedActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkTargetMutex.RLock()
	defer fake.checkTargetMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSharedActor) recordInvocation(key string, args []interface{}) {
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

var _ command.SharedActor = new(FakeSharedActor)
