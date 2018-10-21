package randomiser

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const nGauss = 1000

var gaussianVars [nGauss]float64

func init() {
	fillGaussianVars()
}

func SeedAndTell(s int64) {
	if s < 0 {
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
	fmt.Fprintf(os.Stderr, "Seeder initialised with value %v", s)
}

func ThrowWithProbability(p float64) bool {
	return rand.Float64() <= p
}

func AveragedRandom(average int, maxDeviation int) int {
	randGaussian := gaussianVars[rand.Intn(len(gaussianVars))]
	deviation := float64(maxDeviation) * randGaussian

	return average + int(deviation)
}

func AveragedRandomPartDev(average int, divider int) int {
	return AveragedRandom(average, average/divider)
}

func fillGaussianVars() {
	var (
		s  float64
		v1 float64
		v2 float64
		x  float64
	)

	for i := 0; i < len(gaussianVars); i++ {
		s = 0.
		v1 = 0.
		v2 = 0.
		x = 0.

		for {
			for {
				u1 := rand.Float64()
				u2 := rand.Float64()

				v1 = 2*u1 - 1
				v2 = 2*u2 - 1

				s = v1*v1 + v2*v2

				if s < 1.0 {
					break
				}
			}

			x = v1 * math.Sqrt(-2*math.Log(s)/s)

			if x > -1. && x < 1. {
				break
			}
		}

		gaussianVars[i] = x
	}
}
