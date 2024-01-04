// Code generated by xgen. DO NOT EDIT.

package gofx

// PayeeSyncRequest is The OFX element "PAYEESYNCRQ" is of type "PayeeSyncRequest"
type PayeeSyncRequest struct {
	OFXEXTENSION *OFXExtensionType          `xml:"OFXEXTENSION"`
	PAYEETRNRQ   []*PayeeTransactionRequest `xml:"PAYEETRNRQ"`
	*AbstractSyncRequest
}

// PayeeSyncResponse is The OFX element "PAYEESYNCRS" is of type "PayeeSyncResponse"
type PayeeSyncResponse struct {
	OFXEXTENSION *OFXExtensionType           `xml:"OFXEXTENSION"`
	PAYEETRNRS   []*PayeeTransactionResponse `xml:"PAYEETRNRS"`
	*AbstractSyncResponse
}

// PayeeTransactionRequest is The OFX element "PAYEETRNRQ" is of type "PayeeTransactionRequest"
type PayeeTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType   `xml:"OFXEXTENSION"`
	PAYEERQ      *PayeeRequest       `xml:"PAYEERQ"`
	PAYEEMODRQ   *PayeeModRequest    `xml:"PAYEEMODRQ"`
	PAYEEDELRQ   *PayeeDeleteRequest `xml:"PAYEEDELRQ"`
	*AbstractTransactionRequest
}

// PayeeTransactionResponse is The OFX element "PAYEETRNRS" is of type "PayeeTransactionResponse"
type PayeeTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType    `xml:"OFXEXTENSION"`
	PAYEERS      *PayeeResponse       `xml:"PAYEERS"`
	PAYEEMODRS   *PayeeModResponse    `xml:"PAYEEMODRS"`
	PAYEEDELRS   *PayeeDeleteResponse `xml:"PAYEEDELRS"`
	*AbstractTransactionResponse
}

// PaymentInquiryTransactionRequest is The OFX element "PMTINQTRNRQ" is of type "PaymentInquiryTransactionRequest"
type PaymentInquiryTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType      `xml:"OFXEXTENSION"`
	PMTINQRQ     *PaymentInquiryRequest `xml:"PMTINQRQ"`
	*AbstractTransactionRequest
}

// PaymentInquiryTransactionResponse is The OFX element "PMTINQTRNRS" is of type "PaymentInquiryTransactionResponse"
type PaymentInquiryTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType       `xml:"OFXEXTENSION"`
	PMTINQRS     *PaymentInquiryResponse `xml:"PMTINQRS"`
	*AbstractTransactionResponse
}

// PaymentMailTransactionRequest is The OFX element "PMTMAILTRNRQ" is of type "PaymentMailTransactionRequest"
type PaymentMailTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType   `xml:"OFXEXTENSION"`
	PMTMAILRQ    *PaymentMailRequest `xml:"PMTMAILRQ"`
	*AbstractTransactionRequest
}

// PaymentMailTransactionResponse is The OFX element "PMTMAILTRNRS" is of type "PaymentMailTransactionResponse"
type PaymentMailTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType    `xml:"OFXEXTENSION"`
	PMTMAILRS    *PaymentMailResponse `xml:"PMTMAILRS"`
	*AbstractTransactionResponse
}

// PaymentSyncRequest is The OFX element "PMTSYNCRQ" is of type "PaymentSyncRequest"
type PaymentSyncRequest struct {
	BANKACCTFROM *BankAccount                 `xml:"BANKACCTFROM"`
	OFXEXTENSION *OFXExtensionType            `xml:"OFXEXTENSION"`
	PMTTRNRQ     []*PaymentTransactionRequest `xml:"PMTTRNRQ"`
	*AbstractSyncRequest
}

// PaymentSyncResponse is The OFX element "PMTSYNCRS" is of type "PaymentSyncResponse"
type PaymentSyncResponse struct {
	BANKACCTFROM *BankAccount                  `xml:"BANKACCTFROM"`
	OFXEXTENSION *OFXExtensionType             `xml:"OFXEXTENSION"`
	PMTTRNRS     []*PaymentTransactionResponse `xml:"PMTTRNRS"`
	*AbstractSyncResponse
}

// PaymentTransactionRequest is The OFX element "PMTTRNRQ" is of type "PaymentTransactionRequest"
type PaymentTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType     `xml:"OFXEXTENSION"`
	PMTRQ        *PaymentRequest       `xml:"PMTRQ"`
	PMTMODRQ     *PaymentModRequest    `xml:"PMTMODRQ"`
	PMTCANCRQ    *PaymentCancelRequest `xml:"PMTCANCRQ"`
	*AbstractTransactionRequest
}

// PaymentTransactionResponse is The OFX element "PMTTRNRS" is of type "PaymentTransactionResponse"
type PaymentTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType      `xml:"OFXEXTENSION"`
	PMTRS        *PaymentResponse       `xml:"PMTRS"`
	PMTMODRS     *PaymentModResponse    `xml:"PMTMODRS"`
	PMTCANCRS    *PaymentCancelResponse `xml:"PMTCANCRS"`
	*AbstractTransactionResponse
}

// RecurringPaymentSyncRequest is The OFX element "RECPMTSYNCRQ" is of type "RecurringPaymentSyncRequest"
type RecurringPaymentSyncRequest struct {
	BANKACCTFROM *BankAccount                          `xml:"BANKACCTFROM"`
	OFXEXTENSION *OFXExtensionType                     `xml:"OFXEXTENSION"`
	RECPMTTRNRQ  []*RecurringPaymentTransactionRequest `xml:"RECPMTTRNRQ"`
	*AbstractSyncRequest
}

// RecurringPaymentSyncResponse is The OFX element "RECPMTSYNCRS" is of type "RecurringPaymentSyncResponse"
type RecurringPaymentSyncResponse struct {
	BANKACCTFROM *BankAccount                           `xml:"BANKACCTFROM"`
	OFXEXTENSION *OFXExtensionType                      `xml:"OFXEXTENSION"`
	RECPMTTRNRS  []*RecurringPaymentTransactionResponse `xml:"RECPMTTRNRS"`
	*AbstractSyncResponse
}

// RecurringPaymentTransactionRequest is The OFX element "RECPMTTRNRQ" is of type "RecurringPaymentTransactionRequest"
type RecurringPaymentTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType              `xml:"OFXEXTENSION"`
	RECPMTRQ     *RecurringPaymentRequest       `xml:"RECPMTRQ"`
	RECPMTMODRQ  *RecurringPaymentModRequest    `xml:"RECPMTMODRQ"`
	RECPMTCANCRQ *RecurringPaymentCancelRequest `xml:"RECPMTCANCRQ"`
	*AbstractTransactionRequest
}

// RecurringPaymentTransactionResponse is The OFX element "RECPMTTRNRS" is of type "RecurringPaymentTransactionResponse"
type RecurringPaymentTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType               `xml:"OFXEXTENSION"`
	RECPMTRS     *RecurringPaymentResponse       `xml:"RECPMTRS"`
	RECPMTMODRS  *RecurringPaymentModResponse    `xml:"RECPMTMODRS"`
	RECPMTCANCRS *RecurringPaymentCancelResponse `xml:"RECPMTCANCRS"`
	*AbstractTransactionResponse
}

// PaymentMailSyncRequest is The OFX element "PMTMAILSYNCRQ" is of type "PaymentMailSyncRequest"
type PaymentMailSyncRequest struct {
	INCIMAGES    string                           `xml:"INCIMAGES"`
	USEHTML      string                           `xml:"USEHTML"`
	OFXEXTENSION *OFXExtensionType                `xml:"OFXEXTENSION"`
	PMTMAILTRNRQ []*PaymentMailTransactionRequest `xml:"PMTMAILTRNRQ"`
	*AbstractSyncRequest
}

// PaymentMailSyncResponse is The OFX element "PMTMAILSYNCRS" is of type "PaymentMailSyncResponse"
type PaymentMailSyncResponse struct {
	OFXEXTENSION *OFXExtensionType                 `xml:"OFXEXTENSION"`
	PMTMAILTRNRS []*PaymentMailTransactionResponse `xml:"PMTMAILTRNRS"`
	*AbstractSyncResponse
}
