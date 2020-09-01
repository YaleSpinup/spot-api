package spotinst

import (
	"context"

	"github.com/YaleSpinup/apierror"
	log "github.com/sirupsen/logrus"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

// ListAWSManagedInstances lists existing managed instances
func (m *ManagedInstance) ListAWSManagedInstances(ctx context.Context) ([]*aws.ManagedInstance, error) {
	log.Info("listing aws managed instances")

	output, err := m.Service.CloudProviderAWS().List(ctx, &aws.ListManagedInstancesInput{})
	if err != nil {
		return nil, ErrCode("failed to list managed instances", err)
	}

	return output.ManagedInstances, err
}

// GetAWSManagedInstanceByID gets details of existing managed instance by id
func (m *ManagedInstance) GetAWSManagedInstanceByID(ctx context.Context, id string) (*aws.ManagedInstance, error) {
	if id == "" {
		return nil, apierror.New(apierror.ErrBadRequest, "invalid input", nil)
	}

	log.Infof("getting details about aws managed instance: %s", id)

	output, err := m.Service.CloudProviderAWS().Read(ctx, &aws.ReadManagedInstanceInput{
		ManagedInstanceID: spotinst.String(id),
	})

	if err != nil {
		return nil, ErrCode("failed to read managed instance details", err)
	}

	return output.ManagedInstance, nil
}

// CreateAWSManagedInstance creates a managed instance
func (m *ManagedInstance) CreateAWSManagedInstance(ctx context.Context, input *aws.ManagedInstance) (*aws.ManagedInstance, error) {
	if input == nil {
		return nil, apierror.New(apierror.ErrBadRequest, "invalid input", nil)
	}

	log.Infof("creating aws managed instance with input: %+v", input)

	output, err := m.Service.CloudProviderAWS().Create(ctx, &aws.CreateManagedInstanceInput{
		ManagedInstance: input,
	})

	if err != nil {
		return nil, ErrCode("failed to create managed instance", err)
	}

	return output.ManagedInstance, nil
}

// UpdateAWSManagedInstance updates facets on a managed instance
func (m *ManagedInstance) UpdateAWSManagedInstance(ctx context.Context, input *aws.ManagedInstance) (*aws.ManagedInstance, error) {
	if input == nil {
		return nil, apierror.New(apierror.ErrBadRequest, "invalid input", nil)
	}

	log.Infof("updating aws managed instance with input: %+v", input)

	output, err := m.Service.CloudProviderAWS().Update(ctx, &aws.UpdateManagedInstanceInput{
		ManagedInstance: input,
	})
	if err != nil {
		return nil, ErrCode("failed to update managed instance", err)
	}

	return output.ManagedInstance, nil
}

// DeleteAWSManagedInstanceByID deletes a managed instance by id
func (m *ManagedInstance) DeleteAWSManagedInstanceByID(ctx context.Context, id string) error {
	if id == "" {
		return apierror.New(apierror.ErrBadRequest, "invalid input", nil)
	}

	_, err := m.Service.CloudProviderAWS().Delete(ctx, &aws.DeleteManagedInstanceInput{
		ManagedInstanceID: spotinst.String(id),
	})

	if err != nil {
		return ErrCode("failed to delete managed instance", err)
	}

	return nil
}
