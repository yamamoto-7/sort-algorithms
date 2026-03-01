package main

func QuickSort(arr []int) []int {
	// データの長さが1以下ならそのまま返す
	if len(arr) <= 1 {
		return arr
	}

	// ピボットを選ぶ(中央付近)
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var left []int
	var right []int

	// ピボット以外の要素をループし、
	// ピボットより小さいものを left に
	// それ以外を right に入れる
	for i := 0; i < len(arr); i++ {
		if i == pivotIndex {
			continue
		}
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	// 左右を再帰的にソートする
	sortedLeft := QuickSort(left)
	sortedRight := QuickSort(right)

	// sortedLeft + pivot + sortedRight を結合して返す
	result := append(sortedLeft, pivot)
	result = append(result, sortedRight...)

	return result
}
