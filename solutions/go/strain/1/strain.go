package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](collection []T, predicate func(T) bool) (result []T) {
    for _, element := range collection {
        if predicate(element) {
            result = append(result, element)
        }
    }

    return
}

func Discard[T any](collection []T, predicate func(T) bool) (result []T) {
    for _, element := range collection {
        if !predicate(element) {
            result = append(result, element)
        }
    }

    return
}
// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
