package player

// Avatar is a link to an image
type Avatar struct {
	URL string
}

// Player is a player for the game
type Player struct {
	Name     string
	Age      int
	Avatar   Avatar
	password string
}

// New create a new Player
func New(name string) Player {
	return Player{
		Name: name,
	}
}

//UpdateAvatar update the value of URL
func (a *Avatar) UpdateAvatar(url string) {
	a.URL = url
}
