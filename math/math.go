package math

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type FloatArray struct {
	Array      string      `json:"array"`
	Quantifier json.Number `json:"quantifier,omitempty"`
}

func Minimum(c *fiber.Ctx) error {

	arrayValues, quantifier, err := jsonToStruct(c, false)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	arrayLength := len(arrayValues)

	if quantifier > int64(arrayLength) {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier is larger than the size of the input array")
	} else if quantifier < 1 {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier should be larger than 0 and non-null")
	} else if arrayLength < 1 {
		return fmt.Errorf("array should be larger than 0 and non-null")
	}

	// Appends the JSON array into a new input array to be sorted
	var outputArray []int64

	// Quick sort algorithm using function comparitor
	sort.Slice(arrayValues, func(i, j int) bool { return arrayValues[i] < arrayValues[j] })

	// Writes quantifier length of items to the output array
	for i := 0; int64(i) < quantifier; i++ {
		outputArray = append(outputArray, arrayValues[i])
	}

	// Returns formatted string of output array as http response
	return c.SendString(arrayToString(outputArray, ", "))

}

func Maximum(c *fiber.Ctx) error {

	arrayValues, quantifier, err := jsonToStruct(c, false)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	arrayLength := len(arrayValues)

	if quantifier > int64(arrayLength) {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier is larger than the size of the input array")
	} else if quantifier < 1 {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier should be larger than 0 and non-null")
	} else if arrayLength < 1 {
		return fmt.Errorf("array should be larger than 0 and non-null")
	}

	// Appends the JSON array into a new input array to be sorted
	var outputArray []int64

	// Quick sort algorithm using function comparitor
	sort.Slice(arrayValues, func(i, j int) bool { return arrayValues[i] > arrayValues[j] })

	// Writes quantifier length of items to the output array
	for i := 0; int64(i) < quantifier; i++ {
		outputArray = append(outputArray, arrayValues[i])
	}

	// Returns formatted string of output array as http response
	return c.SendString(arrayToString(outputArray, ", "))
}

func Average(c *fiber.Ctx) error {
	arrayValues, _, err := jsonToStruct(c, true)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	arrayLength := len(arrayValues)

	// Appends the JSON array into a new input array to be sorted
	var sum float64

	for _, value := range arrayValues {
		sum += float64(value)
	}

	avg := sum / float64(arrayLength)

	// Returns formatted string of output array as http response
	return c.SendString(fmt.Sprint(avg))
}

func Median(c *fiber.Ctx) error {
	arrayValues, _, err := jsonToStruct(c, true)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	arrayLength := len(arrayValues)

	// Quick sort algorithm using function comparitor
	sort.Slice(arrayValues, func(i, j int) bool { return arrayValues[i] > arrayValues[j] })

	// Checking to see if the length of the array is even or odd
	isOdd := arrayLength%2 > 0
	var median float64

	if isOdd {
		// This will return an integer even though function returns a int64
		middle := math.Floor(float64(arrayLength) / 2)
		median = float64(arrayValues[int(middle)])
	} else {
		middle := arrayLength / 2
		sum := arrayValues[middle]
		sum += arrayValues[middle+1]
		median = float64(sum) / 2
	}

	// Returns formatted string of output array as http response
	return c.SendString(fmt.Sprint(median))
}

func Percentile(c *fiber.Ctx) error {

	arrayValues, quantifier, err := jsonToStruct(c, false)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if quantifier > 100 {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier percentile is larger than max 100")
	} else if quantifier < 0 {
		// Returns a 500 with error code
		return fmt.Errorf("quantifier percentile is smaller than min 0")
	}

	arrayLength := len(arrayValues)

	// Quick sort algorithm using function comparitor
	sort.Slice(arrayValues, func(i, j int) bool { return arrayValues[i] < arrayValues[j] })

	// Ordinal Rank formula and uses golang math.ceiling function. Returns float but will be an integer
	ordinalRank := int(math.Ceil(((float64(quantifier) / 100) * float64(arrayLength))))

	// Edgecase where we are grabbing the 0th percentile
	if ordinalRank == 0 {
		ordinalRank = 1
	}

	// Returns formatted string as http response by grabbing index of the ordinal rank (x - 1)
	return c.SendString(fmt.Sprint(arrayValues[ordinalRank-1]))
}

func arrayToString(a []int64, delim string) string {
	// Essentially pretty prints the string
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func jsonToStruct(c *fiber.Ctx, omitQuantifer bool) ([]int64, int64, error) {

	// Create new object of FloatArray Struct
	array := new(FloatArray)

	// Prase the input JSON from fiber BASED on the Accept header and default JSON Marshaller
	if err := c.BodyParser(array); err != nil {
		return nil, -1, err
	}

	// Split input array by ,
	arrayStrings := strings.Split(array.Array, ",")

	if len(arrayStrings) < 1 {
		return nil, -1, errors.New("array json is null")
	}

	var arrayValues []int64
	var err error

	// For loop will take every split string value in arrayStrings and try and format it into a int64 type
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

	// If not an api call that needs the quantifier, skip parsing that JSON
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
