// Code generated by xgen. DO NOT EDIT.

package gofx

// LoanStatementTransactionRequest is The OFX element "LOANSTMTTRNRQ" is of type "LoanStatementTransactionRequest"
type LoanStatementTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType     `xml:"OFXEXTENSION"`
	LOANSTMTRQ   *LoanStatementRequest `xml:"LOANSTMTRQ"`
	*AbstractTransactionRequest
}

// LoanStatementTransactionResponse is The OFX element "LOANSTMTTRNRS" is of type "LoanStatementTransactionResponse"
type LoanStatementTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType      `xml:"OFXEXTENSION"`
	LOANSTMTRS   *LoanStatementResponse `xml:"LOANSTMTRS"`
	*AbstractTransactionResponse
}

// AmortizationTransactionRequest is The OFX element "AMRTSTMTTRNRQ" is of type "AmortizationTransactionRequest"
type AmortizationTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType    `xml:"OFXEXTENSION"`
	AMRTSTMTRQ   *AmortizationRequest `xml:"AMRTSTMTRQ"`
	*AbstractTransactionRequest
}

// AmortizationTransactionResponse is The OFX element "AMRTSTMTTRNRS" is of type "AmortizationTransactionResponse"
type AmortizationTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType     `xml:"OFXEXTENSION"`
	AMRTSTMTRS   *AmortizationResponse `xml:"AMRTSTMTRS"`
	*AbstractTransactionResponse
}

// LoanStatementEndTransactionRequest is The OFX element "LOANSTMTENDTRNRQ" is of type "LoanStatementEndTransactionRequest"
type LoanStatementEndTransactionRequest struct {
	OFXEXTENSION  *OFXExtensionType        `xml:"OFXEXTENSION"`
	LOANSTMTENDRQ *LoanStatementEndRequest `xml:"LOANSTMTENDRQ"`
	*AbstractTransactionRequest
}

// LoanStatementEndTransactionResponse is The OFX element "LOANSTMTENDTRNRS" is of type "LoanStatementEndTransactionResponse"
type LoanStatementEndTransactionResponse struct {
	OFXEXTENSION  *OFXExtensionType         `xml:"OFXEXTENSION"`
	LOANSTMTENDRS *LoanStatementEndResponse `xml:"LOANSTMTENDRS"`
	*AbstractTransactionResponse
}

// LoanMailTransactionRequest is The OFX element "LOANMAILTRNRQ" is of type "LoanMailTransactionRequest"
type LoanMailTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	LOANMAILRQ   *LoanMailRequest  `xml:"LOANMAILRQ"`
	*AbstractTransactionRequest
}

// LoanMailTransactionResponse is The OFX element "LOANMAILTRNRS" is of type "LoanMailTransactionResponse"
type LoanMailTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType `xml:"OFXEXTENSION"`
	LOANMAILRS   *LoanMailResponse `xml:"LOANMAILRS"`
	*AbstractTransactionResponse
}

// LoanMailSyncRequest is The OFX element "LOANMAILSYNCRQ" is of type "LoanMailSyncRequest"
type LoanMailSyncRequest struct {
	INCIMAGES     string                        `xml:"INCIMAGES"`
	USEHTML       string                        `xml:"USEHTML"`
	LOANACCTFROM  *LoanAccount                  `xml:"LOANACCTFROM"`
	OFXEXTENSION  *OFXExtensionType             `xml:"OFXEXTENSION"`
	LOANMAILTRNRQ []*LoanMailTransactionRequest `xml:"LOANMAILTRNRQ"`
	*AbstractSyncRequest
}

// LoanMailSyncResponse is The OFX element "LOANMAILSYNCRS" is of type "LoanMailSyncResponse"
type LoanMailSyncResponse struct {
	LOANACCTFROM  *LoanAccount                   `xml:"LOANACCTFROM"`
	OFXEXTENSION  *OFXExtensionType              `xml:"OFXEXTENSION"`
	LOANMAILTRNRS []*LoanMailTransactionResponse `xml:"LOANMAILTRNRS"`
	*AbstractSyncResponse
}
