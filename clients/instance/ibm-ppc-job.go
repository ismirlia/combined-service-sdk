package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/errors"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_jobs"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCJobClient
type IBMPPCJobClient struct {
	IBMPPCClient
}

// NewIBMPPCJobClient
func NewIBMPPCJobClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCJobClient {
	return &IBMPPCJobClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Job
func (f *IBMPPCJobClient) Get(id string) (*models.Job, error) {
	params := p_cloud_jobs.NewPcloudCloudinstancesJobsGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithJobID(id)
	resp, err := f.session.Power.PCloudJobs.PcloudCloudinstancesJobsGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetJobOperationFailed, id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform get Job operation for job id %s", id)
	}
	return resp.Payload, nil
}

// Get All Jobs
func (f *IBMPPCJobClient) GetAll() (*models.Jobs, error) {
	params := p_cloud_jobs.NewPcloudCloudinstancesJobsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudJobs.PcloudCloudinstancesJobsGetall(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf(errors.GetAllJobsOperationFailed, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to perform get all jobs")
	}
	return resp.Payload, nil
}

// Delete a Job
func (f *IBMPPCJobClient) Delete(id string) error {
	params := p_cloud_jobs.NewPcloudCloudinstancesJobsDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithJobID(id)
	_, err := f.session.Power.PCloudJobs.PcloudCloudinstancesJobsDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf(errors.DeleteJobsOperationFailed, id, err)
	}
	return nil
}
