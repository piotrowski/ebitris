package music

import "github.com/piotrowski/ebitris/assets"

type (
	SongName   int
	EffectName int
)

const (
	ReturnOfThe8BitEra SongName = iota + 1
	ArcadeBeat
)

const (
	ExplosionEffect EffectName = iota + 1
)

var songs map[SongName][]byte = map[SongName][]byte{
	ReturnOfThe8BitEra: assets.ReturnOnThe8BitEra,
	ArcadeBeat:         assets.ArcadeBeat,
}

var effects map[EffectName][]byte = map[EffectName][]byte{
	ExplosionEffect: assets.ExplosionEffect,
}
