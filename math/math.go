package math

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type FloatArray struct {
	Array      []float64 `json:"array"`
	Quantifier int       `json:"quantifier"`
}

func Minimum(c *fiber.Ctx) error {

	array := new(FloatArray)
	array.Array = append(array.Array, 1, 1, -1, 6, 5, 6, 7, 8, 9, 10)
	array.Quantifier = 30

	// if err := c.BodyParser(array); err != nil {
	// 	c.Status(503).SendString(err.Error())
	// 	return err
	// }

	arrayLength := len(array.Array)

	if array.Quantifier > arrayLength {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier is larger than the size of the input array")
	} else if array.Quantifier < 1 {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier should be larger than 0 and non-null")
	} else if arrayLength < 1 {
		return fmt.Errorf("array should be larger than 0 and non-null")
	}

	c.JSON(array)

	// Appends the JSON array into a new input array to be sorted
	var inputArray []float64
	var outputArray []float64
	inputArray = append(inputArray, array.Array...)

	// Quick sort algorithm
	sort.Float64s(inputArray)

	// Writes quantifier length of items to the output array
	for i := 0; i < array.Quantifier; i++ {
		outputArray = append(outputArray, inputArray[i])
	}

	// Returns formatted string of output array as http response
	return c.SendString(arrayToString(outputArray, ", "))
}

func Maximum(c *fiber.Ctx) error {

	array := new(FloatArray)
	array.Array = append(array.Array, 1, 1, -1, 6, 5, 6, 7, 8, 9, 10)
	array.Quantifier = 1

	// if err := c.BodyParser(array); err != nil {
	// 	c.Status(503).SendString(err.Error())
	// 	return err
	// }

	arrayLength := len(array.Array)

	if array.Quantifier > arrayLength {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier is larger than the size of the input array")
	} else if array.Quantifier < 1 {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier should be larger than 0 and non-null")
	} else if arrayLength < 1 {
		return fmt.Errorf("array should be larger than 0 and non-null")
	}

	c.JSON(array)

	// Appends the JSON array into a new input array to be sorted
	var inputArray []float64
	var outputArray []float64
	inputArray = append(inputArray, array.Array...)

	// Quick sort algorithm
	sort.Float64s(inputArray)

	// Writes quantifier length of items to the output array
	for i := len(inputArray) - 1; i > len(inputArray)-array.Quantifier-1; i-- {
		outputArray = append(outputArray, inputArray[i])
	}

	// Returns formatted string of output array as http response
	return c.SendString(arrayToString(outputArray, ", "))
}

func Average(c *fiber.Ctx) error {
	array := new(FloatArray)
	array.Array = append(array.Array, 1, 1, -1, 6, 5, 6, 7, 8, 9, 10)

	// if err := c.BodyParser(array); err != nil {
	// 	c.Status(503).SendString(err.Error())
	// 	return err
	// }

	arrayLength := len(array.Array)

	if array.Quantifier > arrayLength {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier is larger than the size of the input array")
	} else if array.Quantifier < 1 {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier should be larger than 0 and non-null")
	} else if arrayLength < 1 {
		return fmt.Errorf("array should be larger than 0 and non-null")
	}

	c.JSON(array)

	// Appends the JSON array into a new input array to be sorted
	var sum float64

	for _, value := range array.Array {
		sum += value
	}

	avg := sum / float64(arrayLength)

	// Returns formatted string of output array as http response
	return c.SendString(fmt.Sprint(avg))
}

func Median(c *fiber.Ctx) error {
	array := new(FloatArray)
	array.Array = append(array.Array, 1, 1, -1, 6, 5, 6, 7, 8, 9, 10, 11, 13, 301)

	// if err := c.BodyParser(array); err != nil {
	// 	c.Status(503).SendString(err.Error())
	// 	return err
	// }

	arrayLength := len(array.Array)

	if arrayLength < 1 {
		// Returns a 500 with error code
		return fmt.Errorf("array should be larger than 0 and non-null")
	}

	c.JSON(array)

	// Appends the JSON array into a new input array to be sorted
	var inputArray []float64
	inputArray = append(inputArray, array.Array...)

	// Quick sort algorithm
	sort.Float64s(inputArray)

	// Checking to see if the length of the array is even or odd
	isOdd := arrayLength%2 > 0
	var median float64

	if isOdd {
		// This will return an integer even though function returns a float64
		middle := math.Floor(float64(arrayLength) / 2)
		median = inputArray[int(middle)]
	} else {
		middle := arrayLength / 2
		sum := inputArray[middle]
		sum += inputArray[middle+1]
		median = sum / 2
	}

	// Returns formatted string of output array as http response
	return c.SendString(fmt.Sprint(median))
}

func Percentile(c *fiber.Ctx) error {
	array := new(FloatArray)
	array.Array = append(array.Array, 1, 1, -1, 6, 5, 6, 7, 8, 9, 10, 11, 13, 301)
	array.Quantifier = 100

	// if err := c.BodyParser(array); err != nil {
	// 	c.Status(503).SendString(err.Error())
	// 	return err
	// }

	arrayLength := len(array.Array)

	if arrayLength < 1 {
		return fmt.Errorf("array should be larger than 0 and non-null")
	}

	c.JSON(array)

	// Appends the JSON array into a new input array to be sorted
	var inputArray []float64
	inputArray = append(inputArray, array.Array...)

	// Quick sort algorithm
	sort.Float64s(inputArray)

	// Ordinal Rank formula and uses golang math.ceiling function. Returns float but will be an integer
	ordinalRank := int(math.Ceil(((float64(array.Quantifier) / 100) * float64(arrayLength))))

	// Edgecase where we are grabbing the 0th percentile
	if ordinalRank == 0 {
		ordinalRank = 1
	}

	// Returns formatted string as http response by grabbing index of the ordinal rank (x - 1)
	return c.SendString(fmt.Sprint(inputArray[ordinalRank-1]))
}

func arrayToString(a []float64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func jsonToStruct(c *fiber.Ctx, omitQuantifer bool) ([]int64, int64, error) {

	array := new(FloatArray)

	if err := c.BodyParser(array); err != nil {
		return nil, -1, err
	}

	arrayStrings := strings.Split(array.Array, ",")

	if len(arrayStrings) < 1 {
		return nil, -1, errors.New("array json is null")
	}

	var arrayValues []int64
	var err error

	for _, value := range arrayStrings {
		value = strings.Trim(value, " ")
		parsedInt, _ := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return nil, -1, errors.New(fmt.Sprint("Could not parse: ", value, " - ", err.Error()))
		} else if value == "" {
			return nil, -1, errors.New("input array is malformed")
		}

		arrayValues = append(arrayValues, parsedInt)
	}

	if !omitQuantifer {

		if quantifier, err := array.Quantifier.Int64(); err != nil {
			return nil, -1, err
		} else {
			return arrayValues, quantifier, nil
		}
	} else {
		return arrayValues, -1, nil
	}
}
