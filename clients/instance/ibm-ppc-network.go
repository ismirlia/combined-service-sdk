package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_networks"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCNetworkClient
type IBMPPCNetworkClient struct {
	IBMPPCClient
}

// NewIBMPPCNetworkClient
func NewIBMPPCNetworkClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCNetworkClient {
	return &IBMPPCNetworkClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Network
func (f *IBMPPCNetworkClient) Get(id string) (*models.Network, error) {
	params := p_cloud_networks.NewPcloudNetworksGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetNetworkOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Network %s", id)
	}
	return resp.Payload, nil
}

// Get All Networks
func (f *IBMPPCNetworkClient) GetAll() (*models.Networks, error) {
	params := p_cloud_networks.NewPcloudNetworksGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get Network for cloud instance %s with error %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Network for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a Network
func (f *IBMPPCNetworkClient) Create(body *models.NetworkCreate) (*models.Network, error) {
	params := p_cloud_networks.NewPcloudNetworksPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postok, postcreated, err := f.session.Power.PCloudNetworks.PcloudNetworksPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateNetworkOperationFailed, body.Name, err)
	}
	if postok != nil && postok.Payload != nil {
		return postok.Payload, nil
	}
	if postcreated != nil && postcreated.Payload != nil {
		return postcreated.Payload, nil
	}
	return nil, fmt.Errorf("failed to perform Create Network Operation for Network %s", body.Name)
}

// Update a Network
func (f *IBMPPCNetworkClient) Update(id string, body *models.NetworkUpdate) (*models.Network, error) {
	params := p_cloud_networks.NewPcloudNetworksPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to perform Update Network Operation for Network %s with error %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Update Network Operation for Network %s", id)
	}
	return resp.Payload, nil
}

// Get All Public Networks
func (f *IBMPPCNetworkClient) GetAllPublic() (*models.Networks, error) {
	filterQuery := "type=\"pub-vlan\""
	params := p_cloud_networks.NewPcloudNetworksGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithFilter(&filterQuery)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Public Networks for cloud instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Public Networks for cloud instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Delete a Network
func (f *IBMPPCNetworkClient) Delete(id string) error {
	params := p_cloud_networks.NewPcloudNetworksDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id)
	_, err := f.session.Power.PCloudNetworks.PcloudNetworksDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PPC Network %s: %w", id, err)
	}
	return nil
}

// Get All Ports on a Network
func (f *IBMPPCNetworkClient) GetAllPorts(id string) (*models.NetworkPorts, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Network Ports for Network %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Network Ports for Network %s", id)
	}
	return resp.Payload, nil
}

// Get a Port on a Network
func (f *IBMPPCNetworkClient) GetPort(id string, networkPortID string) (*models.NetworkPort, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithPortID(networkPortID)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsGet(params, f.session.AuthInfo(f.cloudInstanceID))

	if err != nil {
		return nil, fmt.Errorf("failed to Get PPC Network Port %s for Network %s: %w", networkPortID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PPC Network Port %s for Network %s", networkPortID, id)
	}
	return resp.Payload, nil
}

// Create a Port on a Network
func (f *IBMPPCNetworkClient) CreatePort(id string, body *models.NetworkPortCreate) (*models.NetworkPort, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateNetworkPortOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform Create Network Port Operation for Network %s", id)
	}
	return resp.Payload, nil
}

// Delete a Port on a Network
func (f *IBMPPCNetworkClient) DeletePort(id string, networkPortID string) error {
	params := p_cloud_networks.NewPcloudNetworksPortsDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithPortID(networkPortID)
	_, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to delete the network port %s for network %s with error %w", networkPortID, id, err)
	}
	return nil
}

// Update a Port on a Network
func (f *IBMPPCNetworkClient) UpdatePort(id, networkPortID string, body *models.NetworkPortUpdate) (*models.NetworkPort, error) {
	params := p_cloud_networks.NewPcloudNetworksPortsPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithNetworkID(id).
		WithPortID(networkPortID).WithBody(body)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPortsPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to update the port %s and Network %s with error %w", networkPortID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to update the port %s and Network %s", networkPortID, id)
	}
	return resp.Payload, nil
}
