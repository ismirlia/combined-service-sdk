package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_system_pools"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCSystemPoolClient
type IBMPPCSystemPoolClient struct {
	IBMPPCClient
}

// NewIBMPPCSystemPoolClient
func NewIBMPPCSystemPoolClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCSystemPoolClient {
	return &IBMPPCSystemPoolClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get the System Pools
// Deprecated: Use GetSystemPools()
func (f *IBMPPCSystemPoolClient) Get(id string) (models.SystemPools, error) {
	params := p_cloud_system_pools.NewPcloudSystempoolsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(id)
	resp, err := f.session.Power.PCloudSystemPools.PcloudSystempoolsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetSystemPoolsOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Get System Pools Operation for cloud instance id %s", id)
	}
	return resp.Payload, nil
}

// Get the System Pools
func (f *IBMPPCSystemPoolClient) GetSystemPools() (models.SystemPools, error) {
	params := p_cloud_system_pools.NewPcloudSystempoolsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudSystemPools.PcloudSystempoolsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetSystemPoolsOperationFailed, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Get System Pools Operation for cloud instance id %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}
