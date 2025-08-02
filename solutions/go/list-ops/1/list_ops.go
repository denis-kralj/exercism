package listops

import "fmt"
// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
    fmt.Println("list", s)
	for _, ele := range s {
        fmt.Println("ele", ele)
        fmt.Println("initial before", initial)
        initial = fn(initial, ele)
        fmt.Println("initial after", initial)
    }
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	return s.Reverse().Foldl(func(first int,second int) int { return fn(second, first) }, initial)
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}

    for _, ele := range s {
        if fn(ele) {
            result = append(result, ele)
        }
    }

    return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
    result := IntList{}

    for _, ele := range s {
        result = append(result, fn(ele))
    }

    return result
}

func (s IntList) Reverse() IntList {
    result := IntList{}

	for i := len(s)-1; i >=0; i-- {
		result = append(result, s[i])
    }

    return result
}

func (s IntList) Append(lst IntList) IntList {
    s = append(s, lst...)
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
    for _, list := range lists {
        s = append(s, list...)
    }
	return s
}
