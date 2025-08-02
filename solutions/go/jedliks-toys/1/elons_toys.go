package elon

import "fmt"

func (target *Car) Drive() {
    if target.battery >= target.batteryDrain {
	    target.battery -= target.batteryDrain
    	target.distance+= target.speed
    }
}

func (target *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %v meters", target.distance)
}

func (target *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %v%%", target.battery)
}

func (target *Car) CanFinish(trackDistance int) bool {
    return trackDistance <= (target.battery / target.batteryDrain * target.speed) + target.distance
}
