// Code generated by xgen. DO NOT EDIT.

package schema

// ChallengeRequest is The OFX element "CHALLENGERQ" is of type "ChallengeRequest"
type ChallengeRequest struct {
	USERID   *IdType `xml:"USERID"`
	FICERTID string  `xml:"FICERTID"`
}

// ChallengeResponse is The OFX element "CHALLENGERS" is of type "ChallengeResponse"
type ChallengeResponse struct {
	USERID   string `xml:"USERID"`
	NONCE    string `xml:"NONCE"`
	FICERTID string `xml:"FICERTID"`
}

// PinChangeRequest is The OFX element "PINCHRQ" is of type "PinChangeRequest"
type PinChangeRequest struct {
	USERID      string `xml:"USERID"`
	NEWUSERPASS string `xml:"NEWUSERPASS"`
}

// PinChangeResponse is The OFX element "PINCHRS" is of type "PinChangeResponse"
type PinChangeResponse struct {
	USERID    string `xml:"USERID"`
	DTCHANGED string `xml:"DTCHANGED"`
}

// SignonRequest is The OFX element "SONRQ" is of type "SignonRequest"
type SignonRequest struct {
	DTCLIENT           string                `xml:"DTCLIENT"`
	USERID             string                `xml:"USERID"`
	USERPASS           string                `xml:"USERPASS"`
	USERKEY            string                `xml:"USERKEY"`
	ACCESSTOKEN        string                `xml:"ACCESSTOKEN"`
	GENUSERKEY         string                `xml:"GENUSERKEY"`
	LANGUAGE           string                `xml:"LANGUAGE"`
	FI                 *FinancialInstitution `xml:"FI"`
	SESSCOOKIE         string                `xml:"SESSCOOKIE"`
	APPID              string                `xml:"APPID"`
	APPVER             string                `xml:"APPVER"`
	APPKEY             string                `xml:"APPKEY"`
	CLIENTUID          string                `xml:"CLIENTUID"`
	USERCRED1          string                `xml:"USERCRED1"`
	USERCRED2          string                `xml:"USERCRED2"`
	AUTHTOKEN          string                `xml:"AUTHTOKEN"`
	ACCESSKEY          string                `xml:"ACCESSKEY"`
	MFACHALLENGEANSWER []*MFAChallengeAnswer `xml:"MFACHALLENGEANSWER"`
	OFXEXTENSION       *OFXExtensionType     `xml:"OFXEXTENSION"`
}

// SignonResponse is The OFX element "SONRS" is of type "SignonResponse"
type SignonResponse struct {
	STATUS       *Status               `xml:"STATUS"`
	DTSERVER     string                `xml:"DTSERVER"`
	USERKEY      string                `xml:"USERKEY"`
	TSKEYEXPIRE  string                `xml:"TSKEYEXPIRE"`
	LANGUAGE     string                `xml:"LANGUAGE"`
	DTPROFUP     string                `xml:"DTPROFUP"`
	DTUSERUP     string                `xml:"DTUSERUP"`
	DTACCTUP     string                `xml:"DTACCTUP"`
	FI           *FinancialInstitution `xml:"FI"`
	SESSCOOKIE   string                `xml:"SESSCOOKIE"`
	ACCESSKEY    string                `xml:"ACCESSKEY"`
	OFXEXTENSION *OFXExtensionType     `xml:"OFXEXTENSION"`
}

// MFAChallengeRequest is The OFX element "MFACHALLENGERQ" is of type "MFAChallengeRequest"
type MFAChallengeRequest struct {
	DTCLIENT string `xml:"DTCLIENT"`
}

// MFAChallengeResponse is The OFX element "MFACHALLENGERS" is of type "MFAChallengeResponse"
type MFAChallengeResponse struct {
	MFACHALLENGE []*MFAChallenge `xml:"MFACHALLENGE"`
}