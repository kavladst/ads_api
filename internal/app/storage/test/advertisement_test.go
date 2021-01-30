package test

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"gorm.io/gorm"

	"github.com/kavladst/ads_api/internal/app/configuration"
	"github.com/kavladst/ads_api/internal/app/storage"
)

func TestStorage_GetAd(t *testing.T) {
	testCases := []struct {
		name          string
		adId          uuid.UUID
		adsInDB       []storage.Ad
		isError       bool
		expectedAd    storage.Ad
		expectedError error
	}{
		{
			name:       "get existing ad",
			adId:       validAds[0].ID,
			adsInDB:    []storage.Ad{validAds[0]},
			isError:    false,
			expectedAd: validAds[0],
		},
		{
			name:          "get not existing ad",
			adId:          validAds[0].ID,
			adsInDB:       []storage.Ad{validAds[1]},
			isError:       true,
			expectedError: gorm.ErrRecordNotFound,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testStorage, teardown := setup(testCase.adsInDB)
			defer teardown()
			testAd, err := testStorage.GetAd(testCase.adId)
			if testCase.isError {
				assert.Equal(t, err, testCase.expectedError)
			} else {
				assert.Equal(t, err, nil)
				assertEqualAds(t, []storage.Ad{testAd}, []storage.Ad{testCase.expectedAd})
			}
		})
	}
}

func TestStorage_GetAds(t *testing.T) {
	type testCaseType struct {
		name          string
		page          int
		perPage       int
		sortBy        string
		isReversed    bool
		adsInDB       []storage.Ad
		isError       bool
		expectedError error
	}
	var testCases []testCaseType
	for page := 1; page < 5; page++ {
		for perPage := 1; perPage < 5; perPage++ {
			for _, sortBy := range []string{"price"} {
				for _, isReversed := range []bool{true, false} {
					testCase := testCaseType{
						name:          fmt.Sprintf("pagination %d %d %s %t", page, perPage, sortBy, isReversed),
						page:          page,
						perPage:       perPage,
						sortBy:        sortBy,
						isReversed:    isReversed,
						adsInDB:       validAds,
						isError:       false,
						expectedError: nil,
					}
					testCase.adsInDB = make([]storage.Ad, len(validAds))
					copy(testCase.adsInDB, validAds)
					testCases = append(testCases, testCase)
				}
			}
		}
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testStorage, teardown := setup(testCase.adsInDB)
			defer teardown()
			testAds, err := testStorage.GetAds(testCase.page, testCase.perPage, testCase.sortBy, testCase.isReversed)
			if testCase.isError {
				assert.Equal(t, err, testCase.expectedError)
			} else {
				assert.Equal(t, err, nil)
				expectedAds, err := getExpectedListAds(
					testCase.adsInDB, testCase.page, testCase.perPage, testCase.sortBy, testCase.isReversed,
				)
				assert.Equal(t, err, nil)
				assertEqualAds(t, testAds, expectedAds)
			}
		})
	}
}

var validAds = []storage.Ad{
	{
		ID:          uuid.MustParse("7c60388c-ae4d-4854-8ee2-fbf82422a05f"),
		Name:        "name_1",
		Description: "description_1",
		Price:       10,
		PhotosUrls:  []string{},
		CreatedAt:   time.Time{},
	},
	{
		ID:          uuid.MustParse("4ccbdfd8-088c-4cd2-b7df-07e6f6093536"),
		Name:        "name_2",
		Description: "description_2",
		Price:       20,
		PhotosUrls:  []string{},
		CreatedAt:   time.Time{},
	},
}

func setup(ads []storage.Ad) (*storage.Storage, func()) {
	config, err := configuration.New()
	if err != nil {
		log.Fatal(err)
	}
	testStorage, err := storage.New(config)
	if err != nil {
		log.Fatal(err)
	}

	testStorage.DB.Create(&ads)
	return testStorage, func() {
		testStorage.DB.Exec("TRUNCATE TABLE ads")
	}
}

func assertEqualAds(t *testing.T, ads1, ads2 []storage.Ad) {
	assert.Equal(t, len(ads1), len(ads2))

	adsCopy1 := make([]storage.Ad, len(ads1))
	copy(adsCopy1, ads1)
	adsCopy2 := make([]storage.Ad, len(ads2))
	copy(adsCopy2, ads2)
	for i := 0; i < len(ads1); i++ {
		adsCopy1[i].CreatedAt = adsCopy2[i].CreatedAt
	}
	assert.Equal(t, adsCopy1, adsCopy2)
}

func getExpectedListAds(ads []storage.Ad, page int, perPage int, sortedBy string, isReversed bool) ([]storage.Ad, error) {
	if sortedBy != "price" && sortedBy != "created_at" {
		return nil, errors.New("sort by must be \"price\" or \"created_at\"")
	}
	if page <= 0 {
		return nil, errors.New("page must be positive int")
	}
	if perPage <= 0 {
		return nil, errors.New("per page must be positive int")
	}

	resAds := make([]storage.Ad, len(ads))
	copy(resAds, ads)

	sort.Slice(resAds, func(i, j int) bool {
		var res bool
		if sortedBy == "price" {
			res = resAds[i].Price < resAds[j].Price
		} else if sortedBy == "created_at" {
			res = resAds[i].CreatedAt.Before(resAds[j].CreatedAt)
		} else {
			panic(errors.New("sort by must be \"price\" or \"created_at\""))
		}
		if isReversed {
			res = !res
		}
		return res
	})
	offset := (page - 1) * perPage
	if offset < len(resAds) {
		lastIndex := offset + perPage
		if len(resAds) < lastIndex {
			lastIndex = len(resAds)
		}
		return resAds[offset:lastIndex], nil
	}
	return []storage.Ad{}, nil
}
