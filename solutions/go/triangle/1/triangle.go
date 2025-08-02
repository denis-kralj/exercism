// Contains methods for detecting triangle types
package triangle

type Kind int

const (
    NaT = Kind(1) // not a triangle
    Equ = Kind(2) // equilateral
    Iso = Kind(3) // isosceles
    Sca = Kind(4) // scalene
)

// For a given set of sides, determines if it is a triangle, and if so, what kind
func KindFromSides(a, b, c float64) Kind {
    unfitForTriangle := a <= 0 || b <= 0 || c <= 0 || a + b < c || c + b < a || a + c < b

    switch {
        case unfitForTriangle: return NaT
        case a == b && b == c: return Equ
    	case a == b || b == c || a == c: return Iso
    	default: return Sca
    }
}
