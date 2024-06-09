package game

// Color : Int constant mapping
const (
	Red          = 0
	Blue         = 1
	Green        = 2
	Pink         = 3
	Orange       = 4
	Yellow       = 5
	Black        = 6
	White        = 7
	Purple       = 8
	Brown        = 9
	Cyan         = 10
	Lime         = 11
	Maroon       = 12
	Rose         = 13
	Banana       = 14
	Gray         = 15
	Tan          = 16
	Coral        = 17
	Salmon       = 18
	Bordeaux     = 19
	Olive        = 20
	Turqoise     = 21
	Mint         = 22
	Lavender     = 23
	Nougat       = 24
	Peach        = 25
	Wasabi       = 26
	HotPink      = 27
	Petrol       = 28
	Lemon        = 29
	SignalOrange = 30
	Teal         = 31
	Blurple      = 32
	Sunrise      = 33
	Ice          = 34
)

// ColorStrings for lowercase, possibly for translation if needed
var ColorStrings = map[string]int{
	"red":          Red,
	"blue":         Blue,
	"green":        Green,
	"pink":         Pink,
	"orange":       Orange,
	"yellow":       Yellow,
	"black":        Black,
	"white":        White,
	"purple":       Purple,
	"brown":        Brown,
	"cyan":         Cyan,
	"lime":         Lime,
	"maroon":       Maroon,
	"rose":         Rose,
	"banana":       Banana,
	"gray":         Gray,
	"tan":          Tan,
	"coral":        Coral,
	"salmon":       Salmon,
	"bordeaux":     Bordeaux,
	"olive":        Olive,
	"turqoise":     Turqoise,
	"mint":         Mint,
	"lavender":     Lavender,
	"nougat":       Nougat,
	"peach":        Peach,
	"wasabi":       Wasabi,
	"hotpink":      HotPink,
	"petrol":       Petrol,
	"lemon":        Lemon,
	"signalorange": SignalOrange,
	"teal":         Teal,
	"blurple":      Blurple,
	"sunrise":      Sunrise,
	"ice":          Ice,
}

var ColorVanillaStrings = map[string]int{
	"red":    Red,
	"blue":   Blue,
	"green":  Green,
	"pink":   Pink,
	"orange": Orange,
	"yellow": Yellow,
	"black":  Black,
	"white":  White,
	"purple": Purple,
	"brown":  Brown,
	"cyan":   Cyan,
	"lime":   Lime,
	"maroon": Maroon,
	"rose":   Rose,
	"banana": Banana,
	"gray":   Gray,
	"tan":    Tan,
	"coral":  Coral,
}

var ColorTorStrings = map[string]int{
	"salmon":       Salmon,
	"bordeaux":     Bordeaux,
	"olive":        Olive,
	"turqoise":     Turqoise,
	"mint":         Mint,
	"lavender":     Lavender,
	"nougat":       Nougat,
	"peach":        Peach,
	"wasabi":       Wasabi,
	"hotpink":      HotPink,
	"petrol":       Petrol,
	"lemon":        Lemon,
	"signalorange": SignalOrange,
	"teal":         Teal,
	"blurple":      Blurple,
	"sunrise":      Sunrise,
	"ice":          Ice,
}

// GetColorStringForInt does what it sounds like
func GetColorStringForInt(colorint int) string {
	for str, idx := range ColorStrings {
		if idx == colorint {
			return str
		}
	}
	return ""
}

// IsColorString determines if a string is actually one of our colors
func IsColorString(test string) bool {
	_, ok := ColorStrings[test]
	return ok
}
