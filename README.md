##Average

Highly optimized functions for statistical analysis.

StdDev returns both standard deviation and mean, just in case the mean is useful. The standard deviation function is extremely highly optimized.

`TrimmedMeanInt` is particularly useful for when a real-world average is needed.

JumpPoint returns the original index and value at the point of greatest increase. This was originally developed while writing Frankenstein DBR in an effort to separate pages into two groups depending on their level of content, but in the end it wasn't used.

##Index

    func MeanInt(ar []int) int
    func MedianInt(ar []int) int
    func ModeInt(ar []int)
    func TrimmedMeanInt(ar []int) int
    func MinInt(ar []int) int
    func MaxInt(ar []int) (m int)
    func SumInt(ar []int) (n int)
    
    func MeanFloat(ar []float64) float64
    func MedianFloat(ar []float64)
    func ModeFloat(ar []float64) float64
    func TrimmedMeanFloat(ar []float64)
    func MinFloat(ar []float64) float64
    func MaxFloat(ar []float64) (m float64)
    func SumFloat(ar []float64) (n float64)
    
    func StdDev(a []float64) (float64, float64)
    func JumpPoint(ar []int) (int, int)
