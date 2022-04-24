package calculator

import "testing"

func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        150,
			expectedAmount:        130,
		},
		{
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        200,
			expectedAmount:        160,
		},
		{
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        350,
			expectedAmount:        290,
		},
		{
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        50,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
		amount := calculator.Calculate(tc.purchaseAmount)

		if amount != tc.expectedAmount {
			t.Errorf("Expected %v, got %v", tc.expectedAmount, amount) // Error = Log + Fail
		}
	}
}

func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(150)

	if 130 != amount {
		t.Fail()
	}
}

func TestDiscountMulpliedByTwoApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(200)

	if 160 != amount {
		t.Fail()
	}
}

func TestDiscountMulpliedByThreeApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(350)

	if 290 != amount {
		t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calculate(50)

	if 50 != amount {
		// t.Logf("expected 50, got %v", amount)
		// t.Fail()

		t.Errorf("expected 50, got %v. Failed the discount was not expected to be applied", amount) // Error = Log + Fail
	}
}
