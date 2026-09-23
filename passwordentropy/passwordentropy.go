/*
Create a Go function that calculates the **information entropy** of a password in bits.
Password entropy measures the unpredictability and brute-force resistance of a password
based on its length and the diversity of character sets used.

Formula:
    E = L * log2(R)

Where:
    L = Length of the password
    R = Total size of the character pool based on active categories:
        - Lowercase letters ('a' - 'z'): +26
        - Uppercase letters ('A' - 'Z'): +26
        - Digits ('0' - '9'):            +10
        - Printable symbols (ASCII 32-126 excluding alphanumeric): +32

The result must be rounded to 2 decimal places.
If the password is empty or contains no recognized characters, return 0.00.
*/

package main

import (
	"fmt"
	"math"
)

// PasswordEntropy calculates the information entropy in bits for a given password,
// rounded to 2 decimal places.
func PasswordEntropy(password string) float64 {
	// Step 1: Guard against empty passwords.
	// A password with length 0 has 0 bits of entropy.
	length := len(password)
	if length == 0 {
		return 0.00
	}

	// Step 2: Track which character sets are present in the password.
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	// Step 3: Scan each character to classify its category.
	for i := 0; i < length; i++ {
		ch := password[i]
		switch {
		case ch >= 'a' && ch <= 'z':
			hasLower = true
		case ch >= 'A' && ch <= 'Z':
			hasUpper = true
		case ch >= '0' && ch <= '9':
			hasDigit = true
		case ch >= 32 && ch <= 126:
			// Printable ASCII characters (32 to 126) that are not alphanumeric
			// fall into the symbol/special character set.
			hasSpecial = true
		}
	}

	// Step 4: Determine the total character pool size (R).
	poolSize := 0
	if hasLower {
		poolSize += 26 // a-z
	}
	if hasUpper {
		poolSize += 26 // A-Z
	}
	if hasDigit {
		poolSize += 10 // 0-9
	}
	if hasSpecial {
		poolSize += 32 // printable punctuation, spaces, and symbols
	}

	// If no valid recognized printable characters exist, entropy is 0.
	if poolSize == 0 {
		return 0.00
	}

	// Step 5: Compute information entropy using Shannon's formula:
	// E = L * log2(R)
	entropy := float64(length) * math.Log2(float64(poolSize))

	// Step 6: Round to 2 decimal places.
	// math.Round rounds to the nearest integer.
	return math.Round(entropy*100) / 100
}

 func main() {
        tests := []string{
            "",
            "7",
            "123456",
            "password",
            "Password123",
            "P@ssw0rd!2026",
            "correct horse battery staple!",
        }

        for _, pw := range tests {
            fmt.Printf("%-32q -> %6.2f bits\n", pw, PasswordEntropy(pw))
        }
    }

