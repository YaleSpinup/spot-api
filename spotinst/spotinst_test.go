package spotinst

import (
	"reflect"
	"testing"

	"github.com/YaleSpinup/spot-api/common"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance"
)

// mockElastigroupClient is a fake elastigroup client
type mockElastigroupClient struct {
	elastigroup.Service
	t   *testing.T
	err error
}

// mockManagedInstanceClient is a fake managedinstance client
type mockManagedInstanceClient struct {
	managedinstance.Service
	t   *testing.T
	err error
}

func newmockElastigroupClient(t *testing.T, err error) elastigroup.Service {
	return &mockElastigroupClient{
		t:   t,
		err: err,
	}
}

func newmockManagedInstanceClient(t *testing.T, err error) managedinstance.Service {
	return &mockManagedInstanceClient{
		t:   t,
		err: err,
	}
}

func TestNewElastigroupSession(t *testing.T) {
	e := NewElastigroupSession(common.Account{})
	if to := reflect.TypeOf(e).String(); to != "spotinst.Elastigroup" {
		t.Errorf("expected type to be 'spotinst.Elastigroup', got %s", to)
	}
}

func TestNewManagedInstanceSession(t *testing.T) {
	m := NewManagedInstanceSession(common.Account{})
	if to := reflect.TypeOf(m).String(); to != "spotinst.ManagedInstance" {
		t.Errorf("expected type to be 'spotinst.ManagedInstance', got %s", to)
	}
}
