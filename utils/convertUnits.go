package utils

import (
	"fmt"

	u "github.com/bcicen/go-units"
)

func ConvertUnits(amount float64, fromUnitStr, toUnitStr string) (u.Value, error) {
	// Look up the unit objects by their string names
	fromUnit, err := u.Find(fromUnitStr)
	if err != nil {
		return u.Value{}, fmt.Errorf("unknown unit: %s", fromUnitStr)
	}

	toUnit, err := u.Find(toUnitStr)
	if err != nil {
		return u.Value{}, fmt.Errorf("unknown unit: %s", toUnitStr)
	}

	// Create a new value with the amount and from unit
	val := u.NewValue(amount, fromUnit)

	// Try to convert to the target unit
	convertedVal, err := val.Convert(toUnit)
	if err != nil {
		return u.Value{}, fmt.Errorf("cannot convert from %s to %s", fromUnitStr, toUnitStr)
	}

	return convertedVal, nil
}
