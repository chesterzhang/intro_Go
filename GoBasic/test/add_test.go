package test

import (
	"testing"
)

func TestAdd(t *testing.T) {

	testCases := []struct{
		a int
		b int
		c int
	}{
		//{5,3,4 },
		//		//{13,5,12},
		//		//{ 109,60,91},
		{3,2,1 },
		{10,5,5},
		{ 1000000,1,999999},
	}

	for _, testCase := range testCases{
		if result:=Add(testCase.b,testCase.c); result!=testCase.a{
			t.Errorf("Expected get %d, but got: %d ", testCase.a, result)
		}
	}
}

func BenchmarkAdd(b *testing.B) {

	b.ResetTimer()
	for i := 0; i<b.N;i++  {
		if result:=Add(3,4); result !=7  {
			b.Errorf("Expected get %d, but got: %d ", 7 , result)
		}
	}
}