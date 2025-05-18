package converter

import (
	"errors"
	"fmt"
	"strings"
)

type Conversor struct {
	From      string
	FromValue float64
	To        string
}

var formulas = map[string]map[string]func(float64) float64{
	"meters": {
		"kilometers": func(from float64) float64 { return from / 1000 },
	},
	"kilometers": {
		"meters": func(from float64) float64 { return from * 1000 },
	},
	"grams": {
		"kilograms": func(from float64) float64 { return from / 1000 },
	},
	"kilograms": {
		"grams": func(from float64) float64 { return from * 1000 },
	},
}

var unitCategories = map[string]string{
	"meters":     "distance",
	"kilometers": "distance",
	"grams":      "weight",
	"kilograms":  "weight",
}

func NewConversor(from string, fromValue float64, to string) (Conversor, error) {
	from = strings.ToLower(from)
	to = strings.ToLower(to)

	if _, ok := unitCategories[from]; !ok {
		return Conversor{}, errors.New("unit 'from' not found")
	}

	if _, ok := unitCategories[to]; !ok {
		return Conversor{}, errors.New("unit 'to' not found")
	}

	if unitCategories[from] != unitCategories[to] {
		return Conversor{}, fmt.Errorf("%s only can be converted in %s", from, unitsFromCategory(from))
	}

	return Conversor{
		From:      from,
		FromValue: fromValue,
		To:        to,
	}, nil
}

func (c *Conversor) Result() (float64, error) {
	if err := c.formulaExists(c.From, c.To); err != nil {
		return .0, err
	}

	return formulas[c.From][c.To](c.FromValue), nil
}

func (c *Conversor) formulaExists(from, to string) error {
	errMsg := errors.New("formula not found")

	if _, exists := formulas[from]; !exists {
		return errMsg
	}

	if _, exists := formulas[from][to]; !exists {
		return errMsg
	}

	return nil
}

func ConversorAvailableFormulas() map[string][]string {
	listFormulas := make(map[string][]string)

	for from, target := range formulas {
		for to := range target {
			listFormulas[from] = append(listFormulas[from], to)
		}
	}

	return listFormulas
}

func unitsFromCategory(unit string) []string {
	unitsList := []string{}

	for unt, cat := range unitCategories {
		if (cat == unitCategories[unit]) && unt != unit {
			unitsList = append(unitsList, unt)
		}
	}

	return unitsList
}
