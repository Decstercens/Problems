func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j := m-1, n-1
    for pos := n + m - 1; pos >= 0; pos-- {
        if i >= 0 && j >= 0 {
            if nums1[i] < nums2[j] {
                nums1[pos] = nums2[j]
                j--
            } else {
                nums1[pos] = nums1[i]
                i--
            }
        } else {
            if i < 0 {
                nums1[pos] = nums2[j]
                j--
            } else {
                nums1[pos] = nums1[i]
                i--
            }
        }
    }
}
