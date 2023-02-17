package ibmppcsession

import (
	"fmt"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/runtime"
)

func TestNewIBMPPCSession(t *testing.T) {
	bearerTokenAuth := &core.BearerTokenAuthenticator{BearerToken: "sample"}
	o1 := &IBMPPCOptions{
		Authenticator: bearerTokenAuth,
		UserAccount:   "1234",
		Region:        "dal",
		Zone:          "dal12",
		URL:           "dal.ppc-aas.test.cloud.ibm.com",
	}
	type args struct {
		o *IBMPPCOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *IBMPPCSession
		wantErr bool
	}{
		{
			name: "Simple BearerTokenAuthenticator",
			args: args{
				o: o1,
			},
			want: &IBMPPCSession{
				Options:   o1,
				CRNFormat: "crn:v1:staging:public:ppc-aas:dal12:a/1234:%s::",
			},
		},
		{
			name: "Without Options",
			args: args{
				o: nil,
			},
			wantErr: true,
		},
		{
			name: "Invalid Authenticator",
			args: args{
				o: &IBMPPCOptions{
					Authenticator: &core.BearerTokenAuthenticator{},
				},
			},
			wantErr: true,
		},
		{
			name: "Without Authenticator",
			args: args{
				o: &IBMPPCOptions{
					UserAccount: "1234",
					Zone:        "dal12",
					URL:         "dal.ppc-aas.test.cloud.ibm.com",
				},
			},
			wantErr: true,
		},
		{
			name: "Without UserAccount",
			args: args{
				o: &IBMPPCOptions{
					Authenticator: bearerTokenAuth,
					Zone:          "dal12",
					URL:           "dal.ppc-aas.test.cloud.ibm.com",
				},
			},
			wantErr: true,
		},
		{
			name: "Without Zone",
			args: args{
				o: &IBMPPCOptions{
					Authenticator: bearerTokenAuth,
					UserAccount:   "1234",
					URL:           "dal.ppc-aas.test.cloud.ibm.com",
				},
			},
			wantErr: true,
		},
		{
			name: "Without URL and Region",
			args: args{
				o: &IBMPPCOptions{
					Authenticator: bearerTokenAuth,
					UserAccount:   "1234",
					Zone:          "dal12",
				},
			},
			want: &IBMPPCSession{
				Options:   o1,
				CRNFormat: "crn:v1:bluemix:public:ppc-aas:dal12:a/1234:%s::",
			},
		},
		{
			name: "Without URL but with region",
			args: args{
				o: &IBMPPCOptions{
					Authenticator: bearerTokenAuth,
					UserAccount:   "1234",
					Zone:          "dal12",
					Region:        "dal",
				},
			},
			want: &IBMPPCSession{
				Options:   o1,
				CRNFormat: "crn:v1:bluemix:public:ppc-aas:dal12:a/1234:%s::",
			},
		},
		{
			name: "Simple URL with https",
			args: args{
				o: &IBMPPCOptions{
					Authenticator: &core.NoAuthAuthenticator{},
					UserAccount:   "1234",
					Region:        "dal",
					Zone:          "dal12",
					URL:           "https://dal.ppc-aas.test.cloud.ibm.com",
				},
			},
			want: &IBMPPCSession{
				Options: &IBMPPCOptions{
					Authenticator: &core.NoAuthAuthenticator{},
				},
				CRNFormat: "crn:v1:staging:public:ppc-aas:dal12:a/1234:%s::",
			},
		},
		{
			name: "Simple URL with http",
			args: args{
				o: &IBMPPCOptions{
					Authenticator: &core.NoAuthAuthenticator{},
					UserAccount:   "1234",
					Region:        "dal",
					Zone:          "dal12",
					URL:           "http://dal.ppc-aas.test.cloud.ibm.com",
				},
			},
			want: &IBMPPCSession{
				Options: &IBMPPCOptions{
					Authenticator: &core.NoAuthAuthenticator{},
				},
				CRNFormat: "crn:v1:staging:public:ppc-aas:dal12:a/1234:%s::",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIBMPPCSession(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIBMPPCSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if same, msg := isIBMPPCSessionEqual(got, tt.want); !same {
				t.Error(msg)
			}
		})
	}
}

func TestNewIBMPPCSessionViaEnv(t *testing.T) {
	os.Setenv("IBMCLOUD_POWER_ASP_ENDPOINT", "ppc-aas.test.cloud.ibm.com")
	type args struct {
		o *IBMPPCOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *IBMPPCSession
		wantErr bool
	}{
		{
			name: "URL from Env Without Region",
			args: args{
				o: &IBMPPCOptions{
					Authenticator: &core.NoAuthAuthenticator{},
					UserAccount:   "1234",
					Zone:          "dal12",
				},
			},
			want: &IBMPPCSession{
				Options: &IBMPPCOptions{
					Authenticator: &core.NoAuthAuthenticator{},
				},
				CRNFormat: "crn:v1:staging:public:ppc-aas:dal12:a/1234:%s::",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIBMPPCSession(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIBMPPCSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if same, msg := isIBMPPCSessionEqual(got, tt.want); !same {
				t.Error(msg)
			}
		})
	}
}

func isIBMPPCSessionEqual(x *IBMPPCSession, y *IBMPPCSession) (bool, string) {
	if x == nil && y == nil {
		return true, ""
	}
	if (x == nil && y != nil) || (x != nil && y == nil) {
		return true, fmt.Sprintf("NewIBMPPCSession() = %v, want %v", x, y)
	}
	if x.Options.Authenticator != y.Options.Authenticator {
		return false, fmt.Sprintf("NewIBMPPCSession() Authenticator = %v, want %v", x.Options.Authenticator, y.Options.Authenticator)
	}
	if x.CRNFormat != y.CRNFormat {
		return false, fmt.Sprintf("NewIBMPPCSession() CRNFormat = %v, want %v", x.CRNFormat, y.CRNFormat)
	}
	return true, ""
}

func TestIBMPPCSession_AuthInfo(t *testing.T) {
	type args struct {
		cloudInstanceID string
	}
	type wantHeaders struct {
		auth string
		crn  string
	}
	tests := []struct {
		name        string
		s           *IBMPPCSession
		args        args
		wantHeaders wantHeaders
		wantErr     bool
	}{
		{
			name: "Auth",
			s: &IBMPPCSession{
				Options: &IBMPPCOptions{
					Authenticator: &core.BearerTokenAuthenticator{BearerToken: "sample"},
					Zone:          "dal12",
					URL:           "dal.ppc-aas.test.cloud.ibm.com",
				},
				CRNFormat: "crn...:%s::",
			},
			args: args{
				cloudInstanceID: "1234",
			},
			wantHeaders: wantHeaders{
				auth: "Bearer sample",
				crn:  "crn...:1234::",
			},
		},
		{
			name: "Incorrect Auth",
			s: &IBMPPCSession{
				Options: &IBMPPCOptions{
					Authenticator: &core.IamAuthenticator{ApiKey: "sample"},
					Zone:          "dal12",
					URL:           "dal.ppc-aas.test.cloud.ibm.com",
				},
				CRNFormat: "crn...:%s::",
			},
			args: args{
				cloudInstanceID: "1234",
			},
			wantHeaders: wantHeaders{},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.AuthInfo(tt.args.cloudInstanceID)
			if got == nil {
				t.Errorf("NewAuth() = %v", got)
			}
			// Authenticate the request
			r := &runtime.TestClientRequest{}
			err := got.AuthenticateRequest(r, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthInfo().AuthenticateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Test the header values
			auth := r.Headers.Get("Authorization")
			if auth != tt.wantHeaders.auth {
				t.Errorf("AuthInfo().AuthenticateRequest() Authorization = %v, want %v", auth, tt.wantHeaders.auth)
			}
			crn := r.Headers.Get("CRN")
			if crn != tt.wantHeaders.crn {
				t.Errorf("AuthInfo().AuthenticateRequest() CRN = %v, want %v", crn, tt.wantHeaders.crn)
			}
		})
	}
}
