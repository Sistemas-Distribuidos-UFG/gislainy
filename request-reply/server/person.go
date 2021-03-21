package main

import (
	"fmt"
)

// PersonLegalAge 
type PersonLegalAge struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Sex string `json:"sex"`
}

type PersonWeight struct {
	Sex string `json:"sex"`
	Height float64 `json:"height"`
}

func CheckLegalAge(person PersonLegalAge) Response {

	var response Response;

	isFemaleAndLegalAge := person.Sex == "Female" && person.Age >= 21
	isMaleAndLegalAge := person.Sex == "Male" && person.Age >= 18

	if (isFemaleAndLegalAge || isMaleAndLegalAge) {
		response.Message = fmt.Sprintf("%s is already of legal age", person.Name)
	} else {
		response.Message = fmt.Sprintf("%s is not legal age", person.Name)
	}

	return response
}

// CalculateIdealWeight return your ideal weight 
func CalculateIdealWeight(person PersonWeight) Response {
	var response Response;
	var weight float64
	switch person.Sex {
	case "Male":
		weight = (72.7 * person.Height) - 58.0
	case "Female":
		weight = (62.1 * person.Height) - 44.7
	}

	response.Message = fmt.Sprintf("Your ideal weight is %.3f", weight)

	return response
}