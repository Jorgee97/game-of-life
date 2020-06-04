package main

import (
	"reflect"
	"testing"
)

func TestDeadState(t *testing.T) {
	width, height := 3, 3
	expected := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	actual := DeadState(width, height)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("DeadState failed, expected %v but got %v", expected, actual)
	}
}

func TestNextBoardState(t *testing.T) {
	tests := []struct {
		board [][]int
		want  [][]int
	}{
		{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{[][]int{{1, 0, 1}, {0, 1, 0}, {0, 1, 1}}, [][]int{{0, 1, 0}, {1, 0, 0}, {0, 1, 1}}},
		{[][]int{{0, 0, 1}, {0, 1, 1}, {0, 0, 0}}, [][]int{{0, 1, 1}, {0, 1, 1}, {0, 0, 0}}},
		{[][]int{
			{1, 1, 0, 0, 1},
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 1},
			{1, 1, 0, 1, 0},
			{1, 1, 1, 1, 1},
		},
			[][]int{
				{1, 1, 1, 0, 0},
				{0, 0, 1, 1, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0},
				{1, 0, 0, 1, 1},
			},
		},
	}

	for _, tc := range tests {
		got := NextBoardState(tc.board)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("Nextboard failed, expected: %v but got: %v", tc.want, got)
		}
	}
}
