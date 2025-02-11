package service

import (
	"sort"

	"gymshark/packcalculator/internal/model"
)

// DefaultPackCalculator implementation
type DefaultPackCalculator struct{}

func (pc *DefaultPackCalculator) Calculate(orderAmount int, packSizes []int) map[int]int {
	// Initialize variables to track best solution
	bestResult := make(map[int]int)
	bestTotalItems := int(^uint(0) >> 1) // Max int value
	bestPackCount := int(^uint(0) >> 1)  // Max int value

	// Generate all possible combinations using backtracking
	var findCombination func(remaining int, startIdx int, currentResult map[int]int)
	findCombination = func(remaining int, startIdx int, currentResult map[int]int) {
		// Calculate current solution metrics
		currentTotalItems := 0
		currentPackCount := 0
		for size, count := range currentResult {
			currentTotalItems += size * count
			currentPackCount += count
		}

		// Check if this is a valid solution
		if remaining <= 0 {
			if currentTotalItems >= orderAmount && // Must fulfill order
				(currentTotalItems < bestTotalItems || // Less items
					(currentTotalItems == bestTotalItems && currentPackCount < bestPackCount)) { // Same items, fewer packs
				bestTotalItems = currentTotalItems
				bestPackCount = currentPackCount
				bestResult = make(map[int]int)
				for k, v := range currentResult {
					bestResult[k] = v
				}
			}
			return
		}

		// Try adding each pack size
		for i := startIdx; i < len(packSizes); i++ {
			size := packSizes[i]
			currentResult[size]++
			findCombination(remaining-size, i, currentResult)
			currentResult[size]--
			if currentResult[size] == 0 {
				delete(currentResult, size)
			}
		}
	}

	findCombination(orderAmount, 0, make(map[int]int))
	return bestResult
}

// PackServiceImpl implementation
type PackServiceImpl struct {
	calculator model.PackCalculatorStrategy
	packSizes  []int
}

func NewPackService(calculator model.PackCalculatorStrategy, packSizes []int) model.PackService {
	sorted := make([]int, len(packSizes))
	copy(sorted, packSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	return &PackServiceImpl{
		calculator: calculator,
		packSizes:  sorted,
	}
}

func (ps *PackServiceImpl) CalculatePacks(orderAmount int) model.PackResponse {
	packs := ps.calculator.Calculate(orderAmount, ps.packSizes)
	totalItems := 0
	for size, count := range packs {
		totalItems += size * count
	}

	return model.PackResponse{
		OrderAmount: orderAmount,
		Packs:       packs,
		TotalItems:  totalItems,
	}
}
