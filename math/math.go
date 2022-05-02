package math

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type IntArray struct {
	Array      []int `json:"array"`
	Quantifier int   `json:"Quantifier"`
}

func Minimum(c *fiber.Ctx) error {
	var array IntArray
	c.JSON(array)

	var inputArray []int

	inputArray = append(inputArray, array.Array...)

	sort.Ints(inputArray)
	var minimums []int
	for _, value := range inputArray {
		if len(minimums) > 0 {
			// If the minimum array is greater than 0 (after the first time)

			// If the current value of inputArray is greater than the minimum value, then break
			if minimums[0] < value {
				break
			} else {
				// If the current value of inputArray is equal to the minimum value, then append
				minimums = append(minimums, value)
			}
		} else {
			// First value gets appended to minimums
			minimums = append(minimums, value)
		}
	}

	return c.SendString(arrayToString(minimums, ","))
}

func Maximum(c *fiber.Ctx) error {
	return c.SendString("Hello, Maximum")
}

func Average(c *fiber.Ctx) error {
	return c.SendString("Hello, Average")
}

func Median(c *fiber.Ctx) error {
	return c.SendString("Hello, Median")
}

func Percentile(c *fiber.Ctx) error {
	return c.SendString("Hello, Percentile")
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
