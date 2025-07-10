package raw

import (
	"filter-service/commom"
	"log"
)

type RawService struct {
	rawRepository *RawRepository
}

func NewRawService(repository *RawRepository) *RawService {
	return &RawService{
		rawRepository: repository,
	}
}

func (s *RawService) SaveToBD(raw, valid string) error {
	rawData := RawData{
		Data:  raw,
		Valid: valid,
	}
	_, err := s.rawRepository.Save(&rawData)
	if err != nil {
		log.Println(commom.ErrSaveToBDRAw, err)
		return err
	}
	return nil
}
