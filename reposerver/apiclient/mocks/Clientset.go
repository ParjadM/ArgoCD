// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	io "github.com/argoproj/argo-cd/engine/pkg/utils/io"
	apiclient "github.com/argoproj/argo-cd/reposerver/apiclient"

	mock "github.com/stretchr/testify/mock"
)

// Clientset is an autogenerated mock type for the Clientset type
type Clientset struct {
	mock.Mock
}

// NewRepoServerClient provides a mock function with given fields:
func (_m *Clientset) NewRepoServerClient() (io.Closer, apiclient.RepoServerServiceClient, error) {
	ret := _m.Called()

	var r0 io.Closer
	if rf, ok := ret.Get(0).(func() io.Closer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Closer)
		}
	}

	var r1 apiclient.RepoServerServiceClient
	if rf, ok := ret.Get(1).(func() apiclient.RepoServerServiceClient); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apiclient.RepoServerServiceClient)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
