package saml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSignedRequest(t *testing.T) {
	assert := assert.New(t)
	sp := ServiceProviderSettings{
		PublicCertPath:              "certs/idp.crt",
		PrivateKeyPath:              "certs/idp.key",
		IDPSSOURL:                   "http://www.onelogin.net",
		IDPSSODescriptorURL:         "http://www.onelogin.net",
		IDPPublicCertPath:           "certs/idp.crt",
		AssertionConsumerServiceURL: "http://localhost:8000/auth/saml/name",
		SPSignRequest:               true,
	}
	err := sp.Init()
	assert.NoError(err)

	// Construct an AuthnRequest
	authnRequest := sp.GetAuthnRequest()
	signedXML, err := authnRequest.SignedString(sp.PrivateKeyPath)
	assert.NoError(err)
	assert.NotEmpty(signedXML)

	err = VerifyRequestSignature(signedXML, sp.PublicCertPath)
	assert.NoError(err)
}

func TestGetUnsignedRequest(t *testing.T) {
	assert := assert.New(t)
	sp := ServiceProviderSettings{
		IDPSSOURL:                   "http://www.onelogin.net",
		IDPSSODescriptorURL:         "http://www.onelogin.net",
		IDPPublicCertPath:           "certs/idp.crt",
		AssertionConsumerServiceURL: "http://localhost:8000/auth/saml/name",
		SPSignRequest:               false,
	}
	err := sp.Init()
	assert.NoError(err)

	// Construct an AuthnRequest
	authnRequest := sp.GetAuthnRequest()
	assert.NoError(err)
	assert.NotEmpty(authnRequest)
}
