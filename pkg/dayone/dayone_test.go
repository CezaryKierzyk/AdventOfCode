package dayone_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/CezaryKierzyk/AdventOfCode2024/pkg/dayone"
)

var (
	validData [2][]int = [2][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}
	validSorted [2][]int = [2][]int{
		{1, 2, 3, 3, 3, 4},
		{3, 3, 3, 4, 5, 9},
	}
	validSimilarIndices [2][]int = [2][]int{
		{0, 4, 5},
		{1, 3, 5},
	}

	testIndices             []int = []int{1, 2, 3}
	validReducedLeft        []int = []int{3, 3, 3}
	validLeftSmallestIndex        = 3
	validRightSmallestIndex       = 1
	validDistance                 = 11
	validSimilarity               = 31
)

func TestParseInput(t *testing.T) {
	t.Log("Testing ParseInput function")
	path := filepath.Join("testdata", "in_ex.txt")
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open test file:\n\tReason: %v \n", err)
	}
	defer file.Close()
	data := dayone.ParseInput(path)
	if !(len(data) != len(validData) ||
		len(data[0]) != len(validData[0]) ||
		len(data[1]) != len(validData[1]) ||
		reflect.DeepEqual(data, validData)) {
		t.Fatalf("File parsed incorrectly or data malformed:\n\tData: \t\t%v\n\tShould be: \t%v", data, validData)
	}

}

func TestFindSmallest(t *testing.T) {
	t.Log("Testing FindSmallest function")
	leftSmallestIndex := dayone.FindSmallest(validData[0])
	rightSmallestIndex := dayone.FindSmallest(validData[1])
	if rightSmallestIndex != validRightSmallestIndex || leftSmallestIndex != validLeftSmallestIndex {
		t.Fatalf(
			"Wrong smallest number index on left or right side of array:\n\t"+
				"left smallest: %v should be: %v\n\t"+
				"right smallest: %v should be: %v",
			leftSmallestIndex, validLeftSmallestIndex,
			rightSmallestIndex, validRightSmallestIndex,
		)
	}
}

func TestRemoveAtIndex(t *testing.T) {
	t.Log("Testing RemoveAtIndex function")
	reducedLeft := make([]int, len(validData[0]))
	reducedRight := make([]int, len(validData[0]))

	copy(reducedLeft, validData[0])
	copy(reducedRight, validData[1])

	reducedLeft = dayone.RemoveAtIndex(reducedLeft, 0)
	reducedRight = dayone.RemoveAtIndex(reducedRight, 0)

	if reducedLeft[0] != 4 || reducedRight[0] != 3 {
		t.Fatalf(
			"Failed to delete value at index 0 in left or right slice:\n\t"+
				"reduced left: %v unreduced left %v\n\t"+
				"reduced right: %v unreduced right: %v",
			reducedLeft, validData[0],
			reducedRight, validData[1],
		)
	}
}

func TestHasIndex(t *testing.T) {
	t.Log("Testing HasIndex function")
	for i := range testIndices {
		if !dayone.HasIndex(testIndices, testIndices[i]) {
			t.Fatalf("Index: %v not found in: %v where should be present", testIndices[i], testIndices)
		}
	}
}

func TestRemoveManyIndices(t *testing.T) {
	t.Log("Test RemoveManyIndices function")
	reducedLeft := make([]int, len(validData[0]))
	copy(reducedLeft, validData[0])
	reducedLeft = dayone.RemoveManyIndices(reducedLeft, testIndices)
	if !reflect.DeepEqual(reducedLeft, validReducedLeft) {
		t.Fatalf("Reduced slice: %v should equal: %v", reducedLeft, validReducedLeft)
	}
}

func TestSortInput(t *testing.T) {
	t.Log("Testing SortInput function")
	sorted := dayone.SortInput(validData)
	if !reflect.DeepEqual(sorted, validSorted) {
		t.Fatalf("Failed to sort data:\n\t Func sorted:\t %v\n\t Presorted:\t%v", sorted, validSorted)
	}
}

func TestCalculateDistance(t *testing.T) {
	t.Log("Testing CalculateDistance function")
	distance := dayone.CalculateDistance(validSorted)
	if distance != validDistance {
		t.Fatalf("Calculated distance: %v, should equal: %v", distance, validDistance)
	}
}

func TestFindSimilar(t *testing.T) {
	t.Log("Testing FindSimilar function")
	indices := dayone.FindSimilar(3, validData)
	if !reflect.DeepEqual(indices, validSimilarIndices) {
		t.Fatalf("Found indices: %v, should equal: %v", indices, validSimilarIndices)
	}
}

func TestCountWithRemove(t *testing.T) {
	t.Log("Testing CountWithRemove function")
	similarity := dayone.CountWithRemove(validData)
	if similarity != validSimilarity {
		t.Fatalf("Calculated similarity: %v, should equal: %v", similarity, validSimilarity)
	}
}
