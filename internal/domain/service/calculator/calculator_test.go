package calculatorservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	t.Parallel()

	var defaultPDFPackSizes = []int{250, 500, 1000, 2000, 5000}
	var defaultEmailPackSizes = []int{53, 23, 31}

	testCases := []struct {
		name           string
		packSizes      []int
		amount         int
		expectedResult map[int]int
	}{
		{
			name:           "Test Case Zero Amount",
			packSizes:      defaultPDFPackSizes,
			amount:         0,
			expectedResult: nil,
		},
		{
			name:           "Test Case Zero Pack Sizes",
			packSizes:      nil,
			amount:         500_000,
			expectedResult: nil,
		},
		{
			name:      "Test Case PDF Example - 1",
			packSizes: defaultPDFPackSizes,
			amount:    1,
			expectedResult: map[int]int{
				250: 1,
			},
		},
		{
			name:      "Test Case PDF Example - 2",
			packSizes: defaultPDFPackSizes,
			amount:    250,
			expectedResult: map[int]int{
				250: 1,
			},
		},
		{
			name:      "Test Case PDF Example - 3",
			packSizes: defaultPDFPackSizes,
			amount:    251,
			expectedResult: map[int]int{
				500: 1,
			},
		},
		{
			name:      "Test Case PDF Example - 4",
			packSizes: defaultPDFPackSizes,
			amount:    501,
			expectedResult: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			name:      "Test Case PDF Example - 5",
			packSizes: defaultPDFPackSizes,
			amount:    12_001,
			expectedResult: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
		{
			name:      "Test Case Email Example - 1",
			packSizes: defaultEmailPackSizes,
			amount:    500_000,
			expectedResult: map[int]int{
				23: 2,
				31: 7,
				53: 9429,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			svc := CalculatorService{}
			result := svc.calculate(tc.amount, tc.packSizes)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
