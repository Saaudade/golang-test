package calculator

import (
	"testing"
)

type DiscountRepositoryMock struct{}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

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
			name:                  "should apply 20",
			minimumPurchaseAmount: 100,
			discount:              20,
			purchaseAmount:        150,
			expectedAmount:        130,
		},
		{
			name:                  "should apply 40",
			minimumPurchaseAmount: 100,
			purchaseAmount:        200,
			expectedAmount:        160,
		},
		{
			name:                  "should apply 60",
			minimumPurchaseAmount: 100,
			purchaseAmount:        350,
			expectedAmount:        290,
		},
		{
			name:                  "should not apply",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			expectedAmount:        50,
		},
		{
			name:                  "zero minimum purchase amount",
			minimumPurchaseAmount: 0,
			purchaseAmount:        10,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			discountRepositoryMock := DiscountRepositoryMock{}

			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)

			if err != nil {
				t.Fatalf("could not instantiate the calculator %s", err.Error()) // FailNow+log
			}

			amount := calculator.Calculate(tc.purchaseAmount)

			if amount != tc.expectedAmount {
				t.Errorf("Expected %v, got %v", tc.expectedAmount, amount) // Error = Log + Fail
			}
		})
	}
}

func TestDiscountCalculatorShouldFailWithZeroMinimumAmount(t *testing.T) {
	discountRepositoryMock := DiscountRepositoryMock{}

	_, err := NewDiscountCalculator(0, discountRepositoryMock)
	if err == nil {
		t.Fatalf("should not create the calculator with zero purchase amount %s", err.Error()) // FailNow+log
	}
}
