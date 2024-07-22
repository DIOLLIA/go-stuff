package main

type Player struct {
	dress      Dress
	playerType string // Can be T or CT
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getSingleDressInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
