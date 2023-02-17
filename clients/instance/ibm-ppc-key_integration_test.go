package instance_test

import (
	"context"
	"testing"

	utl "github.com/IBM-Cloud/ppc-aas-go-sdk/internal/testutils"

	client "github.com/IBM-Cloud/ppc-aas-go-sdk/clients/instance"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
	"github.com/stretchr/testify/require"
)

func TestKey(t *testing.T) {
	if utl.DisableTesting {
		return
	}

	// create session and client
	utl.SSHKeyPreCheck(t)
	session := utl.TestSession(t)
	sshKeyClient := client.NewIBMPPCKeyClient(context.Background(), session, utl.CloudInstanceID)

	// CREATE
	body := &models.SSHKey{
		Name:   &utl.SSHKeyName,
		SSHKey: &utl.SSHKeyRSA,
	}
	createResp, err := sshKeyClient.Create(body)
	require.Nil(t, err)
	utl.TestMessage("CREATE SSH Key", utl.SSHKeyName, *createResp)

	// DELETE
	defer func() {
		err := sshKeyClient.Delete(utl.SSHKeyName)
		require.Nil(t, err)
		utl.TestMessage("DELETE SSH Key", utl.SSHKeyName, nil)
	}()

	// GET
	getResp, err := sshKeyClient.Get(utl.SSHKeyName)
	require.Nil(t, err)
	utl.TestMessage("GET SSH Key", utl.SSHKeyName, *getResp)
	// verify variables match
	require.Equal(t, *getResp.Name, utl.SSHKeyName)
	require.Equal(t, *getResp.SSHKey, utl.SSHKeyRSA)

	// GET ALL
	getAllResp, err := sshKeyClient.GetAll()
	require.Nil(t, err)
	utl.TestMessage("GET All SSH Keys", "", *getAllResp)
}
