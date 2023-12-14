package utils

import "fmt"

// PhysicalPinToBCM converts a physical pin number to the corresponding BCM pin for Raspberry Pi 2, 3 and 4.
func PhysicalPinToBCM(pin int) (int, error) {
	switch pin {
	case 3:
		return 2, nil
	case 5:
		return 3, nil
	case 7:
		return 4, nil
	case 8:
		return 14, nil
	case 10:
		return 15, nil
	case 11:
		return 17, nil
	case 12:
		return 18, nil
	case 13:
		return 27, nil
	case 15:
		return 22, nil
	case 16:
		return 23, nil
	case 18:
		return 24, nil
	case 19:
		return 10, nil
	case 21:
		return 9, nil
	case 22:
		return 25, nil
	case 23:
		return 11, nil
	case 24:
		return 8, nil
	case 26:
		return 7, nil
	case 29:
		return 5, nil
	case 31:
		return 6, nil
	case 32:
		return 12, nil
	case 33:
		return 13, nil
	case 35:
		return 19, nil
	case 36:
		return 16, nil
	case 37:
		return 26, nil
	case 38:
		return 20, nil
	case 40:
		return 21, nil
	default:
		return 0, fmt.Errorf("Invalid physical pin number: %d", pin)
	}
}
