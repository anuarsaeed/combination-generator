package combgen

type combinations struct {
	clientInput [] string
	maxAllowedCharacters int
	nrOfCombinations int
	combinationList []string
}

type nonRepeatCombinations struct {
	combinations
}

type charRepeatCombinations struct {
	combinations
}

// Set max allowed character size.
func (comb *combinations) setMaxAllowedCharSize(charRepeat bool) {
	maxAllowedCharSize := comb.maxAllowedCharacters
	if maxAllowedCharSize == 0 || maxAllowedCharSize > len(comb.clientInput){
		comb.maxAllowedCharacters = len(comb.clientInput)
	}else {
		comb.maxAllowedCharacters = maxAllowedCharSize
	}
}

// Interface Calculation method.
func (noRepComb nonRepeatCombinations) calculateNrOfCombinations() int {
	nrOfChars := len(noRepComb.clientInput)	
	result := getFactorial(nrOfChars) / getFactorial(nrOfChars - noRepComb.maxAllowedCharacters)
	return result
}

// Intrface Calculation method.
func (noRepComb nonRepeatCombinations) calculateCombinations() [] string {
	var defaultElement = ""
	permute(&noRepComb, defaultElement, noRepComb.clientInput)
	return noRepComb.combinationList
}



// TODO functionality
func (charRepComb charRepeatCombinations) calculateNrOfCombinations() int {
	return 0
}
func (charRepComb charRepeatCombinations) calculateCombinations() [] string {
	return [] string {}
}

