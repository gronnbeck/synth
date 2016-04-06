package synthesize

import "testing"

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
