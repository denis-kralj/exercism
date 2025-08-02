package darts

func Score(x, y float64) int {
    // https://math.stackexchange.com/questions/1307832/how-to-tell-if-x-y-coordinate-is-within-a-circle
    dartCircleRadiusSquared := (x - 0.0)*(x - 0.0) + (y - 0.0)*(y - 0.0)
    switch {
        case dartCircleRadiusSquared <= 1: return 10
    	case dartCircleRadiusSquared <= 25: return 5
    	case dartCircleRadiusSquared <= 100: return 1
    	default: return 0
    }
}
