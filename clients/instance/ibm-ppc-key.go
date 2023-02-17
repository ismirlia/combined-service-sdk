package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_tenants_ssh_keys"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCKeyClient
type IBMPPCKeyClient struct {
	IBMPPCClient
}

// NewIBMPPCKeyClient
func NewIBMPPCKeyClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCKeyClient {
	return &IBMPPCKeyClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a SSH Key
func (f *IBMPPCKeyClient) Get(id string) (*models.SSHKey, error) {
	var tenantid = f.session.Options.UserAccount
	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithTenantID(tenantid).WithSshkeyName(id)
	resp, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetPPCKeyOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PPC Key %s", id)
	}
	return resp.Payload, nil
}

// Get All SSH Keys
func (f *IBMPPCKeyClient) GetAll() (*models.SSHKeys, error) {
	var tenantid = f.session.Options.UserAccount
	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithTenantID(tenantid)
	resp, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PPC Keys: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PPC Keys")
	}
	return resp.Payload, nil
}

// Create a SSH Key
func (f *IBMPPCKeyClient) Create(body *models.SSHKey) (*models.SSHKey, error) {
	var tenantid = f.session.Options.UserAccount
	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithTenantID(tenantid).WithBody(body)
	postok, postcreated, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreatePPCKeyOperationFailed, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	if postcreated != nil && postcreated.Payload != nil {
		return postcreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create PPC Key")
}

// Delete a SSH Key
func (f *IBMPPCKeyClient) Delete(id string) error {
	var tenantid = f.session.Options.UserAccount
	params := p_cloud_tenants_ssh_keys.NewPcloudTenantsSshkeysDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithTenantID(tenantid).WithSshkeyName(id)
	_, err := f.session.Power.PCloudTenantsSSHKeys.PcloudTenantsSshkeysDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeletePPCKeyOperationFailed, id, err)
	}
	return nil
}
