/*
Exercise 5: Climate Data Analysis
Topics Covered: Go Arrays, Go Strings, Go Type Casting, Go Functions, Go Conditions,
Go Loops
Case Study:
You are tasked with analyzing climate data from multiple cities. The data includes the
city name, average temperature (°C), and rainfall (mm).
1. Data Input: Create a slice of structs to store data for each city. Input data can be
hardcoded or taken from the user.
2. Highest and Lowest Temperature: Write functions to find the city with the
highest and lowest average temperatures. Use conditions for comparison.
3. Average Rainfall: Calculate the average rainfall across all cities using loops. Use
type casting if necessary.
4. Filter Cities by Rainfall: Use loops to display cities with rainfall above a certain
threshold. Prompt the user to enter the threshold value.
5. Search by City Name: Allow users to search for a city by name and display its
data.
Bonus:
• Add error handling for invalid city names and invalid input for thresholds.
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

type City struct {
	Name     string
	AvgTemp  float64
	Rainfall float64
}

func findHighestTemperature(cities []City) (City, error) {
	if len(cities) == 0 {
		return City{}, errors.New("no data available")
	}
	highestTempCity := cities[0]
	for _, city := range cities {
		if city.AvgTemp > highestTempCity.AvgTemp {
			highestTempCity = city
		}
	}
	return highestTempCity, nil
}

func findLowestTemperature(cities []City) (City, error) {
	if len(cities) == 0 {
		return City{}, errors.New("no data available")
	}
	lowestTempCity := cities[0]
	for _, city := range cities {
		if city.AvgTemp < lowestTempCity.AvgTemp {
			lowestTempCity = city
		}
	}
	return lowestTempCity, nil
}

func calculateAverageRainfall(cities []City) float64 {
	if len(cities) == 0 {
		return 0
	}
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []City, threshold float64) {
	fmt.Printf("Cities with rainfall above %.1f mm:\n", threshold)
	found := false
	for _, city := range cities {
		if city.Rainfall > threshold {
			fmt.Printf("- %s: %.1f mm\n", city.Name, city.Rainfall)
			found = true
		}
	}
	if !found {
		fmt.Println("No cities found with rainfall above the threshold.")
	}
}

func searchCityByName(cities []City, name string) (City, error) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return city, nil
		}
	}
	return City{}, errors.New("city not found")
}

func main() {
	cities := []City{
		{"New York", 22.3, 120.5},
		{"Mumbai", 29.4, 254.0},
		{"Tokyo", 18.5, 134.6},
		{"Paris", 15.2, 89.3},
		{"Cape Town", 17.8, 65.0},
	}

	var choice int
	for {
		fmt.Println("Welcome to the Climate Data Analysis System!")
		fmt.Println("Choose an option:")
		fmt.Println("1. Find city with highest temperature")
		fmt.Println("2. Find city with lowest temperature")
		fmt.Println("3. Calculate average rainfall")
		fmt.Println("4. Filter cities by rainfall threshold")
		fmt.Println("5. Search for a city by name")
		fmt.Println("6. Exit")
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			highest, err := findHighestTemperature(cities)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("City with highest temperature: %s (%.1f°C)\n", highest.Name, highest.AvgTemp)
			}
		case 2:
			lowest, err := findLowestTemperature(cities)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("City with lowest temperature: %s (%.1f°C)\n", lowest.Name, lowest.AvgTemp)
			}
		case 3:
			avgRainfall := calculateAverageRainfall(cities)
			fmt.Printf("Average rainfall across all cities: %.2f mm\n", avgRainfall)
		case 4:
			var threshold float64
			fmt.Print("Enter rainfall threshold (mm): ")

			if _, err := fmt.Scanln(&threshold); err != nil {
				fmt.Println("Invalid input. Please enter a numeric value.")
				break
			}

			if threshold < 0 {
				fmt.Println("Rainfall threshold cannot be negative.")
				break
			}
			filterCitiesByRainfall(cities, threshold)
		case 5:
			var cityName string
			fmt.Print("Enter city name: ")
			fmt.Scanln(&cityName)
			city, err := searchCityByName(cities, cityName)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("City: %s, Avg Temp: %.1f°C, Rainfall: %.1f mm\n", city.Name, city.AvgTemp, city.Rainfall)
			}
		case 6:
			fmt.Println("Exiting the Climate Data Analysis System. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please choose a Valid Option from Menu.")
		}
	}
}
