package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
  gasPedal uint16
  brakePedal uint16
  steeringWheel int16
  topSpeedKmH float64
}

func (c car) kmh() float64 {
  return float64 (c.gasPedal) * (c.topSpeedKmH / usixteenbitmax)
}

func (c car) mph() float64 {
  return float64 (c.gasPedal) * (c.topSpeedKmH / usixteenbitmax/ kmh_multiple)
}

func (c *car) newtopSpeedKmH(newSpeed float64) {
  c.topSpeedKmH = newSpeed
}

func newerTopSpeed(c car, speed float64) car {
  c.topSpeedKmH = speed
  return c
}

func main() {
  aCar := car{ gasPedal: 22341,
               brakePedal: 0,
               steeringWheel: 12561,
               topSpeedKmH: 255.0 }

  fmt.Println(aCar.gasPedal)
  fmt.Println(aCar.kmh())
  fmt.Println(aCar.mph())

  // aCar.newtopSpeedKmH(500)

  aCar = newerTopSpeed(aCar, 500)

  fmt.Println(aCar.kmh())
  fmt.Println(aCar.mph())
}
