package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgLayerPrepTimeMinutes int) int {
    if avgLayerPrepTimeMinutes == 0 {
		avgLayerPrepTimeMinutes = 2
    } else if avgLayerPrepTimeMinutes < 0 {
    	panic("can't have negative prep time")
    }

	return len(layers) * avgLayerPrepTimeMinutes
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (noodleGrams int, sauceLiters float64) {
    for _, layer := range layers {
        if layer == "noodles" {
            noodleGrams += 50
        }
    	if layer == "sauce" {
            sauceLiters += .2
        }
    }

    return
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendLayers []string, myLayers []string) {
    myLayers[len(myLayers)-1] = friendLayers[len(friendLayers) - 1]
}
// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantitiesForTwo []float64, number int) []float64 {
    result := make([]float64, len(quantitiesForTwo), len(quantitiesForTwo))
    for i, quantity := range quantitiesForTwo {
        result[i] = quantity*float64(number)/2.0
    }

    return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
