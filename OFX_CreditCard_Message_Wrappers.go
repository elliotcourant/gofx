// Code generated by xgen. DO NOT EDIT.

package gofx

// CreditCardStatementEndTransactionRequest is The OFX element "CCSTMTENDTRNRQ" is of type "CreditCardStatementEndTransactionRequest"
type CreditCardStatementEndTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType              `xml:"OFXEXTENSION"`
	CCSTMTENDRQ  *CreditCardStatementEndRequest `xml:"CCSTMTENDRQ"`
	*AbstractTransactionRequest
}

// CreditCardStatementEndTransactionResponse is The OFX element "CCSTMTENDTRNRS" is of type "CreditCardStatementEndTransactionResponse"
type CreditCardStatementEndTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType               `xml:"OFXEXTENSION"`
	CCSTMTENDRS  *CreditCardStatementEndResponse `xml:"CCSTMTENDRS"`
	*AbstractTransactionResponse
}

// CreditCardStatementTransactionRequest is The OFX element "CCSTMTTRNRQ" is of type "CreditCardStatementTransactionRequest"
type CreditCardStatementTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType           `xml:"OFXEXTENSION"`
	CCSTMTRQ     *CreditCardStatementRequest `xml:"CCSTMTRQ"`
	*AbstractTransactionRequest
}

// CreditCardStatementTransactionResponse is The OFX element "CCSTMTTRNRS" is of type "CreditCardStatementTransactionResponse"
type CreditCardStatementTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType            `xml:"OFXEXTENSION"`
	CCSTMTRS     *CreditCardStatementResponse `xml:"CCSTMTRS"`
	*AbstractTransactionResponse
}
