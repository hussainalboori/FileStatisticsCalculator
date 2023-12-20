package main
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)
func main() {
	// Check if a file name is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name as an argument.")
		return
	}
	fileName := os.Args[1]
	// Calculate and display the average from the file
	average, err := findAverage(fileName)
	if err != nil {
		fmt.Printf("Error finding average: %v\n", err)
		return
	}
	// Calculate and display the median from the file
	median, err := findMedian(fileName)
	if err != nil {
		fmt.Printf("Error finding median: %v\n", err)
		return
	}
	// Calculate and display the variance from the file
	variance, err := findVariance(fileName)
	if err != nil {
		fmt.Printf("Error finding variance: %v\n", err)
		return
	}
	// Calculate and display the standard deviation from the file
	standardDeviation, err := findStandardDeviation(fileName)
	if err != nil {
		fmt.Printf("Error finding standard deviation: %v\n", err)
		return
	}
	// Print the calculated values
	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", standardDeviation)
}
// findAverage calculates the average of numbers in a file
func findAverage(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	count := 0
	// Read each line of the file
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		sum += num
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, fmt.Errorf("no numbers found in the file")
	}
	average := float64(sum) / float64(count)
	roundedAverage := int(math.Round(average))
	return roundedAverage, nil
}
// findMedian calculates the median of numbers in a file
func findMedian(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []int
	// Read each line of the file
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers found in the file")
	}
	sort.Ints(numbers) // Sort the numbers in ascending order
	if len(numbers)%2 == 1 {
		median := numbers[len(numbers)/2]
		return median, nil
	}
	median := int(math.Round(float64(numbers[len(numbers)/2-1]+numbers[len(numbers)/2]) / 2.0))
	return median, nil
}
// findVariance calculates the variance of numbers in a file
func findVariance(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []float64
	// Read each line of the file
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers found in the file")
	}
	mean := calculateMean(numbers)
	var sumOfSquares float64
	for _, num := range numbers {
		diff := num - mean
		sumOfSquares += diff * diff
	}
	variance := sumOfSquares / float64(len(numbers))
	roundedVariance := int(math.Round(variance))
	return roundedVariance, nil
}
// findStandardDeviation calculates the standard deviation of numbers in a file
func findStandardDeviation(fileName string) (int, error) {
	file, err :=os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []float64
	// Read each line of the file
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers found in the file")
	}
	mean := calculateMean(numbers)
	var sumOfSquares float64
	for _, num := range numbers {
		diff := num - mean
		sumOfSquares += diff * diff
	}
	variance := sumOfSquares / float64(len(numbers))
	standardDeviation := int(math.Round(math.Sqrt(variance)))
	return standardDeviation, nil
}
// calculateMean calculates the mean (average) of a list of numbers
func calculateMean(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}