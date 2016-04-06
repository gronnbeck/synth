package synthesize

import (
	"encoding/json"
	"testing"
)

func Test_leftContains_simple_true(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = "test"

	r := map[string]interface{}{}
	r["test"] = "test"
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Should return true but returned false")
		t.Fail()
	}
}

func Test_leftContains_simple_false(t *testing.T) {
	l := map[string]interface{}{}
	l["tester"] = "test"

	r := map[string]interface{}{}
	r["test"] = "test"
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if contains {
		t.Log("Should return false but returned true")
		t.Fail()
	}
}

func Test_leftContains_complex_true(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = map[string]interface{}{
		"test": "test",
	}

	r := map[string]interface{}{}
	r["test"] = map[string]interface{}{
		"test": "test",
	}
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Should return true but returned false")
		t.Fail()
	}
}

func Test_leftContains_complex_false(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = map[string]interface{}{
		"test": "test",
	}

	r := map[string]interface{}{}
	r["test"] = "test"
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if contains {
		t.Log("Should return false but returned true")
		t.Fail()
	}
}

func Test_leftContains_string_array_true(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = []string{
		"hello",
		"world",
	}

	r := map[string]interface{}{}
	r["test"] = []string{
		"hello",
		"world",
	}
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Expected array to work but failed")
		t.Fail()
	}
}

func Test_leftContains_number_array_true(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = []float64{
		3.3,
		3,
	}

	r := map[string]interface{}{}
	r["test"] = []float64{
		3.3,
		3,
	}
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Expected array to work but failed")
		t.Fail()
	}
}

func Test_leftContains_complex_array(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = []map[string]interface{}{
		map[string]interface{}{
			"test": "test",
		},
	}

	r := map[string]interface{}{}
	r["test"] = []map[string]interface{}{
		map[string]interface{}{
			"test": "test",
		},
	}
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Expected complex arrays to work but failed")
		t.Fail()
	}
}

func Test_ExpectedResponse_Comparison(t *testing.T) {
	input := `{"hello": "world", "integer": 4, "float": 3.5}`

	right := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &right)

	if err != nil {
		t.Fatalf("Should be able to parse json but it failed")
	}

	left := map[string]interface{}{
		"hello":   "world",
		"integer": 4.0,
		"float":   3.5,
	}

	contains := leftContains(left, right)

	if !contains {
		t.Fatal("Left should be contained in right")
	}
}

func Test_ExpectedResponse_Comparison_Complex(t *testing.T) {
	input := `{"complex": {"hello": "world"}}`

	right := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &right)

	if err != nil {
		t.Fatalf("Should be able to parse json but it failed")
	}

	left := map[string]interface{}{
		"complex": map[string]interface{}{
			"hello": "world",
		},
	}

	contains := leftContains(left, right)

	if !contains {
		t.Fatal("Left should be contained in right")
	}
}

func Test_ExpectedResponse_Comparison_Array(t *testing.T) {
	input := `{"array": [1, 2, 3]}`

	right := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &right)

	if err != nil {
		t.Fatalf("Should be able to parse json but it failed")
	}

	left := map[string]interface{}{
		"array": []interface{}{1.0, 2.0, 3.0},
	}

	contains := leftContains(left, right)

	if !contains {
		t.Fatal("Left should be contained in right")
	}
}

func Test_ExpectedResponse_Comparison_NotEqual(t *testing.T) {
	input := `{"hello2": "world"}`

	right := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &right)

	if err != nil {
		t.Fatalf("Should be able to parse json but it failed")
	}

	left := map[string]interface{}{
		"hello": "world",
	}

	contains := leftContains(left, right)

	if contains {
		t.Fatal("Left should not be contained in right")
	}
}
