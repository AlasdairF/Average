package avg

import (
 "github.com/AlasdairF/Sort/Float64"
 "github.com/AlasdairF/Sort/Int"
 "github.com/AlasdairF/Sort/IntInt"
 "math"
)

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

func MedianInt(ar []int) int {
	if len(ar) == 0 {
		return 0
	}
	sortInt.Asc(ar)
	return ar[len(ar) / 2]
}

func MedianFloat(ar []float64) float64 {
	if len(ar) == 0 {
		return 0
	}
	sortFloat64.Asc(ar)
	return ar[len(ar) / 2]
}

func ModeInt(ar []int) int {
	if len(ar) == 0 {
		return 0
	}
	mc := make(map[int]uint32)
	for _, v := range ar {
		mc[v]++
	}
	var mode int
	var best uint32
	for i, v := range mc {
		if v > best {
			best = v
			mode = i
		}
	}
	return mode
}

func ModeFloat(ar []float64) float64 {
	if len(ar) == 0 {
		return 0
	}
	mc := make(map[float64]uint32)
	for _, v := range ar {
		mc[v]++
	}
	var mode float64
	var best uint32
	for i, v := range mc {
		if v > best {
			best = v
			mode = i
		}
	}
	return mode
}

func TrimmedMeanInt(ar []int) int {
	newar := make([]int, len(ar))
	var l int
	for _, v := range ar {
		if v > 0 {
			newar[l] = v
			l++
		}
	}
	newar = newar[0:l]
	sortInt.Asc(newar)
	a := l / 5
	return MeanInt(newar[a:l-a])
}

func TrimmedMeanFloat(ar []float64) float64 {
	newar := make([]float64, len(ar))
	var l int
	for _, v := range ar {
		if v > 0 {
			newar[l] = v
			l++
		}
	}
	newar = newar[0:l]
	sortFloat64.Asc(newar)
	a := l / 5
	return MeanFloat(newar[a:l-a])
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

func SumInt(ar []int) (n int) {
	for _, v := range ar {
		n += v
	}
	return
}

func SumFloat(ar []float64) (n float64) {
	for _, v := range ar {
		n += v
	}
	return
}

func StdDev(a []float64) (float64, float64) {
	mean := MeanFloat(a)
	var d float64
	for _, v := range a {
		v -= mean
		d += v * v
	}
	stddev := math.Sqrt(d / float64(len(a)))
	return stddev, mean
}

// Returns the key & value just after the greatest increase
func JumpPoint(ar []int) (int, int) {
	if len(ar) == 0 {
		return 0, 0
	}
	sortInt.Asc(ar)
	a := make([]sortIntInt.KeyVal, len(ar))
	last := ar[0]
	for i, v := range ar {
		a[i] = sortIntInt.KeyVal{i, v - last}
		last = v
	}
	sortIntInt.StableDesc(a)
	return a[0].K, ar[a[0].K]
}
