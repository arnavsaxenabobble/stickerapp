package service_test

import (
	"stickerapp/sticker/delivery/http/service"
	"stickerapp/sticker/domain"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock structure of sticker repository interface
type MockRepository struct {
	mock.Mock
}

// implement mock methods of sticker repository
func (mock *MockRepository) FindByName(name string, limit int) []domain.Sticker {
	// set expectations to be called
	args := mock.Called()
	// basis the first argument, return a domain.Sticker array
	result := args.Get(0)
	return result.([]domain.Sticker)
}

// implement mock methods of sticker repository
func (mock *MockRepository) FindAll(limit int) []domain.Sticker {
	// set expectations to be called
	args := mock.Called()
	// basis the first argument, return a domain.Sticker array
	result := args.Get(0)
	return result.([]domain.Sticker)
}

// test the method FindByName
func TestValidateFindByName(t *testing.T) {
	// create an instance of sticker repository with mock repository
	mockRepo := new(MockRepository)
	// create a mock variable that will be used to return from service function
	var stickers domain.Sticker
	time := time.Now()
	stickers.Clicks = 10
	stickers.CreatedAt = time
	stickers.ID = 123
	stickers.StickerName = "mockSticker"
	stickers.Trending = 1

	// set expectations
	mockRepo.On("FindByName").Return([]domain.Sticker{stickers})

	// create new instance of sticker service with mock repository
	testService := service.New(mockRepo)
	// set default configuration using viper to set CDN related attributes
	viper.SetDefault("CDN_URL", "mockurl")
	viper.SetDefault("CDN_ID", "mockid")
	// internally tests findCDN method as well
	result := testService.FindByName("mockSticker", 10)
	// assert that object recieved is not null
	assert.NotNil(t, result)
	// Behavioural Expectations
	mockRepo.AssertExpectations(t)
	// finally check if the qualified CDN path is returned for the image
	assert.Equal(t, "mockurlmockid/mockSticker", result.Images[0])
}

// test the method FindAll
func TestValidateFindAll(t *testing.T) {
	// create an instance of sticker repository with mock repository
	mockRepo := new(MockRepository)

	// create a mock variable that will be used to return from service function
	var stickers domain.Sticker
	time := time.Now()
	stickers.Clicks = 10
	stickers.CreatedAt = time
	stickers.ID = 123
	stickers.StickerName = "mockSticker"
	stickers.Trending = 1

	// set expectations
	mockRepo.On("FindAll").Return([]domain.Sticker{stickers})

	// create new instance of sticker service with mock repository
	testService := service.New(mockRepo)
	// set default configuration using viper to set CDN related attributes
	viper.SetDefault("CDN_URL", "mockurl")
	viper.SetDefault("CDN_ID", "mockid")
	// internally tests findCDN method as well
	result := testService.FindAll(10)
	// Check if the object recieved is not null
	assert.NotNil(t, result)
	// Behavioural Expectations
	mockRepo.AssertExpectations(t)
	// finally check if the qualified CDN path is returned for the image
	assert.Equal(t, "mockurlmockid/mockSticker", result.Images[0])
}
