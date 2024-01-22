package quizTypes

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Words struct{}

const QuizTypeWords = "words"

var _ Runnable = &Words{}

func NewWordsQuiz() *Words {
	return &Words{}
}

func (w Words) Run(difficulty int) bool {
	quizWord := getRandomWord()
	quizString := quizWord[:len(quizWord)-1]
	answer := getAnswerString(quizString)
	expectedAnswer := quizWord[len(quizWord)-1:]

	return evaluateSolutionString(expectedAnswer, answer, quizString)
}

func getRandomWord() string {
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Int() % len(wordList)
	return wordList[randomIndex]
}

func getAnswerString(quizString string) string {
	fmt.Print(quizString)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSuffix(input, "\n")
}

func evaluateSolutionString(expectedAnswer string, actualAnswer string, quizString string) bool {
	success := true

	for expectedAnswer != actualAnswer {
		success = false
		fmt.Println("wrong, try again!")
		actualAnswer = getAnswerString(quizString)
	}

	fmt.Println("correct")

	return success
}

// todo: move word list to config file
var wordList = []string{
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
	//"Fels",
	"Elf",
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
}
