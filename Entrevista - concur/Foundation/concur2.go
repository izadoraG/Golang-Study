package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//Print the best rated hotel names, a best rated hotel

//has a rating higher than 7.0 and more than 1000 reviews;

// Print the hotel names that offers parking as an amenity;

//Sort the dataset by the highest rating and secondly by
//the most amenities an hotel has, new keys in the sorted dataset are allowed
//but previous keys should not be changed. Expected output is to print the whole
//ordered dataset in a JSON format.

type Hotel struct {
	Name    string  `json:"name"`
	City    string  `json:"city"`
	State   string  `json:"state"`
	Country string  `json:"country"`
	Amenity string  `json:"amenity"`
	Rating  float32 `json:"rating"`
	Reviews string  `json:"reviews"`
}

func countAmenity(amenity string) int {
	if amenity == "" {
		return 0
	}
	return len(strings.Split(amenity, "|"))
}

func main() {
	var hotels []Hotel
	jsonData := `[
		{
			"name": "Blue Three Miami",
			"city": "Miami",
			"state": "Florida",
			"country": "US",
			"amenities": "Free Wifi|Garden|Cleaning services|Laundry|Breakfast",
			"rating": 7.3,
			"reviews": "1236"
		},
		{
			"name": "University Inn",
			"city": "San Jose",
			"state": "California",
			"country": "US",
			"amenities": "Free Wifi|Cleaning services|Bike tours|Garden|Breakfast|Parking",
			"rating": 7.3,
			"reviews": "5251"
		},
		{
			"name": "Sunset Lodge",
			"city": "Long Beach",
			"state": "California",
			"country": "US",
			"amenities": "Parking|Cleaning services|Laundry|Garden|Breakfast",
			"rating": 7.6,
			"reviews": "859"
		},
		{
			"name": "Four Seasons",
			"city": "Orlando",
			"state": "Florida",
			"country": "US",
			"amenities": "Parking",
			"rating": 6.7,
			"reviews": "1037"
		},
		{
			"name": "Sleep Inn",
			"city": "Austin",
			"state": "Texas",
			"country": "US",
			"amenities": "",
			"rating": 6.2,
			"reviews": "none"
		},
		{
			"name": "Mountain Plaza",
			"city": "Denver",
			"state": "Colorado",
			"country": "US",
			"amenities": "Cleaning services|Laundry|Garden|Breakfast|Wifi|Bike tours",
			"rating": 7.8,
			"reviews": "3876"
		}
	]`

	json.Unmarshal([]byte(jsonData), &hotels)

	bestRating := float32(0)
	hotelTop := ""

	for _, hotel := range hotels {
		if hotel.Rating > bestRating {
			bestRating = hotel.Rating
			hotelTop = hotel.Name

		}
	}
	fmt.Println(hotelTop, bestRating)

	for _, hotel := range hotels {
		num, err := strconv.Atoi(hotel.Reviews)
		if err != nil {
			continue
		}
		if hotel.Rating > 7.0 && num > 1000 {
			fmt.Println(hotel.Name)
		}
	}

	sort.SliceStable(hotels, func(i, j int) bool {
		if hotels[i].Rating == hotels[j].Rating {
			return countAmenity(hotels[i].Amenity) > countAmenity(hotels[j].Amenity)
		}
		return hotels[i].Rating > hotels[j].Rating
	})
	result, err := json.MarshalIndent(hotels, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))

}
