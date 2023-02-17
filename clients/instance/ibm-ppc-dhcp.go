package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_service_d_h_c_p"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// NewIBMPPCDhcpClient
type IBMPPCDhcpClient struct {
	IBMPPCClient
}

// NewIBMPPCDhcpClient
func NewIBMPPCDhcpClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCDhcpClient {
	return &IBMPPCDhcpClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Create a DHCP server
func (f *IBMPPCDhcpClient) Create(body *models.DHCPServerCreate) (*models.DHCPServer, error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postaccepted, err := f.session.Power.PCloudServicedhcp.PcloudDhcpPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateDchpOperationFailed, f.cloudInstanceID, err)
	}
	if postaccepted != nil && postaccepted.Payload != nil {
		return postaccepted.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create DHCP")
}

// Get a DHCP server
func (f *IBMPPCDhcpClient) Get(id string) (*models.DHCPServerDetail, error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithDhcpID(id)
	resp, err := f.session.Power.PCloudServicedhcp.PcloudDhcpGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetDhcpOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get DHCP %s", id)
	}
	return resp.Payload, nil
}

// Get All DHCP servers
func (f *IBMPPCDhcpClient) GetAll() (models.DHCPServers, error) {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudServicedhcp.PcloudDhcpGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all DHCP servers: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all DHCP servers")
	}
	return resp.Payload, nil
}

// Delete a DHCP server
func (f *IBMPPCDhcpClient) Delete(id string) error {
	params := p_cloud_service_d_h_c_p.NewPcloudDhcpDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithDhcpID(id)
	_, err := f.session.Power.PCloudServicedhcp.PcloudDhcpDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteDhcpOperationFailed, id, err)
	}
	return nil
}
