package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Add calculations.
func Add(c *gin.Context) {
	var numbers Numbers
	err := c.Bind(&numbers)
	if err != nil {

	}
	total := 0
	for i := 0; i < len(numbers.Input); i++ {
		intValue, _ := strconv.Atoi(numbers.Input[i])
		total += intValue
	}
	result := strconv.Itoa(total)
	c.JSON(200, gin.H{
		"result": result,
	})
}

// Substract calculations.
func Substract(c *gin.Context) {
	var numbers Numbers
	err := c.Bind(&numbers)
	if err != nil {

	}
	total := 0
	for i := 0; i < len(numbers.Input); i++ {
		intValue, _ := strconv.Atoi(numbers.Input[i])
		total -= intValue
	}
	result := strconv.Itoa(total)
	c.JSON(200, gin.H{
		"result": result,
	})
}

// Times calculations.
func Times(c *gin.Context) {
	var numbers Numbers
	err := c.Bind(&numbers)
	if err != nil {

	}
	var total int
	for i := 0; i < len(numbers.Input); i++ {
		intValue, _ := strconv.Atoi(numbers.Input[i])
		if i == 0 {
			total = intValue
		}
		if i > 0 {
			total = total * intValue
		}
	}
	fmt.Println(total)
	result := strconv.Itoa(total)
	c.JSON(200, gin.H{
		"result": result,
	})
}

func divide(a *int, b int) {
	if b == 0 {
		return
	}
	*a = *a / b
}

// Divide calculations.
func Divide(c *gin.Context) {
	var numbers Numbers
	err := c.Bind(&numbers)
	if err != nil {

	}
	var total int
	for i := 0; i < len(numbers.Input); i++ {
		intValue, _ := strconv.Atoi(numbers.Input[i])
		if i == 0 {
			total = intValue
		}
		if i > 0 {
			divide(&total, intValue)
		}
	}
	result := strconv.Itoa(total)
	c.JSON(200, gin.H{
		"result": result,
	})
}

func factorial(a int) int {
	if a == 1 {
		return a
	}
	return a * factorial(a-1)
}

// Factorial calculations.
func Factorial(c *gin.Context) {
	var number Number
	err := c.Bind(&number)
	if err != nil {

	}
	input, _ := strconv.Atoi(number.Input)
	fmt.Println(input)
	result := strconv.Itoa(factorial(input))
	c.JSON(200, gin.H{
		"result": result,
	})
}

// Numbers of the process.
type Numbers struct {
	Input []string `json:"input"`
}

// Number of the process.
type Number struct {
	Input string `json:"input"`
}
