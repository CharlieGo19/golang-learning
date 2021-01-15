package heroes

// Hero is a big old poo
type Hero struct {
	ID         int
	HeroName   string
	HeroAlias  string
	HeroRating int
}

var (
	heroez []*Hero // we use pointer, so we can manipulate this data directly without having to copy the actual user.
	nextID int     = 1
)
