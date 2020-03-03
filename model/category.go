package model

import (
	"fmt"
)

type Category struct {
	ID          int
	URL         string
	ImageURL    string
	Title       string
	Description string
}

var categories = []Category{
	Category{
		ID:       1,
		URL:      "/shop_details",
		ImageURL: "lemon.png",
		Title:    "Juices and Mixes",
		Description: `Explore our wide assortment of juices and mixes expected by 
		today's lemonade stand clientelle. Now featuring a full line of 
		organic juices that are, guaranteed to be obtained from trees that 
		have never been treated with pesticides or artificial fertilizers.`,
	},
	Category{
		ID:       2,
		URL:      ".",
		ImageURL: "kiwi.png",
		Title:    "Cups, Straws, and Other Supplies",
		Description: `From paper cups to bio-degradable plastic to straws 
		and napkins, LSS is your source for the sundries that keep your 
		stand running smoothly.`,
	},
	Category{
		ID:       3,
		URL:      ".",
		ImageURL: "pineapple.png",
		Title:    "Signs and Advertising",
		Description: `Sure, you could just wait for people to find your stand
		along the side of the road, but if you want to take it to the next
		level, our premium line of advertising supplies.`,
	},
}

func GetCategories() []Category {
	return categories
}

func GetCategory(id int) (*Category, error) {
	for _, category := range categories {
		if category.ID == id {
			return &category, nil
		}
	}

	return nil, fmt.Errorf("CategoryID: %v not found", id)
}
