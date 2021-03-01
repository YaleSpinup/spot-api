package spotinst

import (
	"github.com/YaleSpinup/spot-api/common"
	log "github.com/sirupsen/logrus"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

// Elastigroup is a wrapper around the spotinst elastigroup service and the provider type (eg. aws, azure, gcp, etc)
type Elastigroup struct {
	Service elastigroup.Service
}

// ManagedInstance is a wrapper around the spotinst managedinstance service and the provider type (eg. aws, azure, gcp, etc)
type ManagedInstance struct {
	Service    managedinstance.Service
	DefaultVPC string
}

// NewElastigroupSession creates a new elastigroup session
func NewElastigroupSession(account common.Account) Elastigroup {
	e := Elastigroup{}
	log.Info("creating new spotinst session for elastigroups in SpotInst")
	sess := session.New(&spotinst.Config{
		Credentials: credentials.NewStaticCredentials(account.Token, account.Id),
	})
	e.Service = elastigroup.New(sess)
	return e
}

// NewManagedInstanceSession creates a new elastigroup session
func NewManagedInstanceSession(account common.Account) ManagedInstance {
	m := ManagedInstance{}
	log.Info("creating new spotinst session for managed instances in SpotInst")
	sess := session.New(&spotinst.Config{
		Credentials: credentials.NewStaticCredentials(account.Token, account.Id),
	})
	m.Service = managedinstance.New(sess)
	m.DefaultVPC = account.DefaultVPC
	return m
}
