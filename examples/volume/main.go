package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/ppc-aas-go-sdk/clients/instance"
	ps "github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
	"github.com/IBM/go-sdk-core/v5/core"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".ppc-aas.test.cloud.ibm.com"

	// volume inputs
	piID := " < POWER INSTANCE ID > "
	name := " < NAME OF THE volume > "
	size := 20.0
	vtype := "tier1"
	sharable := true
	replicationEnabled := false

	authenticator := &core.BearerTokenAuthenticator{
		BearerToken: token,
	}
	// authenticator := &core.IamAuthenticator{
	// 	ApiKey: "< API KEY >",
	// 	// Uncomment for test environment
	// 	URL: "https://iam.test.cloud.ibm.com",
	// }

	// Create the session
	options := &ps.IBMPPCOptions{
		Authenticator: authenticator,
		UserAccount:   accountID,
		Zone:          zone,
		URL:           url,
		Debug:         true,
	}
	session, err := ps.NewIBMPPCSession(options)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPPCVolumeClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	body := &models.CreateDataVolume{
		Name:               &name,
		Size:               &size,
		DiskType:           vtype,
		Shareable:          &sharable,
		ReplicationEnabled: &replicationEnabled,
	}
	createRespOk, err := powerClient.CreateVolume(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	volumeID := *createRespOk.VolumeID
	getResp, err := powerClient.Get(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getallResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getallResp)

	getVolRCRelationships, err := powerClient.GetVolumeRemoteCopyRelationships(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getVolRCRelationships)

	getVolFCMappings, err := powerClient.GetVolumeFlashCopyMappings(volumeID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", getVolFCMappings)

	err = powerClient.DeleteVolume(volumeID)
	if err != nil {
		log.Fatal(err)
	}
}
