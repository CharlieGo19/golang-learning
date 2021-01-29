package organisation

import (
	"errors"
	"strings"
)

// A Hero is what? Very subjective topic...
type Hero interface {
	InitHero(firstName string, secondName string, alias string, age int) SuperHero
	GetHeroName() (string, error)
	GetHeroAge() (int, error)
	SetAlias(alias string)
	GetAlias() (string, error)
	SetInstaHandle(instaHandle string)
	GetInstaHandle() string
}

// SuperHero is not going to be commented, haha golang, f u and your aggrisive best practise enforcement!
type SuperHero struct {
	firstName   string
	secondName  string
	alias       string
	instaHandle string
	age         int
}

func InitHero(firstName string, secondName string, age int) SuperHero {
	return SuperHero{
		firstName:  firstName,
		secondName: secondName,
		age:        age,
	}
}

func (hero *SuperHero) GetHeroName() (*string, error) {
	var returnName string
	if len(hero.firstName) != 0 && len(hero.secondName) != 0 {
		returnName = hero.firstName + " " + hero.secondName
		return &returnName, nil
	}
	return nil, errors.New("the name is empty (first or second), nothing to return")
}

func (hero *SuperHero) GetHeroAge() (*int, error) {
	if hero.age > 0 {
		return &hero.age, nil
	}
	return nil, errors.New("age either doesn't exist or is unsuitable")
}

// SetAlias sets alias of SuperHero
// Note: we use pointer based recieve - because golang creates a copy with each change.
// Therefore, if we did not do this, the changes would not take effect.
func (hero *SuperHero) SetAlias(alias string) error {
	if len(alias) != 0 {
		hero.alias = alias
		return nil
	}
	return errors.New("alias cannot be an empty string")

}

// GetAlias - just to demonstrate, here we do not use a pointer.
// As this is a read only function, no changes are made, therefore no pointer is required.
// However it is good practise to be consistent and use pointers across the board.
func (hero SuperHero) GetAlias() (*string, error) {
	if len(hero.alias) != 0 {
		return &hero.alias, nil
	}
	return nil, errors.New("alias is empty, nothing to return")
}

func (hero *SuperHero) SetInstaHandle(instaHandle string) error {
	if len(instaHandle) == 0 {
		return errors.New("instagram handle cannot be an empty string")
	} else if !strings.HasPrefix(instaHandle, "@") {
		return errors.New("instagram handle must start with an @ symbol")
	}

	hero.instaHandle = instaHandle
	return nil
}

func (hero *SuperHero) GetInstaHandle() (*string, error) {
	if len(hero.instaHandle) != 0 {
		return &hero.instaHandle, nil
	}
	return nil, errors.New("instagram handle is empty, nothing to return")
}
