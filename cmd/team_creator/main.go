package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ilmaruk/gofms/models"
	"github.com/ilmaruk/gofms/randomiser"
	"math/rand"
	"os"
	"strings"
)

const defaultGoalkeepers = 3
const defaultDefenders = 8
const defaultDefensiveMidfielders = 3
const defaultMidfielders = 8
const defaultAttackingMidfielders = 3
const defaultAttackers = 5

const defaultAverageMainSkill = 14
const defaultAverageMidSkill = 11
const defaultAverageSecondarySkill = 7

const defaultAverageStamina = 60
const defaultAverageAggression = 30

type config struct {
	numGoalkeepers          int
	numDefenders            int
	numDefensiveMidfielders int
	numMidfielders          int
	numAttackingMidfielders int
	numAttackers            int
	averageStamina          int
	averageAggression       int
	averageMainSkill        int
	averageMidSkill         int
	averageSecondarySkill   int
}

func main() {
	seed := flag.Int64("seed", -1, "Randomiser seed")

	numGoalkeepers := flag.Int("numGk", defaultGoalkeepers, "Number of goalkeepers")
	numDefenders := flag.Int("numDf", defaultDefenders, "Number of defenders")
	numDefensiveMidfielders := flag.Int("numDm", defaultDefensiveMidfielders, "Number of defensive midfielders")
	numMidfielders := flag.Int("numMf", defaultMidfielders, "Number of midfielders")
	numAttackingMidfielders := flag.Int("numAm", defaultAttackingMidfielders, "Number of attacking midfielders")
	numAttackers := flag.Int("numAt", defaultAttackers, "Number of attackers")

	avgStamina := flag.Int("avgStamina", defaultAverageStamina, "Average stamina")
	avgAggression := flag.Int("avgAggression", defaultAverageAggression, "Average aggression")

	avgMainSkill := flag.Int("avgMainSkill", defaultAverageMainSkill, "Average main skill")
	avgMidSkill := flag.Int("avgMidSkill", defaultAverageMidSkill, "Average mid skill")
	avgSecondarySkill := flag.Int("avgSecondarySkill", defaultAverageSecondarySkill, "Average secondary skill")
	flag.Parse()

	randomiser.SeedAndTell(*seed)

	cfg := config{
		numGoalkeepers:          *numGoalkeepers,
		numDefenders:            *numDefenders,
		numDefensiveMidfielders: *numDefensiveMidfielders,
		numMidfielders:          *numMidfielders,
		numAttackingMidfielders: *numAttackingMidfielders,
		numAttackers:            *numAttackers,
		averageStamina:          *avgStamina,
		averageAggression:       *avgAggression,
		averageMainSkill:        *avgMainSkill,
		averageMidSkill:         *avgMidSkill,
		averageSecondarySkill:   *avgSecondarySkill,
	}

	team := generateTeam(cfg)

	b, err := json.MarshalIndent(team, "", "  ")
	if err != nil {
		fmt.Printf("ERROR: %d\n", err)
		os.Exit(1)
	}

	fmt.Println(string(b))
}

func generateTeam(cfg config) models.Team {
	team := models.Team{Name: "Foo FC"}

	team.Players = append(team.Players, generatePlayers(cfg.numGoalkeepers, generateGoalkeeper, cfg)...)
	team.Players = append(team.Players, generatePlayers(cfg.numDefenders, generateDefender, cfg)...)
	team.Players = append(team.Players, generatePlayers(cfg.numDefensiveMidfielders, generateDefensiveMidfielder, cfg)...)
	team.Players = append(team.Players, generatePlayers(cfg.numMidfielders, generateMidfielder, cfg)...)
	team.Players = append(team.Players, generatePlayers(cfg.numAttackingMidfielders, generateAttackingMidfielder, cfg)...)
	team.Players = append(team.Players, generatePlayers(cfg.numAttackers, generateAttacker, cfg)...)

	return team
}

func generateRandomName() string {
	vowelish := []string{"a", "o", "e", "i", "u"}
	vowelishNotAtBeginning := []string{"ew", "ow", "oo", "oa", "oi", "oe", "ae", "ua"}
	consonantish := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "y",
		"z", "br", "cl", "gr", "st", "jh", "tr", "ty", "dr", "kr", "ry", "bt", "sh", "ch", "pr"}
	consonantishNotAtBeginning := []string{"mn", "nh", "rt", "rs", "rst", "dn", "nd", "ds", "bt", "bs", "bl", "sk",
		"vr", "ks", "sy", "ny", "vr", "sht", "ck"}

	firstNameAbbreviation := string(rune(int('A') + rand.Intn(25)))

	lastWasVowel := false
	name := firstNameAbbreviation + " "

	if randomiser.ThrowWithProbability(50.) {
		lastWasVowel = true
		name += strings.ToUpper(randomStringFromList(vowelish))
	} else {
		lastWasVowel = false
		name += strings.ToUpper(randomStringFromList(consonantish))
	}

	howMany := 3 + rand.Intn(3)
	for i := 0; i < howMany; i++ {
		if lastWasVowel {
			if randomiser.ThrowWithProbability(50.) {
				name += randomStringFromList(consonantish)
			} else {
				name += randomStringFromList(consonantishNotAtBeginning)
			}
		} else {
			if randomiser.ThrowWithProbability(75.) {
				name += randomStringFromList(vowelish)
			} else {
				name += randomStringFromList(vowelishNotAtBeginning)
			}
		}

		lastWasVowel = !lastWasVowel
	}

	// TODO: Eventually, capitalize the first letter of the surname

	return name
}

func randomStringFromList(l []string) string {
	return l[rand.Intn(len(l))]
}

func preferredSide() string {
	random := rand.Intn(150)

	if random <= 8 {
		return "RLC"
	} else if random <= 13 {
		return "RL"
	} else if random <= 23 {
		return "RC"
	} else if random <= 33 {
		return "LC"
	} else if random <= 73 {
		return "R"
	} else if random <= 13 {
		return "L"
	} else {
		return "C"
	}
}

func nationality() string {
	nationalities := []string{"arg", "aus", "bra", "bul", "cam", "cro", "den", "eng", "fra", "ger", "hol", "ire",
		"isr", "ita", "jap", "nig", "nor", "saf", "spa", "usa"}
	return randomStringFromList(nationalities)
}

type playerGenerator func(cfg config) models.Player

func generatePlayers(amount int, strategy playerGenerator, cfg config) []models.Player {
	players := make([]models.Player, amount)
	for i := 0; i < amount; i++ {
		players[i] = strategy(cfg)
	}
	return players
}

func generateGenericPlayer(cfg config) models.Player {
	player := models.NewPlayer()
	player.Name = generateRandomName()
	player.Age = randomiser.AveragedRandom(23, 7)
	player.PreferredSide = preferredSide()
	player.Nationality = nationality()
	player.Stamina = randomiser.AveragedRandomPartDev(cfg.averageStamina, 2)
	player.Aggression = randomiser.AveragedRandomPartDev(cfg.averageAggression, 3)
	// TODO: generate stuff

	return player
}

func generateGoalkeeper(cfg config) models.Player {
	player := generateGenericPlayer(cfg)
	player.Stopping = randomiser.AveragedRandomPartDev(cfg.averageMainSkill, 3)
	player.Tackling = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill/2, 2)
	player.Passing = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill/2, 2)
	player.Shooting = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill/2, 2)
	return player
}

func generateDefender(cfg config) models.Player {
	player := generateGenericPlayer(cfg)
	player.Tackling = randomiser.AveragedRandomPartDev(cfg.averageMainSkill, 3)
	player.Stopping = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill/2, 2)
	player.Passing = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill, 2)
	player.Shooting = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill, 2)
	return player
}

func generateDefensiveMidfielder(cfg config) models.Player {
	player := generateGenericPlayer(cfg)
	player.Passing = randomiser.AveragedRandomPartDev(cfg.averageMidSkill, 3)
	player.Tackling = randomiser.AveragedRandomPartDev(cfg.averageMidSkill, 3)
	player.Stopping = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill/2, 2)
	player.Shooting = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill, 2)
	return player
}

func generateMidfielder(cfg config) models.Player {
	player := generateGenericPlayer(cfg)
	player.Passing = randomiser.AveragedRandomPartDev(cfg.averageMainSkill, 3)
	player.Stopping = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill/2, 2)
	player.Tackling = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill, 2)
	player.Shooting = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill, 2)
	return player
}

func generateAttackingMidfielder(cfg config) models.Player {
	player := generateGenericPlayer(cfg)
	player.Passing = randomiser.AveragedRandomPartDev(cfg.averageMidSkill, 3)
	player.Shooting = randomiser.AveragedRandomPartDev(cfg.averageMidSkill, 3)
	player.Tackling = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill, 2)
	player.Stopping = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill/2, 2)
	return player
}

func generateAttacker(cfg config) models.Player {
	player := generateGenericPlayer(cfg)
	player.Shooting = randomiser.AveragedRandomPartDev(cfg.averageMainSkill, 3)
	player.Stopping = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill/2, 2)
	player.Tackling = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill, 2)
	player.Passing = randomiser.AveragedRandomPartDev(cfg.averageSecondarySkill, 2)
	return player
}
