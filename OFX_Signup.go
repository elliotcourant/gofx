// Code generated by xgen. DO NOT EDIT.

package gofx

// AbstractServiceAction ...
type AbstractServiceAction struct {
}

// ServiceEnum is The OFX element "SERVICEENUM" is of type "ServiceEnum"
type ServiceEnum string

// ServiceAdd is The OFX element "SVCADD" is of type "ServiceAdd"
type ServiceAdd struct {
	BANKACCTTO *BankAccount        `xml:"BANKACCTTO"`
	CCACCTTO   *CreditCardAccount  `xml:"CCACCTTO"`
	INVACCTTO  *InvestmentAccount  `xml:"INVACCTTO"`
	PRESACCTTO *PresentmentAccount `xml:"PRESACCTTO"`
	*AbstractServiceAction
}

// ServiceChange is The OFX element "SVCCHG" is of type "ServiceChange"
type ServiceChange struct {
	BANKACCTFROM *BankAccount        `xml:"BANKACCTFROM"`
	CCACCTFROM   *CreditCardAccount  `xml:"CCACCTFROM"`
	INVACCTFROM  *InvestmentAccount  `xml:"INVACCTFROM"`
	PRESACCTFROM *PresentmentAccount `xml:"PRESACCTFROM"`
	BANKACCTTO   *BankAccount        `xml:"BANKACCTTO"`
	CCACCTTO     *CreditCardAccount  `xml:"CCACCTTO"`
	INVACCTTO    *InvestmentAccount  `xml:"INVACCTTO"`
	PRESACCTTO   *PresentmentAccount `xml:"PRESACCTTO"`
	*AbstractServiceAction
}

// ServiceDelete is The OFX element "SVCDEL" is of type "ServiceDelete"
type ServiceDelete struct {
	BANKACCTFROM *BankAccount        `xml:"BANKACCTFROM"`
	CCACCTFROM   *CreditCardAccount  `xml:"CCACCTFROM"`
	INVACCTFROM  *InvestmentAccount  `xml:"INVACCTFROM"`
	PRESACCTFROM *PresentmentAccount `xml:"PRESACCTFROM"`
	*AbstractServiceAction
}

// TempPassType is The OFX element "TEMPPASS" is of type "TempPassType"
type TempPassType string
