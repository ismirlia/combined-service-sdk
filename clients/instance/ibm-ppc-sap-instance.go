package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_s_a_p"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCSAPnstanceClient
type IBMPPCSAPnstanceClient struct {
	IBMPPCClient
}

// NewIBMPPCSAPnstanceClient
func NewIBMPPCSAPnstanceClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCSAPnstanceClient {
	return &IBMPPCSAPnstanceClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Create a SAP Instance
func (f *IBMPPCSAPnstanceClient) Create(body *models.SAPCreate) (*models.PVMInstanceList, error) {
	params := p_cloud_s_a_p.NewPcloudSapPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCCreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postok, postcreated, postAccepted, err := f.session.Power.PCloudsap.PcloudSapPost(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Create SAP Instance: %w", err)
	}
	if postok != nil && len(postok.Payload) > 0 {
		return &postok.Payload, nil
	}
	if postcreated != nil && len(postcreated.Payload) > 0 {
		return &postcreated.Payload, nil
	}
	if postAccepted != nil && len(postAccepted.Payload) > 0 {
		return &postAccepted.Payload, nil
	}
	return nil, fmt.Errorf("failed to Create SAP Instance")
}

// Get a SAP Profile
func (f *IBMPPCSAPnstanceClient) GetSAPProfile(id string) (*models.SAPProfile, error) {
	params := p_cloud_s_a_p.NewPcloudSapGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSapProfileID(id)
	resp, err := f.session.Power.PCloudsap.PcloudSapGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get sap profile %s : %w", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get sap profile %s", id)
	}
	return resp.Payload, nil
}

// Get All SAP Profiles
func (f *IBMPPCSAPnstanceClient) GetAllSAPProfiles(cloudInstanceID string) (*models.SAPProfiles, error) {
	params := p_cloud_s_a_p.NewPcloudSapGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudsap.PcloudSapGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to get all sap profiles for power instance %s: %w", cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all sap profiles for power instance %s", cloudInstanceID)
	}
	return resp.Payload, nil
}
