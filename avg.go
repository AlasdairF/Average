package avg

import (
 "sort"
)

type keyValFloat struct {
	k int
	v float32
}
type sorterFloat []keyValFloat
func (a sorterFloat) Len() int           { return len(a) }
func (a sorterFloat) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorterFloat) Less(i, j int) bool { return a[i].v > a[j].v } // descending order

func MeanInt(ar []int) int {
	var avg int
	for _, v := range ar {
		avg += v
	}
	return avg / len(ar)
}

func MeanFloat(ar []float64) float64 {
	var avg float64
	for _, v := range ar {
		avg += v
	}
	return avg / float64(len(ar))
}

func TrimmedMeanInt(ar []int) int {
	sort.Ints(ar)
	a := len(ar)/5
	var avg int
	for i:=a; i<len(ar)-a; i++ {
		avg += ar[i]
	}
	return avg / (len(ar) - (a*2))
}

func TrimmedMeanFloat(ar []float64) float64 {
	sort.Float64s(ar)
	a := len(ar)/5
	var avg float64
	for i:=a; i<len(ar)-a; i++ {
		avg += ar[i]
	}
	return avg / float64(len(ar) - (a*2))
}

func Min(ar []int) int {
	var m = 9223372036854775807
	for _, v := range ar {
		if v < m {
			m = v
		}
	}
	return m
}

func Max(ar []int) (m int) {
	for _, v := range ar {
		if v > m {
			m = v
		}
	}
	return
}

// Returns the value at which the ints jump from low to high value, by highest percentage, then deducts 1 so greater than can be used
func JumpPoint(ar []int) int {
	// Remove anything that is 0
	n := make([]int, len(ar))
	var count int
	for i:=0; i<len(ar); i++ {
		if ar[i]>0 {
			n[count] = ar[i]
			count++
		}
	}
	n = n[0:count]
	// Sort list by percentage difference
	a := make(sorterFloat, len(n)-1)
	for i:=0; i<len(n)-1; i++ {
		a[i] = keyValFloat{i, float32(n[i+1]) / float32(n[i])}
	}
	sort.Sort(a)
	// If the first jump point is in the latter 2/3 but the second jump point is in the first 1/3 then use the second jump point, but only if its no less than half the size of the first
	mid := len(n)/3
	if a[0].k>=mid && a[1].k<mid && a[1].k*2>a[0].k {
		return n[a[1].k]
	} else {
		return n[a[0].k]
	}
}
