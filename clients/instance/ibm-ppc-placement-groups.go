package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_placement_groups"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCPlacementGroupClient
type IBMPPCPlacementGroupClient struct {
	IBMPPCClient
}

// NewIBMPPCPlacementGroupClient
func NewIBMPPCPlacementGroupClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCPlacementGroupClient {
	return &IBMPPCPlacementGroupClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a PPC Placement Group
func (f *IBMPPCPlacementGroupClient) Get(id string) (*models.PlacementGroup, error) {
	params := p_cloud_placement_groups.NewPcloudPlacementgroupsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPlacementGroupID(id)
	resp, err := f.session.Power.PCloudPlacementGroups.PcloudPlacementgroupsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetPlacementGroupOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get Placement Group %s", id)
	}
	return resp.Payload, nil
}

// Get All Placement Groups
func (f *IBMPPCPlacementGroupClient) GetAll() (*models.PlacementGroups, error) {
	params := p_cloud_placement_groups.NewPcloudPlacementgroupsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudPlacementGroups.PcloudPlacementgroupsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Get All Placement Groups: %w", err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all Placement Groups")
	}
	return resp.Payload, nil
}

// Create a Placement Group
func (f *IBMPPCPlacementGroupClient) Create(body *models.PlacementGroupCreate) (*models.PlacementGroup, error) {
	params := p_cloud_placement_groups.NewPcloudPlacementgroupsPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postok, err := f.session.Power.PCloudPlacementGroups.PcloudPlacementgroupsPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.CreatePlacementGroupOperationFailed, f.cloudInstanceID, err)
	}
	if postok == nil || postok.Payload == nil {
		return nil, fmt.Errorf("failed to Create Placement Group")
	}
	return postok.Payload, nil
}

// Delete a Placement Group
func (f *IBMPPCPlacementGroupClient) Delete(id string) error {
	params := p_cloud_placement_groups.NewPcloudPlacementgroupsDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPlacementGroupID(id)
	_, err := f.session.Power.PCloudPlacementGroups.PcloudPlacementgroupsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeletePlacementGroupOperationFailed, id, err)
	}
	return nil
}

// Add an Instance to a Placement Group
func (f *IBMPPCPlacementGroupClient) AddMember(id string, body *models.PlacementGroupServer) (*models.PlacementGroup, error) {
	params := p_cloud_placement_groups.NewPcloudPlacementgroupsMembersPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPlacementGroupID(id).
		WithBody(body)
	postok, err := f.session.Power.PCloudPlacementGroups.PcloudPlacementgroupsMembersPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.AddMemberPlacementGroupOperationFailed, *body.ID, id, err)
	}
	if postok == nil || postok.Payload == nil {
		return nil, fmt.Errorf("failed to Add Member for instance %s and placement group %s", *body.ID, id)
	}
	return postok.Payload, nil
}

// Remove an Instance to a Placement Group
func (f *IBMPPCPlacementGroupClient) DeleteMember(id string, body *models.PlacementGroupServer) (*models.PlacementGroup, error) {
	params := p_cloud_placement_groups.NewPcloudPlacementgroupsMembersDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPlacementGroupID(id).
		WithBody(body)
	delok, err := f.session.Power.PCloudPlacementGroups.PcloudPlacementgroupsMembersDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.DeleteMemberPlacementGroupOperationFailed, *body.ID, id, err)
	}
	if delok == nil || delok.Payload == nil {
		return nil, fmt.Errorf("failed to Delete Member for instance %s and placement group %s", *body.ID, id)
	}
	return delok.Payload, nil
}
