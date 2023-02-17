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

	// ssh inputs
	name := " < NAME OF THE ssh > "
	piID := " < POWER INSTANCE ID > "
	ssh := " <ssh ID> "

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
	powerClient := v.NewIBMPPCKeyClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	getAllResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[0]****************** %+v \n", *getAllResp)

	body := &models.SSHKey{
		Name:   &name,
		SSHKey: &ssh,
	}
	createRespOk, err := powerClient.Create(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

	sshID := *createRespOk.Name
	getResp, err := powerClient.Get(sshID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	err = powerClient.Delete(sshID)
	if err != nil {
		log.Fatal(err)
	}
}
