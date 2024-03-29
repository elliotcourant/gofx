// Code generated by xgen. DO NOT EDIT.

package gofx

// AssetClassEnum is The OFX element "ASSETCLASSENUM" is of type "AssetClassEnum"
type AssetClassEnum string

// CallTypeEnum is The OFX element "CALLTYPEENUM" is of type "CallTypeEnum"
type CallTypeEnum string

// CouponFrequencyEnum is The OFX element "COUPONFREQENUM" is of type "CouponFrequencyEnum"
type CouponFrequencyEnum string

// DebtClassEnum is The OFX element "DEBTCLASSENUM" is of type "DebtClassEnum"
type DebtClassEnum string

// DebtEnum is The OFX element "DEBTENUM" is of type "DebtEnum"
type DebtEnum string

// DebtInfo is The OFX element "DEBTINFO" is of type "DebtInfo"
type DebtInfo struct {
	PARVALUE     string `xml:"PARVALUE"`
	DEBTTYPE     string `xml:"DEBTTYPE"`
	DEBTCLASS    string `xml:"DEBTCLASS"`
	COUPONRT     string `xml:"COUPONRT"`
	DTCOUPON     string `xml:"DTCOUPON"`
	COUPONFREQ   string `xml:"COUPONFREQ"`
	CALLPRICE    string `xml:"CALLPRICE"`
	YIELDTOCALL  string `xml:"YIELDTOCALL"`
	DTCALL       string `xml:"DTCALL"`
	CALLTYPE     string `xml:"CALLTYPE"`
	YIELDTOMAT   string `xml:"YIELDTOMAT"`
	DTMAT        string `xml:"DTMAT"`
	ASSETCLASS   string `xml:"ASSETCLASS"`
	FIASSETCLASS string `xml:"FIASSETCLASS"`
	*AbstractSecurityInfo
}

// EmptyType is The OFX element "EMPTY" is of type "EmptyType"
type EmptyType string

// FinancialInstitutionMutualFundAssetClass is The OFX element "FIMFASSETCLASS" is of type "FinancialInstitutionMutualFundAssetClass"
type FinancialInstitutionMutualFundAssetClass struct {
	FIPORTION []*FinancialInstitutionPortion `xml:"FIPORTION"`
}

// FinancialInstitutionPortion is The OFX element "FIPORTION" is of type "FinancialInstitutionPortion"
type FinancialInstitutionPortion struct {
	FIASSETCLASS string `xml:"FIASSETCLASS"`
	PERCENT      string `xml:"PERCENT"`
}

// MutualFundAssetClass is The OFX element "MFASSETCLASS" is of type "MutualFundAssetClass"
type MutualFundAssetClass struct {
	PORTION []*Portion `xml:"PORTION"`
}

// MutualFundInfo is The OFX element "MFINFO" is of type "MutualFundInfo"
type MutualFundInfo struct {
	MFTYPE         string                                    `xml:"MFTYPE"`
	YIELD          string                                    `xml:"YIELD"`
	DTYIELDASOF    string                                    `xml:"DTYIELDASOF"`
	MFASSETCLASS   *MutualFundAssetClass                     `xml:"MFASSETCLASS"`
	FIMFASSETCLASS *FinancialInstitutionMutualFundAssetClass `xml:"FIMFASSETCLASS"`
	*AbstractSecurityInfo
}

// MutualFundTypeEnum is The OFX element "MFTENUM" is of type "MutualFundTypeEnum"
type MutualFundTypeEnum string

// OptionInfo is The OFX element "OPTINFO" is of type "OptionInfo"
type OptionInfo struct {
	OPTTYPE      string      `xml:"OPTTYPE"`
	STRIKEPRICE  string      `xml:"STRIKEPRICE"`
	DTEXPIRE     string      `xml:"DTEXPIRE"`
	SHPERCTRCT   string      `xml:"SHPERCTRCT"`
	SECID        *SecurityId `xml:"SECID"`
	ASSETCLASS   string      `xml:"ASSETCLASS"`
	FIASSETCLASS string      `xml:"FIASSETCLASS"`
	*AbstractSecurityInfo
}

// OptionTypeEnum is The OFX element "OPTTYPEENUM" is of type "OptionTypeEnum"
type OptionTypeEnum string

// OtherInfo is The OFX element "OTHERINFO" is of type "OtherInfo"
type OtherInfo struct {
	TYPEDESC     string `xml:"TYPEDESC"`
	ASSETCLASS   string `xml:"ASSETCLASS"`
	FIASSETCLASS string `xml:"FIASSETCLASS"`
	*AbstractSecurityInfo
}

// Portion is The OFX element "PORTION" is of type "Portion"
type Portion struct {
	ASSETCLASS string `xml:"ASSETCLASS"`
	PERCENT    string `xml:"PERCENT"`
}

// Rating is The OFX element "RATING" is of type "Rating"
type Rating string

// GeneralSecurityInfo is The OFX element "SECINFO" is of type "GeneralSecurityInfo"
type GeneralSecurityInfo struct {
	SECID     *SecurityId `xml:"SECID"`
	SECNAME   string      `xml:"SECNAME"`
	TICKER    string      `xml:"TICKER"`
	FIID      string      `xml:"FIID"`
	RATING    string      `xml:"RATING"`
	UNITPRICE string      `xml:"UNITPRICE"`
	DTASOF    string      `xml:"DTASOF"`
	CURRENCY  *Currency   `xml:"CURRENCY"`
	MEMO      string      `xml:"MEMO"`
}

// AbstractSecurityInfo ...
type AbstractSecurityInfo struct {
	SECINFO *GeneralSecurityInfo `xml:"SECINFO"`
}

// SecurityList is The OFX element "SECLIST" is of type "SecurityList"
type SecurityList struct {
	MFINFO    []*MutualFundInfo `xml:"MFINFO"`
	STOCKINFO []*StockInfo      `xml:"STOCKINFO"`
	OPTINFO   []*OptionInfo     `xml:"OPTINFO"`
	DEBTINFO  []*DebtInfo       `xml:"DEBTINFO"`
	OTHERINFO []*OtherInfo      `xml:"OTHERINFO"`
}

// StockEnum is The OFX element "STOCKENUM" is of type "StockEnum"
type StockEnum string

// StockInfo is The OFX element "STOCKINFO" is of type "StockInfo"
type StockInfo struct {
	STOCKTYPE    string `xml:"STOCKTYPE"`
	YIELD        string `xml:"YIELD"`
	DTYIELDASOF  string `xml:"DTYIELDASOF"`
	ASSETCLASS   string `xml:"ASSETCLASS"`
	FIASSETCLASS string `xml:"FIASSETCLASS"`
	*AbstractSecurityInfo
}

// TickerType is The OFX element "TICKER" is of type "TickerType"
type TickerType string

// TypeDescriptionType is The OFX element "TYPEDESC" is of type "TypeDescriptionType"
type TypeDescriptionType string
