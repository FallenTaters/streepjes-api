package shared

type Club int

const (
	ClubUnknown    Club = iota // Unknown
	ClubParabool               // Parabool
	ClubGladiators             // Gladiators
)

//go:generate enumer -json -linecomment -type Club
