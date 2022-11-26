package richestcustomerwealth

import (
	"fmt"
	"testing"
)

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, customer := range accounts {
		wealth := 0
		for _, balance := range customer {
			wealth += balance
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
func TestMaximumWealth(t *testing.T) {
	fmt.Println(maximumWealth([][]int{{1, 2}}))
}
