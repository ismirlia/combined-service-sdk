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

	// session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	zone := " < ZONE > "
	accountID := " < ACCOUNT ID > "
	url := region + ".ppc-aas.test.cloud.ibm.com"

	// dhcp inputs
	piID := " < POWER INSTANCE ID > "

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
	powerClient := v.NewIBMPPCDhcpClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	dhcpServer, err := powerClient.Create(&models.DHCPServerCreate{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *dhcpServer)

	dhcpId := *dhcpServer.ID
	getResp, err := powerClient.Get(dhcpId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[2]****************** %+v \n", *getResp)

	getAllResp, err := powerClient.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[3]****************** %+v \n", getAllResp)

	err = powerClient.Delete(dhcpId)
	if err != nil {
		log.Fatal(err)
	}
}
