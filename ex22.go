package main

import "fmt"

func kelvinToCelsius (k float64) float64 {
  k -= 273.15
  return k
}

func celsiusToFahrenheit (c float64) float64 {
  c = (c * 9.0 / 5.0) + 32.0
  return c
}

func main () {
  kelvin := 294.0
  celsius := kelvinToCelsius(kelvin)
  fmt.Println(kelvin, "K is", celsius, "celsius");
}
