// Code generated by xgen. DO NOT EDIT.

package gofx

// ProfileTransactionRequest is The OFX element "PROFTRNRQ" is of type "ProfileTransactionRequest"
type ProfileTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	PROFRQ       *ProfileRequest   `xml:"PROFRQ"`
	*AbstractTransactionRequest
}

// ProfileTransactionResponse is The OFX element "PROFTRNRS" is of type "ProfileTransactionResponse"
type ProfileTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	PROFRS       *ProfileResponse  `xml:"PROFRS"`
	*AbstractTransactionResponse
}
