package instance_test

import (
	"context"
	"testing"

	utl "github.com/IBM-Cloud/ppc-aas-go-sdk/internal/testutils"

	client "github.com/IBM-Cloud/ppc-aas-go-sdk/clients/instance"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
	"github.com/stretchr/testify/require"
)

func TestSnapshot(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and clients
	utl.SnapshotPreCheck(t)
	session := utl.TestSession(t)
	instanceClient := client.NewIBMPPCInstanceClient(context.Background(), session, utl.CloudInstanceID)
	snapshotClient := client.NewIBMPPCSnapshotClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE
	body := &models.SnapshotCreate{
		Name:        &utl.SnapshotName,
		Description: utl.SnapshotDescription,
	}
	createResp, err := instanceClient.CreatePvmSnapShot(utl.SnapshotInstanceID, body)
	require.Nil(t, err)
	snapshotID := createResp.SnapshotID
	utl.TestMessage("CREATE Snapshot", *snapshotID, *createResp)

	// DELETE
	defer func() {
		err = snapshotClient.Delete(*snapshotID)
		require.Nil(t, err)
		utl.TestMessage("DELETE Snapshot", *snapshotID, nil)
	}()

	// GET
	getResp, err := snapshotClient.Get(*snapshotID)
	require.Nil(t, err)
	utl.TestMessage("GET Snapshot", *snapshotID, *getResp)

	// GET ALL
	getAllResp, err := snapshotClient.GetAll()
	require.Nil(t, err)
	utl.TestMessage("GET All Snapshots", "", *getAllResp)
}
