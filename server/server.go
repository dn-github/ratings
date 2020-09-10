package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/dn-github/ratings/pb"
)

type ratingsImpl struct{}

func NewRatingsImpl() *ratingsImpl {
	return &ratingsImpl{}
}

// Ratings ...
func (r *ratingsImpl) Ratings(ctx context.Context, book *pb.Book) (*pb.Rating, error) {
	randRating := int64(rand.Intn(5) + 1)
	logMessage := fmt.Sprintf("Rating for %s is %d", book.Name, randRating)
	log.Println(logMessage)
	return &pb.Rating{
		Rating: randRating,
	}, nil
}
