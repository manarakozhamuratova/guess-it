package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var k []float64
	for reader.Scan() {
		t := reader.Text()
		n, err := strconv.Atoi(t)
		if err != nil {
			fmt.Print(err)
			return
		}
		if len(k) > 2 && (n > int(average(k))*2 || n < int(average(k)*0.3)) {
			n = int(average(k))
		}
		k = append(k, float64(n))
		answer(k)
	}
}

func average(num []float64) float64 {
	var sum float64
	for _, number := range num {
		sum += number
	}
	return math.Round(sum / float64(len(num)))
}

func variance(n []float64, avg float64) float64 {
	var sum float64
	for _, i := range n {
		sum += math.Pow(i-avg, 2)
	}
	sd := sum / float64(len(n))
	return sd
}

func std(n []float64, avg float64) float64 {
	return math.Sqrt(float64(variance(n, avg)))
}

func answer(num []float64) {
	avg := average(num)
	div := std(num, avg)
	upper := avg + div*1.28
	lower := avg - div*1.28
	fmt.Println(lower, upper)
}
