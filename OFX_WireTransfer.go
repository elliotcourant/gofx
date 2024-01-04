// Code generated by xgen. DO NOT EDIT.

package schema

// ExtBankDescription is The OFX element "EXTBANKDESC" is of type "ExtBankDescription"
type ExtBankDescription struct {
	NAME       string `xml:"NAME"`
	BANKID     string `xml:"BANKID"`
	ADDR1      string `xml:"ADDR1"`
	ADDR2      string `xml:"ADDR2"`
	ADDR3      string `xml:"ADDR3"`
	CITY       string `xml:"CITY"`
	STATE      string `xml:"STATE"`
	POSTALCODE string `xml:"POSTALCODE"`
	COUNTRY    string `xml:"COUNTRY"`
	PHONE      string `xml:"PHONE"`
}

// WireBeneficiary is The OFX element "WIREBENEFICIARY" is of type "WireBeneficiary"
type WireBeneficiary struct {
	NAME       string       `xml:"NAME"`
	BANKACCTTO *BankAccount `xml:"BANKACCTTO"`
	MEMO       string       `xml:"MEMO"`
}

// WireDestinationBank is The OFX element "WIREDESTBANK" is of type "WireDestinationBank"
type WireDestinationBank struct {
	EXTBANKDESC *ExtBankDescription `xml:"EXTBANKDESC"`
}