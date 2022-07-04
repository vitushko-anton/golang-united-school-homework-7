package coverage

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T) {

	testTable := []struct {
		people   People
		expected int
		err      string
	}{
		{
			people:   People{},
			expected: 0,
			err:      "Incorrect result. Expectation",
		},
		{
			people:   People{{}, {}, {}},
			expected: 3,
			err:      "Incorrect result. Expectation",
		},
	}

	for _, value := range testTable {
		if value.people.Len() != value.expected {
			t.Errorf("%s %d got %d", value.err, value.expected, value.people.Len())
		}
	}
}

func TestLess(t *testing.T) {
	var people People
	people = People{
		{
			"BBB",
			"BBB",
			time.Time{},
		},
		{
			"BBB",
			"BBB",
			time.Time{}.Add(5 * time.Minute),
		},
		{
			"AAA",
			"BBB",
			time.Time{},
		},
		{
			"AAA",
			"AAA",
			time.Time{},
		},
	}

	testTable := []struct {
		i   int
		j   int
		err string
	}{
		{
			i:   0,
			j:   1,
			err: "Incorrect result. Less by Birthday",
		},
		{
			i:   0,
			j:   2,
			err: "Incorrect result. Less by FirstName",
		},
		{
			i:   2,
			j:   3,
			err: "Incorrect result. Less by LastName",
		},
	}

	for _, value := range testTable {
		if people.Less(value.i, value.j) {
			t.Errorf("%s", value.err)
		}
	}
}

func TestSwap(t *testing.T) {
	var people People
	people = People{
		{
			"BBB",
			"BBB",
			time.Time{},
		},
		{
			"AAA",
			"AAA",
			time.Time{},
		},
	}
	people0 := people[0]
	people1 := people[1]
	people.Swap(0, 1)

	if people[0] != people1 || people[1] != people0 {
		t.Errorf("Wrong swap")
	}

}

/////

func TestNew(t *testing.T) {
	actual, err := New("---")
	if actual != nil || !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("Wrong string error")
	}

	actual, err = New("1 1 \n 2")
	if actual != nil || err.Error() != "Rows need to be the same length" {
		t.Errorf("Wrong matrix error")
	}

	actual, err = New("1 1 \n 2 3")
	expected := &Matrix{2, 2, []int{1, 1, 2, 3}}

	if actual.cols != expected.cols || actual.rows != expected.rows || !reflect.DeepEqual(actual.data, expected.data) {
		t.Errorf("Wrong empty matrix")
	}
}

func TestRows(t *testing.T) {
	matrix := &Matrix{2, 2, []int{4, 5, 6, 7}}
	expected := [][]int{{4, 5}, {6, 7}}

	actual := matrix.Rows()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrong rows matrix")
	}
}

func TestCols(t *testing.T) {
	matrix := &Matrix{3, 3, []int{1, 10, 100, 2, 20, 200, 3, 30, 300}}
	expected := [][]int{{1, 2, 3}, {10, 20, 30}, {100, 200, 300}}

	actual := matrix.Cols()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Wrong cols matrix")
	}
}

func TestSet(t *testing.T) {
	matrix := &Matrix{3, 3, []int{1, 10, 100, 2, 20, 200, 3, 30, 300}}
	expectedData := []int{1, 10, 100, 2, 20, 200, 3, 30, 10000000}

	actual := matrix.Set(2, 2, 10000000)

	if !actual || !reflect.DeepEqual(matrix.data, expectedData) {
		t.Errorf("Wrong set matrix")
	}

	actual = matrix.Set(-1, 2, 10000000)
	if actual {
		t.Errorf("Wrong set error")
	}
}
