package fullstory

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func writeKeyValue(k string, v interface{}) string {
	switch val := v.(type) {
	case bool:
		return fmt.Sprintf("%s_bool: %t,\n", k, val)
	case float32, float64:
		return fmt.Sprintf("%s_real: %f,\n", k, val)
	case int, int32, int64:
		return fmt.Sprintf("%s_int: %d,\n", k, val)
	case string:
		return fmt.Sprintf("%s_str: %q,\n", k, val)
	case time.Time:
		return fmt.Sprintf("%s_date: new Date(%q),\n", k, val.Format(time.RFC3339))

	// list types
	case []bool:
		bools := make([]string, len(v.([]bool)))
		for i, b := range v.([]bool) {
			bools[i] = strconv.FormatBool(b)
		}
		return fmt.Sprintf("%s_bools: [%s],\n", k, strings.Join(bools, ","))

	case []float32:
		reals := make([]string, len(v.([]float32)))
		for i, r := range v.([]float32) {
			reals[i] = fmt.Sprintf("%f", r)
		}
		return fmt.Sprintf("%s_reals: [%s],\n", k, strings.Join(reals, ","))

	case []float64:
		reals := make([]string, len(v.([]float64)))
		for i, r := range v.([]float64) {
			reals[i] = fmt.Sprintf("%f", r)
		}
		return fmt.Sprintf("%s_reals: [%s],\n", k, strings.Join(reals, ","))

	case []int:
		ints := make([]string, len(v.([]int)))
		for i, i2 := range v.([]int) {
			ints[i] = fmt.Sprintf("%d", i2)
		}
		return fmt.Sprintf("%s_ints: [%s],\n", k, strings.Join(ints, ","))

	case []int32:
		ints := make([]string, len(v.([]int32)))
		for i, i2 := range v.([]int32) {
			ints[i] = fmt.Sprintf("%d", i2)
		}
		return fmt.Sprintf("%s_ints: [%s],\n", k, strings.Join(ints, ","))

	case []int64:
		ints := make([]string, len(v.([]int64)))
		for i, i2 := range v.([]int64) {
			ints[i] = fmt.Sprintf("%d", i2)
		}
		return fmt.Sprintf("%s_ints: [%s],\n", k, strings.Join(ints, ","))

	case []string:
		strs := make([]string, len(v.([]string)))
		for i, s := range v.([]string) {
			strs[i] = fmt.Sprintf("%q", s)
		}
		return fmt.Sprintf("%s_strs: [%s],\n", k, strings.Join(strs, ","))

	case []time.Time:
		dates := make([]string, len(v.([]time.Time)))
		for i, d := range v.([]time.Time) {
			dates[i] = fmt.Sprintf("new Date(%q)", d.Format(time.RFC3339))
		}
		return fmt.Sprintf("%s_dates: [%s]n", k, strings.Join(dates, ","))

	}

	return ""
}

type Vars map[string]interface{}

func (v Vars) String() string {
	buf := strings.Builder{}
	for k, v := range v {
		buf.WriteString(writeKeyValue(k, v))
	}
	return buf.String()
}
