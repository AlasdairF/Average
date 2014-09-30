package avg

import (
 "sort"
)

type keyVal struct {
	k int
	v int
}
type sorter []keyVal
func (a sorter ) Len() int           { return len(a) }
func (a sorter ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter ) Less(i, j int) bool { return a[i].v > a[j].v } // descending order

func MeanInt(ar []int) int {
	if len(ar) == 0 {
		return 0
	}
	var avg int
	for _, v := range ar {
		avg += v
	}
	return avg / len(ar)
}

func MeanFloat(ar []float64) float64 {
	if len(ar) == 0 {
		return 0
	}
	var avg float64
	for _, v := range ar {
		avg += v
	}
	return avg / float64(len(ar))
}

func TrimmedMeanInt(ar []int) int {
	sort.Ints(ar)
	a := len(ar)/5
	return MeanInt(ar[a:len(ar)-a])
}

func TrimmedMeanFloat(ar []float64) float64 {
	sort.Float64s(ar)
	a := len(ar)/5
	return MeanFloat(ar[a:len(ar)-a])
}

func MinInt(ar []int) int {
	if len(ar) == 0 {
		return 0
	}
	var m int = 9223372036854775807
	for _, v := range ar {
		if v < m {
			m = v
		}
	}
	return m
}

func MinFloat(ar []float64) float64 {
	if len(ar) == 0 {
		return 0
	}
	var m float64 = 9223372036854775807
	for _, v := range ar {
		if v < m {
			m = v
		}
	}
	return m
}

func MaxInt(ar []int) (m int) {
	for _, v := range ar {
		if v > m {
			m = v
		}
	}
	return
}

func MaxFloat(ar []float64) (m float64) {
	for _, v := range ar {
		if v > m {
			m = v
		}
	}
	return
}

func TotalInt(ar []int) (n int) {
	for _, v := range ar {
		n += v
	}
	return
}

func TotalFloat(ar []float64) (n float64) {
	for _, v := range ar {
		n += v
	}
	return
}

// Returns the key & value just after the greatest increase
func JumpPoint(ar []int) (int, int) {
	if len(ar) == 0 {
		return 0, 0
	}
	sort.Ints(ar)
	a := make(sorter, len(ar))
	last := ar[0]
	for i, v := range ar {
		a[i] = keyVal{i, v - last}
		last = v
	}
	sort.Stable(a)
	return a[0].k, ar[a[0].k]
}
