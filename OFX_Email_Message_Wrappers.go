// Code generated by xgen. DO NOT EDIT.

package schema

// GetMimeTransactionRequest is The OFX element "GETMIMETRNRQ" is of type "GetMimeTransactionRequest"
type GetMimeTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	GETMIMERQ    *GetMimeRequest   `xml:"GETMIMERQ"`
	*AbstractTransactionRequest
}

// GetMimeTransactionResponse is The OFX element "GETMIMETRNRS" is of type "GetMimeTransactionResponse"
type GetMimeTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	GETMIMERS    *GetMimeResponse  `xml:"GETMIMERS"`
	*AbstractTransactionResponse
}

// MailSyncRequest is The OFX element "MAILSYNCRQ" is of type "MailSyncRequest"
type MailSyncRequest struct {
	INCIMAGES    string                    `xml:"INCIMAGES"`
	USEHTML      string                    `xml:"USEHTML"`
	OFXEXTENSION *OFXExtensionType         `xml:"OFXEXTENSION"`
	MAILTRNRQ    []*MailTransactionRequest `xml:"MAILTRNRQ"`
	*AbstractSyncRequest
}

// MailSyncResponse is The OFX element "MAILSYNCRS" is of type "MailSyncResponse"
type MailSyncResponse struct {
	OFXEXTENSION *OFXExtensionType          `xml:"OFXEXTENSION"`
	MAILTRNRS    []*MailTransactionResponse `xml:"MAILTRNRS"`
	*AbstractSyncResponse
}

// MailTransactionRequest is The OFX element "MAILTRNRQ" is of type "MailTransactionRequest"
type MailTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	MAILRQ       *MailRequest      `xml:"MAILRQ"`
	*AbstractTransactionRequest
}

// MailTransactionResponse is The OFX element "MAILTRNRS" is of type "MailTransactionResponse"
type MailTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	MAILRS       *MailResponse     `xml:"MAILRS"`
	*AbstractTransactionResponse
}
