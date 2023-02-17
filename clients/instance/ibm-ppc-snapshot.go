package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_snapshots"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCSnapshotClient
type IBMPPCSnapshotClient struct {
	IBMPPCClient
}

// NewIBMPPCSnapshotClient
func NewIBMPPCSnapshotClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCSnapshotClient {
	return &IBMPPCSnapshotClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Snapshot
func (f *IBMPPCSnapshotClient) Get(id string) (*models.Snapshot, error) {
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSnapshotID(id)
	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get PPC Snapshot %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PPC Snapshot %s", id)
	}
	return resp.Payload, nil
}

// Delete a Snapshot
func (f *IBMPPCSnapshotClient) Delete(id string) error {
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSnapshotID(id)
	_, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to Delete PPC Snapshot %s: %w", id, err)
	}
	return nil
}

// Update a Snapshot
func (f *IBMPPCSnapshotClient) Update(id string, body *models.SnapshotUpdate) (models.Object, error) {
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSnapshotID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsPut(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Update PPC Snapshot %s: %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update PPC Snapshot %s", id)
	}
	return resp.Payload, nil
}

// Get All Snapshots
func (f *IBMPPCSnapshotClient) GetAll() (*models.Snapshots, error) {
	params := p_cloud_snapshots.NewPcloudCloudinstancesSnapshotsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudSnapshots.PcloudCloudinstancesSnapshotsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PPC Snapshots: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PPC Snapshots")
	}
	return resp.Payload, nil
}

// Create or Restore a Snapshot
func (f *IBMPPCSnapshotClient) Create(instanceID, snapshotID, restoreFailAction string) (*models.Snapshot, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsRestorePostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(instanceID).
		WithSnapshotID(snapshotID).WithRestoreFailAction(&restoreFailAction)
	resp, err := f.session.Power.PCloudpVMInstances.PcloudPvminstancesSnapshotsRestorePost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to restore PPC Snapshot %s of the instance %s: %w", snapshotID, instanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to restore PPC Snapshot %s of the instance %s", snapshotID, instanceID)
	}
	return resp.Payload, nil
}
