package rental

import "gothinking/refactoring/videostore5/movie"

/*
Rental表示一个顾客租借电影

*/
type Rental struct {
	aMovie     movie.Movie
	daysRented int //租期
}

func (r *Rental) Init(movie movie.Movie, daysRented int) {
	r.aMovie = movie
	r.daysRented = daysRented
}

func (r Rental) GetDaysRented() int {
	return r.daysRented
}

func (r Rental) GetMovie() movie.Movie {
	return r.aMovie
}

/*
func (r Rental) GetCharge() float64 {
	thisAmount := 0.0
	// determine amounts for each record ,  各种片子价格不同
	aMovie := r.GetMovie()
	switch aMovie.GetPriceCode() {
	case movie.REGULAR:
		thisAmount += 2
		if r.daysRented > 2 {
			thisAmount += float64(r.daysRented-2.0) * 1.5
		}
	case movie.NEW_RELEASE:
		thisAmount += float64(r.daysRented) * 3
	case movie.CHILDREN:
		thisAmount += 1.5
		if r.daysRented > 3 {
			thisAmount += float64(r.daysRented-3) * 1.5
		}
	}
	return thisAmount
}
*/

func (r Rental) GetCharge() float64 {
	return r.GetMovie().GetCharge(r.daysRented)
}

/*
func (r Rental) GetFrequentRenterPoints() int {
	if r.GetMovie().GetPriceCode() == movie.NEW_RELEASE && r.GetDaysRented() > 1 {
		return 2
	} else {
		return 1
	}
}
*/

func (r Rental) GetFrequentRenterPoints() int {
	return r.GetMovie().GetFrequentRenterPoints(r.daysRented)
}
