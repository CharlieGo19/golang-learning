package models

// Hero is a big old poo
type Hero struct {
	ID         int
	HeroName   string
	HeroAlias  string
	HeroRating int
}

var (
	heroes []*Hero // we use pointer, so we can manipulate this data directly without having to copy the actual user.
	nextID int     = 1
)

func GetHeroes() []*Hero {
	return heroes
}

// AddHero there is a paradigm in Go, where you always return an error with something that could go wrong.
func AddHero(hero Hero) (Hero, error) {
	hero.ID = nextID
	nextID++
	heroes = append(heroes, &hero)
	return hero, nil
}
