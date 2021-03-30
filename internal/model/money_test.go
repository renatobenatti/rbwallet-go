package model

import "testing"

func TestNew(t *testing.T) {
	m := NewMoney(50.51)

	if m.Amount() != 50.51 {
		t.Errorf("expect '%.2f', result '%.2f'", 50.51, m.Amount())
	}
}

func TestAdd(t *testing.T) {
	testsAdd := []struct {
		money      *Money
		moneyToAdd *Money
		expected   float64
	}{
		{
			money:      NewMoney(10.10),
			moneyToAdd: NewMoney(11.00),
			expected:   21.10,
		},
		{
			money:      NewMoney(100.11),
			moneyToAdd: NewMoney(0.01),
			expected:   100.12,
		},
		{
			money:      NewMoney(0.10),
			moneyToAdd: NewMoney(1102.00),
			expected:   1102.10,
		},
	}

	for _, tt := range testsAdd {
		t.Run("Add money", func(t *testing.T) {
			result := tt.money.Add(tt.moneyToAdd)
			if !result.Equals(tt.expected) {
				t.Errorf("result %.2f expect %.2f", result.Amount(), tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	testsAdd := []struct {
		money           *Money
		moneyToSubtract *Money
		expected        float64
	}{
		{
			money:           NewMoney(111.10),
			moneyToSubtract: NewMoney(111.10),
			expected:        0.00,
		},
		{
			money:           NewMoney(100.11),
			moneyToSubtract: NewMoney(0.01),
			expected:        100.10,
		},
		{
			money:           NewMoney(1102.10),
			moneyToSubtract: NewMoney(1102.00),
			expected:        0.10,
		},
	}

	for _, tt := range testsAdd {
		t.Run("Add money", func(t *testing.T) {
			result := tt.money.Subtract(tt.moneyToSubtract)
			if !result.Equals(tt.expected) {
				t.Errorf("result %.2f expect %.2f", result.Amount(), tt.expected)
			}
		})
	}
}
