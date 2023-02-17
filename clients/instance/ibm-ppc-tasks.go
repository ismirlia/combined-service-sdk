package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/helpers"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/client/p_cloud_tasks"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// IBMPPCTaskClient
type IBMPPCTaskClient struct {
	IBMPPCClient
}

// NewIBMPPCTaskClient
func NewIBMPPCTaskClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCTaskClient {
	return &IBMPPCTaskClient{
		*NewIBMPPCClient(ctx, sess, cloudInstanceID),
	}
}

// Get a Task
func (f *IBMPPCTaskClient) Get(id string) (*models.Task, error) {
	params := p_cloud_tasks.NewPcloudTasksGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCGetTimeOut).
		WithTaskID(id)
	resp, err := f.session.Power.PCloudTasks.PcloudTasksGet(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the task %s: %w", id, err)
	}
	return resp.Payload, nil
}

// Delete a Task
func (f *IBMPPCTaskClient) Delete(id string) error {
	params := p_cloud_tasks.NewPcloudTasksDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PPCDeleteTimeOut).
		WithTaskID(id)
	_, err := f.session.Power.PCloudTasks.PcloudTasksDelete(params, f.session.AuthInfo(f.cloudInstanceID))
	if err != nil {
		return fmt.Errorf("failed to delete the task id ... %w", err)
	}
	return nil
}
