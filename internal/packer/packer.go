// Package packer implements the optimal pack calculation algorithm.
//
// Rules (in priority order):
//  1. Only whole packs can be sent. Packs cannot be broken open.
//  2. Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order.
//  3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.
package packer

import (
	"container/heap"
	"math"
	"sort"
)

// Calculate returns the optimal pack breakdown for the given order quantity
// using the provided pack sizes.
//
// Algorithm:
//  1. Dijkstra over residues mod minPack finds the minimum reachable T >= order.
//  2. Greedy reconstruction (largest pack first, with reachability check) finds
//     the fewest packs that sum to exactly T.
func Calculate(order int, packSizes []int) map[int]int {

	if order <= 0 || len(packSizes) == 0 {
		return nil
	}

	sizes := make([]int, len(packSizes))

	// Copy the pack sizes to the sizes slice.
	copy(sizes, packSizes)

	// Sort the pack sizes in descending order.
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	// Build reachability table via Dijkstra’s algorithm over residues.
	dist := dijkstraAlgorithm(sizes)

	// Find minimum T >= order.
	T := minTotal(order, sizes[len(sizes)-1], dist)

	// Reconstruct with fewest packs.
	return reconstruct(T, sizes, dist)
}

const maxInteger = math.MaxInt

// ── Dijkstra over residues ────────────────────────────────────────────────────
//
// dist[r] = smallest value reachable (as a non-negative integer linear
// combination of sizes) whose value mod minPack == r.
// This is Dijkstra on an implicit graph with minPack nodes; each node r has
// edges r → (r+s) % minPack with weight s, for every pack size s.

func dijkstraAlgorithm(sizes []int) []int {

	minPack := sizes[len(sizes)-1]

	dist := make([]int, minPack)

	for i := range dist {
		dist[i] = maxInteger
	}

	dist[0] = 0

	minHeap := NewMinHeap()

	minHeap.Push(item{value: 0, residue: 0})

	heap.Init(minHeap)

	for minHeap.Len() > 0 {

		cur := heap.Pop(minHeap).(item)

		if cur.value > dist[cur.residue] {
			continue // stale entry
		}

		for _, s := range sizes {

			next := cur.value + s

			nr := next % minPack

			if next < dist[nr] {
				dist[nr] = next
				heap.Push(minHeap, item{value: next, residue: nr})
			}
		}
	}
	return dist
}

// For each residue r, dist[r] is the smallest base value with that residue.
// To reach a value >= order with residue r, we may need to add k*minPack to
// dist[r] for some k >= 0.  We pick the smallest such value across all r.
func minTotal(order, minPack int, dist []int) int {
	best := maxInteger
	for r, base := range dist {
		if base == maxInteger {
			continue
		}
		candidate := base
		if candidate < order {
			// ceil((order - base) / minPack) * minPack
			diff := order - base
			steps := (diff + minPack - 1) / minPack
			candidate = base + steps*minPack
		}

		// Verify: residue must be preserved (adding k*minPack never changes it).
		if candidate%minPack != r {
			continue
		}
		if candidate < best {
			best = candidate
		}
	}
	return best
}

// Greedy largest-first, but before committing to `count` packs of size s we
// reduce count until the remainder is itself reachable (dist[remainder % minPack]
// <= remainder).  This avoids dead ends where a large pack leaves a remainder
// that cannot be expressed with the remaining sizes.
func reconstruct(total int, sizes []int, dist []int) map[int]int {
	minPack := sizes[len(sizes)-1]
	result := make(map[int]int)
	remaining := total

	for i, s := range sizes {
		if remaining <= 0 {
			break
		}
		if i == len(sizes)-1 {
			// Last (smallest) pack: must divide evenly because T was chosen
			// to be reachable, and all previous steps preserved reachability.
			count := remaining / s
			if count > 0 {
				result[s] = count
				remaining -= count * s
			}
			break
		}

		count := remaining / s
		// Walk count down until the remainder is reachable.
		for count >= 0 {
			rem := remaining - count*s
			if rem == 0 || isReachable(rem, minPack, dist) {
				break
			}
			count--
		}
		if count > 0 {
			result[s] = count
			remaining -= count * s
		}
	}

	return result
}

// isReachable reports whether value can be expressed as a non-negative
// integer linear combination of the pack sizes, using the Dijkstra dist table.
func isReachable(value, minPack int, dist []int) bool {
	r := value % minPack
	return dist[r] != maxInteger && dist[r] <= value
}
