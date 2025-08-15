package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minimumBus(families []int) int {
	sort.Ints(families) // urutkan biar pairing optimal
	count := 0
	low := 0
	high := len(families) - 1

	for low <= high {
		if families[low]+families[high] <= 4 {
			low++
		}
		high--
		count++
	}

	return count
}

func main() {
	var n int
	fmt.Print("Input the number of families: ")
	fmt.Scan(&n)

	fmt.Print("Input the number of members in the family (separated by a space): ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	if len(parts) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	families := make([]int, 0, n)
	for _, p := range parts {
		m, _ := strconv.Atoi(p)
		families = append(families, m)
	}

	fmt.Printf("Minimum bus required is: %d\n", minimumBus(families))
}
