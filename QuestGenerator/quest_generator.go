package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Quest struct {
	name      string
	questType string
	area      string
	amount    int
}

type Target struct {
	name       string
	targetType string
	area       []string
	upperBound int
}

func main() {

	filePath := "quest_history.txt"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			return
		}
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lastQuest string

	for scanner.Scan() {
		lastQuest = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return
	}

	lastQuestType := ""
	lastQuestArea := ""

	if lastQuest != "" {
		parts := strings.Split(lastQuest, " ")
		lastQuestType = parts[1]
		lastQuestArea = parts[2]
	}

	targets := []Target{
		{name: "Amelia", targetType: "Friendly", area: []string{"Field", "Mountain", "Forest"}, upperBound: 2},
		{name: "Mary", targetType: "Friendly", area: []string{"Snow", "Beach", "Tundra"}, upperBound: 2},
		{name: "Ozymandias", targetType: "Friendly", area: []string{"Taiga", "Mushroom", "Swamp"}, upperBound: 2},
		{name: "Maximilian", targetType: "Friendly", area: []string{"Badlands", "Jungle", "Grove"}, upperBound: 2},
		{name: "Goblins", targetType: "Enemy", area: []string{"Mushroom", "Field", "Grove"}, upperBound: 20},
		{name: "Wolves", targetType: "Enemy", area: []string{"Taiga", "Snow", "Tundra"}, upperBound: 10},
		{name: "Cyclops", targetType: "Enemy", area: []string{"Beach", "Jungle", "Swamp"}, upperBound: 6},
		{name: "Wrym", targetType: "Enemy", area: []string{"Mountain", "Forest", "Badlands"}, upperBound: 2},
		{name: "Dandelions", targetType: "Flower", area: []string{"Mountain", "Beach", "Taiga"}, upperBound: 14},
		{name: "Daisies", targetType: "Flower", area: []string{"Mushroom", "Snow", "Jungle"}, upperBound: 12},
		{name: "Sunflowers", targetType: "Flower", area: []string{"Forest", "Field", "Tundra"}, upperBound: 16},
		{name: "Roses", targetType: "Flower", area: []string{"Grove", "Swamp", "Badlands"}, upperBound: 8},
	}

	targetIndex := rand.Intn(len(targets))
	target := targets[targetIndex]

	areaIndex := rand.Intn(3)
	area := target.area[areaIndex]

	for lastQuestType == target.targetType || lastQuestArea == area {
		targetIndex = rand.Intn(len(targets))
		target = targets[targetIndex]

		areaIndex = rand.Intn(3)
		area = target.area[areaIndex]
	}

	amount := target.upperBound/2 + rand.Intn(target.upperBound/2)

	quest := Quest{
		name:      target.name,
		questType: target.targetType,
		area:      area,
		amount:    amount,
	}

	questStr := fmt.Sprintf("%s %s %s %d\n", quest.name, quest.questType, quest.area, quest.amount)
	if _, err := file.WriteString(questStr); err != nil {
		return
	}

	switch quest.questType {
	case "Enemy":
		fmt.Printf("Kill %d %s in the %s area.\n", quest.amount, quest.name, quest.area)
	case "Friendly":
		fmt.Printf("Escort %s to the %s area.\n", quest.name, quest.area)
	case "Flower":
		fmt.Printf("Pick %d %s in the %s area.\n", quest.amount, quest.name, quest.area)
	}
}
