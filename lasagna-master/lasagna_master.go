package lasagna

func PreparationTime(layers []string, avgPrepTimePerLayer int) int {
	if avgPrepTimePerLayer == 0 {
		avgPrepTimePerLayer = 2
	}

	return len(layers) * avgPrepTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	n := 0
	s := 0.0
	for _, l := range layers {
		if l == "noodles" {
			n += 50
		} else if l == "sauce" {
			s += 0.2
		}
	}

	return n, s
}

func AddSecretIngredient(friendsList, myList []string) {
	for mk, mv := range myList {
		if mv == "?" {
			myList[mk] = friendsList[len(friendsList)-1]
		}
	}
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var r []float64
	for _, q := range quantities {
		r = append(r, q*float64(portions)/2)
	}
	return r
}
