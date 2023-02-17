package instance_test

import (
	"context"
	"testing"

	utl "github.com/IBM-Cloud/ppc-aas-go-sdk/internal/testutils"

	client "github.com/IBM-Cloud/ppc-aas-go-sdk/clients/instance"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
	"github.com/stretchr/testify/require"
)

func TestVolume(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and client
	utl.VolumePreCheck(t)
	session := utl.TestSession(t)
	volumeClient := client.NewIBMPPCVolumeClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE
	volumeBody := &models.CreateDataVolume{
		Name:      &utl.VolumeName,
		Size:      &utl.VolumeSize,
		DiskType:  utl.VolumeType,
		Shareable: &utl.VolumeShareable,
	}
	createRespOk, err := volumeClient.CreateVolume(volumeBody)
	require.Nil(t, err)
	VolumeID := *createRespOk.VolumeID
	utl.TestMessage("CREATE Volume", VolumeID, *createRespOk)

	// DELETE
	defer func() {
		err = volumeClient.DeleteVolume(VolumeID)
		require.Nil(t, err)
		utl.TestMessage("DELETE Volume", VolumeID, nil)
	}()

	// GET
	getResp, err := volumeClient.Get(VolumeID)
	require.Nil(t, err)
	utl.TestMessage("GET Volume", VolumeID, *getResp)
	// verify variables match
	require.Equal(t, *getResp.Name, utl.VolumeName)
	require.Equal(t, *getResp.Size, utl.VolumeSize)
	require.Equal(t, getResp.DiskType, utl.VolumeType)
	require.Equal(t, *getResp.Shareable, utl.VolumeShareable)

	// GET All
	getallResp, err := volumeClient.GetAll()
	require.Nil(t, err)
	utl.TestMessage("GET ALL Volumes", "", *getallResp)
}
