// Code generated by xgen. DO NOT EDIT.

package schema

// BillerInfo is The OFX element "BILLERINFO" is of type "BillerInfo"
type BillerInfo struct {
	BILLPUB            string              `xml:"BILLPUB"`
	BILLERID           string              `xml:"BILLERID"`
	NAME               string              `xml:"NAME"`
	ADDR1              string              `xml:"ADDR1"`
	ADDR2              string              `xml:"ADDR2"`
	ADDR3              string              `xml:"ADDR3"`
	CITY               string              `xml:"CITY"`
	STATE              string              `xml:"STATE"`
	POSTALCODE         string              `xml:"POSTALCODE"`
	COUNTRY            string              `xml:"COUNTRY"`
	SIC                string              `xml:"SIC"`
	PHONE              string              `xml:"PHONE"`
	PAYMENTINSTRUMENTS *PaymentInstruments `xml:"PAYMENTINSTRUMENTS"`
	ACCTFORMAT         string              `xml:"ACCTFORMAT"`
	ACCTEDITMASK       string              `xml:"ACCTEDITMASK"`
	HELPMESSAGE        string              `xml:"HELPMESSAGE"`
	RESTRICT           string              `xml:"RESTRICT"`
	LOGO               string              `xml:"LOGO"`
	VALIDATE           string              `xml:"VALIDATE"`
	BILLERINFOURL      string              `xml:"BILLERINFOURL"`
}

// PaymentInstrument is The OFX element "PAYMENTINSTRUMENT" is of type "PaymentInstrument"
type PaymentInstrument struct {
	PMTINSTRUMENTTYPE string `xml:"PMTINSTRUMENTTYPE"`
	BRAND             string `xml:"BRAND"`
}

// PaymentInstruments is The OFX element "PAYMENTINSTRUMENTS" is of type "PaymentInstruments"
type PaymentInstruments struct {
	PAYMENTINSTRUMENT []*PaymentInstrument `xml:"PAYMENTINSTRUMENT"`
}

// PaymentInstrumentEnum is The OFX element "PMTINSTRUMENTENUM" is of type "PaymentInstrumentEnum"
type PaymentInstrumentEnum string
