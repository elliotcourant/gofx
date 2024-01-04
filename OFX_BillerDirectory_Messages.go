// Code generated by xgen. DO NOT EDIT.

package schema

// FindBillerRequest is The OFX element "FINDBILLERRQ" is of type "FindBillerRequest"
type FindBillerRequest struct {
	DTUPDATE        string `xml:"DTUPDATE"`
	BILLERID        string `xml:"BILLERID"`
	NAME            string `xml:"NAME"`
	ADDR1           string `xml:"ADDR1"`
	ADDR2           string `xml:"ADDR2"`
	ADDR3           string `xml:"ADDR3"`
	CITY            string `xml:"CITY"`
	STATE           string `xml:"STATE"`
	POSTALCODE      string `xml:"POSTALCODE"`
	COUNTRY         string `xml:"COUNTRY"`
	SIC             string `xml:"SIC"`
	CONSUPOSTALCODE string `xml:"CONSUPOSTALCODE"`
	INCIMAGES       string `xml:"INCIMAGES"`
}

// FindBillerResponse is The OFX element "FINDBILLERRS" is of type "FindBillerResponse"
type FindBillerResponse struct {
	DTUPDATE   string        `xml:"DTUPDATE"`
	BILLERINFO []*BillerInfo `xml:"BILLERINFO"`
}