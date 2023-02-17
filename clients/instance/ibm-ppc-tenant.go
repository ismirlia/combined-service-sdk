package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_tenants"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCTenantClient
type IBMPPCTenantClient struct {
	IBMPPCClient
}

// NewIBMPPCTenantClient
func NewIBMPPCTenantClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCTenantClient {
	return &IBMPPCTenantClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Tenant
func (f *IBMPPCTenantClient) Get(tenantid string) (*models.Tenant, error) {
	params := p_cloud_tenants.NewPcloudTenantsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithTenantID(tenantid)
	resp, err := f.session.Power.PCloudTenants.PcloudTenantsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant %s with error %w", tenantid, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get tenant %s", tenantid)
	}
	return resp.Payload, nil
}

// Get own Tenant
func (f *IBMPPCTenantClient) GetSelfTenant() (*models.Tenant, error) {
	params := p_cloud_tenants.NewPcloudTenantsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithTenantID(f.session.Options.UserAccount)
	resp, err := f.session.Power.PCloudTenants.PcloudTenantsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get self tenant with error %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get self tenant")
	}
	return resp.Payload, nil
}
