package geneticshuman

import (
	"fmt"
	"strings"

	"github.com/Flokey82/genetics"
)

// FiveFactor represents a simplified version of the five factor personality model.
type FiveFactor struct {
	Openness          int // openness to experience
	Conscientiousness int // conscientiousness
	Extraversion      int // extraversion
	Agreeableness     int // agreeableness
	Neuroticism       int // neuroticism
}

// String returns a string representation of the five factor personality.
func (s FiveFactor) String() string {
	return fmt.Sprintf(
		"O: %d, C: %d, E: %d, A: %d, N: %d",
		s.Openness, s.Conscientiousness, s.Extraversion, s.Agreeableness, s.Neuroticism,
	)
}

// GetFiveFactor returns the five factor personality of a person based on their genes.
func GetFiveFactor(g *genetics.Genes) FiveFactor {
	return FiveFactor{
		Openness:          g.Get(GOpenness),
		Conscientiousness: g.Get(GConscientiousness),
		Extraversion:      g.Get(GExtraversion),
		Agreeableness:     g.Get(GAgreeableness),
		Neuroticism:       g.Get(GNeuroticism),
	}
}

type Trait int

// The various traits of a person.
const (
	TraitDeceptive = 1 << iota
	TraitHonest
	TraitAggressive
	TraitCalm
	TraitCowardly
	TraitBrave
	TraitAmbitious
	TraitContent
	TraitCareless
	TraitCareful
	TraitParanoid
	TraitTrusting
	TraitCruel
	TraitKind
	TraitMax
)

var traitToString = map[Trait]string{
	TraitDeceptive:  "deceptive",
	TraitHonest:     "honest",
	TraitAggressive: "aggressive",
	TraitCalm:       "calm",
	TraitCowardly:   "cowardly",
	TraitBrave:      "brave",
	TraitAmbitious:  "ambitious",
	TraitContent:    "content",
	TraitCareless:   "careless",
	TraitCareful:    "careful",
	TraitParanoid:   "paranoid",
	TraitTrusting:   "trusting",
	TraitCruel:      "cruel",
	TraitKind:       "kind",
}

var traitToOpposite = map[Trait]Trait{
	TraitDeceptive:  TraitHonest,
	TraitHonest:     TraitDeceptive,
	TraitAggressive: TraitCalm,
	TraitCalm:       TraitAggressive,
	TraitCowardly:   TraitBrave,
	TraitBrave:      TraitCowardly,
	TraitAmbitious:  TraitContent,
	TraitContent:    TraitAmbitious,
	TraitCareless:   TraitCareful,
	TraitCareful:    TraitCareless,
	TraitParanoid:   TraitTrusting,
	TraitTrusting:   TraitParanoid,
	TraitCruel:      TraitKind,
	TraitKind:       TraitCruel,
}

func (t Trait) String() string {
	if tStr, ok := traitToString[t]; ok {
		return tStr
	}
	// concat all traits that are set.
	var tstrs []string
	for i := Trait(1); i < TraitMax; i <<= 1 {
		if t.HasTrait(i) {
			tstrs = append(tstrs, traitToString[i])
		}
	}
	return strings.Join(tstrs, ", ")
}

func (t Trait) HasTrait(trait Trait) bool {
	return t&trait != 0
}

func (t Trait) CountOpposites(other Trait) int {
	// Count how many opposite traits the other has.
	var count int
	for i := Trait(1); i < TraitMax; i <<= 1 {
		if t.HasTrait(i) && other.HasTrait(traitToOpposite[i]) {
			count++
		}
	}
	return count
}

func (t Trait) CountCommon(other Trait) int {
	// Count how many traits we have in common with the other.
	var count int
	for i := Trait(1); i < TraitMax; i <<= 1 {
		if t.HasTrait(i) && other.HasTrait(i) {
			count++
		}
	}
	return count
}

func (t Trait) Compare(other Trait) float64 {
	// Compare how many traits we have in common and how many traits we have that are opposite.
	// The more traits we have in common, the more we like each other.
	// The more traits we have that are opposite, the more we dislike each other.

	// Count how many traits we have in common.
	common := t.CountCommon(other)
	// Count how many opposite traits we have.
	opposites := t.CountOpposites(other)

	diff := common - opposites
	return float64(diff) / float64(max(common, opposites))
}

const (
	TraitMax  = 15
	TraitHigh = TraitMax / 3 * 2
	TraitMid  = TraitMax / 2
	TraitLow  = TraitMax / 3
)

// GetTraits returns the traits of a person based on their five factor personality.
func GetTraits(ff FiveFactor) Trait {
	var traits Trait

	// Each trait has some preconditions that need to be met for a chance of manifesting.
	// Depending on how much we exceed the threshold, the chance of the trait manifesting
	// increases.

	// Low agreeableness and low conscientiousness might lead to deception.
	if ff.Agreeableness < TraitLow && ff.Conscientiousness < TraitLow {
		traits |= TraitDeceptive
	} else if ff.Agreeableness > TraitHigh && ff.Conscientiousness > TraitHigh {
		traits |= TraitHonest
	}

	// Low agreeableness, high neuroticism and either high extraversion or high openness might lead to aggression.
	if ff.Agreeableness < TraitLow && ff.Neuroticism > TraitHigh && (ff.Extraversion > TraitHigh || ff.Openness > TraitHigh) {
		traits |= TraitAggressive
	} else if ff.Agreeableness > TraitMid && ff.Neuroticism < TraitLow {
		traits |= TraitCalm
	}

	// High extraversion and high agreeableness might lead to bravery.
	if ff.Extraversion > TraitHigh && ff.Agreeableness > TraitHigh {
		traits |= TraitBrave
	} else if (ff.Extraversion < TraitLow || ff.Openness < TraitLow) && ff.Agreeableness < TraitLow {
		traits |= TraitCowardly
	}

	// High conscientiousness and high extraversion might lead to ambition.
	if ff.Conscientiousness > TraitHigh && ff.Extraversion > TraitHigh {
		traits |= TraitAmbitious
	} else if ff.Conscientiousness < TraitLow && ff.Extraversion < TraitLow {
		traits |= TraitContent
	}

	// High conscientiousness and high neuroticism might lead to carefulness.
	if ff.Conscientiousness > TraitHigh && ff.Neuroticism > TraitHigh {
		traits |= TraitCareful
	} else if ff.Conscientiousness < TraitLow && ff.Neuroticism < TraitLow {
		traits |= TraitCareless
	}

	// High neuroticism, low agreeableness and high openness or high concienciousness might lead to paranoia,
	// while low neuroticism, high agreeableness and high openness or low conscientiousness might lead to trust.
	if ff.Neuroticism > TraitHigh && ff.Agreeableness < TraitLow && (ff.Openness > TraitHigh || ff.Conscientiousness > TraitHigh) {
		traits |= TraitParanoid
	} else if ff.Neuroticism < TraitLow && ff.Agreeableness > TraitHigh && (ff.Openness > TraitHigh || ff.Conscientiousness < TraitLow) {
		traits |= TraitTrusting
	}

	// Low agreeableness and high neuroticism or low conscientiousness might lead to cruelty.
	if ff.Agreeableness < TraitLow && (ff.Neuroticism > TraitHigh || ff.Conscientiousness < TraitLow) {
		traits |= TraitCruel
	} else if ff.Agreeableness > TraitHigh && (ff.Conscientiousness > TraitHigh || ff.Extraversion > TraitHigh) {
		traits |= TraitKind
	}

	return traits
}
