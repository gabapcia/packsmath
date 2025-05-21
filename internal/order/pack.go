package order

import "context"

// PackStorage defines the interface for retrieving available pack sizes.
//
//go:generate moq -pkg mock -out mock/pack_storage.go . PackStorage
type PackStorage interface {
	// ListPackSizes returns the list of available pack sizes from storage.
	ListPackSizes(ctx context.Context) ([]int, error)
}

// resolvePackCombination determines the most efficient way to fulfill an order using given pack sizes
//
// The function returns a map where the keys are pack sizes and the values are the number of packs used.
// It guarantees:
//  1. The least total number of items sent (>= order).
//  2. Among those, the fewest number of packs.
//
// If no exact combination is possible, it selects the smallest single pack that can fulfill the order
func resolvePackCombination(order int, packSizes []int) map[int]int {
	type result struct {
		totalItems int
		numPacks   int
		counts     map[int]int
	}

	memo := make(map[int]*result)

	var dfs func(remaining int) *result
	dfs = func(remaining int) *result {
		if res, ok := memo[remaining]; ok {
			return res
		}

		var best *result

		for _, size := range packSizes {
			next := remaining - size
			if next < 0 {
				continue
			}

			sub := dfs(next)
			if sub == nil {
				continue
			}

			total := sub.totalItems + size
			counts := make(map[int]int)

			for k, v := range sub.counts {
				counts[k] = v
			}

			counts[size]++

			curr := &result{
				totalItems: total,
				numPacks:   sub.numPacks + 1,
				counts:     counts,
			}

			if best == nil ||
				curr.totalItems < best.totalItems ||
				(curr.totalItems == best.totalItems && curr.numPacks < best.numPacks) {
				best = curr
			}
		}

		// base case: exactly fulfilled
		if remaining == 0 {
			best = &result{totalItems: 0, numPacks: 0, counts: make(map[int]int)}
		}

		memo[remaining] = best
		return best
	}

	best := dfs(order)
	if best == nil {
		// fallback: find smallest pack that covers the order
		minOver := -1
		var overCounts map[int]int
		for _, size := range packSizes {
			if size >= order && (minOver == -1 || size < minOver) {
				minOver = size
				overCounts = map[int]int{size: 1}
			}
		}
		return overCounts
	}
	return best.counts
}

// PackOrder returns the optimal pack distribution for the given order quantity
//
// It retrieves available pack sizes from the underlying PackStorage implementation and uses resolvePackCombination to calculate the result
func (s *service) PackOrder(ctx context.Context, order int) (map[int]int, error) {
	sizes, err := s.packStorage.ListPackSizes(ctx)
	if err != nil {
		return nil, err
	}

	packs := resolvePackCombination(order, sizes)
	return packs, nil
}
