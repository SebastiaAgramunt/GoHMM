package main

import "fmt"


func main() {

	StartProbability := map[string]float64{"Rainy": 0.6, "Sunny": 0.4}
	TransitionProbability := map[string]map[string]float64{"Rainy" : {"Rainy": 0.7, "Sunny": 0.3},"Sunny" : {"Rainy": 0.4, "Sunny": 0.6}}
	EmissionProbability := map[string]map[string]float64{"Rainy" : {"walk": 0.1, "shop": 0.4, "clean": 0.5}, "Sunny" : {"walk": 0.6, "shop": 0.3, "clean": 0.1}}


	fmt.Println(StartProbability)
	fmt.Println(TransitionProbability)
	fmt.Println(EmissionProbability)

	seq := CreateSequence([]string{"Rainy","Sunny"},[]string{"walk","shop","clean"})
	fmt.Println(seq)

	fmt.Println(CreateRandomMatrix(4,4))
	fmt.Println(CreateRandomArray(4))

	fmt.Println(CreateRandomMatrix(4,4)[1][2])

}
