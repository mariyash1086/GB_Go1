package fibonachi_test

import (
	"HW10/fibonachi"
	"testing"
)

func TestFibonachiProcessSmoke(t *testing.T) {
	result, err := fibonachi.FibonachiProcess(10)
	if err != nil {
		t.Fatalf("Func expexted to return no error")
	}
	expected := 55.0
	if result != expected {
		t.Fatalf("FibonachiProcess(1) expected to return %f, but return %f.", expected, result)
	}
}

func TestFibonachiProcessTable(t *testing.T) {

	type table struct {
		x, result float64
	}

	arr := []table{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 100000}, //ошибка
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
		{15, 610},
		{16, 987},
		{17, 1597},
		{18, 2584},
		{19, 4181},
		{20, 6765},
		{21, 10946},
		{22, 17711},
	}
	for _, elem := range arr {

		result, err := fibonachi.FibonachiProcess(elem.x)
		if err != nil {
			t.Errorf("Func expexted to return no error")
			continue
		}

		expected := elem.result
		if result != expected {
			t.Fatalf("FibonachiProcess %f expected to return %f, but return %f.", elem.x, expected, result)
		}
	}
}
