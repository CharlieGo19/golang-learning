package organisation

import "fmt"

// HomeWorld is a type alias, it is just that, a reference.
// It DOES SHARE method sets.
type BankStonks = int

// BankStonks is a type definition.
// It DOES NOT SHARE method sets.
// Therefore we can define our own methods on this.
// That's the difference between an alias and a type definition.
// It is still treated as an int.
// Therefore a type definition copies the FIELDS of an object...
// ...And an alias copies the FIELDS AND METHODS.
type Identity string

// SuperNaturalEntity will be implemented by Villian / AntiHero
type SuperNaturalEntity interface {
	ReturnFullName() string
	ReturnFullAlias() string
}

type Identifiable interface {
	Bumhole() string
}

type Name struct {
	firstName  string
	secondName string
}

// Villian has CatchPhrase to distinguish itself from AntiHero
// Can embed interface in structs & interfaces!
type Villian struct {
	Name
	alias       string
	CatchPhrase string
	Identifiable
}

// AntiHero has BenisSize to distinguish itself from Villian
type AntiHero struct {
	Name
	alias     string
	BenisSize int
}

func (n Name) ReturnFullName() string {
	return fmt.Sprintf("%s %s", n.firstName, n.secondName)
}

func (v Villian) ReturnFullAlias() string {
	return fmt.Sprintf("%s\n", v.alias)
}

func (a AntiHero) ReturnFullAlias() string {
	return fmt.Sprintf("%s\n", a.alias)
}

func NewEntity(first, second, entity string) SuperNaturalEntity {
	if entity == "villian" {
		return Villian{
			Name: Name{
				firstName:  first,
				secondName: second,
			},
		}
	}
	return AntiHero{
		Name: Name{
			firstName:  first,
			secondName: second,
		},
	}
}
