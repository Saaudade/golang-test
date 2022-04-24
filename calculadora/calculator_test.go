package calculator

import "testing"

func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{
			name:                  "Should apply 20",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        150,
			expectedAmount:        130,
		},
		{
			name:                  "Should apply 40",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        200,
			expectedAmount:        160,
		},
		{
			name:                  "Should apply 60",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        350,
			expectedAmount:        290,
		},
		{
			name:                  "Should not apply",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        10,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
			amount := calculator.Calculate(tc.purchaseAmount)

			if amount != tc.expectedAmount {
				t.Errorf("Expected %v, got %v", tc.expectedAmount, amount) // Error = Log + Fail
			}
		})
	}
}
