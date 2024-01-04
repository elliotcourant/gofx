// Code generated by xgen. DO NOT EDIT.

package schema

// AbstractWireResponse ...
type AbstractWireResponse struct {
}

// WireCancellationResponse is The OFX element "WIRECANRS" is of type "WireCancellationResponse"
type WireCancellationResponse struct {
	SRVRTID string `xml:"SRVRTID"`
	*AbstractWireResponse
}

// AbstractWireRequest ...
type AbstractWireRequest struct {
}

// WireRequest is The OFX element "WIRERQ" is of type "WireRequest"
type WireRequest struct {
	BANKACCTFROM    *BankAccount         `xml:"BANKACCTFROM"`
	WIREBENEFICIARY *WireBeneficiary     `xml:"WIREBENEFICIARY"`
	WIREDESTBANK    *WireDestinationBank `xml:"WIREDESTBANK"`
	TRNAMT          string               `xml:"TRNAMT"`
	DTDUE           string               `xml:"DTDUE"`
	PAYINSTRUCT     string               `xml:"PAYINSTRUCT"`
	*AbstractWireRequest
}

// WireCancellationRequest is The OFX element "WIRECANRQ" is of type "WireCancellationRequest"
type WireCancellationRequest struct {
	SRVRTID string `xml:"SRVRTID"`
	*AbstractWireRequest
}

// WireResponse is The OFX element "WIRERS" is of type "WireResponse"
type WireResponse struct {
	CURDEF          string               `xml:"CURDEF"`
	SRVRTID         string               `xml:"SRVRTID"`
	BANKACCTFROM    *BankAccount         `xml:"BANKACCTFROM"`
	WIREBENEFICIARY *WireBeneficiary     `xml:"WIREBENEFICIARY"`
	WIREDESTBANK    *WireDestinationBank `xml:"WIREDESTBANK"`
	TRNAMT          string               `xml:"TRNAMT"`
	DTDUE           string               `xml:"DTDUE"`
	PAYINSTRUCT     string               `xml:"PAYINSTRUCT"`
	DTXFERPRJ       string               `xml:"DTXFERPRJ"`
	DTPOSTED        string               `xml:"DTPOSTED"`
	FEE             string               `xml:"FEE"`
	CONFMSG         string               `xml:"CONFMSG"`
	*AbstractWireResponse
}