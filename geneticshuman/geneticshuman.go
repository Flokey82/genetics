package geneticshuman

import (
	"fmt"
	"strings"

	"github.com/Flokey82/genetics"
)

// Proposed gene layout
//
//  _______________________ 2 gender
// || _____________________ 2 eye color
// ||||  __________________ 3 hair color         ___________________________ 4 Openness
// |||| ||| _______________ 4 complexion        ||||  ______________________ 4 Conscientiousness
// |||| |||| ||| __________ 3 height            |||| |||| __________________ 4 Extraversion
// |||| |||| |||| || ______ 3 mass              |||| |||| ||||  ____________ 4 Agreeableness
// |||| |||| |||| |||| | __ 3 growth            |||| |||| |||| ||||  _______ 4 Neuroticism
// |||| |||| |||| |||| ||||                     |||| |||| |||| |||| ||||
// xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx
//                          |||| |||| |||| ||||                          |||| ||||
// 4 strength _________________  |||| |||| ||||                           ________ unused
// 4 intelligence __________________  |||| ||||
// 4 dexterity __________________________  ||||
// 4 resilience ______________________________

var (
	GGender = genetics.Gene{
		NumBits: 2,
		Offset:  62,
	}
	GEyeColor = genetics.Gene{
		NumBits: 2,
		Offset:  60,
	}
	GHairColor = genetics.Gene{
		NumBits: 3,
		Offset:  57,
	}
	GComplexion = genetics.Gene{
		NumBits: 4,
		Offset:  53,
	}
	GHeight = genetics.Gene{
		NumBits: 3,
		Offset:  50,
	}
	GMass = genetics.Gene{
		NumBits: 3,
		Offset:  47,
	}
	GGrowth = genetics.Gene{
		NumBits: 3,
		Offset:  44,
	}
	GStrength = genetics.Gene{
		NumBits: 4,
		Offset:  40,
	}
	GIntelligence = genetics.Gene{
		NumBits: 4,
		Offset:  36,
	}
	GDexterity = genetics.Gene{
		NumBits: 4,
		Offset:  32,
	}
	GResilience = genetics.Gene{
		NumBits: 4,
		Offset:  28,
	}
	GOpenness = genetics.Gene{
		NumBits: 4,
		Offset:  24,
	}
	GConscientiousness = genetics.Gene{
		NumBits: 4,
		Offset:  20,
	}
	GExtraversion = genetics.Gene{
		NumBits: 4,
		Offset:  16,
	}
	GAgreeableness = genetics.Gene{
		NumBits: 4,
		Offset:  12,
	}
	GNeuroticism = genetics.Gene{
		NumBits: 4,
		Offset:  8,
	}
)

type Gender int

const (
	GenderMale   Gender = 0x1
	GenderFemale Gender = 0x3
)

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "  male"
	case GenderFemale:
		return "female"
	default:
		return "x"
	}
}

func SetGender(g *genetics.Genes, val Gender) {
	g.Set(GGender, int(val))
}

func GetGender(g *genetics.Genes) Gender {
	return Gender(g.Get(GGender))
}

type EyeColor int

const (
	EyeColorRed   EyeColor = 0x0
	EyeColorBlue  EyeColor = 0x1
	EyeColorGreen EyeColor = 0x2
	EyeColorBrown EyeColor = 0x3
)

func (g EyeColor) String() string {
	switch g {
	case EyeColorRed:
		return "  red"
	case EyeColorBlue:
		return " blue"
	case EyeColorGreen:
		return "green"
	case EyeColorBrown:
		return "brown"
	default:
		return "x"
	}
}

func SetEyeColor(g *genetics.Genes, val EyeColor) {
	g.Set(GEyeColor, int(val))
}

func GetEyeColor(g *genetics.Genes) EyeColor {
	return EyeColor(g.Get(GEyeColor))
}

type HairColor int

const (
	HairColorBlonde   HairColor = 0x0
	HairColorRed      HairColor = 0x1
	HairColorBrown    HairColor = 0x2
	HairColorBlack    HairColor = 0x3
	HairColorCurlMask HairColor = 0x4
)

func SetHairColor(g *genetics.Genes, val HairColor, curls bool) {
	if curls {
		val &= HairColorCurlMask
	}
	g.Set(GHairColor, int(val))
}

func GetHairColor(g *genetics.Genes) (HairColor, bool) {
	c := HairColor(g.Get(GHairColor))
	return c & (HairColorCurlMask - 1), c&HairColorCurlMask != 0
}

func GetHairColorStr(g *genetics.Genes) string {
	b, curl := GetHairColor(g)
	prfx := "      "
	if curl {
		prfx = "curly "
	}
	switch b {
	case HairColorBlonde:
		return prfx + "blonde"
	case HairColorRed:
		return prfx + "   red"
	case HairColorBrown:
		return prfx + " brown"
	case HairColorBlack:
		return prfx + " black"
	default:
		return prfx + "x"
	}
}

type Attrs struct {
	Complexion int
	Height     int
	Mass       int
	Growth     int
}

func (s Attrs) String() string {
	return fmt.Sprintf(
		"CMPLX: %d, HEIGH: %d, MASS: %d, GROW: %d",
		s.Complexion, s.Height, s.Mass, s.Growth,
	)
}

func GetAttrs(g *genetics.Genes) Attrs {
	return Attrs{
		Complexion: g.Get(GComplexion),
		Height:     g.Get(GHeight),
		Mass:       g.Get(GMass),
		Growth:     g.Get(GGrowth),
	}
}

type Stats struct {
	Strength     int
	Intelligence int
	Dexterity    int
	Resilience   int
}

func (s Stats) String() string {
	return fmt.Sprintf(
		"Str: %d, Int: %d, Dex: %d, Res: %d",
		s.Strength, s.Intelligence, s.Dexterity, s.Resilience,
	)
}

func GetStats(g *genetics.Genes) Stats {
	return Stats{
		Strength:     g.Get(GStrength),
		Intelligence: g.Get(GIntelligence),
		Dexterity:    g.Get(GDexterity),
		Resilience:   g.Get(GResilience),
	}
}

type FiveFactor struct {
	Openness          int
	Conscientiousness int
	Extraversion      int
	Agreeableness     int
	Neuroticism       int
}

func (s FiveFactor) String() string {
	return fmt.Sprintf(
		"O: %d, C: %d, E: %d, A: %d, N: %d",
		s.Openness, s.Conscientiousness, s.Extraversion, s.Agreeableness, s.Neuroticism,
	)
}

func GetFiveFactor(g *genetics.Genes) FiveFactor {
	return FiveFactor{
		Openness:          g.Get(GOpenness),
		Conscientiousness: g.Get(GConscientiousness),
		Extraversion:      g.Get(GExtraversion),
		Agreeableness:     g.Get(GAgreeableness),
		Neuroticism:       g.Get(GNeuroticism),
	}
}

func String(g genetics.Genes) string {
	var strs []string
	strs = append(strs, fmt.Sprintf("gender: %s", GetGender(&g)))
	strs = append(strs, fmt.Sprintf("Eyes: %s", GetEyeColor(&g)))
	strs = append(strs, fmt.Sprintf("Hair: %s", GetHairColorStr(&g)))
	strs = append(strs, fmt.Sprintf("Attrs: %s", GetAttrs(&g).String()))
	strs = append(strs, fmt.Sprintf("Stats: %s", GetStats(&g).String()))
	strs = append(strs, fmt.Sprintf("FiveFactor: %s", GetFiveFactor(&g).String()))
	return strings.Join(strs, ", ")
}
