package main

import "fmt"

func winIdx(value int) int {
	if value == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	arr := [3][2]string{
		{"html", "c#"},
		{"c#", "pyton"},
		{"pyton", "c#"},
	}

	myMap := make(map[string]int)
	score := []int{0, 0, 1}
	bestTeam := ""
	bestTeamScore := 0

	for idx, value := range arr {
		myMap[value[winIdx(score[idx])]] += 3
		if myMap[value[winIdx(score[idx])]] > bestTeamScore {
			bestTeamScore = myMap[value[winIdx(score[idx])]]
			bestTeam = value[winIdx(score[idx])]
		}
	}

	fmt.Println(bestTeam, bestTeamScore)
}
