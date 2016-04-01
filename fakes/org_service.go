// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/rainmaker"
)

type OrgService struct {
	ListStub        func(token string) (rainmaker.OrganizationsList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		token string
	}
	listReturns struct {
		result1 rainmaker.OrganizationsList
		result2 error
	}
}

func (fake *OrgService) List(token string) (rainmaker.OrganizationsList, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		token string
	}{token})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(token)
	} else {
		return fake.listReturns.result1, fake.listReturns.result2
	}
}

func (fake *OrgService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *OrgService) ListArgsForCall(i int) string {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].token
}

func (fake *OrgService) ListReturns(result1 rainmaker.OrganizationsList, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 rainmaker.OrganizationsList
		result2 error
	}{result1, result2}
}