package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_storage_capacity"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCStorageCapacityClient
type IBMPPCStorageCapacityClient struct {
	IBMPPCClient
}

// NewIBMPPCStorageCapacityClient
func NewIBMPPCStorageCapacityClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCStorageCapacityClient {
	return &IBMPPCStorageCapacityClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get All Storage Capacities for all available Storage Pools in a Region
func (f *IBMPPCStorageCapacityClient) GetAllStoragePoolsCapacity() (*models.StoragePoolsCapacity, error) {
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityPoolsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityPoolsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage pools: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage pools")
	}
	return resp.Payload, nil
}

// Get Storage Capacity for a Storage Pool in a Region
func (f *IBMPPCStorageCapacityClient) GetStoragePoolCapacity(storagePool string) (*models.StoragePoolCapacity, error) {
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityPoolsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithStoragePoolName(storagePool)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityPoolsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for storage pool %s: %w", storagePool, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for storage pool %s", storagePool)
	}
	return resp.Payload, nil
}

// Get Storage capacity for a Storage Type in a Region
func (f *IBMPPCStorageCapacityClient) GetStorageTypeCapacity(storageType string) (*models.StorageTypeCapacity, error) {
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityTypesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithStorageTypeName(storageType)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityTypesGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for storage type %s: %w", storageType, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for storage type %s", storageType)
	}
	return resp.Payload, nil
}

// Get Storage Capacity for all available Storage Types in a Region
func (f *IBMPPCStorageCapacityClient) GetAllStorageTypesCapacity() (*models.StorageTypesCapacity, error) {
	params := p_cloud_storage_capacity.NewPcloudStoragecapacityTypesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudStorageCapacity.PcloudStoragecapacityTypesGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage types %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the capacity for all storage types")
	}
	return resp.Payload, nil
}
