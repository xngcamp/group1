package main

import (
	"camp/movieData/customer"
	"camp/movieData/movie"
	"camp/movieData/rental"
	"fmt"
)

func main() {
	aCustomer := new(customer.Customer)
	aCustomer.Init("Tom")

	rental1 := &rental.Rental{
		AMovie: movie.RegularMovie{
			movie.Param{Title: "The Godfather", PriceCode: movie.REGULAR},
		},
		DaysRented: 5,
	}
	aCustomer.AddRental(rental1)

	rental2 := &rental.Rental{
		AMovie: movie.NewIssueMovie{
			movie.Param{Title: "Fast & Furious", PriceCode: movie.NEW_ISSUE},
		},
		DaysRented: 1,
	}
	aCustomer.AddRental(rental2)

	result := aCustomer.Statement()
	fmt.Println(result)
}