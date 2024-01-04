// Code generated by xgen. DO NOT EDIT.

package schema

// BillPubInfo is The OFX element "BILLPUBINFO" is of type "BillPubInfo"
type BillPubInfo struct {
	BILLPUB string `xml:"BILLPUB"`
	BILLID  string `xml:"BILLID"`
}

// ExtendedPayee is The OFX element "EXTDPAYEE" is of type "ExtendedPayee"
type ExtendedPayee struct {
	PAYEEID   string `xml:"PAYEEID"`
	IDSCOPE   string `xml:"IDSCOPE"`
	NAME      string `xml:"NAME"`
	DAYSTOPAY string `xml:"DAYSTOPAY"`
}

// ExtendedPayment is The OFX element "EXTDPMT" is of type "ExtendedPayment"
type ExtendedPayment struct {
	EXTDPMTFOR string                  `xml:"EXTDPMTFOR"`
	EXTDPMTCHK string                  `xml:"EXTDPMTCHK"`
	EXTDPMTDSC string                  `xml:"EXTDPMTDSC"`
	EXTDPMTINV *ExtendedPaymentInvoice `xml:"EXTDPMTINV"`
}

// ExtendedPaymentForEnum is The OFX element "EXTDPMTFORENUM" is of type "ExtendedPaymentForEnum"
type ExtendedPaymentForEnum string

// ExtendedPaymentInvoice is The OFX element "EXTDPMTINV" is of type "ExtendedPaymentInvoice"
type ExtendedPaymentInvoice struct {
	INVOICE []*Invoice `xml:"INVOICE"`
}

// ExtType is The OFX element "EXT" is of type "ExtType"
type ExtType string

// IdScopeEnum is The OFX element "IDSCOPEENUM" is of type "IdScopeEnum"
type IdScopeEnum string

// PaymentInfo is The OFX element "PMTINFO" is of type "PaymentInfo"
type PaymentInfo struct {
	BANKACCTFROM *BankAccount       `xml:"BANKACCTFROM"`
	TRNAMT       string             `xml:"TRNAMT"`
	PAYEEID      string             `xml:"PAYEEID"`
	PAYEE        *Payee             `xml:"PAYEE"`
	PAYEELSTID   string             `xml:"PAYEELSTID"`
	BANKACCTTO   *BankAccount       `xml:"BANKACCTTO"`
	EXTDPMT      []*ExtendedPayment `xml:"EXTDPMT"`
	PAYACCT      string             `xml:"PAYACCT"`
	DTDUE        string             `xml:"DTDUE"`
	MEMO         string             `xml:"MEMO"`
	BILLREFINFO  string             `xml:"BILLREFINFO"`
	BILLPUBINFO  *BillPubInfo       `xml:"BILLPUBINFO"`
}

// PaymentProcessingStatus is The OFX element "PMTPRCSTS" is of type "PaymentProcessingStatus"
type PaymentProcessingStatus struct {
	PMTPRCCODE string `xml:"PMTPRCCODE"`
	DTPMTPRC   string `xml:"DTPMTPRC"`
}

// PaymentProcessStatusEnum is The OFX element "PMTPROCESSSTATUSENUM" is of type "PaymentProcessStatusEnum"
type PaymentProcessStatusEnum string