package tests

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/entities/gatheringPlace/query"
	repositoryPlaces "cmd/app/entities/gatheringPlace/repository"
	"cmd/app/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

var db *sql.DB

func setUpConnection() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "admin"
		password = "admin"
		dbname   = "foodate"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func closeConnection() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	setUpConnection()
	_ = m.Run()
	closeConnection()
}
func TestCreatingGatheringPlace(t *testing.T) {
	//set up
	var databasePlacesRepository = repositoryPlaces.NewPlacesDatabaseRepository(db)
	var currentPlace = gatheringPlace.NewGatheringPlace()
	currentPlace.CuisineType = gatheringPlace.FastFood
	var ctx = context.Background()

	//main part
	_, errCreating := databasePlacesRepository.Create(ctx, currentPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	placeFromRepository, errFinding := databasePlacesRepository.FindByID(ctx, currentPlace.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing equality of found and created entities
	assert.Equal(t, currentPlace, placeFromRepository)
	errDeleting := databasePlacesRepository.Delete(ctx, currentPlace)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
}

func TestFindingByCriteriaGatheringPlace(t *testing.T) {
	//set up
	var databasePlacesRepository = repositoryPlaces.NewPlacesDatabaseRepository(db)
	var firstPlace = gatheringPlace.NewGatheringPlace()
	firstPlace.CuisineType = gatheringPlace.FastFood
	var ctx = context.Background()
	_, errCreating := databasePlacesRepository.Create(ctx, firstPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	var secondPlace = gatheringPlace.NewGatheringPlace()
	secondPlace.CuisineType = gatheringPlace.Eastern
	_, errCreating = databasePlacesRepository.Create(ctx, secondPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	var thirdPlace = gatheringPlace.NewGatheringPlace()
	thirdPlace.CuisineType = gatheringPlace.FastFood
	_, errCreating = databasePlacesRepository.Create(ctx, thirdPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	var findingCriteria = query.FindCriteria{CuisineType: gatheringPlace.FastFood}

	//main part
	placesWithFastFood, errFinding := databasePlacesRepository.FindByCriteria(ctx, findingCriteria)
	findingCriteria.CuisineType = gatheringPlace.Eastern
	placesWithEastern, errFinding := databasePlacesRepository.FindByCriteria(ctx, findingCriteria)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing
	// first - check, if a placesWithFastFood contains right restaurants
	assert.True(t, slices.Contains(placesWithFastFood, *firstPlace))
	assert.True(t, slices.Contains(placesWithFastFood, *thirdPlace))
	assert.False(t, slices.Contains(placesWithFastFood, *secondPlace))
	// second - check, if a placesWithEastern contains right restaurants
	assert.True(t, slices.Contains(placesWithEastern, *secondPlace))
	assert.False(t, slices.Contains(placesWithFastFood, *secondPlace))
	errDeleting := databasePlacesRepository.Delete(ctx, firstPlace)
	errDeleting = databasePlacesRepository.Delete(ctx, secondPlace)
	errDeleting = databasePlacesRepository.Delete(ctx, thirdPlace)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
}

func TestUpdatingGatheringPlace(t *testing.T) {
	//set up
	var databasePlacesRepository = repositoryPlaces.NewPlacesDatabaseRepository(db)
	var firstPlace = gatheringPlace.NewGatheringPlace()
	firstPlace.CuisineType = gatheringPlace.FastFood
	var ctx = context.Background()
	_, errCreating := databasePlacesRepository.Create(ctx, firstPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	//main part
	firstPlace.Address = models.Address{Country: "Romania", City: "Bucharest"}
	_, errUpdating := databasePlacesRepository.Update(ctx, firstPlace)
	if errUpdating != nil {
		t.Fatalf("Error in updating place: %v", errUpdating)
	}
	placeFromRepository, errFinding := databasePlacesRepository.FindByID(ctx, firstPlace.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing
	//checking the equality of found and updated entity
	assert.Equal(t, firstPlace, placeFromRepository)
	errDeleting := databasePlacesRepository.Delete(ctx, firstPlace)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
}
