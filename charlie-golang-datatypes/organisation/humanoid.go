package organisation

import "fmt"

type ID int

type EuropeanHuman struct {
	humanName
	europeanUnionIdentifier
}

type CanzukHuman struct {
	humanName
	canzukIdentifier
}

type HillBilly struct {
	humanName
	unitedStatesIdentifier
}

type Humanoid interface {

	// Name
	HumanName
	// ID No. (how the gov. track you lul)
	IdentifiableCitizen
}

type IdentifiableCitizen interface {
	ReturnID() ID
	ReturnCountry() string
}

type HumanName interface {
	ReturnName() string
}

type humanName struct {
	firstName  string
	secondName string
}

type europeanUnionIdentifier struct {
	id            ID
	countryOrigin string
}

type canzukIdentifier struct {
	id            ID
	nhsID         int
	countryOrigin string
}

type unitedStatesIdentifier struct {
	id                ID
	numberOfGunsOwned int
	weight            int
	countryOrigin     string
}

func (eui europeanUnionIdentifier) ReturnID() ID {
	return eui.id
}

func (eui europeanUnionIdentifier) ReturnCountry() string {
	return eui.countryOrigin
}

func (ci canzukIdentifier) ReturnID() ID {
	return ci.id
}

func (ci canzukIdentifier) ReturnCountry() string {
	return ci.countryOrigin
}

func (usi unitedStatesIdentifier) ReturnID() ID {
	return usi.id
}

func (usi unitedStatesIdentifier) ReturnCountry() string {
	return usi.countryOrigin
}

func (hn humanName) ReturnName() string {
	return fmt.Sprintf("%s %s", hn.firstName, hn.secondName)
}

func (ci *CanzukHuman) SetNHSid(id int) {
	ci.nhsID = id
}

func (ci canzukIdentifier) GetNHSid() int {
	return ci.nhsID
}

func NewHuman(id ID, locale, country, firstName, secondName string) Humanoid {
	if locale == "eu" {
		return EuropeanHuman{
			humanName: humanName{
				firstName:  firstName,
				secondName: secondName,
			},
			europeanUnionIdentifier: europeanUnionIdentifier{
				id:            id,
				countryOrigin: country,
			},
		}
	} else if locale == "canzuk" {
		return CanzukHuman{
			humanName: humanName{
				firstName:  firstName,
				secondName: secondName,
			},
			canzukIdentifier: canzukIdentifier{
				id:            id,
				countryOrigin: country,
			},
		}
	}
	return HillBilly{
		humanName: humanName{
			firstName:  firstName,
			secondName: secondName,
		},
		unitedStatesIdentifier: unitedStatesIdentifier{
			id:            id,
			countryOrigin: country,
		},
	}
}
