package service

import (
	"stickerapp/sticker/domain"
	"stickerapp/sticker/repository"

	"github.com/spf13/viper"
)

// interface for sticker serivce
type StickerService interface {
	FindByName(name string, limit int) domain.CDN
	FindAll(limit int) domain.CDN
}

// stuct for service
type stickerService struct {
	stickerRepository repository.StickerRepository
}

// constructor for sticker service
func New(repo repository.StickerRepository) StickerService {
	return &stickerService{
		stickerRepository: repo,
	}
}

// find all method to use repository to find all stickers
func (service *stickerService) FindAll(limit int) domain.CDN {
	result := service.stickerRepository.FindAll(limit)
	return findCDN(result)
}

// find by name method to use repository to find sticker by name
func (service *stickerService) FindByName(name string, limit int) domain.CDN {
	result := service.stickerRepository.FindByName(name, limit)
	return findCDN(result)
}

// find CDN method to find CDN of the sticker image
// it is a private method so no need to add in interface
func findCDN(stickers []domain.Sticker) domain.CDN {
	// get cdn url and id from config files
	cdnUrl := viper.GetString("CDN_URL")
	cdnId := viper.GetString("CDN_ID")
	// create a new variable of the struct cdn to store parsed values
	var cdn domain.CDN
	// iterate over stickers data and calculate CDN
	for _, sticker := range stickers {
		cdn.Images = append(cdn.Images, cdnUrl+cdnId+"/"+sticker.StickerName)
	}
	return cdn
}
