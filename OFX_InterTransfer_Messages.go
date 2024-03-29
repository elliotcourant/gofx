// Code generated by xgen. DO NOT EDIT.

package gofx

// InterCancellationResponse is The OFX element "INTERCANRS" is of type "InterCancellationResponse"
type InterCancellationResponse struct {
	SRVRTID string `xml:"SRVRTID"`
	*AbstractInterResponse
}

// InterModResponse is The OFX element "INTERMODRS" is of type "InterModResponse"
type InterModResponse struct {
	SRVRTID    string                    `xml:"SRVRTID"`
	XFERINFO   *TransferInfo             `xml:"XFERINFO"`
	XFERPRCSTS *TransferProcessingStatus `xml:"XFERPRCSTS"`
	*AbstractInterResponse
}

// AbstractInterRequest ...
type AbstractInterRequest struct {
}

// InterRequest is The OFX element "INTERRQ" is of type "InterRequest"
type InterRequest struct {
	XFERINFO *TransferInfo `xml:"XFERINFO"`
	*AbstractInterRequest
}

// InterModRequest is The OFX element "INTERMODRQ" is of type "InterModRequest"
type InterModRequest struct {
	SRVRTID  string        `xml:"SRVRTID"`
	XFERINFO *TransferInfo `xml:"XFERINFO"`
	*AbstractInterRequest
}

// InterCancellationRequest is The OFX element "INTERCANRQ" is of type "InterCancellationRequest"
type InterCancellationRequest struct {
	SRVRTID string `xml:"SRVRTID"`
	*AbstractInterRequest
}

// AbstractInterResponse ...
type AbstractInterResponse struct {
}

// InterResponse is The OFX element "INTERRS" is of type "InterResponse"
type InterResponse struct {
	CURDEF     string                    `xml:"CURDEF"`
	SRVRTID    string                    `xml:"SRVRTID"`
	XFERINFO   *TransferInfo             `xml:"XFERINFO"`
	DTXFERPRJ  string                    `xml:"DTXFERPRJ"`
	DTPOSTED   string                    `xml:"DTPOSTED"`
	REFNUM     string                    `xml:"REFNUM"`
	RECSRVRTID string                    `xml:"RECSRVRTID"`
	XFERPRCSTS *TransferProcessingStatus `xml:"XFERPRCSTS"`
	*AbstractInterResponse
}

// RecurringInterCancellationResponse is The OFX element "RECINTERCANRS" is of type "RecurringInterCancellationResponse"
type RecurringInterCancellationResponse struct {
	CANPENDING string `xml:"CANPENDING"`
	*AbstractRecurringInterResponse
}

// RecurringInterModResponse is The OFX element "RECINTERMODRS" is of type "RecurringInterModResponse"
type RecurringInterModResponse struct {
	RECURRINST *RecurringInstructions `xml:"RECURRINST"`
	INTERRS    *InterResponse         `xml:"INTERRS"`
	MODPENDING string                 `xml:"MODPENDING"`
	*AbstractRecurringInterResponse
}

// AbstractRecurringInterRequest ...
type AbstractRecurringInterRequest struct {
}

// RecurringInterRequest is The OFX element "RECINTERRQ" is of type "RecurringInterRequest"
type RecurringInterRequest struct {
	RECURRINST *RecurringInstructions `xml:"RECURRINST"`
	INTERRQ    *InterRequest          `xml:"INTERRQ"`
	*AbstractRecurringInterRequest
}

// RecurringInterModRequest is The OFX element "RECINTERMODRQ" is of type "RecurringInterModRequest"
type RecurringInterModRequest struct {
	RECSRVRTID string                 `xml:"RECSRVRTID"`
	RECURRINST *RecurringInstructions `xml:"RECURRINST"`
	INTERRQ    *InterRequest          `xml:"INTERRQ"`
	MODPENDING string                 `xml:"MODPENDING"`
	*AbstractRecurringInterRequest
}

// RecurringInterCancellationRequest is The OFX element "RECINTERCANRQ" is of type "RecurringInterCancellationRequest"
type RecurringInterCancellationRequest struct {
	RECSRVRTID string `xml:"RECSRVRTID"`
	CANPENDING string `xml:"CANPENDING"`
	*AbstractRecurringInterRequest
}

// AbstractRecurringInterResponse ...
type AbstractRecurringInterResponse struct {
	RECSRVRTID string `xml:"RECSRVRTID"`
}

// RecurringInterResponse is The OFX element "RECINTERRS" is of type "RecurringInterResponse"
type RecurringInterResponse struct {
	RECURRINST *RecurringInstructions `xml:"RECURRINST"`
	INTERRS    *InterResponse         `xml:"INTERRS"`
	*AbstractRecurringInterResponse
}
