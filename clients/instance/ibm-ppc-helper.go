package instance

import (
	"context"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ibmppcsession"
)

// Helper methods that will be used by the client classes

// IBMPPCHelperClient
type IBMPPCClient struct {
	session         *ibmppcsession.IBMPPCSession
	cloudInstanceID string
	ctx             context.Context
}

// NewIBMPPCClient
func NewIBMPPCClient(ctx context.Context, sess *ibmppcsession.IBMPPCSession, cloudInstanceID string) *IBMPPCClient {
	return &IBMPPCClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		ctx:             ctx,
	}
}
