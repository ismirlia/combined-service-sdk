package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_volume_groups"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCVolumeGroupClient
type IBMPPCVolumeGroupClient struct {
	IBMPPCClient
}

// NewIBMPPCVolumeGroupClient
func NewIBMPPCVolumeGroupClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCVolumeGroupClient {
	return &IBMPPCVolumeGroupClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Volume-Group
func (f *IBMPPCVolumeGroupClient) Get(id string) (*models.VolumeGroup, error) {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeGroupID(id)
	resp, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVolumeGroupOperationFailed, id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get volume-group %s", id)
	}
	return resp.Payload, nil
}

// GetAll Volume-Group
func (f *IBMPPCVolumeGroupClient) GetAll() (*models.VolumeGroups, error) {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all volume-groups for Cloud Instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all volume-groups for Cloud Instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get a Volume-Group with Details
func (f *IBMPPCVolumeGroupClient) GetDetails(id string) (*models.VolumeGroupDetails, error) {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsGetDetailsParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeGroupID(id)
	resp, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsGetDetails(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVolumeGroupDetailsOperationFailed, id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get volume-group %s details", id)
	}
	return resp.Payload, nil
}

// GetAll Volume-Group with Details
func (f *IBMPPCVolumeGroupClient) GetAllDetails() (*models.VolumeGroupsDetails, error) {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsGetallDetailsParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsGetallDetails(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all volume-groups details for Cloud Instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all volume-groups details for Cloud Instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a Volume Group
func (f *IBMPPCVolumeGroupClient) CreateVolumeGroup(body *models.VolumeGroupCreate) (*models.VolumeGroupCreateResponse, error) {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	respOk, respPartial, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVolumeGroupOperationFailed, f.cloudInstanceID, err)
	}
	if respOk != nil && respOk.Payload != nil {
		return respOk.Payload, nil
	}
	if respPartial != nil && respPartial.Payload != nil {
		return respPartial.Payload, nil
	}
	return nil, fmt.Errorf("failed to create volume-group")
}

// Update a Volume Group
func (f *IBMPPCVolumeGroupClient) UpdateVolumeGroup(id string, body *models.VolumeGroupUpdate) error {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeGroupID(id).WithBody(body)
	resp, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.CreateVolumeGroupOperationFailed, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return fmt.Errorf("failed to update volume-group %s", id)
	}
	return nil
}

// Delete a Volume Group
func (f *IBMPPCVolumeGroupClient) DeleteVolumeGroup(id string) error {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeGroupID(id)
	_, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteVolumeGroupOperationFailed, id, err)
	}
	return nil
}

// Get live details of a Volume Group
func (f *IBMPPCVolumeGroupClient) GetVolumeGroupLiveDetails(id string) (*models.VolumeGroupStorageDetails, error) {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsStorageDetailsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeGroupID(id)
	resp, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsStorageDetailsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetLiveVolumeGroupDetailsOperationFailed, id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get live details of volume-group %s", id)
	}
	return resp.Payload, nil
}

// Performs action on Volume Group
func (f *IBMPPCVolumeGroupClient) VolumeGroupAction(id string, body *models.VolumeGroupAction) (models.Object, error) {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsActionPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeGroupID(id).WithBody(body)
	resp, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsActionPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.VolumeGroupActionOperationFailed, id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform action on volume-group %s", id)
	}
	return resp.Payload, nil
}

// Get remote copy relationships of the volume belonging to volume group
func (f *IBMPPCVolumeGroupClient) GetVolumeGroupRemoteCopyRelationships(id string) (*models.VolumeGroupRemoteCopyRelationships, error) {
	params := p_cloud_volume_groups.NewPcloudVolumegroupsRemoteCopyRelationshipsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeGroupID(id)
	resp, err := f.session.Power.PCloudVolumeGroups.PcloudVolumegroupsRemoteCopyRelationshipsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVolumeGroupRemoteCopyRelationshipsOperationFailed, id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get remote copy relationships of the volumes belonging to volume group %s", id)
	}
	return resp.Payload, nil
}
