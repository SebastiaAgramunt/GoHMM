package main

import ("math/rand"
	"time"
)

type Sequence struct {
	HiddenStates []string
	ObservableStates []string

	StartProbability []float64
	TransitionProbability [][]float64
	EmissionProbability [][]float64
}

func CreateSequence(HiddenStates []string, ObservableStates []string) *Sequence{

	NHidden := len(HiddenStates)
	NObservable := len(ObservableStates)

	return &Sequence{HiddenStates:HiddenStates,
	ObservableStates:ObservableStates,
	StartProbability: CreateRandomArray(NHidden),
	TransitionProbability:CreateRandomMatrix(NHidden,NHidden),
	EmissionProbability:CreateRandomMatrix(NHidden,NObservable)}
}

func CreateRandomArray(numElem int) []float64{
	rand.Seed(time.Now().UTC().UnixNano())
	m := make([]float64,numElem)
	var k,sum float64

	for i,_ := range(m){
		k = rand.Float64()
		sum += k
		m[i] = k
	}
	for i,_ := range(m){
		m[i] = m[i]/sum
	}
	return m
}

func CreateRandomMatrix(numRows int, numCols int) [][]float64{
	rand.Seed(time.Now().UTC().UnixNano())
	m := make([][]float64, numRows)
	var k,sum float64
	for i := 0; i < numRows; i++ {
		m[i] = make([]float64, numCols)
	}

	for i := 0; i < numRows; i++{
		for j := 0; j< numCols; j++{
			k = rand.Float64()
			sum += k
			m[i][j] = k
		}
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			m[i][j] = m[i][j]/sum
		}
	}
	return m
}