package customer

import (
	"errors"
	"testing"
)

func TestCustomerNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		exceptedErr error
	}
	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			exceptedErr: ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Saharat Muksarn",
			exceptedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name)
			if !errors.Is(err, tc.exceptedErr) {
				t.Errorf("excepted %v ,but got %v", tc.exceptedErr, err)
			}
		})
	}
}
