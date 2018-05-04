package main

import "fmt"

func main() {

	StartProbability := map[string]float64{"Rainy": 0.6, "Sunny": 0.4}
	TransitionProbability := map[string]map[string]float64{"Rainy" : {"Rainy": 0.7, "Sunny": 0.3},"Sunny" : {"Rainy": 0.4, "Sunny": 0.6}}


	fmt.Println(StartProbability)
	fmt.Println(TransitionProbability)
}
