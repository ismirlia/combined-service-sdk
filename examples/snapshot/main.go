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

	piID := " < POWER INSTANCE ID > "
	instance_id := " < INSTANCE ID > "
	snap_name := " < SNAPSHOT NAME > "
	description := " < DESCRIPTION > "

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

	powerClient := v.NewIBMPPCInstanceClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	snapshotBody := &models.SnapshotCreate{Name: &snap_name, Description: description}
	createRespSnapOk, err := powerClient.CreatePvmSnapShot(instance_id, snapshotBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v \n", *createRespSnapOk)

	powerSnapClient := v.NewIBMPPCSnapshotClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	getAllSnapResp, err := powerSnapClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getAllSnapResp)

	getSnapResp, err := powerSnapClient.Get(*createRespSnapOk.SnapshotID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", *getSnapResp)

	err = powerSnapClient.Delete(*createRespSnapOk.SnapshotID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[4]****************** %+v \n", err)

}
