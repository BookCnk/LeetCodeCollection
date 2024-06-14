package main

func romanToInt(s string) int {
	// Map to store the Roman numerals and their integer values
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	// Iterate over the string in reverse order
	for i := len(s) - 1; i >= 0; i-- {
		currentValue := roman[s[i]]

		// If the current value is less than the previous value, subtract it
		if currentValue < prevValue {
			total -= currentValue
		} else {
			// Otherwise, add it
			total += currentValue
		}
		// Update the previous value
		prevValue = currentValue
	}

	return total
}
