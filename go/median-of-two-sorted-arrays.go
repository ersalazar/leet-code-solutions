package leetcodesolutions

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	na := len(nums1)
	nb := len(nums2)
	n := na + nb

	var solve func(A, B []int, startA, endA, startB, endB, k int) float64
	solve = func(A, B []int, startA, endA, startB, endB, k int) float64 {
		if startA > endA {
			return float64(B[k-startA])
		}
		if startB > endB {
			return float64(A[k-startB])
		}

		indexA := (endA + startA) / 2
		indexB := (endB + startB) / 2

		valueA := A[indexA]
		valueB := B[indexB]

		if k > (indexA + indexB) {
			if valueA <= valueB {
				return solve(A, B, indexA+1, endA, startB, endB, k)
			} else {
				return solve(A, B, startA, endA, indexB+1, endB, k)
			}
		} else {
			if valueA <= valueB {
				return solve(A, B, startA, endA, startB, indexB-1, k)
			} else {
				return solve(A, B, startA, indexA-1, startB, endB, k)
			}
		}
	}

	if n%2 == 0 {
		return (solve(nums1, nums2, 0, na-1, 0, nb-1, n/2-1) +
			solve(nums1, nums2, 0, na-1, 0, nb-1, n/2)) / 2.0
	} else {
		return solve(nums1, nums2, 0, na-1, 0, nb-1, n/2)
	}
}
