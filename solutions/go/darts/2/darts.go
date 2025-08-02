package darts

func Score(x, y float64) int {
    // https://math.stackexchange.com/questions/1307832/how-to-tell-if-x-y-coordinate-is-within-a-circle
    switch dartCircleRadiusSquared := x * x + y * y; {
        case dartCircleRadiusSquared <= 1: return 10
    	case dartCircleRadiusSquared <= 25: return 5
    	case dartCircleRadiusSquared <= 100: return 1
    	default: return 0
    }
}
