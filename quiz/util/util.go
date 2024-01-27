package util

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// todo: refactor this mess

func GetAnswerInt(question string) int {
	fmt.Print(question)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0
	}
	input = strings.TrimSuffix(input, "\n")

	integer, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return integer
}

func EvaluateSolutionInt(expectedAnswer int, actualAnswer int, question string) bool {

	success := true

	for expectedAnswer != actualAnswer {
		success = false
		fmt.Println("Falsch! Versuche es erneut.")
		actualAnswer = GetAnswerInt(question)
	}

	fmt.Println("Richtig!")

	return success
}

func GetRandomWord(filterFunc func(string) bool) string {
	rand.Seed(time.Now().Unix())
	var words = WordList.Filter(filterFunc)
	randomIndex := rand.Int() % len(words)
	return words[randomIndex]
}

func GetAnswerString(question string, hideQuestionFunc func()) string {
	fmt.Print(question)

	hideQuestionFunc()

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSuffix(input, "\n")
}

func EvaluateSolutionString(expectedAnswer string, actualAnswer string, question string, hideQuestionFunc func()) bool {
	success := true

	for expectedAnswer != actualAnswer {
		success = false
		fmt.Println("Falsch! Versuche es erneut.")
		actualAnswer = GetAnswerString(question, hideQuestionFunc)
	}

	fmt.Println("Richtig!")

	return success
}

func FilterWordsFunc(filter string) func(string) bool {
	return func(a string) bool {
		if filter == "" {
			return true
		}

		for _, f := range strings.Split(filter, ",") {
			if strings.Contains(a, f) {
				return true
			}
		}

		return false
	}
}

type words []string

func (ws words) Contains(search string) bool {
	for _, w := range ws {
		if w == search {
			return true
		}
	}

	return false
}

func (ws words) Filter(f func(string) bool) words {
	var returnWords words
	for _, w := range ws {
		if f(w) {
			returnWords = append(returnWords, w)
		}
	}
	return returnWords
}

// todo: move word list to config file

var WordList = words{
	"Apfel",
	"Birne",
	"Daumen",
	"Elefant",
	"Flöte",
	"Mama",
	"Oma",
	"Opa",
	"Papa",
	"Lama",
	"Limo",
	"Esel",
	"Amsel",
	"Salami",
	"See",
	"Fee",
	"Fluss",
	"Affe",
	"Fels",
	"Feld",
	"Elf",
	"Elb",
	"Museum",
	"Sofa",
	"Elfe",
	"Familie",
	"Film",
	"Flummi",
	"Sessel",
	"alle",
	"Mai",
	"Moos",
	"Ananas",
	"Null",
	"Name",
	"Lineal",
	"Nase",
	"Sonne",
	"Mond",
	"fallen",
	"messen",
	"Messer",
	"Gabel",
	"Löffel",
	"Nuss",
	"Mann",
	"malen",
	"essen",
	"fassen",
	"lassen",
	"lesen",
	"nennen",
	"sammeln",
	"offen",
	"minus",
	"Ofen",
	"Unfall",
	"Unsinn",
	"anfassen",
	"Waffel",
	"Wal",
	"Badewanne",
	"Welle",
	"Wolf",
	"Wolle",
	"Arm",
	"Bein",
	"Hand",
	"Fuß",
	"Kopf",
	"Ferien",
	"Flur",
	"rennen",
	"lernen",
	"fressen",
	"rollen",
	"rufen",
	"werfen",
	"wissen",
	"wollen",
	"fror",
	"froh",
	"anrufen",
	"immer",
	"leer",
	"Meer",
	"Murmel",
	"Nummer",
	"Rasen",
	"Rassel",
	"Rolle",
	"Roller",
	"Rose",
	"Rosine",
	"Sommer",
	"Wasser",
	"unser",
	"warm",
	"warum",
	"Tomate",
	"Gurke",
	"Alter",
	"Antenne",
	"Antwort",
	"Ast",
	"Eltern",
	"Ente",
	"ernst",
	"Ernte",
	"etwas",
	"falten",
	"Fenster",
	"Fest",
	"Fett",
	"flattern",
	"Foto",
	"frisst",
	"Futter",
	"interessant",
	"Laterne",
	"Luft",
	"Lust",
	"Mantel",
	"Meter",
	"Minute",
	"Mitte",
	"Moment",
	"Monat",
	"Monster",
	"Motor",
	"Mut",
	"Mutter",
	"Vater",
	"Großmutter",
	"Großvater",
	"Natur",
	"Nest",
	"Note",
	"nett",
	"oft",
	"Ort",
	"Ostern",
	"raten",
	"Ratte",
	"retten",
	"rot",
	"Saft",
	"Salat",
	"selten",
	"sofort",
	"Tafel",
	"Tal",
	"Tanne",
	"Tante",
	"Tasse",
	"Tee",
	"Telefon",
	"Teller",
	"Test",
	"Toilette",
	"toll",
	"Tonne",
	"Tor",
	"Torte",
	"tot",
	"treffen",
	"treten",
	"Trommel",
	"Turm",
	"turnen",
	"warten",
	"Welt",
	"Wetter",
	"Winter",
	"Wort",
	"Wurst",
	"Wut",
	"turnen",
	"Weihnachten",
	"nehmen",
	"erwarten",
	"Interesse",
	"Information",
	"Liter",
	"Metall",
	"Wörter",
	"Sätze",
	"Satz",
	"Astronaut",
	"Auto",
	"Laus",
	"laut",
	"Mauer",
	"Maus",
	"rau",
	"Raum",
	"sauer",
	"trauen",
	"Traum",
	"Maulwurf",
	"Automat",
	"laufen",
	"raufen",
	"Taufe",
	"Frau",
	"Wurm",
	"Ampel",
	"Apfelsine",
	"April",
	"aufpassen",
	"impfen",
	"Lampe",
	"Lappen",
	"Lippe",
	"Lupe",
	"Mops",
	"Opossum",
	"paar",
	"Palme",
	"Pappe",
	"Partner",
	"Partnerin",
	"Pass",
	"passen",
	"passt",
	"Pause",
	"Panne",
	"Perle",
	"Pinsel",
	"Pirat",
	"Pilot",
	"Plan",
	"plus",
	"Post",
	"Popo",
	"Pottwal",
	"prima",
	"Pulli",
	"Puppe",
	"pusten",
	"pustet",
	"Raupe",
	"Suppe",
	"Tapir",
	"Tipp",
	"Topf",
	"Treppe",
	"Tulpe",
	"Wappen",
	"Wespe",
	"Wippe",
	"wippen",
	"allein",
	"Ameise",
	"Ei",
	"Eier",
	"Eimer",
	"ein",
	"einmal",
	"Eis",
	"Eisen",
	"eislaufen",
	"Feier",
	"feiern",
	"fein",
	"frei",
	"Leine",
	"leise",
	"leiser",
	"Leiter",
	"mein",
	"Meise",
	"meist",
	"meistens",
	"Preis",
	"Reifen",
	"Reim",
	"reimen",
	"Reis",
	"Reise",
	"reisen",
	"reiten",
	"Reiter",
	"Seife",
	"Seil",
	"sein",
	"seit",
	"Seite",
	"Teil",
	"teilen",
	"weil",
	"Wein",
	"weinen",
	"weit",
	"weiter",
}
