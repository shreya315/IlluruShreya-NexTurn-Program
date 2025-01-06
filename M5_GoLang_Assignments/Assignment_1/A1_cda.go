package main

import (
	"fmt"
	"strings"
)

type CityInfo struct {
	CityName string
	AvgTemp  float64
	Rainfall float64
}

func main() {
	cityData := []CityInfo{
		{"Hyderabad", 12.5, 690},
		{"Vijayawada", 18.9, 420},
		{"Karaikal", 23.3, 700},
		{"Pune", 19.3, 1350},
		{"nellore", 16.5, 890},
		{"Karnool", 20.6, 540},
	}

	hotCity, maxTemp := getCityWithHighestTemp(cityData)
	coldCity, minTemp := getCityWithLowestTemp(cityData)

	fmt.Printf("Hottest city: %s with %.2f°C\n", hotCity, maxTemp)
	fmt.Printf("Coldest city: %s with %.2f°C\n", coldCity, minTemp)

	avgRainfall := calculateAverageRainfall(cityData)
	fmt.Printf("Overall average rainfall: %.2f mm\n", avgRainfall)

	var rainThreshold float64
	fmt.Println("Enter rainfall limit (mm):")
	fmt.Scan(&rainThreshold)
	filterCitiesByRain(cityData, rainThreshold)

	var searchName string
	fmt.Println("Enter a city name to find details:")
	fmt.Scan(&searchName)
	findCityByName(cityData, searchName)
}

func getCityWithHighestTemp(data []CityInfo) (string, float64) {
	highestTemp := data[0].AvgTemp
	highestCity := data[0].CityName

	for _, city := range data {
		if city.AvgTemp > highestTemp {
			highestTemp = city.AvgTemp
			highestCity = city.CityName
		}
	}
	return highestCity, highestTemp
}

func getCityWithLowestTemp(data []CityInfo) (string, float64) {
	lowestTemp := data[0].AvgTemp
	lowestCity := data[0].CityName

	for _, city := range data {
		if city.AvgTemp < lowestTemp {
			lowestTemp = city.AvgTemp
			lowestCity = city.CityName
		}
	}
	return lowestCity, lowestTemp
}

func calculateAverageRainfall(data []CityInfo) float64 {
	var totalRainfall float64
	for _, city := range data {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(data))
}

func filterCitiesByRain(data []CityInfo, threshold float64) {
	fmt.Printf("Cities with annual rainfall exceeding %.2f mm:\n", threshold)
	for _, city := range data {
		if city.Rainfall > threshold {
			fmt.Printf("%s - %.2f mm\n", city.CityName, city.Rainfall)
		}
	}
}

func findCityByName(data []CityInfo, name string) {
	name = strings.Title(strings.ToLower(name))
	for _, city := range data {
		if strings.Title(strings.ToLower(city.CityName)) == name {
			fmt.Printf("Details for %s:\n", city.CityName)
			fmt.Printf("Average Temp: %.2f°C\n", city.AvgTemp)
			fmt.Printf("Rainfall: %.2f mm\n", city.Rainfall)
			return
		}
	}
	fmt.Println("No details found for the specified city.")
}
