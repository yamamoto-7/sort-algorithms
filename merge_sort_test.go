package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "正常系: マージソート成功 1",
			input:    []int{5, 3, 8, 1},
			expected: []int{1, 3, 5, 8},
		},
		{
			name:     "正常系: マージソート成功 2",
			input:    []int{5, 3, 8, 1, 9, 2, 100, 7},
			expected: []int{1, 2, 3, 5, 7, 8, 9, 100},
		},
		{
			name:     "正常系: 空配列を渡した場合",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "正常系: 数字の重複がある場合",
			input:    []int{4, 2, 2, 1},
			expected: []int{1, 2, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, result)
			}
		})
	}
}
