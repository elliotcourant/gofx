// Code generated by xgen. DO NOT EDIT.

package schema

// BillStatusModRequest is The OFX element "BILLSTATUSMODRQ" is of type "BillStatusModRequest"
type BillStatusModRequest struct {
	BILLID        string             `xml:"BILLID"`
	BILLSTATUS    *BillStatus        `xml:"BILLSTATUS"`
	BILLPMTSTATUS *BillPaymentstatus `xml:"BILLPMTSTATUS"`
}

// BillStatusModResponse is The OFX element "BILLSTATUSMODRS" is of type "BillStatusModResponse"
type BillStatusModResponse struct {
	BILLID        string             `xml:"BILLID"`
	BILLSTATUS    *BillStatus        `xml:"BILLSTATUS"`
	BILLPMTSTATUS *BillPaymentstatus `xml:"BILLPMTSTATUS"`
}

// BillTableStructureRequest is The OFX element "BILLTBLSTRUCTRQ" is of type "BillTableStructureRequest"
type BillTableStructureRequest struct {
	BILLID              string `xml:"BILLID"`
	BILLDETAILTABLETYPE string `xml:"BILLDETAILTABLETYPE"`
}

// BillTableStructureResponse is The OFX element "BILLTBLSTRUCTRS" is of type "BillTableStructureResponse"
type BillTableStructureResponse struct {
	BILLID              string              `xml:"BILLID"`
	BILLDETAILTABLETYPE string              `xml:"BILLDETAILTABLETYPE"`
	COLDEF              []*ColumnDefinition `xml:"COLDEF"`
}

// PresentmentDetailRequest is The OFX element "PRESDETAILRQ" is of type "PresentmentDetailRequest"
type PresentmentDetailRequest struct {
	BILLID              string `xml:"BILLID"`
	BILLDETAILTABLETYPE string `xml:"BILLDETAILTABLETYPE"`
}

// PresentmentDetailResponse is The OFX element "PRESDETAILRS" is of type "PresentmentDetailResponse"
type PresentmentDetailResponse struct {
	PRESDETAIL []*PresentmentDetail `xml:"PRESDETAIL"`
}

// PresentmentListRequest is The OFX element "PRESLISTRQ" is of type "PresentmentListRequest"
type PresentmentListRequest struct {
	BILLPUB              string   `xml:"BILLPUB"`
	DTSTART              string   `xml:"DTSTART"`
	DTEND                string   `xml:"DTEND"`
	DTDUEBY              string   `xml:"DTDUEBY"`
	BILLERID             string   `xml:"BILLERID"`
	BILLID               string   `xml:"BILLID"`
	BILLTYPE             []string `xml:"BILLTYPE"`
	BILLSTATUSCODE       []string `xml:"BILLSTATUSCODE"`
	BILLPMTSTATUSCODE    []string `xml:"BILLPMTSTATUSCODE"`
	NOTIFYWILLING        string   `xml:"NOTIFYWILLING"`
	INCLUDEDETAIL        string   `xml:"INCLUDEDETAIL"`
	INCLUDEBILLSTATUS    string   `xml:"INCLUDEBILLSTATUS"`
	INCLUDEBILLPMTSTATUS string   `xml:"INCLUDEBILLPMTSTATUS"`
	INCLUDESTATUSHIST    string   `xml:"INCLUDESTATUSHIST"`
	INCLUDECOUNTS        string   `xml:"INCLUDECOUNTS"`
	INCLUDESUMMARY       string   `xml:"INCLUDESUMMARY"`
}

// PresentmentListResponse is The OFX element "PRESLISTRS" is of type "PresentmentListResponse"
type PresentmentListResponse struct {
	BILLPUB    string             `xml:"BILLPUB"`
	USERID     string             `xml:"USERID"`
	DTSTART    string             `xml:"DTSTART"`
	DTEND      string             `xml:"DTEND"`
	PRESLIST   *PresentmentList   `xml:"PRESLIST"`
	PRESCOUNTS *PresentmentCounts `xml:"PRESCOUNTS"`
}

// PresentmentMailRequest is The OFX element "PRESMAILRQ" is of type "PresentmentMailRequest"
type PresentmentMailRequest struct {
	PRESACCTFROM *PresentmentAccount `xml:"PRESACCTFROM"`
	MAIL         *Mail               `xml:"MAIL"`
}

// PresentmentMailResponse is The OFX element "PRESMAILRS" is of type "PresentmentMailResponse"
type PresentmentMailResponse struct {
	PRESACCTFROM *PresentmentAccount `xml:"PRESACCTFROM"`
	MAIL         *Mail               `xml:"MAIL"`
}

// PresentmentNotifyRequest is The OFX element "PRESNOTIFYRQ" is of type "PresentmentNotifyRequest"
type PresentmentNotifyRequest struct {
	PRESDELIVERYID *PresentmentDeliveryId `xml:"PRESDELIVERYID"`
}

// PresentmentNotifyResponse is The OFX element "PRESNOTIFYRS" is of type "PresentmentNotifyResponse"
type PresentmentNotifyResponse struct {
	PRESDELIVERYID *PresentmentDeliveryId `xml:"PRESDELIVERYID"`
}
