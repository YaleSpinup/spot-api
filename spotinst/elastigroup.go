package spotinst

import (
	"context"

	"github.com/YaleSpinup/spot-api/apierror"
	log "github.com/sirupsen/logrus"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

func (e *Elastigroup) ListAWSElastigroups(ctx context.Context) ([]*aws.Group, error) {
	log.Info("listing aws elastigroups")

	// List all groups.
	output, err := e.Service.CloudProviderAWS().List(ctx, &aws.ListGroupsInput{})
	if err != nil {
		return nil, ErrCode("failed to list buckets", err)
	}

	return output.Groups, err
}

func (e *Elastigroup) GetAWSElastigroupByID(ctx context.Context, id string) (*aws.Group, error) {
	if id == "" {
		return nil, apierror.New(apierror.ErrBadRequest, "invalid input", nil)
	}

	output, err := e.Service.CloudProviderAWS().Read(ctx, &aws.ReadGroupInput{
		GroupID: spotinst.String(id),
	})

	if err != nil {
		return nil, ErrCode("failed to read elastigroup details", err)
	}

	return output.Group, nil
}

func (e *Elastigroup) CreateAWSElastigroup(ctx context.Context, group *aws.Group) (*aws.Group, error) {
	if group == nil {
		return nil, apierror.New(apierror.ErrBadRequest, "invalid input", nil)
	}

	output, err := e.Service.CloudProviderAWS().Create(ctx, &aws.CreateGroupInput{
		Group: group,
	})

	if err != nil {
		return nil, ErrCode("failed tocreate elastigroup", err)
	}

	return output.Group, nil
}
