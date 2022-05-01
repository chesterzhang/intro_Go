package triangletest

import "math"

func CalTriangle(a,b int) int   {
	c:=int(math.Sqrt(float64(a*a+b*b)))
	return c
}