package spotinst

import (
	"reflect"
	"testing"

	"github.com/YaleSpinup/spot-api/common"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
)

// mockElastigroupClient is a fake elastigroup client
type mockElastigroupClient struct {
	elastigroup.Service
	t   *testing.T
	err error
}

func newmockElastigroupClient(t *testing.T, err error) elastigroup.Service {
	return &mockElastigroupClient{
		t:   t,
		err: err,
	}
}

func TestNewElastigroupSession(t *testing.T) {
	e := NewElastigroupSession(common.Account{})
	if to := reflect.TypeOf(e).String(); to != "spotinst.Elastigroup" {
		t.Errorf("expected type to be 'elastigroup.Elastigroup', got %s", to)
	}
}
