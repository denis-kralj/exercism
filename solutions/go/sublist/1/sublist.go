package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
    if areEqual(l1, l2) {
        return RelationEqual
    } else if isSuperset(l1, l2) {
    	return RelationSuperlist
    } else if isSuperset(l2, l1) {
    	return RelationSublist
    }

    return RelationUnequal
}

func areEqual(l1, l2 []int) bool {
    if len(l1) != len(l2) {
        return false
    }

    for i, v := range l1 {
        if v != l2[i] {
            return false
        }
    }

    return true
}

func isSuperset(l1, l2 []int) bool {
    if len(l1) <= len(l2) {
        return false
    }

    for i:=0; i < len(l1) - len(l2) + 1; i++ {
		if areEqual(l1[i:i+len(l2)], l2) {
            return true
        }
    }

    return false
}
