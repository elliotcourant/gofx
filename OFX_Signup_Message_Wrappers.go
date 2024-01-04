// Code generated by xgen. DO NOT EDIT.

package schema

// AccountInfoTransactionRequest is The OFX element "ACCTINFOTRNRQ" is of type "AccountInfoTransactionRequest"
type AccountInfoTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType   `xml:"OFXEXTENSION"`
	ACCTINFORQ   *AccountInfoRequest `xml:"ACCTINFORQ"`
	*AbstractTransactionRequest
}

// AccountInfoTransactionResponse is The OFX element "ACCTINFOTRNRS" is of type "AccountInfoTransactionResponse"
type AccountInfoTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType    `xml:"OFXEXTENSION"`
	ACCTINFORS   *AccountInfoResponse `xml:"ACCTINFORS"`
	*AbstractTransactionResponse
}

// AccountSyncRequest is The OFX element "ACCTSYNCRQ" is of type "AccountSyncRequest"
type AccountSyncRequest struct {
	OFXEXTENSION *OFXExtensionType            `xml:"OFXEXTENSION"`
	ACCTTRNRQ    []*AccountTransactionRequest `xml:"ACCTTRNRQ"`
	*AbstractSyncRequest
}

// AccountSyncResponse is The OFX element "ACCTSYNCRS" is of type "AccountSyncResponse"
type AccountSyncResponse struct {
	OFXEXTENSION *OFXExtensionType             `xml:"OFXEXTENSION"`
	ACCTTRNRS    []*AccountTransactionResponse `xml:"ACCTTRNRS"`
	*AbstractSyncResponse
}

// AccountTransactionRequest is The OFX element "ACCTTRNRQ" is of type "AccountTransactionRequest"
type AccountTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	ACCTRQ       *AccountRequest   `xml:"ACCTRQ"`
	*AbstractTransactionRequest
}

// AccountTransactionResponse is The OFX element "ACCTTRNRS" is of type "AccountTransactionResponse"
type AccountTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	ACCTRS       *AccountResponse  `xml:"ACCTRS"`
	*AbstractTransactionResponse
}

// UserInfoTransactionRequest is The OFX element "USERINFOTRNRQ" is of type "UserInfoTransactionRequest"
type UserInfoTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	USERINFORQ   *UserInfoRequest  `xml:"USERINFORQ"`
	*AbstractTransactionRequest
}

// UserInfoTransactionResponse is The OFX element "USERINFOTRNRS" is of type "UserInfoTransactionResponse"
type UserInfoTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	USERINFORS   *UserInfoResponse `xml:"USERINFORS"`
	*AbstractTransactionResponse
}

// ChangeUserInfoSyncRequest is The OFX element "CHGUSERINFOSYNCRQ" is of type "ChangeUserInfoSyncRequest"
type ChangeUserInfoSyncRequest struct {
	OFXEXTENSION     *OFXExtensionType                   `xml:"OFXEXTENSION"`
	CHGUSERINFOTRNRQ []*ChangeUserInfoTransactionRequest `xml:"CHGUSERINFOTRNRQ"`
	*AbstractSyncRequest
}

// ChangeUserInfoSyncResponse is The OFX element "CHGUSERINFOSYNCRS" is of type "ChangeUserInfoSyncResponse"
type ChangeUserInfoSyncResponse struct {
	OFXEXTENSION     *OFXExtensionType                    `xml:"OFXEXTENSION"`
	CHGUSERINFOTRNRS []*ChangeUserInfoTransactionResponse `xml:"CHGUSERINFOTRNRS"`
	*AbstractSyncResponse
}

// ChangeUserInfoTransactionRequest is The OFX element "CHGUSERINFOTRNRQ" is of type "ChangeUserInfoTransactionRequest"
type ChangeUserInfoTransactionRequest struct {
	OFXEXTENSION  *OFXExtensionType      `xml:"OFXEXTENSION"`
	CHGUSERINFORQ *ChangeUserInfoRequest `xml:"CHGUSERINFORQ"`
	*AbstractTransactionRequest
}

// ChangeUserInfoTransactionResponse is The OFX element "CHGUSERINFOTRNRS" is of type "ChangeUserInfoTransactionResponse"
type ChangeUserInfoTransactionResponse struct {
	OFXEXTENSION  *OFXExtensionType       `xml:"OFXEXTENSION"`
	CHGUSERINFORS *ChangeUserInfoResponse `xml:"CHGUSERINFORS"`
	*AbstractTransactionResponse
}

// EnrollTransactionRequest is The OFX element "ENROLLTRNRQ" is of type "EnrollTransactionRequest"
type EnrollTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	ENROLLRQ     *EnrollRequest    `xml:"ENROLLRQ"`
	*AbstractTransactionRequest
}

// EnrollTransactionResponse is The OFX element "ENROLLTRNRS" is of type "EnrollTransactionResponse"
type EnrollTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	ENROLLRS     *EnrollResponse   `xml:"ENROLLRS"`
	*AbstractTransactionResponse
}