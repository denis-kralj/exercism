package gross

const dozen int = 12
const (
    quarter_of_a_dozen = dozen / 4
    half_of_a_dozen = dozen / 2
    small_gross = dozen * 10
    gross = dozen * dozen
    great_gross = gross * dozen
)
// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["dozen"] = dozen
    units["quarter_of_a_dozen"] = quarter_of_a_dozen
    units["half_of_a_dozen"] = half_of_a_dozen
    units["small_gross"] = small_gross
    units["gross"] = gross
    units["great_gross"] = great_gross

    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]

    if exists {
        bill[item] += value
    }

    return exists
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]

    if !exists {
        return false
    }
    
    currentItemCount, exists := bill[item]

    if !exists || currentItemCount < value {
        return false
    }

    bill[item] -= value
    if bill[item] == 0 {
        delete(bill, item)
    }

    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (value int, exists bool) {
	value, exists = bill[item]
    return
}
