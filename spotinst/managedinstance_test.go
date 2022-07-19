package spotinst

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance"
	spotiface "github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
)

func (m *mockManagedInstanceClient) List(ctx context.Context, inp *spotiface.ListManagedInstancesInput) (*spotiface.ListManagedInstancesOutput, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &spotiface.ListManagedInstancesOutput{ManagedInstances: []*spotiface.ManagedInstance{{ID: aws.String("smi-435")}}}, nil
}

func TestGetAWSManagedInstanceByID(t *testing.T) {
	t.Log("TODO")
}

func TestCreateAWSManagedInstance(t *testing.T) {
	t.Log("TODO")
}

func TestDeleteAWSManagedInstanceByID(t *testing.T) {
	t.Log("TODO")
}

func TestManagedInstance_ListAWSManagedInstances(t *testing.T) {
	type fields struct {
		Service    managedinstance.Service
		DefaultVPC string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		m       *ManagedInstance
		args    args
		fields  fields
		want    []*spotiface.ManagedInstance
		wantErr bool
	}{
		{
			name:    "success case",
			args:    args{ctx: context.TODO()},
			fields:  fields{Service: newmockManagedInstanceClient(t, nil)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ManagedInstance{
				Service:    tt.fields.Service,
				DefaultVPC: "",
			}
			got, err := m.ListAWSManagedInstances(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ManagedInstance.ListAWSManagedInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ManagedInstance.ListAWSManagedInstances() = %v, want %v", got, tt.want)
			}
		})
	}
}
