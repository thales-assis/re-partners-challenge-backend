package packer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	t.Parallel()

	var defaultPDFPackageSizes = []int{250, 500, 1000, 2000, 5000}
	var defaultEmailPackageSizes = []int{23, 31, 53}

	testCases := []struct {
		name           string
		packageSizes   []int
		amount         int
		expectedResult map[int]int
	}{
		{
			name:           "Test Case Zero Amount",
			packageSizes:   defaultPDFPackageSizes,
			amount:         0,
			expectedResult: nil,
		},
		{
			name:           "Test Case Zero Package Sizes",
			packageSizes:   nil,
			amount:         500_000,
			expectedResult: nil,
		},
		{
			name:         "Test Case PDF Example - 1",
			packageSizes: defaultPDFPackageSizes,
			amount:       1,
			expectedResult: map[int]int{
				250: 1,
			},
		},
		{
			name:         "Test Case PDF Example - 2",
			packageSizes: defaultPDFPackageSizes,
			amount:       250,
			expectedResult: map[int]int{
				250: 1,
			},
		},
		{
			name:         "Test Case PDF Example - 3",
			packageSizes: defaultPDFPackageSizes,
			amount:       251,
			expectedResult: map[int]int{
				500: 1,
			},
		},
		{
			name:         "Test Case PDF Example - 4",
			packageSizes: defaultPDFPackageSizes,
			amount:       501,
			expectedResult: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			name:         "Test Case PDF Example - 5",
			packageSizes: defaultPDFPackageSizes,
			amount:       12_001,
			expectedResult: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
		{
			name:         "Test Case Email Example - 1",
			packageSizes: defaultEmailPackageSizes,
			amount:       500_000,
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

			result := Calculate(tc.amount, tc.packageSizes)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
