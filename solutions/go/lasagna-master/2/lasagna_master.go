package lasagna

func PreparationTime(layers []string, avgLayerPrepTimeMinutes int) int {
    if avgLayerPrepTimeMinutes == 0 {
		avgLayerPrepTimeMinutes = 2
    }

	return len(layers) * avgLayerPrepTimeMinutes
}

func Quantities(layers []string) (noodleGrams int, sauceLiters float64) {
    for _, layer := range layers {
        switch {
            case layer == "noodles": noodleGrams += 50
        	case layer == "sauce": sauceLiters += .2
        }
    }

    return
}

func AddSecretIngredient(friendLayers []string, myLayers []string) {
    myLayers[len(myLayers)-1] = friendLayers[len(friendLayers) - 1]
}

func ScaleRecipe(quantitiesForTwo []float64, number int) []float64 {
    result := make([]float64, len(quantitiesForTwo))

    for i, quantity := range quantitiesForTwo {
        result[i] = quantity*float64(number)/2.0
    }

    return result
}
