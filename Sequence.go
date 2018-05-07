package main

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
	m := make([]float64,numElem)
	return m
}

func CreateRandomMatrix(numRows int, numCols int) [][]float64{
	m := make([][]float64, numRows)
	for i := 0; i < numRows; i++ {
		m[i] = make([]float64, numCols)
	}
	return m
}