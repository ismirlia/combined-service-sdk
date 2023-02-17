package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_volumes"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCVolumeClient
type IBMPPCVolumeClient struct {
	IBMPPCClient
}

// NewIBMPPCVolumeClient
func NewIBMPPCVolumeClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCVolumeClient {
	return &IBMPPCVolumeClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Volume
func (f *IBMPPCVolumeClient) Get(id string) (*models.Volume, error) {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeID(id)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVolumeOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Volume %s", id)
	}
	return resp.Payload, nil
}

// Get All Volumes
func (f *IBMPPCVolumeClient) GetAll() (*models.Volumes, error) {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes for Cloud Instance %s: %w", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes for Cloud Instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Get All Affinity Volumes
func (f *IBMPPCVolumeClient) GetAllAffinityVolumes(affinity string) (*models.Volumes, error) {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithAffinity(&affinity)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes with affinity %s for Cloud Instance %s: %w", affinity, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes with affinity %s for Cloud Instance %s", affinity, f.cloudInstanceID)
	}
	return resp.Payload, nil
}

// Create a VolumeV2
func (f *IBMPPCVolumeClient) CreateVolumeV2(body *models.MultiVolumesCreate) (*models.Volumes, error) {
	params := p_cloud_volumes.NewPcloudV2VolumesPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	resp, err := f.session.Power.PCloudVolumes.PcloudV2VolumesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVolumeV2OperationFailed, *body.Name, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create Volume v2")
	}
	return resp.Payload, nil
}

// Create a Volume
func (f *IBMPPCVolumeClient) CreateVolume(body *models.CreateDataVolume) (*models.Volume, error) {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreateVolumeOperationFailed, *body.Name, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Create Volume")
	}
	return resp.Payload, nil
}

// Update a Volume
func (f *IBMPPCVolumeClient) UpdateVolume(id string, body *models.UpdateVolume) (*models.Volume, error) {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.UpdateVolumeOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update Volume %s", id)
	}
	return resp.Payload, nil
}

// Delete a Volume
func (f *IBMPPCVolumeClient) DeleteVolume(id string) error {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeID(id)
	_, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteVolumeOperationFailed, id, err)
	}
	return nil
}

// Attach a Volume to an Instance
func (f *IBMPPCVolumeClient) Attach(id, volumename string) error {
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithVolumeID(volumename)
	_, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.AttachVolumeOperationFailed, volumename, err)
	}
	return nil
}

// Detach a Volume from an Instance
func (f *IBMPPCVolumeClient) Detach(id, volumename string) error {
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithVolumeID(volumename)
	_, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DetachVolumeOperationFailed, volumename, err)
	}
	return nil
}

// Get All Volumes attached to an Instance
func (f *IBMPPCVolumeClient) GetAllInstanceVolumes(id string) (*models.Volumes, error) {
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all Volumes for PPC Instance %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Volumes for PPC Instance %s", id)
	}
	return resp.Payload, nil
}

// Set a Volume as the Boot Volume for an Instance
func (f *IBMPPCVolumeClient) SetBootVolume(id, volumename string) error {
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesSetbootPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithVolumeID(volumename)
	_, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesSetbootPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to set the boot volume %s for instance %s", volumename, id)
	}
	return nil
}

// Check if a Volume is attached to an Instance
func (f *IBMPPCVolumeClient) CheckVolumeAttach(id, volumeID string) (*models.Volume, error) {
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithVolumeID(volumeID)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s: %w", volumeID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s", volumeID, id)
	}
	return resp.Payload, nil
}

// Update a Volume attached to an Instance
func (f *IBMPPCVolumeClient) UpdateVolumeAttach(id, volumeID string, body *models.PVMInstanceVolumeUpdate) error {
	params := p_cloud_volumes.NewPcloudPvminstancesVolumesPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithVolumeID(volumeID).WithBody(body)
	resp, err := f.session.Power.PCloudVolumes.PcloudPvminstancesVolumesPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s: %w", volumeID, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return fmt.Errorf("failed to validate that the volume %s is attached to the pvminstance %s", volumeID, id)
	}
	return nil
}

// Performs action on volume
func (f *IBMPPCVolumeClient) VolumeAction(id string, body *models.VolumeAction) error {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesActionPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeID(id).WithBody(body)
	_, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesActionPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to perform action on volume %s with error %w", id, err)
	}
	return nil
}

// Get remote copy relationship of a volume
func (f *IBMPPCVolumeClient) GetVolumeRemoteCopyRelationships(id string) (*models.VolumeRemoteCopyRelationship, error) {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesRemoteCopyRelationshipGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeID(id)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesRemoteCopyRelationshipGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVolumeRemoteCopyRelationshipsOperationFailed, id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get remote copy relationships of a volume %s", id)
	}
	return resp.Payload, nil
}

// Get a list of flashcopy mappings of a given volume
func (f *IBMPPCVolumeClient) GetVolumeFlashCopyMappings(id string) (models.FlashCopyMappings, error) {
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesFlashCopyMappingsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithVolumeID(id)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesFlashCopyMappingsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetVolumeFlashCopyMappingOperationFailed, id, f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get flash copy mapping of a volume %s", id)
	}
	return resp.Payload, nil
}
