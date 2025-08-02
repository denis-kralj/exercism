package strain

func Keep[T any](collection []T, predicate func(T) bool) (result []T) {
    for _, element := range collection {
        if predicate(element) {
            result = append(result, element)
        }
    }

    return
}

func Discard[T any](collection []T, predicate func(T) bool) (result []T) {
    return Keep(collection, func(element T) bool { return !predicate(element) })
}