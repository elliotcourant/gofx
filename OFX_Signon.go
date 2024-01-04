// Code generated by xgen. DO NOT EDIT.

package schema

// FinancialInstitution is The OFX element "FI" is of type "FinancialInstitution"
type FinancialInstitution struct {
	ORG string `xml:"ORG"`
	FID string `xml:"FID"`
}

// ApplicationIdType is The OFX element "APPID" is of type "ApplicationIdType"
type ApplicationIdType string

// FinancialInstitutionCertificateIdType is The OFX element "FICERTID" is of type "FinancialInstitutionCertificateIdType"
type FinancialInstitutionCertificateIdType string

// IntegerType is The OFX element "INT" is of type "IntegerType"
type IntegerType string

// NumberNCEType is The OFX element "NONCE" is of type "NumberNCEType"
type NumberNCEType string

// PasswordType is The OFX element "PASSWORD" is of type "PasswordType"
type PasswordType string

// SessionCookieType is The OFX element "SESSCOOKIE" is of type "SessionCookieType"
type SessionCookieType string

// UserKeyType is The OFX element "USERKEY" is of type "UserKeyType"
type UserKeyType string

// AccessKeyType is The OFX element "ACCESSKEY" is of type "AccessKeyType".
type AccessKeyType string

// MFAChallenge is The OFX element "MFACHALLENGE" is of type "MFAChallenge"
type MFAChallenge struct {
	MFAPHRASEID    string `xml:"MFAPHRASEID"`
	MFAPHRASELABEL string `xml:"MFAPHRASELABEL"`
}

// MFAChallengeAnswer is The OFX element "MFACHALLENGEANSWER" is of type "MFAChallengeAnswer".
type MFAChallengeAnswer struct {
	MFAPHRASEID string `xml:"MFAPHRASEID"`
	MFAPHRASEA  string `xml:"MFAPHRASEA"`
}

// MFAPhraseType is The OFX element "MFAPHRASEA" is of type "MFAPhraseType".
//
//	The OFX element "MFAPHRASELABEL" is of type "MFAPhraseType".
type MFAPhraseType string

// UserCredType is The OFX element "USERCRED1" is of type "UserCredType".
//
//	The OFX element "USERCRED2" is of type "UserCredType".
//	The effective size of USERCRED1 and USERCRED2 is A-32.  If Type 1
//	  security is used, then the actual field length is A-171.
type UserCredType string

// AuthTokenType is The OFX element "AUTHTOKEN" is of type "AuthTokenType".
//
//	The effective size of AUTHTOKEN is A-32.  If Type 1
//	  security is used, then the actual field length is A-171.
type AuthTokenType string

// ClientUidType is The OFX element "CLIENTUID" is of type "ClientUidType"
type ClientUidType string

// AccessTokenType is The OFX element "ACCESSTOKEN" is of type "AccessTokenType"
type AccessTokenType string

// AppKeyType is The OFX element "APPKEY" is of type "AppKeyType"
type AppKeyType string