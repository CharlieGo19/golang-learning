package main

import (
	"fmt"
	"os"

	"charlie.learning.datatypes/custom/organisation"
)

func main() {

	var geralt organisation.SuperHero = organisation.InitHero("Geralt", "of Rivia", 95)
	var mattMurdock organisation.SuperHero = organisation.InitHero("Matt", "Murdock", 27)
	// This is a villian, but to demonstrate it has implemented the interface 'SuperNaturalEntity' I have set its type as such.
	var jackNapier organisation.SuperNaturalEntity = organisation.NewEntity("Jack", "Napier", "villian")
	var frankCastle organisation.SuperNaturalEntity = organisation.NewEntity("Frank", "Castle", "antihero")

	err := geralt.SetAlias("The Witcher")
	checkIfNil(err)
	err = geralt.SetInstaHandle("@henrycavill")
	checkIfNil(err)
	err = mattMurdock.SetAlias("Daredevil")
	checkIfNil(err)
	err = mattMurdock.SetInstaHandle("@_charliecox")
	checkIfNil(err)
	// Note: I did extra error stuff for demo, I don't think you'd normally do this for getters.
	gName, err := geralt.GetHeroName()
	checkIfNil(err)
	gAlias, err := geralt.GetAlias()
	checkIfNil(err)
	gAge, err := geralt.GetHeroAge()
	checkIfNil(err)
	gInsta, err := geralt.GetInstaHandle()
	checkIfNil(err)
	jackNapier.ReturnFullAlias()
	// We do some wizardy here, ignore it for now, it's known as type casting if you're interest.
	// This is so we can access the BenisSize.
	// This extremely long winded, but it's to demonstrate embedded structs & the 'implicit' interface application (I don't like it).
	// This is to distinguish frankCastle from jackNapier, although they're both at this point of type SuperNaturalEntity...
	// ... Only Frank can become an antihero and Jack a villian, because of BenisSize and CatchPhrase respectively.
	// If you tried to cast, because the memory layout is different, it woudldn't copy over the existing data to the new type and it would...
	// return false; however you would be able to reassign data to the 'newJack' if you tried what I did with frank, which is bizzare behaviour.
	newFrank, _ := frankCastle.(organisation.AntiHero)
	newFrank.BenisSize = 100
	fmt.Printf("%s has a benis size of %d inches! Hot damn, how does that boy walk.\n", newFrank.ReturnFullName(), newFrank.BenisSize)
	fmt.Printf("The character %s aka %s is %d years old. The actors insta can be found at: %s.\n", *gName, *gAlias, *gAge, *gInsta)

}

func checkIfNil(err error) {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}
