package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	// Preencher nums1 de trás para frente
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// Se ainda houver elementos em nums2, copiamos para nums1
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
