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

	log.Infof("getting details about aws elastigroup: %s", id)

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

	log.Infof("creating aws elastigroup with input: %+v", group)

	output, err := e.Service.CloudProviderAWS().Create(ctx, &aws.CreateGroupInput{
		Group: group,
	})

	if err != nil {
		return nil, ErrCode("failed tocreate elastigroup", err)
	}

	return output.Group, nil
}

func (e *Elastigroup) DeleteAWSElastigroupByID(ctx context.Context, id string) error {
	if id == "" {
		return apierror.New(apierror.ErrBadRequest, "invalid input", nil)
	}

	_, err := e.Service.CloudProviderAWS().Delete(ctx, &aws.DeleteGroupInput{
		GroupID: spotinst.String(id),
		StatefulDeallocation: &aws.StatefulDeallocation{
			ShouldDeleteImages:            spotinst.Bool(true),
			ShouldDeleteNetworkInterfaces: spotinst.Bool(true),
			ShouldDeleteVolumes:           spotinst.Bool(true),
			ShouldDeleteSnapshots:         spotinst.Bool(true),
		},
	})

	if err != nil {
		return ErrCode("failed to read elastigroup details", err)
	}

	return nil
}
