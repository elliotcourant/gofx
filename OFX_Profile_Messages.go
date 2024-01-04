// Code generated by xgen. DO NOT EDIT.

package schema

// ProfileRequest is The OFX element "PROFRQ" is of type "ProfileRequest"
type ProfileRequest struct {
	CLIENTROUTING string `xml:"CLIENTROUTING"`
	DTPROFUP      string `xml:"DTPROFUP"`
}

// ProfileResponse is The OFX element "PROFRS" is of type "ProfileResponse"
type ProfileResponse struct {
	MSGSETLIST     *MessageSetList `xml:"MSGSETLIST"`
	SIGNONINFOLIST *SignonInfoList `xml:"SIGNONINFOLIST"`
	DTPROFUP       string          `xml:"DTPROFUP"`
	FINAME         string          `xml:"FINAME"`
	ADDR1          string          `xml:"ADDR1"`
	ADDR2          string          `xml:"ADDR2"`
	ADDR3          string          `xml:"ADDR3"`
	CITY           string          `xml:"CITY"`
	STATE          string          `xml:"STATE"`
	POSTALCODE     string          `xml:"POSTALCODE"`
	COUNTRY        string          `xml:"COUNTRY"`
	CSPHONE        string          `xml:"CSPHONE"`
	TSPHONE        string          `xml:"TSPHONE"`
	FAXPHONE       string          `xml:"FAXPHONE"`
	URL            string          `xml:"URL"`
	EMAIL          string          `xml:"EMAIL"`
}