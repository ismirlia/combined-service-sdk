package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_instances"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCCloudInstanceClient
type IBMPPCCloudInstanceClient struct {
	IBMPPCClient
}

// NewIBMPPCCloudInstanceClient
func NewIBMPPCCloudInstanceClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCCloudInstanceClient {
	return &IBMPPCCloudInstanceClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Cloud Instance
func (f *IBMPPCCloudInstanceClient) Get(id string) (*models.CloudInstance, error) {
	params := p_cloud_instances.NewPcloudCloudinstancesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(id)
	resp, err := f.session.Power.PCloudInstances.PcloudCloudinstancesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetCloudInstanceOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Cloud Instance %s", id)
	}
	return resp.Payload, nil
}

// Update a Cloud Instance
func (f *IBMPPCCloudInstanceClient) Update(id string, body *models.CloudInstanceUpdate) (*models.CloudInstance, error) {
	params := p_cloud_instances.NewPcloudCloudinstancesPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(id).WithBody(body)
	resp, err := f.session.Power.PCloudInstances.PcloudCloudinstancesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateCloudInstanceOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to update the Cloud instance %s", id)
	}
	return resp.Payload, nil
}

// Delete a Cloud Instance
func (f *IBMPPCCloudInstanceClient) Delete(id string) error {
	params := p_cloud_instances.NewPcloudCloudinstancesDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(id)
	_, err := f.session.Power.PCloudInstances.PcloudCloudinstancesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteCloudInstanceOperationFailed, id, err)
	}
	return nil
}
