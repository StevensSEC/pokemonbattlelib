package pokemonbattlelib

var battleItemStats = map[Item]int{
	ItemXAccuracy: StatAccuracy,
	ItemXAttack:   StatAtk,
	ItemXDefense:  StatDef,
	ItemXSpAtk:    StatSpAtk,
	ItemXSpDef:    StatSpAtk,
	ItemXSpeed:    StatSpeed,
}

//go:generate go run data/gen.go
