package complexnumbers

import "math"
// Define the Number type here.
type Number struct {
    realPart float64
    imaginaryPart float64
}

func (n Number) Real() float64 {
	return n.realPart
}

func (n Number) Imaginary() float64 {
	return n.imaginaryPart
}

func (n1 Number) Add(n2 Number) Number {
	return Number{ realPart: n1.Real() + n2.Real(), imaginaryPart: n1.Imaginary() + n2.Imaginary() }
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{ realPart: n1.Real() - n2.Real(), imaginaryPart: n1.Imaginary() - n2.Imaginary() }
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{realPart:n1.Real()*n2.Real() - n1.Imaginary()*n2.Imaginary(), imaginaryPart: n1.Imaginary()*n2.Real() + n1.Real()*n2.Imaginary()}
}

func (n Number) Times(factor float64) Number {
	return Number{realPart:n.Real()*factor, imaginaryPart:n.Imaginary()*factor}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
        realPart: (n1.Real()*n2.Real()+n1.Imaginary()*n2.Imaginary())/(n2.Real()*n2.Real()+n2.Imaginary()*n2.Imaginary()),
        imaginaryPart: (n1.Imaginary()*n2.Real()-n1.Real()*n2.Imaginary())/(n2.Real()*n2.Real()+n2.Imaginary()*n2.Imaginary()),
    }
}

func (n Number) Conjugate() Number {
	return Number{realPart: n.Real(), imaginaryPart: n.Imaginary()*-1}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.Real()*n.Real()+n.Imaginary()*n.Imaginary())
}

func (n Number) Exp() Number {
	return Number{
        realPart: math.Pow(math.E, n.Real())*math.Cos(n.Imaginary()), 
        imaginaryPart: math.Pow(math.E, n.Real())* math.Sin(n.Imaginary()),
    }
}
