package main

func MergeSort(arr []int) []int {
	// データの長さが1以下ならそのまま返す
	if len(arr) <= 1 {
		return arr
	}

	// 中央インデックスを求める
	centerIndex := len(arr) / 2

	// 左右に分割する
	left := arr[:centerIndex]
	right := arr[centerIndex:]

	// 左右を再帰的にソートする
	sortedLeft := MergeSort(left)
	sortedRight := MergeSort(right)

	// merge関数で結合して返す
	return merge(sortedLeft, sortedRight)
}

// すでにソート済みの left と right を1つのスライスにマージする。
func merge(left, right []int) []int {
	result := []int{}
	i := 0
	j := 0

	// 左右の要素を順番に比較し、小さいほうをresultに追加する
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 片方に余った要素を全てresultに追加する
	if i < len(left) {
		result = append(result, left[i:]...)
	} else if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}
