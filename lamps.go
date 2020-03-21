package lamps

func LampSwitch(lampsNum, timesPressed int) int {
	lamps := make([]bool, lampsNum, lampsNum)
	//foreach time pressed
	for i := 0; i < timesPressed; i++ {
		//foreach lamp that its index is a multiplication of current pressed times, flip its condition
		for lampIndex, condition := range lamps {
			if (lampIndex+1)%(i+1) == 0 {
				lamps[lampIndex] = !condition
			}
		}
	}
	//count the number of "on" lamps
	count := 0
	for _, condition := range lamps {
		if condition {
			count++
		}
	}
	return count
}
