// Code generated by xgen. DO NOT EDIT.

package schema

// SecurityListTransactionRequest is The OFX element "SECLISTTRNRQ" is of type "SecurityListTransactionRequest"
type SecurityListTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType    `xml:"OFXEXTENSION"`
	SECLISTRQ    *SecurityListRequest `xml:"SECLISTRQ"`
	*AbstractTransactionRequest
}

// SecurityListTransactionResponse is The OFX element "SECLISTTRNRS" is of type "SecurityListTransactionResponse"
type SecurityListTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	SECLISTRS    string            `xml:"SECLISTRS"`
	*AbstractTransactionResponse
}