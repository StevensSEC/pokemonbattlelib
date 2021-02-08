package pokemonbattlelib

import "fmt" 
type ElementalType uint32
var effect int

const (
	Normal ElementalType = 65536 >> iota
	Fight //66048 100000000000000000
	Flying  
	Poison  
	Ground  
	Rock  
	Bug
	Ghost
	Steel
	Fire
	Water
	Grass
	Electric
	Psychic
	Ice
	Dragon
	Dark
)

func contains(arr []uint32, str uint32) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}
 
func getEffect(move, def ElementalType) int {
    var match =uint32(move| def)
    fmt.Println("match spot: ", uint32(move| def))
    fmt.Printf("%b", int32(132096))
    noEffect := []uint32{132096,66560,16896,40960,133120,8208,9}
    halfEffect := []uint32{135168,131584,0,0,0} //need more num
    doubleEffect := []uint32{0,0,0} //need more num

    if contains(noEffect,uint32(132096))==true{
        effect=0
        fmt.Println("no effect")
    }
    if contains(halfEffect,match)==true{
        effect=50
        fmt.Println("half effect")
    }
    if contains(doubleEffect,match)==true{
        effect=150
        fmt.Println("double effect")
    }
    // else{
    //     effect=100
    //     fmt.Println("normal effect")
    // }
    return effect
}

