package infra

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("Expected %d, got %d", -2, ans)
	}
}

// часто в go прогоняют код через табличные тесты
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{2, 2, 1},
		{2, 3, 1},
	}
	for _, tt := range tests {
		testament := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testament, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("Expected %d, got %d", tt.want, ans)
			}
		})

	}
}
