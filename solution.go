package square

import "math"

type MySideNum int

const SidesTriangle MySideNum = 3
const SidesSquare MySideNum = 4
const SidesCircle MySideNum = 0

func CalcSquare(sideLen float64, sidesNum MySideNum) float64 {
	var result = 0.0

	switch sidesNum {
		case SidesTriangle: result = calcForTriangle(sideLen)
		case SidesSquare: result = calcForSquare(sideLen)
		case SidesCircle: result = calcForCircle(sideLen)
	}

	return result
}

func calcForTriangle(sideLen float64) float64 {
	return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
}

func calcForSquare(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}

func calcForCircle(sideLen float64) float64 {
	return math.Pi * math.Pow(sideLen, 2)
}

