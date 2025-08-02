package clock

import "fmt"
// Define the Clock type here.
type Clock int

func New(h, m int) Clock {
    val := (h*60+m) % (24*60)
    
    if val <= 0 {
        val+=24*60
    }

    return Clock(val)
}

func (c Clock) Add(m int) Clock {
	c += Clock(m)
    h := (int(c) / 60) % 24
    m = int(c) % 60
    return New(h,m)
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-1*m)
}

func (c Clock) String() string {
    h := (int(c) / 60) % 24
    m := int(c) % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
