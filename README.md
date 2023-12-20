# File Statistics Calculator

This Go program calculates various statistics of numbers stored in a file.

## Prerequisites
- Go 1.13 or higher

## Usage
To run the program, provide a file name as an argument when executing the compiled binary.

```shell
$ go run main.go <file_name>
```

## Features
The program calculates the following statistics:

- **Average**: Calculates the average of the numbers in the file.
- **Median**: Calculates the median of the numbers in the file.
- **Variance**: Calculates the variance of the numbers in the file.
- **Standard Deviation**: Calculates the standard deviation of the numbers in the file.

## Implementation Details
The program reads the file line by line and performs the necessary calculations using the provided functions.

### findAverage(fileName string) (int, error)
Calculates the average of the numbers in the file.

### findMedian(fileName string) (int, error)
Calculates the median of the numbers in the file.

### findVariance(fileName string) (int, error)
Calculates the variance of the numbers in the file.

### findStandardDeviation(fileName string) (int, error)
Calculates the standard deviation of the numbers in the file.

### calculateMean(numbers []float64) float64
Calculates the mean (average) of a list of numbers.

## Error Handling
If any errors occur during file reading or calculations, the program returns an error message.

## Example
Suppose we have a file named `numbers.txt` with the following contents:
```
3
1
4
1
5
9
```
To calculate the statistics of the numbers in this file, we can run the program as follows:

```shell
$ ./file-statistics-calculator numbers.txt
```

The program will output:
```
Average: 4
Median: 4
Variance: 6
Standard Deviation: 2
```

## License
This program is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.