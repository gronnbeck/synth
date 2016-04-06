package synthesize

import "reflect"

func leftContains(left map[string]interface{}, right map[string]interface{}) bool {
	isLeftContains := true
	for k, v := range left {

		if reflect.TypeOf(v) != reflect.TypeOf(right[k]) {
			return false
		}

		switch v.(type) {
		case map[string]interface{}:
			isLeftContains = isLeftContains &&
				leftContains(v.(map[string]interface{}), right[k].(map[string]interface{}))
		default:
			isLeftContains = isLeftContains && reflect.DeepEqual(v, right[k])
		}
	}
	return isLeftContains
}
