package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
)

func main() {
	content := readFile()

	fmt.Println("Average:", average(content))
	fmt.Println("Median:", median(content))
	variance := variance(content)
	fmt.Println("Variance:", (variance))
	fmt.Println("Standard Deviation:", (stdev(variance)))
}

func readFile() []int {
	if len(os.Args) < 2 {
		log.Fatal("Not enough input arguments!")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var dataNumbers []int

	for scanner.Scan() {
		if scanner.Text() != "" {
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			dataNumbers = append(dataNumbers, number)
		}
	}
	return dataNumbers
}

// Calculate average
func average(dataNumbers []int) int {
	var answer float64
	var sum float64
	var count float64

	for i := 0; i < len(dataNumbers); i++ {
		sum += float64(dataNumbers[i])
		count += 1
	}
	answer = (math.Round(sum / count))
	return int(answer)
}

// Calculate median
func median(dataNumbers []int) int {
	sort.Ints(dataNumbers) // sort numbers from smallest to largest
	var medianAnswer float64

	if len(dataNumbers)%2 != 0 { // if the number of data points is odd, the median is the middle value.
		medianAnswer = float64(dataNumbers[(len(dataNumbers) / 2)])
	} else { // If the number of data points is even, the median is the average of the two middle values.
		numberOne := float64(dataNumbers[len(dataNumbers)/2] - 2)
		numberTwo := float64(dataNumbers[(len(dataNumbers)/2)] - 1)
		medianAnswer = ((numberOne + numberTwo) / 2)
	}
	return int(math.Round(medianAnswer))
}

// Calculate variance
func variance(dataNumbers []int) float64 {
	// 1. Find the mean of the data set. Add up all the values and divide by the total number of values.
	sum := big.NewInt(0)
	for i := 0; i < len(dataNumbers); i++ {
		sum = sum.Add(sum, big.NewInt(int64(dataNumbers[i])))
	}
	mean := new(big.Float).Quo(new(big.Float).SetInt(sum), big.NewFloat(float64(len(dataNumbers))))

	// 2. For each value in the data set, subtract the mean and then square the result.
	var tempSlice []*big.Float
	for _, nb := range dataNumbers {
		tempNb := new(big.Float).Sub(new(big.Float).SetInt64(int64(nb)), mean)
		tempNb = tempNb.Mul(tempNb, tempNb)
		tempSlice = append(tempSlice, tempNb)
	}

	// 3. Add up all the squared differences in tempSlice
	sumTwo := new(big.Float).SetInt64(0)
	for _, nb := range tempSlice {
		sumTwo = sumTwo.Add(sumTwo, nb)
	}

	// 4. Divide the result from step 3 by the total number of values. This is the variance.
	varianceNumber := new(big.Float).Quo(sumTwo, big.NewFloat(float64(len(dataNumbers))))

	// 5. Convert the number to float64 and then round off.
	varianceNb, _ := varianceNumber.Float64()
	finalVarNb := math.Round(varianceNb)

	return finalVarNb
}

// Take the square root of the variance. This is the standard deviation.
func stdev(variance float64) float64 {
	return math.Round(math.Sqrt(variance))
}
