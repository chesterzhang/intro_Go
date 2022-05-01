package triangletest

import (
	"testing"
)

//func CalTriangle(a,b int) int   {
//	c:=int(math.Sqrt(float64(a*a+b*b)))
//	return c
//}

//表格驱动测试, testing.T 用于传参
func TestTriangle(t *testing.T)  {
	tests := []struct{a,b,c int} {
		{3,4,5},
		{5,12,13},
	}

	for _,tt:=range tests   {
		actual:=CalTriangle(tt.a,tt.b)
		if actual!=tt.c{
			t.Errorf("calTriangle(%d,%d); "+"got %d; expected %d", tt.a,tt.b,actual, tt.c)
		}
	}
}

func BenchmarkCalTriangle(b *testing.B) {
	var x int=3
	var y int=4
	var z int=5

	for i:=0;i<b.N ;i++  {
		actual:=CalTriangle(x,y)
		if actual!=z{
			b.Errorf("calTriangle(%d,%d); "+"got %d; expected %d", x,y,actual, z)
		}
	}
}
