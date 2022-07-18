package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		return (kth(nums1, nums2, l/2) + kth(nums1, nums2, l/2-1)) / 2
	} else {
		return kth(nums1, nums2, l/2)
	}
}

func kth(nums1 []int, nums2 []int, k int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[k])
	}
	if len(nums2) == 0 {
		return float64(nums1[k])
	}
	ia, ib := len(nums1)/2, len(nums2)/2
	ma, mb := nums1[ia], nums2[ib]

	if ia+ib < k {
		if ma > mb {
			return kth(nums1, nums2[ib+1:], k-ib-1)
		} else {
			return kth(nums1[ia+1:], nums2, k-ia-1)
		}
	} else {
		if ma > mb {
			return kth(nums1[:ia], nums2, k)
		} else {
			return kth(nums1, nums2[:ib], k)
		}
	}
}

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 3}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("%f", median)
}
